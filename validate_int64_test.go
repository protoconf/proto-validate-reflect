package proto_validate_reflect

import (
	"reflect"
	"testing"

	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func int64Ptr(f int64) *int64 {
	return &f
}

func TestValidateInt64(t *testing.T) {
	number1 := int64(10)
	number2 := int64(9)
	type args struct {
		value protoreflect.Value
		rules *validate.Int64Rules
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
				rules: &validate.Int64Rules{Const: int64Ptr(number2)},
			},
			want1: []error{ErrorConst},
		},
		{
			name: "int64 lt",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int64Rules{Lt: int64Ptr(number2)},
			},
			want1: []error{ErrorLt},
		},
		{
			name: "int64 lt equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int64Rules{Lt: int64Ptr(number1)},
			},
			want: true,
		},
		{
			name: "int64 lte",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int64Rules{Lte: int64Ptr(number2)},
			},
			want1: []error{ErrorLte},
		},
		{
			name: "int64 lte equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int64Rules{Lte: int64Ptr(number1)},
			},
			want1: []error{ErrorLte},
		},
		// Gt
		{
			name: "int64 gt",
			args: args{
				value: protoreflect.ValueOf(number2),
				rules: &validate.Int64Rules{Gt: int64Ptr(number1)},
			},
			want1: []error{ErrorGt},
		},
		{
			name: "int64 gt equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int64Rules{Gt: int64Ptr(number1)},
			},
			want: true,
		},
		{
			name: "int64 gte",
			args: args{
				value: protoreflect.ValueOf(number2),
				rules: &validate.Int64Rules{Gte: int64Ptr(number1)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "int64 gte equals",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int64Rules{Gte: int64Ptr(number1)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "int64 in set",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int64Rules{In: []int64{number2}},
			},
			want1: []error{ErrorIn},
		},
		{
			name: "int64 not in set",
			args: args{
				value: protoreflect.ValueOf(number1),
				rules: &validate.Int64Rules{NotIn: []int64{number1}},
			},
			want1: []error{ErrorNotIn},
		},
		{
			name: "int64 ignore empty",
			args: args{
				value: protoreflect.ValueOf(int64(0)),
				rules: &validate.Int64Rules{IgnoreEmpty: boolPtr(true)},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateInt64(tt.args.value, tt.args.rules)
			if got != tt.want {
				t.Errorf("ValidateInt64() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidateInt64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
