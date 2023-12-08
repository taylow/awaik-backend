package tui

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveEmoji(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{
			name: "no emoji",
			s:    "foo",
			want: "foo",
		},
		{
			name: "emoji",
			s:    "👋 foo",
			want: "foo",
		},
		{
			name: "emoji with double space",
			s:    "👋  foo",
			want: "foo",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, removeEmoji(tt.s))
		})
	}
}
