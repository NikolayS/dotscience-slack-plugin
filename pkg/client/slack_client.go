package client

import (
	"github.com/nlopes/slack"
	"go.uber.org/zap"

	"github.com/dotmesh-io/dotscience-slack-plugin/pkg/config"
)

type SlackClient struct {
	cfg         config.Config
	slackClient *slack.Client
	logger      *zap.Logger
}

func (c *SlackClient) Notify() error {

	msg := &slack.WebhookMessage{}

	return slack.PostWebhook(c.cfg.SlackURL, msg)
}
