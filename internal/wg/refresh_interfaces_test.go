package wg

import (
	"reflect"
	"testing"
)

func TestParseActiveInterfaces(t *testing.T) {
	tests := []struct {
		name   string
		output string
		want   []InterfaceStatus
	}{
		{
			name:   "single interface",
			output: "wg0",
			want: []InterfaceStatus{
				{Name: "wg0"},
			},
		},
		{
			name:   "multiple interfaces",
			output: "wg0 wg1",
			want: []InterfaceStatus{
				{Name: "wg0"},
				{Name: "wg1"},
			},
		},
		{
			name:   "extra whitespace",
			output: "  wg0\twg1  ",
			want: []InterfaceStatus{
				{Name: "wg0"},
				{Name: "wg1"},
			},
		},
		{
			name:   "empty output",
			output: "",
			want:   []InterfaceStatus{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := parseActiveInterfaces(tt.output)

			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("parseActiveInterfaces(%q) = %#v, want %#v", tt.output, got, tt.want)
			}
		})
	}
}
