package main

import (
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"

	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func ValidateString(value protoreflect.Value, rules *validate.StringRules) (bool, []error) {
	s := value.String()
	var errors []error

	if s == "" && *rules.IgnoreEmpty {
		return true, errors
	}

	if rules.Const != nil && s != *rules.Const {
		errors = append(errors, ErrorConst)
	}

	// Character length checks
	charLen := uint64(len([]rune(s)))
	if rules.Len != nil && charLen != *rules.Len {
		errors = append(errors, ErrorLen)
	}
	if rules.MinLen != nil && charLen < *rules.MinLen {
		errors = append(errors, ErrorMinLen)
	}
	if rules.MaxLen != nil && charLen > *rules.MaxLen {
		errors = append(errors, ErrorMaxLen)
	}

	// Byte length checks
	byteLen := uint64(len(s))
	if rules.LenBytes != nil && byteLen != *rules.LenBytes {
		errors = append(errors, ErrorLenBytes)
	}
	if rules.MinBytes != nil && byteLen < *rules.MinBytes {
		errors = append(errors, ErrorMinBytes)
	}
	if rules.MaxBytes != nil && byteLen > *rules.MaxBytes {
		errors = append(errors, ErrorMaxBytes)
	}

	// Pattern checks
	if rules.Pattern != nil {
		re := regexp.MustCompile(*rules.Pattern)
		if !re.MatchString(s) {
			errors = append(errors, ErrorPattern)
		}
	}

	// Substring checks
	if rules.Prefix != nil && !strings.HasPrefix(s, *rules.Prefix) {
		errors = append(errors, ErrorPrefix)
	}
	if rules.Suffix != nil && !strings.HasSuffix(s, *rules.Suffix) {
		errors = append(errors, ErrorSuffix)
	}
	if rules.Contains != nil && !strings.Contains(s, *rules.Contains) {
		errors = append(errors, ErrorContains)
	}
	if rules.NotContains != nil && strings.Contains(s, *rules.NotContains) {
		errors = append(errors, ErrorNotContains)
	}

	// Membership checks
	if len(rules.In) > 0 && !stringInSlice(s, rules.In) {
		errors = append(errors, ErrorIn)
	}
	if len(rules.NotIn) > 0 && stringInSlice(s, rules.NotIn) {
		errors = append(errors, ErrorNotIn)
	}

	// Well-known format checks
	switch rules.WellKnown.(type) {
	case *validate.StringRules_Email:
		if _, err := mail.ParseAddress(s); err != nil {
			errors = append(errors, ErrorInvalidEmail)
		}
	case *validate.StringRules_Hostname:
		host := strings.ToLower(strings.TrimSuffix(s, "."))

		if len(host) > 253 {
			errors = append(errors, ErrorHostnameTooLong)
		}

		for _, part := range strings.Split(s, ".") {
			if l := len(part); l == 0 || l > 63 {
				errors = append(errors, ErrorHostnameInvalidPart)
				continue
			}

			if part[0] == '-' {
				errors = append(errors, ErrorHostnamePartBeginWithHyphens)
			}

			if part[len(part)-1] == '-' {
				errors = append(errors, ErrorHostnamePartEndsWithHyphens)
			}

			for _, r := range part {
				if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
					errors = append(errors,
						fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r)),
					)
				}
			}
		}
	case *validate.StringRules_Ip:
		if net.ParseIP(s) == nil {
			errors = append(errors, ErrorInvalidIp)
		}
	case *validate.StringRules_Ipv4:

		if net.ParseIP(s).To4() == nil {
			errors = append(errors, ErrorInvalidIpv4)
		}
	case *validate.StringRules_Ipv6:
		if net.ParseIP(s).To16() == nil {
			errors = append(errors, ErrorInvalidIpv6)
		}
	case *validate.StringRules_Uri:
		if _, err := url.ParseRequestURI(s); err != nil {
			errors = append(errors, ErrorInvalidUri)
		}
		//... Handle other rules like UriRef, Address, Uuid, WellKnownRegex, Strict here
	}

	return len(errors) == 0, errors
}

func stringInSlice(s string, slice []string) bool {
	for _, v := range slice {
		if v == s {
			return true
		}
	}
	return false
}
