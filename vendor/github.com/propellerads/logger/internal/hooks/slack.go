package hooks

import (
	"bytes"
	"context"
	"encoding/json"
	"strconv"
	"sync"

	"github.com/hashicorp/go-cleanhttp"
	"github.com/slack-go/slack"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"

	"github.com/propellerads/logger/config"
)

type SlackHook struct {
	Level  zapcore.Level
	Config config.SlackHook

	client  *slack.Client
	encoder zapcore.Encoder

	once sync.Once
}

func NewSlackHook(config config.SlackHook, level zapcore.Level) *SlackHook {
	return &SlackHook{
		Config: config,
		Level:  level,
	}
}

var LevelColorMap = map[zapcore.Level]string{
	zapcore.DebugLevel: "#9b30ff",
	zapcore.InfoLevel:  "#468847",
	zapcore.WarnLevel:  "#c09853",
	zapcore.ErrorLevel: "#f0ad4e",
	zapcore.FatalLevel: "#ff7708",
	zapcore.PanicLevel: "#000000",
}

func (h *SlackHook) Execute(entry zapcore.Entry, fields []zapcore.Field) error {
	h.once.Do(func() {
		h.client = slack.New(h.Config.Token, slack.OptionHTTPClient(cleanhttp.DefaultPooledClient()))
		h.encoder = zapcore.NewJSONEncoder(zapcore.EncoderConfig{
			SkipLineEnding: true,
		})
	})
	if entry.Level < h.Level {
		return nil
	}

	var (
		color string
		ok    bool
		err   error
	)

	if color, ok = LevelColorMap[entry.Level]; !ok {
		color = "#ffffff"
	}

	val := entry.Stack
	if val == "" {
		if val, err = h.fieldsEncodePretty(fields); err != nil {
			return err
		}
	}

	var title string
	if entry.Caller.Defined {
		title = entry.Caller.String()
	}

	attachment := slack.Attachment{
		Title: entry.Level.CapitalString(),
		Text:  entry.Message,
		Color: color,
		Fields: []slack.AttachmentField{
			{
				Title: title,
				Value: val,
				Short: false,
			},
		},
		Ts: json.Number(strconv.FormatInt(entry.Time.Unix(), 10)),
	}

	_, _, err = h.client.PostMessageContext(
		context.Background(),
		h.Config.Channel,
		slack.MsgOptionUsername(h.Config.Username),
		slack.MsgOptionIconEmoji(h.Config.IconEmoji),
		slack.MsgOptionAttachments(attachment),
	)

	return err
}

func (h *SlackHook) fieldsEncodePretty(fields []zapcore.Field) (string, error) {
	var (
		eb  *buffer.Buffer
		pb  bytes.Buffer
		err error
	)

	if eb, err = h.encoder.EncodeEntry(zapcore.Entry{}, fields); err != nil {
		return "", err
	}
	if err = json.Indent(&pb, eb.Bytes(), "", "  "); err != nil {
		return "", err
	}

	return "```" + pb.String() + "```", nil
}
