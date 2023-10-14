package main

import (
	"testing"

	"github.com/protoconf/proto-validate-reflect/cases"
	"google.golang.org/protobuf/proto"
)

func TestValidate(t *testing.T) {
	type args struct {
		msg proto.Message
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "test email",
			args:    args{msg: &cases.StringEmail{Val: "myemail.com"}},
			wantErr: true,
		},
		{
			name:    "test float",
			args:    args{msg: &cases.FloatConst{Val: 99.99}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Validate(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
