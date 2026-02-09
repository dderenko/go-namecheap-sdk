Package diode provides a thread-safe, lock-free, non-blocking io.Writer wrapper

Example:
```
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"diode://"}
	err = zap.RegisterSink("diode", func(url *url.URL) (zap.Sink, error) {
		return logdiode.NewWriter(os.Stdout, 100, time.Second, func(missed int) {
			fmt.Printf("{\"msg\": \"skipped %d lines\"}\n", missed)
		}), nil
	})
	if err != nil {
		return nil, err
	}
```