package config

import (
	"strings"
)

type (
	// Config stores the configuration settings.
	Config struct {
		Project   string `envconfig:"PLUGIN_PROJECT"`
		ProjectID string `envconfig:"PLUGIN_PROJECT_ID"`
		TaskID    string `envconfig:"PLUGIN_TASK_ID"`
		// Optional link to the runner ID
		RunnerID string `envconfig:"PLUGIN_RUNNER_ID"`
		// Workload/pipeline status: [ok, error, terminated]
		Status Status `envconfig:"PLUGIN_STATUS"`
		// Used by the plugin to produce a clickable link
		// back to the project and task
		DotscienceHost string `envconfig:"PLUGIN_DOTSCIENCE_HOST"`

		SlackURL string `envconfig:"PLUGIN_SLACK_URL"`
		IconURL  string `envconfig:"PLUGIN_ICON_URL"`
		Channel  string `envconfig:"PLUGIN_CHANNEL"`

		// User supplied template like:
		// "my build failed, task ID: {{ .TaskID }}, status: {{ .Status }}"
		Template string `envconfig:"PLUGIN_TEMPLATE"`
	}
)

type Status string

func (s Status) String() string {
	return string(s)
}

func (s Status) Success() bool {
	return strings.ToLower(string(s)) == "success"
}

func (s Status) Error() bool {
	return strings.ToLower(string(s)) == "failure"
}

func (s Status) Terminated() bool {
	return strings.ToLower(string(s)) == "terminated"
}
