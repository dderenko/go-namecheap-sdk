package config

import (
	"go.uber.org/zap"
)

type SamplingConfig = zap.SamplingConfig

type Logger struct {
	Level            string         `yaml:"level"`        // Log level. Possible values: `debug`, `info`, `warn`, `error`, `dpanic`, `panic`, `fatal`
	Encoding         string         `yaml:"encoding"`     // If encoding is `console`: logger uses development config. Sampling is disabled.
	TimeEncoder      string         `yaml:"time_encoder"` // Time encoder. Possible values: `iso8601`, `epoch` (DEFAULT), `epoch_ms`, `rfc3339`, `rfc3339_nano`
	Tags             []string       `yaml:"tags"`
	Sampling         SamplingConfig `yaml:"sampling"`      // Warn: initial: 0, thereafter: 0 - disables logging.
	UseDiode         bool           `yaml:"use_diode"`     // Use github.com/propellerads/logdiode writer. Need for small amount of particular cases
	EnableCaller     bool           `yaml:"enable_caller"` // annotate logs with the calling function's file and line number
	EnableStacktrace bool           `yaml:"enable_stacktrace"`
	StacktraceAll    bool           `yaml:"stacktrace_all"` // by default stacktraces are only logged for errors and up
	Slack            SlackHook      `yaml:"slack"`          // Slack hook settings
}

type SlackHook struct {
	Enabled   bool   `yaml:"enabled"`
	Level     string `yaml:"level"`
	Token     string `yaml:"token" json:"token"`
	Channel   string `yaml:"channel" json:"channel"`
	Username  string `yaml:"username" json:"username"`
	IconEmoji string `yaml:"icon_emoji"`
}
