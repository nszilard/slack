package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/nszilard/slack/internal/slack"
	"github.com/nszilard/slack/models"

	"github.com/spf13/cobra"
)

func Test_sendChangelogMessage(t *testing.T) {
	ch := "testChannel"
	c1 := "- feat: some commit"
	c2 := "- chore: another commit"
	expected := fmt.Sprintf("```%v\n%v```", c1, c2)
	changelogCommits = fmt.Sprintf("%v\n%v", c1, c2)

	client.WithChannel(ch)
	err := sendChangelogMessage(changelogCommand, nil)
	if err != nil {
		t.Errorf("send changelog message: unexpected error: %v", err)
	}

	actual := slack.Message{}
	json.Unmarshal(client.HTTPClient.(*mockHTTPClient).pl, &actual)

	if actual.Channel != ch {
		t.Errorf("send changelog message: unexpected channel %q (expected: %q)", actual.Channel, ch)
	}

	if actual.Blocks[4].Text.Text != expected {
		t.Errorf("send changelog message: unexpected commits: %q (expected: %q)", actual.Blocks[4].Text.Text, expected)
	}
}

func Test_checkChangelog(t *testing.T) {
	tests := []struct {
		name    string
		service string
		version string
		commits string
		cmd     *cobra.Command
		args    []string
		wantErr bool
	}{
		{
			name:    "Should error for empty ChangelogService",
			cmd:     mainCmd,
			service: "",
			version: "test",
			commits: "test",
			wantErr: true,
		},
		{
			name:    "Should error for empty ChangelogVersion",
			cmd:     mainCmd,
			service: "test",
			version: "",
			commits: "test",
			wantErr: true,
		},
		{
			name:    "Should error for empty changelogCommits and nil args",
			cmd:     mainCmd,
			service: "test",
			version: "test",
			commits: "",
			wantErr: true,
		},
		{
			name:    "Should error for empty changelogCommits and empty args",
			cmd:     mainCmd,
			service: "test",
			version: "test",
			commits: "",
			args:    []string{},
			wantErr: true,
		},
		{
			name:    "Shouldn't return error if all values are set (commits as env value)",
			cmd:     mainCmd,
			service: "test",
			version: "test",
			commits: "- test: Some test",
		},
		{
			name:    "Shouldn't return error if all values are set (commits as arg value)",
			cmd:     mainCmd,
			service: "test",
			version: "test",
			commits: "",
			args:    []string{"- test: Some commit"},
		},
	}
	for _, tt := range tests {
		changelogService = tt.service
		changelogVersion = tt.version
		os.Setenv(models.ArgChangelogCommitsEnv, tt.commits)

		if err := checkChangelog(tt.cmd, tt.args); (err != nil) != tt.wantErr {
			t.Errorf("Case %q: error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
