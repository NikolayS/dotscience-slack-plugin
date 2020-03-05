package main

import (
	"os"

	"go.uber.org/zap"

	"github.com/dotmesh-io/dotscience-slack-plugin/pkg/client"
	"github.com/dotmesh-io/dotscience-slack-plugin/pkg/config"
	"github.com/dotmesh-io/dotscience-slack-plugin/pkg/logger"
)

func main() {

	log := logger.GetLoggerInstance(zap.InfoLevel)

	conf := config.MustLoad()

	if conf.SlackURL == "" {
		log.Fatal("Slack URL not set, cannot continue")
		os.Exit(1)
	}

	slackClient := client.New(log, conf)

	log.With(
		zap.String("host", conf.SlackURL),
		zap.String("username", conf.Project),
		zap.String("project", conf.ProjectID),
		zap.String("vcs_type", conf.TaskID),
		zap.String("revision", conf.RunnerID),
		zap.String("tag", conf.Status),
		zap.String("branch", conf.DotscienceHost),
	).Info("client initialized, sending Slack notification")

	err := slackClient.Notify()
	if err != nil {
		log.With(zap.Error(err)).Fatal("failed to send Slack notification")
		os.Exit(1)
	}

	log.Info("notification sent, exiting.")
	// done
}
