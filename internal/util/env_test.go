package util

import (
	"os"
	"testing"
)

const envKey = "TEMP_UNIT_TEST_TestFallbackEnvString"

func TestFallbackEnvString(t *testing.T) {
	tests := []struct {
		name     string
		flag     string
		envKey   string
		envValue string
		want     string
	}{
		{
			name:     "Value is set, and env is also set",
			flag:     "test",
			envKey:   envKey,
			envValue: "changed",
			want:     "test",
		},
		{
			name:     "No value, but env is set",
			flag:     "",
			envKey:   envKey,
			envValue: "changed",
			want:     "changed",
		},
		{
			name:     "No value and no env variable",
			flag:     "",
			envKey:   envKey,
			envValue: "",
			want:     "",
		},
	}
	for _, c := range tests {
		defer os.Unsetenv(c.envKey)
		os.Setenv(c.envKey, c.envValue)

		actual := FallbackEnvString(c.flag, c.envKey)

		if actual != c.want {
			t.Errorf("Case %q: expected %q, but got %q", c.name, c.want, actual)
		}
	}
}
