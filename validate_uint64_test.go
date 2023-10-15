package proto_validate_reflect

import (
	"reflect"
	"testing"

	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func uint64Ptr(f uint64) *uint64 {
	return &f
}

func TestValidateUInt64(t *testing.T) {
	number1 := uint64(10)
	number2 := uint64(9)
	type args struct {
		value protoreflect.Value
		rules *validate.UInt64Rules
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
				rules: &validate.UInt64Rules{Const: uint64Ptr(number2)},
			},
			want1: []error{ErrorConst},
		},
		{
			name: "uint64 lt",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt64Rules{Lt: uint64Ptr(number2)},
			},
			want1: []error{ErrorLt},
		},
		{
			name: "uint64 lt equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt64Rules{Lt: uint64Ptr(number1)},
			},
			want: true,
		},
		{
			name: "uint64 lte",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt64Rules{Lte: uint64Ptr(number2)},
			},
			want1: []error{ErrorLte},
		},
		{
			name: "uint64 lte equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt64Rules{Lte: uint64Ptr(number1)},
			},
			want1: []error{ErrorLte},
		},
		// Gt
		{
			name: "uint64 gt",
			args: args{
				value: protoreflect.ValueOf(number2),
				rules: &validate.UInt64Rules{Gt: uint64Ptr(number1)},
			},
			want1: []error{ErrorGt},
		},
		{
			name: "uint64 gt equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt64Rules{Gt: uint64Ptr(number1)},
			},
			want: true,
		},
		{
			name: "uint64 gte",
			args: args{
				value: protoreflect.ValueOf(number2),
				rules: &validate.UInt64Rules{Gte: uint64Ptr(number1)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "uint64 gte equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt64Rules{Gte: uint64Ptr(number1)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "uint64 in set",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt64Rules{In: []uint64{number2}},
			},
			want1: []error{ErrorIn},
		},
		{
			name: "uint64 not in set",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.UInt64Rules{NotIn: []uint64{number1}},
			},
			want1: []error{ErrorNotIn},
		},
		{
			name: "uint64 ignore empty",
			args: args{
				value: protoreflect.ValueOf(uint64(0)),
				rules: &validate.UInt64Rules{IgnoreEmpty: boolPtr(true)},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateUInt64(tt.args.value, tt.args.rules)
			if got != tt.want {
				t.Errorf("ValidateUInt64() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidateUInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
