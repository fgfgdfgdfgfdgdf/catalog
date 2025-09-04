package utils

import "testing"

func TestLowercaseFirst(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "NaMe",
			},
			want: "naMe",
		},
		{
			name: "2",
			args: args{
				s: "Name",
			},
			want: "name",
		},
		{
			name: "3",
			args: args{
				s: "name",
			},
			want: "name",
		},
		{
			name: "4",
			args: args{
				s: "naMe",
			},
			want: "naMe",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LowercaseFirst(tt.args.s); got != tt.want {
				t.Errorf("LowercaseFirst() = %v, want %v", got, tt.want)
			}
		})
	}
}
