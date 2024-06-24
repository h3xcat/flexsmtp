package patterns

import (
	"testing"
)

func TestValidateAddressLiteral(t *testing.T) {
	testCases := []struct {
		input    string
		expected bool
	}{
		// Valid test cases
		{"[192.168.1.1]", true},
		{"[255.255.255.255]", true},
		{"[0.0.0.0]", true},
		{"[127.0.0.1]", true},
		{"[IPv6:2001:0db8:85a3:0000:0000:8a2e:0370:7334]", true},
		{"[IPv6:2001:db8:85a3:0:0:8a2e:370:7334]", true},
		{"[IPv6:2001:db8:85a3::8a2e:370:7334]", true},
		{"[IPv6:::1]", true},
		{"[IPv6:2001:db8::1]", true},
		{"[IPv6:2001:0db8:0000:0042:0000:8a2e:0370:7334]", true},
		{"[IPv6:2001:0db8:85a3:0000:0000:8a2e:192.168.1.1]", true},
		{"[IPv6:2001:db8:85a3:0:0:8a2e:192.168.1.1]", true},
		{"[IPv6:::192.168.1.1]", true},
		{"[IPv6:2001:db8::192.168.1.1]", true},
		{"[IPv6:2001:db8::85a3:192.168.1.1]", true},

		// Invalid test cases
		{"[256.256.256.256]", false},
		{"[192.168.1.999]", false},
		{"[192.168.1.]", false},
		{"[192.168.1]", false},
		{"[192.168.1.1.1]", false},
		{"[IPv6:2001:db8::85a3::8a2e:370:7334]", false},
		{"[IPv6:2001:db8:85a3:0:0:8a2e:370:7334:1234]", false},
		{"[IPv6:2001:db8:85a3:0:0:8a2e:370g:7334]", false},
		{"[IPv6:2001:db8::g]", false},
		{"[IPv6:2001:db8::85a3:192.168.1]", false},
		{"[IPv6:2001:db8::85a3:256.256.256.256]", false},
		{"[IPv6::192.168.1.1:1234]", false},
		{"[foo:]", false},
		{"[example]", false},
		{"[foo bar:content]", false},
		{"[123:content]", false},
		{"192.168.1.1", false},
		{"[192.168.1.1", false},
		{"192.168.1.1]", false},
		{"[IPv6:::1", false},
		{"IPv6:::1]", false},
		{"[example:SomeContent]", false},
		{"[foo:bar]", false},
		{"[tag:12345]", false},
	}

	for _, tc := range testCases {
		t.Run(tc.input, func(t *testing.T) {
			result := ValidateAddressLiteral(tc.input)
			if result != tc.expected {
				t.Errorf("ValidateAddressLiteral(%s) = %v; expected %v", tc.input, result, tc.expected)
			}
		})
	}
}
