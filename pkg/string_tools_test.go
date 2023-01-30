package pkg

import (
	"testing"
)

func Test_randomString(t *testing.T) {
	tests := []struct {
		name   string
		length int
	}{
		{"test1", 10},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RandomString(tt.length)
			if len(got) != tt.length {
				t.Errorf("randomString() = %v, want %d", got, tt.length)
			}
		})
	}
}

func TestCreateSlug(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		text string
		want string
	}{
		{"Free text", "Free text", "free-text"},
		{"Irregular free text", " irregular  free text", "irregular-free-text"},
		{"CamelCase", "CamelCase", "camel-case"},
		{"Non alphanumeric", "Non& alphanumericº", "non-alphanumeric"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSlug(tt.text); got != tt.want {
				t.Errorf("CreateSlug() = %v, want %v", got, tt.want)
			}
		})
	}
}
