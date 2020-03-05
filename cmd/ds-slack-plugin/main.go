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

	plugin := client.New(log, conf)

	log.With(
		zap.String("slack_url", conf.SlackURL),
		zap.String("project", conf.Project),
		zap.String("project", conf.ProjectID),
		zap.String("task_id", conf.TaskID),
		zap.String("runner_id", conf.RunnerID),
		zap.String("status", conf.Status.String()),
		zap.String("dotscience_host", conf.DotscienceHost),
	).Info("client initialized, sending Slack notification")

	err := plugin.Exec()
	if err != nil {
		log.With(zap.Error(err)).Fatal("failed to send Slack notification")
		os.Exit(1)
	}

	log.Info("notification sent, exiting.")
	// done
}
