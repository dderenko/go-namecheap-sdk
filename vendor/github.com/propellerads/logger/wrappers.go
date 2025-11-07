package logger

import (
	"errors"
	"slices"

	"github.com/propellerads/logger/field"
)

type WrapperAmbassadorV0 struct {
	*Logger
}

func (w *WrapperAmbassadorV0) Infow(msg string, keysAndValues ...any) {
	w.Logger.Info(msg, transformToFields(keysAndValues)...)
}

func (w *WrapperAmbassadorV0) Warnw(msg string, keysAndValues ...any) {
	w.Logger.Warn(msg, transformToFields(keysAndValues)...)
}

func (w *WrapperAmbassadorV0) Errorw(msg string, keysAndValues ...any) {
	w.Logger.Error(msg, transformToFields(keysAndValues)...)
}

type WrapperAmbassadorV1 struct {
	*Logger
}

func (l *WrapperAmbassadorV1) WithService(name string) *WrapperAmbassadorV1 {
	l.Logger = l.With(field.Service(name))

	return l
}

func (l *WrapperAmbassadorV1) Info(msg string, list []string) {
	l.Logger.Info(msg, field.Strings("addresses", list))
}

func (l *WrapperAmbassadorV1) Warn(msg string) {
	l.Logger.Warn(msg)
}

func (l *WrapperAmbassadorV1) Error(msg string, err error) {
	l.Logger.Error(msg, field.Error(err))
}

type WrapperAmbassadorV2 struct {
	*Logger
}

func (l *WrapperAmbassadorV2) WithService(name string) {
	l.Logger = l.With(field.Service(name))
}

func (l *WrapperAmbassadorV2) Info(msg string, list []string) {
	l.Logger.Info(msg, field.Strings("addresses", list))
}

func (l *WrapperAmbassadorV2) Warn(msg string) {
	l.Logger.Warn(msg)
}

func (l *WrapperAmbassadorV2) Error(msg string, err error) {
	l.Logger.Error(msg, field.Error(err))
}

type WrapperGrpcClientV0 = WrapperAmbassadorV0
type WrapperGrpcClientV1 = WrapperAmbassadorV1
type WrapperGrpcClientV11 = WrapperAmbassadorV2

type WrapperConsulKvV4 struct {
	*Logger
}

func (l *WrapperConsulKvV4) Error(msg string, featureKey string, err error) {
	l.Logger.Error(msg, field.Error(err), field.FeatureKey(featureKey))
}

type WrapperConsulV4 struct {
	*Logger
}

func (l *WrapperConsulV4) WithModule(name string) {
	l.Logger = l.Logger.With(field.Module(name))
}

func (l *WrapperConsulV4) Info(msg, instance string) {
	l.Logger.Info(msg, field.Instance(instance))
}

func (l *WrapperConsulV4) Fatal(msg, instance string, err error) {
	l.Logger.Fatal(msg, field.Instance(instance), field.Error(err))
}

func (l *WrapperConsulV4) Error(msg, instance string, err error) {
	l.Logger.Error(msg, field.Instance(instance), field.Error(err))
}

type WrapperTracingV0 struct {
	*Logger
}

func (t *WrapperTracingV0) WithModule(name string) {
	t.Logger = t.Logger.With(field.Module(name))
}

func (t *WrapperTracingV0) Info(msg string) {
	t.Logger.Info(msg)
}

func (t *WrapperTracingV0) Error(msg string, err error) {
	t.Logger.Error(msg, field.Error(err))
}

var (
	errOdd   = errors.New("odd number of arguments, expected key-value pairs")
	errNoKey = errors.New("first argument must be a string key")
)

func transformToFields(keysAndValues []any) []field.Field {
	length := len(keysAndValues)
	if length == 0 {
		return nil
	}

	if len(keysAndValues)%2 != 0 {
		return []field.Field{field.Error(errOdd)}
	}

	ff := make([]field.Field, 0, length)
	for pair := range slices.Chunk(keysAndValues, 2) {
		key, ok := pair[0].(string)
		if !ok {
			ff = append(ff, field.Error(errNoKey))
			continue
		}
		ff = append(ff, field.Any(key, pair[1]))
	}

	return ff
}
