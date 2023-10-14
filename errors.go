package main

import "errors"

var (
	ErrorConst                        = errors.New("value does not match the const rule")
	ErrorLt                           = errors.New("value is not less than the specified value")
	ErrorLte                          = errors.New("value is not less than or equal to the specified value")
	ErrorGt                           = errors.New("value is not greater than the specified value")
	ErrorGte                          = errors.New("value is not greater than or equal to the specified value")
	ErrorLen                          = errors.New("value does not match the specified length")
	ErrorMinLen                       = errors.New("value is shorter than the minimum length")
	ErrorMaxLen                       = errors.New("value is longer than the maximum length")
	ErrorLenBytes                     = errors.New("byte length does not match the specified length")
	ErrorMinBytes                     = errors.New("byte length is less than the minimum specified")
	ErrorMaxBytes                     = errors.New("byte length exceeds the maximum specified")
	ErrorPattern                      = errors.New("value does not match the specified pattern")
	ErrorPrefix                       = errors.New("value does not have the specified prefix")
	ErrorSuffix                       = errors.New("value does not have the specified suffix")
	ErrorContains                     = errors.New("value does not contain the specified substring")
	ErrorNotContains                  = errors.New("value contains a prohibited substring")
	ErrorIn                           = errors.New("value is not in the specified set")
	ErrorNotIn                        = errors.New("value is in the not_in set")
	ErrorInvalid                      = errors.New("value is not valid")
	ErrorInvalidEmail                 = errors.New("value is not a valid email address")
	ErrorHostnameTooLong              = errors.New("hostname cannot exceed 253 characters")
	ErrorHostnameInvalidPart          = errors.New("hostname part must be non-empty and cannot exceed 63 characters")
	ErrorHostnamePartBeginWithHyphens = errors.New("hostname parts cannot begin with hyphens")
	ErrorHostnamePartEndsWithHyphens  = errors.New("hostname parts cannot end with hyphens")
	ErrorInvalidIp                    = errors.New("value is not a valid ip address")
	ErrorInvalidIpv4                  = errors.New("value is not a valid ipv4 address")
	ErrorInvalidIpv6                  = errors.New("value is not a valid ipv6 address")
	ErrorInvalidUri                   = errors.New("value is not a valid uri")
)