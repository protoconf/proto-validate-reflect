package proto_validate_reflect

import (
	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ValidateUInt32(value protoreflect.Value, rules *validate.UInt32Rules) (bool, []error) {
	f := uint32(value.Uint())
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
	if len(rules.In) > 0 && !uint32InSlice(f, rules.In) {
		errors = append(errors, ErrorIn)
	}
	if len(rules.NotIn) > 0 && uint32InSlice(f, rules.NotIn) {
		errors = append(errors, ErrorNotIn)
	}

	return len(errors) == 0, errors
}

func uint32InSlice(s uint32, slice []uint32) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
