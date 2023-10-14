package main

import (
	"reflect"
	"testing"

	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func doublePtr(f float64) *float64 {
	return &f
}

func TestValidateDouble(t *testing.T) {
	type args struct {
		value protoreflect.Value
		rules *validate.DoubleRules
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
				value: protoreflect.ValueOf(99.9),
				rules: &validate.DoubleRules{Const: doublePtr(9.99)},
			},
			want1: []error{ErrorConst},
		},
		{
			name: "float lt",
			args: args{
				value: protoreflect.ValueOf(99.9),
				rules: &validate.DoubleRules{Lt: doublePtr(9.99)},
			},
			want1: []error{ErrorLt},
		},
		{
			name: "float lt equals",
			args: args{
				value: protoreflect.ValueOf(99.9),
				rules: &validate.DoubleRules{Lt: doublePtr(99.9)},
			},
			want: true,
		},
		{
			name: "float lte",
			args: args{
				value: protoreflect.ValueOf(99.9),
				rules: &validate.DoubleRules{Lte: doublePtr(9.99)},
			},
			want1: []error{ErrorLte},
		},
		{
			name: "float lte equals",
			args: args{
				value: protoreflect.ValueOf(99.9),
				rules: &validate.DoubleRules{Lte: doublePtr(99.9)},
			},
			want1: []error{ErrorLte},
		},
		// Gt
		{
			name: "float gt",
			args: args{
				value: protoreflect.ValueOf(9.99),
				rules: &validate.DoubleRules{Gt: doublePtr(99.9)},
			},
			want1: []error{ErrorGt},
		},
		{
			name: "float gt equals",
			args: args{
				value: protoreflect.ValueOf(99.9),
				rules: &validate.DoubleRules{Gt: doublePtr(99.9)},
			},
			want: true,
		},
		{
			name: "float gte",
			args: args{
				value: protoreflect.ValueOf(9.99),
				rules: &validate.DoubleRules{Gte: doublePtr(99.9)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "float gte equals",
			args: args{
				value: protoreflect.ValueOf(99.9),
				rules: &validate.DoubleRules{Gte: doublePtr(99.9)},
			},
			want1: []error{ErrorGte},
		},
		{
			name: "float in set",
			args: args{
				value: protoreflect.ValueOf(99.9),
				rules: &validate.DoubleRules{In: []float64{9.9}},
			},
			want1: []error{ErrorIn},
		},
		{
			name: "float not in set",
			args: args{
				value: protoreflect.ValueOf(99.9),
				rules: &validate.DoubleRules{NotIn: []float64{99.9}},
			},
			want1: []error{ErrorNotIn},
		},
		{
			name: "float ignore empty",
			args: args{
				value: protoreflect.ValueOf(0.0),
				rules: &validate.DoubleRules{IgnoreEmpty: boolPtr(true)},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateDouble(tt.args.value, tt.args.rules)
			if got != tt.want {
				t.Errorf("ValidateFloat() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidateFloat() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
