// Package diode provides a thread-safe, lock-free, non-blocking io.Writer wrapper.
package logdiode

import (
	"context"
	"io"
	"sync"
	"time"

	"code.cloudfoundry.org/go-diodes"
)

var bufPool = &sync.Pool{
	New: func() interface{} {
		return make([]byte, 0, 500)
	},
}

type Alerter func(missed int)

type diodeFetcher interface {
	diodes.Diode
	Next() diodes.GenericDataType
}

// Writer is a io.Writer wrapper that uses a diode to make Write lock-free,
// non-blocking and thread safe.
type Writer struct {
	w    io.Writer
	d    diodeFetcher
	c    context.CancelFunc
	done chan struct{}
}

func NewWriter(w io.Writer, size int, pollInterval time.Duration, f Alerter) Writer {
	ctx, cancel := context.WithCancel(context.Background())
	dw := Writer{
		w:    w,
		c:    cancel,
		done: make(chan struct{}),
	}
	if f == nil {
		f = func(int) {}
	}
	d := diodes.NewManyToOne(size, diodes.AlertFunc(f))
	if pollInterval > 0 {
		dw.d = diodes.NewPoller(d,
			diodes.WithPollingInterval(pollInterval),
			diodes.WithPollingContext(ctx),
		)
	} else {
		dw.d = diodes.NewWaiter(
			d,
			diodes.WithWaiterContext(ctx),
		)
	}
	go dw.poll()
	return dw
}

func (dw Writer) Write(p []byte) (n int, err error) {
	p = append(bufPool.Get().([]byte), p...)
	dw.d.Set(diodes.GenericDataType(&p))
	return len(p), nil
}

// Close releases the diode poller and call Close on the wrapped writer if
// io.Closer is implemented.
func (dw Writer) Close() error {
	dw.c()
	<-dw.done
	if w, ok := dw.w.(io.Closer); ok {
		return w.Close()
	}
	return nil
}

func (dw Writer) poll() {
	defer close(dw.done)
	for {
		d := dw.d.Next()
		if d == nil {
			return
		}
		p := *(*[]byte)(d)
		_, _ = dw.w.Write(p)

		// Proper usage of a sync.Pool requires each entry to have approximately
		// the same memory cost. To obtain this property when the stored type
		// contains a variably-sized buffer, we add a hard limit on the maximum buffer
		// to place back in the pool.
		//
		// See https://golang.org/issue/23199
		const maxSize = 1 << 16 // 64KiB
		if cap(p) <= maxSize {
			bufPool.Put(p[:0])
		}
	}
}

func (dw Writer) Sync() error {
	for {
		d, ok := dw.d.TryNext()
		if d == nil || !ok {
			return nil
		}

		p := *(*[]byte)(d)
		_, _ = dw.w.Write(p)
	}
}
