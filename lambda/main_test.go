package main

import (
	"testing"
)

func Test_handle(t *testing.T) {
	type args struct {
		event Event
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "add 12 and 8",
			args: args{event: Event{
				X: 12,
				Y: 8,
			}},
			want: "sum=20",
		},
		{
			name: "add 8 and 12",
			args: args{event: Event{
				X: 8,
				Y: 12,
			}},
			want: "sum=20",
		},
		{
			name: "add 0 and 5",
			args: args{event: Event{
				X: 0,
				Y: 5,
			}},
			want: "sum=5",
		},
		{
			name: "add -3 and 4",
			args: args{event: Event{
				X: -3,
				Y: 4,
			}},
			want: "sum=1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := handle(tt.args.event); got != tt.want {
				t.Errorf("handle() = %v, want %v", got, tt.want)
			}
		})
	}
}
