package main

import (
	"fmt"

	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"
)

func Validate(msg proto.Message) error {
	m := msg.ProtoReflect()
	errors := map[protoreflect.Name][]error{}

	m.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		rules := fieldDescriptorToRules(fd)
		switch x := rules.Type.(type) {
		case *validate.FieldRules_String_:
			if ok, errs := ValidateString(v, x.String_); !ok {
				errors[fd.Name()] = errs
			}
		case *validate.FieldRules_Float:
			if ok, errs := ValidateFloat(v, x.Float); !ok {
				errors[fd.Name()] = errs
			}
		}
		return true
	})
	return fmt.Errorf("%v", errors)
}

func fieldDescriptorToRules(fd protoreflect.FieldDescriptor) *validate.FieldRules {
	ext := fd.Options().(*descriptorpb.FieldOptions)
	b, _ := proto.Marshal(ext)
	opts := &validate.FieldOptions{}
	proto.Unmarshal(b, opts)
	return opts.Rules
}
