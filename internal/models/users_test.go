package models

import (
	"github.com/hayohtee/snippetbox/internal/assert"
	"testing"
)

func TestUserModelExists(t *testing.T) {
	if testing.Short() {
		t.Skip("models: skipping integration test")
	}
	tests := []struct {
		name   string
		userId int
		want   bool
	}{
		{
			name:   "Valid ID",
			userId: 1,
			want:   true,
		},
		{
			name:   "Zero ID",
			userId: 0,
			want:   false,
		},
		{
			name:   "Non-existent ID",
			userId: 2,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDb(t)
			m := UserModel{DB: db}

			exists, err := m.Exists(tt.userId)

			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}
