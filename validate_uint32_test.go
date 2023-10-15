package proto_validate_reflect

import (
	"reflect"
	"testing"

	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func uint32Ptr(f uint32) *uint32 {
	return &f
}

func TestValidateUInt32(t *testing.T) {
	number1 := uint32(10)
	number2 := uint32(9)
	type args struct {
		value protoreflect.Value
		rules *validate.UInt32Rules
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []error
	}{
		{
			name: "const",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt32Rules{Const: uint32Ptr(number2)},
			},
			want1: []error{ErrorConst},
		},
		{
			name: "uint32 lt",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt32Rules{Lt: uint32Ptr(number2)},
			},
			want1: []error{ErrorLt},
		},
		{
			name: "uint32 lt equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt32Rules{Lt: uint32Ptr(number1)},
			},
			want: true,
		},
		{
			name: "uint32 lte",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt32Rules{Lte: uint32Ptr(number2)},
			},
			want1: []error{ErrorLte},
		},
		{
			name: "uint32 lte equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt32Rules{Lte: uint32Ptr(number1)},
			},
			want1: []error{ErrorLte},
		},
		// Gt
		{
			name: "uint32 gt",
			args: args{
				value: protoreflect.ValueOf(number2),
				rules: &validate.UInt32Rules{Gt: uint32Ptr(number1)},
			},
			want1: []error{ErrorGt},
		},
		{
			name: "uint32 gt equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt32Rules{Gt: uint32Ptr(number1)},
			},
			want: true,
		},
		{
			name: "uint32 gte",
			args: args{
				value: protoreflect.ValueOf(number2),
				rules: &validate.UInt32Rules{Gte: uint32Ptr(number1)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "uint32 gte equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt32Rules{Gte: uint32Ptr(number1)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "uint32 in set",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt32Rules{In: []uint32{number2}},
			},
			want1: []error{ErrorIn},
		},
		{
			name: "uint32 not in set",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt32Rules{NotIn: []uint32{number1}},
			},
			want1: []error{ErrorNotIn},
		},
		{
			name: "uint32 ignore empty",
			args: args{
				value: protoreflect.ValueOf(uint32(0)),
				rules: &validate.UInt32Rules{IgnoreEmpty: boolPtr(true)},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateUInt32(tt.args.value, tt.args.rules)
			if got != tt.want {
				t.Errorf("ValidateUInt32() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidateUInt32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
