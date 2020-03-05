package config

type (
	// Config stores the configuration settings.
	Config struct {
		Project   string `envconfig:"PLUGIN_PROJECT"`
		ProjectID string `envconfig:"PLUGIN_PROJECT_ID"`
		TaskID    string `envconfig:"PLUGIN_TASK_ID"`
		// Optional link to the runner ID
		RunnerID string `envconfig:"PLUGIN_RUNNER_ID"`
		// Workload/pipeline status: [ok, error, terminated]
		Status string `envconfig:"PLUGIN_STATUS"`
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
