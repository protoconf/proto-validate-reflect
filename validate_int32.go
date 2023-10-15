package proto_validate_reflect

import (
	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ValidateInt32(value protoreflect.Value, rules *validate.Int32Rules) (bool, []error) {
	f := int32(value.Int())
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
	if len(rules.In) > 0 && !int32InSlice(f, rules.In) {
		errors = append(errors, ErrorIn)
	}
	if len(rules.NotIn) > 0 && int32InSlice(f, rules.NotIn) {
		errors = append(errors, ErrorNotIn)
	}

	return len(errors) == 0, errors
}

func int32InSlice(s int32, slice []int32) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
