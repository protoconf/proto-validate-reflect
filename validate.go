package proto_validate_reflect

import (
	"fmt"
	"strings"

	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

type ValidateError struct {
	name   protoreflect.Name
	errors map[protoreflect.Name][]error
}

func (v *ValidateError) Func(name protoreflect.Name, f func() (bool, []error)) {
	if ok, errs := f(); !ok {
		v.errors[name] = errs
	}
}

func (v *ValidateError) Error() string {
	ret := []string{fmt.Sprintf(`Message "%s" has errors:`, v.name)}
	for field, errs := range v.errors {
		if len(errs) > 0 {
			ret = append(ret, fmt.Sprintf(`field "%s" has the following errors:`, field))
			for _, err := range errs {
				ret = append(ret, fmt.Sprintf("\t%s", err))
			}
		}
	}

	return strings.Join(ret, "\n")
}

func (v *ValidateError) Ok() bool {
	return len(v.errors) == 0
}

func (v *ValidateError) Err() error {
	if !v.Ok() {
		return v
	}
	return nil
}

func Validate(msg proto.Message) (bool, error) {
	b, _ := proto.Marshal(msg)
	dyn := dynamicpb.NewMessage(msg.ProtoReflect().Descriptor())
	proto.Unmarshal(b, dyn)
	return ValidateDynamic(dyn)
}

func ValidateDynamic(m *dynamicpb.Message) (bool, error) {
	errors := &ValidateError{
		name:   m.Descriptor().FullName().Name(),
		errors: map[protoreflect.Name][]error{},
	}

	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		rules := fieldDescriptorToRules(fd)
		switch x := rules.Type.(type) {
		case *validate.FieldRules_String_:
			errors.Func(fd.Name(), func() (bool, []error) {
				return ValidateString(v, x.String_)
			})
		case *validate.FieldRules_Float:
			errors.Func(fd.Name(), func() (bool, []error) {
				return ValidateFloat(v, x.Float)
			})
		case *validate.FieldRules_Double:
			errors.Func(fd.Name(), func() (bool, []error) {
				return ValidateDouble(v, x.Double)
			})
		case *validate.FieldRules_Int32:
			errors.Func(fd.Name(), func() (bool, []error) {
				return ValidateInt32(v, x.Int32)
			})
		case *validate.FieldRules_Int64:
			errors.Func(fd.Name(), func() (bool, []error) {
				return ValidateInt64(v, x.Int64)
			})
		case *validate.FieldRules_Uint32:
			errors.Func(fd.Name(), func() (bool, []error) {
				return ValidateUInt32(v, x.Uint32)
			})
		case *validate.FieldRules_Uint64:
			errors.Func(fd.Name(), func() (bool, []error) {
				return ValidateUInt64(v, x.Uint64)
			})
		}
		return true
	})
	return errors.Ok(), errors.Err()
}

func fieldDescriptorToRules(fd protoreflect.FieldDescriptor) *validate.FieldRules {
	ext := fd.Options().(*descriptorpb.FieldOptions)
	b, _ := proto.Marshal(ext)
	opts := &validate.FieldOptions{}
	proto.Unmarshal(b, opts)
	return opts.Rules
}
