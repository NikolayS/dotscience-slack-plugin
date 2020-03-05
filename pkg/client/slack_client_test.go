package client

import (
	"testing"

	"github.com/dotmesh-io/dotscience-slack-plugin/pkg/config"
)

func Test_templateMessage(t *testing.T) {
	type args struct {
		cfg         *config.Config
		templateStr string
	}

	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "default",
			args: args{
				cfg: &config.Config{
					Project: "foo",
					TaskID:  "101",
					Status:  "OK",
				},
				templateStr: defaultTemplate,
			},
			want:    "Dotscience project 'foo' task has completed.",
			wantErr: false,
		},
		{
			name: "error",
			args: args{
				cfg: &config.Config{
					Project: "foo",
					TaskID:  "101",
					Status:  "error",
				},
				templateStr: defaultTemplate,
			},
			want:    "Dotscience project 'foo' task has encountered an error.",
			wantErr: false,
		},
		{
			name: "terminated",
			args: args{
				cfg: &config.Config{
					Project: "foo",
					TaskID:  "101",
					Status:  "terminated",
				},
				templateStr: defaultTemplate,
			},
			want:    "Dotscience project 'foo' task has been terminated.",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := templateMessage(toTemplatePayload(*tt.args.cfg), tt.args.templateStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("templateMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("templateMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
