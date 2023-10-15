package proto_validate_reflect

import (
	"strings"
	"testing"

	"github.com/protoconf/proto-validate-reflect/cases"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
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
			name:    "test email is valid",
			args:    args{msg: &cases.StringEmail{Val: "my@email.com"}},
			wantErr: false,
		},
		{
			name:    "test float",
			args:    args{msg: &cases.FloatConst{Val: 99.99}},
			wantErr: true,
		},
		{
			name:    "test double",
			args:    args{msg: &cases.DoubleConst{Val: 99.99}},
			wantErr: true,
		},
		{
			name:    "test int32",
			args:    args{msg: &cases.Int32Const{Val: 99}},
			wantErr: true,
		},
		{
			name:    "test int64",
			args:    args{msg: &cases.Int64Const{Val: 99}},
			wantErr: true,
		},
		{
			name:    "test uint32",
			args:    args{msg: &cases.UInt32Const{Val: 99}},
			wantErr: true,
		},
		{
			name:    "test uint64",
			args:    args{msg: &cases.UInt64Const{Val: 99}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Validate(tt.args.msg); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateError_Error(t *testing.T) {
	type fields struct {
		name   protoreflect.Name
		errors map[protoreflect.Name][]error
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test",
			fields: fields{
				name: protoreflect.Name("HelloWorldMessage"),
				errors: map[protoreflect.Name][]error{
					protoreflect.Name("test_field"): {ErrorConst},
				},
			},
			want: strings.Join([]string{
				`Message "HelloWorldMessage" has errors:`,
				`field "test_field" has the following errors:`,
				"\tvalue does not match the const rule"}, "\n"),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &ValidateError{
				name:   tt.fields.name,
				errors: tt.fields.errors,
			}
			if got := v.Error(); got != tt.want {
				t.Errorf("ValidateError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
