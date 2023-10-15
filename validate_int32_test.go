package proto_validate_reflect

import (
	"reflect"
	"testing"

	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func int32Ptr(f int32) *int32 {
	return &f
}

func TestValidateInt32(t *testing.T) {
	number1 := int32(10)
	number2 := int32(9)
	type args struct {
		value protoreflect.Value
		rules *validate.Int32Rules
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
				rules: &validate.Int32Rules{Const: int32Ptr(number2)},
			},
			want1: []error{ErrorConst},
		},
		{
			name: "int32 lt",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int32Rules{Lt: int32Ptr(number2)},
			},
			want1: []error{ErrorLt},
		},
		{
			name: "int32 lt equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int32Rules{Lt: int32Ptr(number1)},
			},
			want: true,
		},
		{
			name: "int32 lte",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int32Rules{Lte: int32Ptr(number2)},
			},
			want1: []error{ErrorLte},
		},
		{
			name: "int32 lte equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int32Rules{Lte: int32Ptr(number1)},
			},
			want1: []error{ErrorLte},
		},
		// Gt
		{
			name: "int32 gt",
			args: args{
				value: protoreflect.ValueOf(number2),
				rules: &validate.Int32Rules{Gt: int32Ptr(number1)},
			},
			want1: []error{ErrorGt},
		},
		{
			name: "int32 gt equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int32Rules{Gt: int32Ptr(number1)},
			},
			want: true,
		},
		{
			name: "int32 gte",
			args: args{
				value: protoreflect.ValueOf(number2),
				rules: &validate.Int32Rules{Gte: int32Ptr(number1)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "int32 gte equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int32Rules{Gte: int32Ptr(number1)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "int32 in set",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int32Rules{In: []int32{number2}},
			},
			want1: []error{ErrorIn},
		},
		{
			name: "int32 not in set",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int32Rules{NotIn: []int32{number1}},
			},
			want1: []error{ErrorNotIn},
		},
		{
			name: "int32 ignore empty",
			args: args{
				value: protoreflect.ValueOf(int32(0)),
				rules: &validate.Int32Rules{IgnoreEmpty: boolPtr(true)},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateInt32(tt.args.value, tt.args.rules)
			if got != tt.want {
				t.Errorf("ValidateInt32() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidateInt32() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
