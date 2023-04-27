package beegoorm

import (
	"testing"
)

func TestFormat(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "right",
			args: args{
				"[select * from user where id = ? or name = ?] - `1`, `dzw`",
			},
			want: "select * from user where id = '1' or name = 'dzw'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := Format(tt.args.str)
			if err != nil {
				t.Errorf("error = %v", err)
				return
			}
			if data != tt.want {
				t.Errorf("\ndata = 1%v1\n want = 1%v1", data, tt.want)
			}
			return
		})

	}

}
