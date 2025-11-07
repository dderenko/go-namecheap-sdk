package logger

import (
	"errors"
	"fmt"
	"net/url"
	"os"
	"time"

	"github.com/propellerads/logdiode"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/propellerads/logger/config"
	"github.com/propellerads/logger/field"
	"github.com/propellerads/logger/internal/hooks"
	"github.com/propellerads/logger/internal/metrics"
)

var (
	ErrEnvVarsNotProvided = errors.New("env vars not provided")
)

// Logger alias for zap.Logger
type (
	Logger        = zap.Logger
	SugaredLogger = zap.SugaredLogger
)

type EnvVars interface {
	Service() string
	Env() string
	ReleaseVersion() string
	CommitHash() string
}

// New creates a new logger with the given configuration.
func New(env EnvVars, config config.Logger) (*Logger, error) {
	if env == nil {
		return nil, ErrEnvVarsNotProvided
	}

	var (
		cfg    zap.Config
		lvl    zapcore.Level
		logger *zap.Logger
		slack  *hooks.SlackHook
		err    error
	)

	if config.Encoding == "console" {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
		cfg.DisableCaller = !config.EnableCaller
		cfg.Sampling.Initial = config.Sampling.Initial
		cfg.Sampling.Thereafter = config.Sampling.Thereafter
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

		if config.Slack.Enabled {
			lvl = zapcore.ErrorLevel
			if config.Slack.Level != "" {
				if err = lvl.Set(config.Slack.Level); err != nil {
					return nil, err
				}
			}
			slack = hooks.NewSlackHook(config.Slack, lvl)
		}
	}
	cfg.DisableStacktrace = !config.EnableStacktrace
	cfg.Encoding = config.Encoding
	cfg.OutputPaths = []string{"stderr"}
	switch config.TimeEncoder {
	case "iso8601":
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	case "epoch":
		cfg.EncoderConfig.EncodeTime = zapcore.EpochTimeEncoder
	case "epoch_ms":
		cfg.EncoderConfig.EncodeTime = zapcore.EpochMillisTimeEncoder
	case "rfc3339":
		cfg.EncoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	case "rfc3339_nano":
		cfg.EncoderConfig.EncodeTime = zapcore.RFC3339NanoTimeEncoder
	}

	if err = lvl.UnmarshalText([]byte(config.Level)); err != nil {
		return nil, err
	}
	cfg.Level.SetLevel(lvl)

	if config.UseDiode {
		cfg.OutputPaths = []string{"diode://"}
		err = zap.RegisterSink("diode", func(url *url.URL) (zap.Sink, error) {
			return logdiode.NewWriter(os.Stderr, 100, time.Second, func(missed int) {
				fmt.Printf(`{"msg": "skipped %d lines"}\n`, missed)
			}), nil
		})
		if err != nil {
			return nil, err
		}
	}

	opts := []zap.Option{
		zap.WrapCore(func(core zapcore.Core) zapcore.Core {
			return newInstrumentedCore(core, env, config, slack)
		}),
		zap.Fields(
			field.Service(env.Service()),
			field.Env(env.Env()),
			field.Version(env.ReleaseVersion()),
			field.Tags(config.Tags),
			field.Commit(env.CommitHash()),
		),
	}

	if config.StacktraceAll {
		opts = append(opts, zap.AddStacktrace(zapcore.InfoLevel), zap.AddStacktrace(zapcore.WarnLevel))
	}

	if logger, err = cfg.Build(opts...); err != nil {
		return nil, err
	}

	if err = metrics.Register(); err != nil {
		return nil, err
	}

	return logger, nil
}

var (
	NewNop         = zap.NewNop
	NewDevelopment = zap.NewDevelopment
	AddStacktrace  = zap.AddStacktrace
)

const (
	DebugLevel  = zapcore.DebugLevel
	InfoLevel   = zapcore.InfoLevel
	WarnLevel   = zapcore.WarnLevel
	ErrorLevel  = zapcore.ErrorLevel
	DPanicLevel = zapcore.DPanicLevel
	PanicLevel  = zapcore.PanicLevel
	FatalLevel  = zapcore.FatalLevel
)
