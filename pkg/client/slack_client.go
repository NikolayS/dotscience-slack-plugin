package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/nlopes/slack"
	"go.uber.org/zap"

	"github.com/dotmesh-io/dotscience-slack-plugin/pkg/config"
)

var defaultTemplate = `{{if .OK }}Dotscience project '{{.Project }}' task has completed.{{else}}
{{if .Error }}Dotscience project '{{.Project }}' task has encountered an error.{{else}}{{end}}
{{if .Terminated }}Dotscience project '{{.Project }}' task has been terminated.{{else}}{{end}}
{{end}}`

// TemplatePayload - passed into the template gen
type TemplatePayload struct {
	config.Config
	// Populated on load, used by the template
	OK         bool
	Terminated bool
	Error      bool
}

// SlackClient - sends notifications to slack chan
type SlackClient struct {
	cfg config.Config

	logger *zap.Logger
}

// New returns new SlackClient notification plugin
func New(logger *zap.Logger, cfg config.Config) *SlackClient {

	return &SlackClient{
		cfg:    cfg,
		logger: logger,
	}
}

func color(build string) string {
	switch build {
	case "error":
		return "#F44336"
	case "terminated":
		return "#2196F3"
	case "ok":
		return "#00C853"
	default:
		return "#9E9E9E"
	}
}

func title(status string) string {
	switch status {
	case "error":
		return "Dotscience pipeline error"
	case "terminated":
		return "Dotscience pipeline terminated"
	case "ok":
		return "Dotscience pipeline completed"
	default:
		return "#9E9E9E"
	}
}

func toTemplatePayload(cfg config.Config) *TemplatePayload {
	return &TemplatePayload{
		Config:     cfg,
		OK:         cfg.Status.OK(),
		Terminated: cfg.Status.Terminated(),
		Error:      cfg.Status.Error(),
	}
}

// Exec - runs the plugin triggering notification
func (c *SlackClient) Exec() error {

	var text string
	var err error

	templatePayload := toTemplatePayload(c.cfg)

	if c.cfg.Template != "" {
		text, err = templateMessage(templatePayload, c.cfg.Template)
	} else if c.cfg.Template == "" || err != nil {
		text, err = templateMessage(templatePayload, defaultTemplate)
		if err != nil {
			return err
		}
	}

	attachements := []slack.Attachment{
		slack.Attachment{
			Fallback:  text,
			Color:     color(c.cfg.Status.String()),
			Title:     title(c.cfg.Status.String()),
			TitleLink: fmt.Sprintf("%s/runner/%s/tasks", c.cfg.DotscienceHost, c.cfg.RunnerID),
			Fields: []slack.AttachmentField{
				slack.AttachmentField{
					Value: text,
					Short: false,
				},
			},
			Actions: []slack.AttachmentAction{
				{
					Name: "Runners",
					URL:  fmt.Sprintf("%s/runner/%s/tasks", c.cfg.DotscienceHost, c.cfg.RunnerID),
				},
			},
			Footer: fmt.Sprintf("%s", c.cfg.DotscienceHost),
			Ts:     json.Number(strconv.Itoa(int(time.Now().Unix()))),
		},
	}

	msg := &slack.WebhookMessage{
		IconURL:     c.cfg.IconURL,
		Channel:     c.cfg.Channel,
		Attachments: attachements,
	}

	return slack.PostWebhook(c.cfg.SlackURL, msg)
}

func templateMessage(cfg *TemplatePayload, templateStr string) (string, error) {
	t, err := template.New("notification").Parse(templateStr)
	if err != nil {
		return "", err
	}

	buf := bytes.NewBuffer([]byte(""))

	err = t.Execute(buf, cfg)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(buf.String()), nil
}
