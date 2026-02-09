package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/propellerads/logger/config"
	"github.com/propellerads/logger/internal/hooks"
	"github.com/propellerads/logger/internal/metrics"
)

type instrumentedCore struct {
	zapcore.Core

	env    EnvVars
	config config.Logger
	hook   *hooks.SlackHook
}

func newInstrumentedCore(
	core zapcore.Core,
	env EnvVars,
	config config.Logger,
	hook *hooks.SlackHook,
) *instrumentedCore {
	return &instrumentedCore{
		Core:   core,
		env:    env,
		config: config,
		hook:   hook,
	}
}

func (c *instrumentedCore) With(fields []zapcore.Field) zapcore.Core {
	return newInstrumentedCore(c.Core.With(fields), c.env, c.config, c.hook)
}

func (c *instrumentedCore) Check(entry zapcore.Entry, checked *zapcore.CheckedEntry) *zapcore.CheckedEntry {
	out := c.Core.Check(entry, checked)
	if out != nil {
		out.AddCore(entry, c)
	} else {
		metrics.LogsSkipped.WithLabelValues(c.env.Service(), entry.Level.String()).Inc()
	}
	return out
}

func (c *instrumentedCore) Write(entry zapcore.Entry, fields []zapcore.Field) error {
	if entry.Level >= zap.ErrorLevel {
		metrics.Logs.WithLabelValues(c.env.Service(), entry.Level.String(), entry.Message).Inc()
	}

	if c.hook != nil {
		return c.hook.Execute(entry, fields)
	}

	return nil
}
