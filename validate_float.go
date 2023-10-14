package main

import (
	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

type FloatRules struct {
	Const       *float32
	Lt          *float32
	Lte         *float32
	Gt          *float32
	Gte         *float32
	In          []float32
	NotIn       []float32
	IgnoreEmpty bool
}

func ValidateFloat(value protoreflect.Value, rules *validate.FloatRules) (bool, []error) {
	f := float32(value.Float())
	var errors []error

	// Check for empty value
	if f == 0 && *rules.IgnoreEmpty {
		return true, errors
	}

	// Check for const rule
	if rules.Const != nil && f != *rules.Const {
		errors = append(errors, ErrorConst)
	}

	// Check for lt rule
	if rules.Lt != nil && f > *rules.Lt {
		errors = append(errors, ErrorLt)
	}

	// Check for lte rule
	if rules.Lte != nil && f >= *rules.Lte {
		errors = append(errors, ErrorLte)
	}

	// Check for gt rule
	if rules.Gt != nil && f < *rules.Gt {
		errors = append(errors, ErrorGt)
	}

	// Check for gte rule
	if rules.Gte != nil && f <= *rules.Gte {
		errors = append(errors, ErrorGte)
	}

	// Membership checks
	if len(rules.In) > 0 && !floatInSlice(f, rules.In) {
		errors = append(errors, ErrorIn)
	}
	if len(rules.NotIn) > 0 && floatInSlice(f, rules.NotIn) {
		errors = append(errors, ErrorNotIn)
	}

	return len(errors) == 0, errors
}

func floatInSlice(s float32, slice []float32) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
