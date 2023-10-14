package main

import (
	"errors"
	"reflect"
	"testing"

	"github.com/protoconf/proto-validate-reflect/validate"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func stringPtr(s string) *string {
	return &s
}

func boolPtr(b bool) *bool {
	return &b
}

func uint64Ptr(u uint64) *uint64 {
	return &u
}
func TestValidateString(t *testing.T) {
	type args struct {
		value protoreflect.Value
		rules *validate.StringRules
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []error
	}{
		{
			name: "email",
			args: args{value: protoreflect.ValueOf("my@company.com"), rules: &validate.StringRules{
				WellKnown: &validate.StringRules_Email{Email: true},
			}},
			want: true,
		},
		{
			name: "invalid email",
			args: args{value: protoreflect.ValueOf("invalid.company.com"), rules: &validate.StringRules{
				WellKnown: &validate.StringRules_Email{Email: true},
			}},
			want1: []error{ErrorInvalidEmail},
		},
		{
			name: "const is not exact",
			args: args{value: protoreflect.ValueOf("abcdefg"), rules: &validate.StringRules{
				Const: stringPtr("abc"),
			}},
			want1: []error{ErrorConst},
		},
		{
			name: "empty",
			args: args{value: protoreflect.ValueOf(""), rules: &validate.StringRules{IgnoreEmpty: boolPtr(true)}},
			want: true,
			// want1: []error{},
		},
		{
			name: "length is not exact",
			args: args{value: protoreflect.ValueOf("abcdefg"), rules: &validate.StringRules{
				Len: uint64Ptr(4),
			}},
			want1: []error{ErrorLen},
		},
		{
			name: "min length is not exact",
			args: args{value: protoreflect.ValueOf("abcdefg"), rules: &validate.StringRules{
				MinLen: uint64Ptr(40),
			}},
			want1: []error{ErrorMinLen},
		},
		{
			name: "max length is not exact",
			args: args{value: protoreflect.ValueOf("abcdefg"), rules: &validate.StringRules{
				MaxLen: uint64Ptr(4),
			}},
			want1: []error{ErrorMaxLen},
		},
		// Bytes
		{
			name: "bytes length is not exact",
			args: args{value: protoreflect.ValueOf("abcdefg"), rules: &validate.StringRules{
				LenBytes: uint64Ptr(4),
			}},
			want1: []error{ErrorLenBytes},
		},
		{
			name: "bytes min length is not exact",
			args: args{value: protoreflect.ValueOf("abcdefg"), rules: &validate.StringRules{
				MinBytes: uint64Ptr(40),
			}},
			want1: []error{ErrorMinBytes},
		},
		{
			name: "bytes max length is not exact",
			args: args{value: protoreflect.ValueOf("abcdefg"), rules: &validate.StringRules{
				MaxBytes: uint64Ptr(4),
			}},
			want1: []error{ErrorMaxBytes},
		},

		{
			name: "pattern not match",
			args: args{value: protoreflect.ValueOf("Hello, World!"), rules: &validate.StringRules{
				Pattern: stringPtr("Abc"),
			}},
			want1: []error{ErrorPattern},
		},
		{
			name: "prefix not match",
			args: args{value: protoreflect.ValueOf("Hello, World!"), rules: &validate.StringRules{
				Prefix: stringPtr("Abc"),
			}},
			want1: []error{ErrorPrefix},
		},
		{
			name: "suffix not match",
			args: args{value: protoreflect.ValueOf("Hello, World!"), rules: &validate.StringRules{
				Suffix: stringPtr("Abc"),
			}},
			want1: []error{ErrorSuffix},
		},
		{
			name: "string contains",
			args: args{value: protoreflect.ValueOf("Hello, World!"), rules: &validate.StringRules{
				Contains: stringPtr("Abc"),
			}},
			want1: []error{ErrorContains},
		},
		{
			name: "string not contains",
			args: args{value: protoreflect.ValueOf("Hello, World!"), rules: &validate.StringRules{
				NotContains: stringPtr("World"),
			}},
			want1: []error{ErrorNotContains},
		},
		{
			name: "invalid hostname",
			args: args{value: protoreflect.ValueOf("hello,world"), rules: &validate.StringRules{
				WellKnown: &validate.StringRules_Hostname{Hostname: true},
			}},
			want1: []error{errors.New(`hostname parts can only contain alphanumeric characters or hyphens, got ","`)},
		},
		{
			name: "hostname too long",
			args: args{value: protoreflect.ValueOf("aleph.bet.gimmel.daled.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf." +
				"aleph.bet.gimmel.daled.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf." +
				"aleph.bet.gimmel.daled.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf"),
				rules: &validate.StringRules{
					WellKnown: &validate.StringRules_Hostname{Hostname: true},
				}},
			want1: []error{ErrorHostnameTooLong},
		},
		{
			name: "hostname invalid part",
			args: args{value: protoreflect.ValueOf("aleph.bet..daled.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf"),
				rules: &validate.StringRules{
					WellKnown: &validate.StringRules_Hostname{Hostname: true},
				}},
			want1: []error{ErrorHostnameInvalidPart},
		},
		{
			name: "hostname invalid part begins with hyphens",
			args: args{value: protoreflect.ValueOf("aleph.bet.-daled.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf"),
				rules: &validate.StringRules{
					WellKnown: &validate.StringRules_Hostname{Hostname: true},
				}},
			want1: []error{ErrorHostnamePartBeginWithHyphens},
		},
		{
			name: "hostname invalid part ends with hyphens",
			args: args{value: protoreflect.ValueOf("aleph.bet.daled-.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf"),
				rules: &validate.StringRules{
					WellKnown: &validate.StringRules_Hostname{Hostname: true},
				}},
			want1: []error{ErrorHostnamePartEndsWithHyphens},
		},
		{
			name: "invalid ip",
			args: args{value: protoreflect.ValueOf("aleph.bet.daled.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf"),
				rules: &validate.StringRules{
					WellKnown: &validate.StringRules_Ip{Ip: true},
				}},
			want1: []error{ErrorInvalidIp},
		},
		{
			name: "invalid ipv4",
			args: args{value: protoreflect.ValueOf("aleph.bet.daled.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf"),
				rules: &validate.StringRules{
					WellKnown: &validate.StringRules_Ipv4{Ipv4: true},
				}},
			want1: []error{ErrorInvalidIpv4},
		},
		{
			name: "invalid ipv6",
			args: args{value: protoreflect.ValueOf("aleph.bet.daled.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf"),
				rules: &validate.StringRules{
					WellKnown: &validate.StringRules_Ipv6{Ipv6: true},
				}},
			want1: []error{ErrorInvalidIpv6},
		},
		{
			name: "invalid uri",
			args: args{value: protoreflect.ValueOf("aleph.bet.daled.hay.vav.zain.het.tet.yod.kaf.lamed.mem.nun.samekh.ain.peh.tzadi.kof.resh.shin.taf"),
				rules: &validate.StringRules{
					WellKnown: &validate.StringRules_Uri{Uri: true},
				}},
			want1: []error{ErrorInvalidUri},
		},
		{
			name: "value in",
			args: args{value: protoreflect.ValueOf("abc"),
				rules: &validate.StringRules{
					In: []string{"hello"},
				}},
			want1: []error{ErrorIn},
		},
		{
			name: "value not in",
			args: args{value: protoreflect.ValueOf("abc"),
				rules: &validate.StringRules{
					NotIn: []string{"abc"},
				}},
			want1: []error{ErrorNotIn},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ValidateString(tt.args.value, tt.args.rules)
			if got != tt.want {
				t.Errorf("ValidateString() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ValidateString() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
