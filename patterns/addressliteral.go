package patterns

import (
	"regexp"
)

var (
	hexdigRegex = `[0-9A-Fa-f]`                         // HEXDIG               = DIGIT / "A" / "B" / "C" / "D" / "E" / "F" ; hexadecimal digit, case-insensitive
	snumRegex   = `(?:25[0-5]|2[0-4]\d|1\d{2}|\d{1,2})` // Snum                 = 1*3DIGIT ; representing a decimal integer value in the range 0 through 255

	ipv4AddressLiteralRegex = snumRegex + `(?:\.` + snumRegex + `){3}` // IPv4-address-literal = Snum 3("."  Snum)

	ipv6HexRegex  = hexdigRegex + `{1,4}`                          // IPv6-hex = 1*4HEXDIG
	ipv6FullRegex = ipv6HexRegex + `(?:\:` + ipv6HexRegex + `){7}` // IPv6-full = IPv6-hex 7(":" IPv6-hex)
	ipv6CompRegex = `(?:` +                                        // IPv6-comp
		`\:\:|` + // "::"
		`\:\:(?:` + ipv6HexRegex + `\:){0,6}` + ipv6HexRegex + `|` + // "::" [*6(IPv6-hex ":")] IPv6-hex
		`(?:` + ipv6HexRegex + `\:){1}\:(?:` + ipv6HexRegex + `\:){0,5}` + ipv6HexRegex + `|` + // 1(IPv6-hex ":") ":" [*5(IPv6-hex ":")] IPv6-hex
		`(?:` + ipv6HexRegex + `\:){2}\:(?:` + ipv6HexRegex + `\:){0,4}` + ipv6HexRegex + `|` + // 2(IPv6-hex ":") ":" [*4(IPv6-hex ":")] IPv6-hex
		`(?:` + ipv6HexRegex + `\:){3}\:(?:` + ipv6HexRegex + `\:){0,3}` + ipv6HexRegex + `|` + // 3(IPv6-hex ":") ":" [*3(IPv6-hex ":")] IPv6-hex
		`(?:` + ipv6HexRegex + `\:){4}\:(?:` + ipv6HexRegex + `\:){0,2}` + ipv6HexRegex + `|` + // 4(IPv6-hex ":") ":" [*2(IPv6-hex ":")] IPv6-hex
		`(?:` + ipv6HexRegex + `\:){5}\:(?:` + ipv6HexRegex + `\:){0,1}` + ipv6HexRegex + `|` + // 5(IPv6-hex ":") ":" [*1(IPv6-hex ":")] IPv6-hex
		`(?:` + ipv6HexRegex + `\:){0,6}` + ipv6HexRegex + `\:\:)` // [*6(IPv6-hex ":")] IPv6-hex "::"
	ipv6v4FullRegex = ipv6HexRegex + `(?:\:` + ipv6HexRegex + `){5}\:` + ipv4AddressLiteralRegex // IPv6v4-full = IPv6-hex 5(":" IPv6-hex) ":" IPv4-address-literal
	ipv6v4CompRegex = `(?:` +                                                                    // IPv6v4-comp
		`\:\:(?:` + ipv6HexRegex + `\:){0,5}` + ipv4AddressLiteralRegex + `|` + // "::" [*5(IPv6-hex ":")] IPv4-address-literal
		`(?:` + ipv6HexRegex + `\:){1}\:(?:` + ipv6HexRegex + `\:){0,4}` + ipv4AddressLiteralRegex + `|` + // 1(IPv6-hex ":") ":" [*4(IPv6-hex ":")] IPv4-address-literal
		`(?:` + ipv6HexRegex + `\:){2}\:(?:` + ipv6HexRegex + `\:){0,3}` + ipv4AddressLiteralRegex + `|` + // 2(IPv6-hex ":") ":" [*3(IPv6-hex ":")] IPv4-address-literal
		`(?:` + ipv6HexRegex + `\:){3}\:(?:` + ipv6HexRegex + `\:){0,2}` + ipv4AddressLiteralRegex + `|` + // 3(IPv6-hex ":") ":" [*2(IPv6-hex ":")] IPv4-address-literal
		`(?:` + ipv6HexRegex + `\:){4}\:(?:` + ipv6HexRegex + `\:){0,1}` + ipv4AddressLiteralRegex + `|` + // 4(IPv6-hex ":") ":" [*1(IPv6-hex ":")] IPv4-address-literal
		`(?:` + ipv6HexRegex + `\:){5}\:` + ipv4AddressLiteralRegex + `)` // 5(IPv6-hex ":") ":" IPv4-address-literal

	ipv6AddrRegex = `(` + ipv6FullRegex + `|` + ipv6CompRegex + `|` + ipv6v4FullRegex + `|` + ipv6v4CompRegex + `)` // IPv6-addr = IPv6-full / IPv6-comp / IPv6v4-full / IPv6v4-comp

	ipv4AddressLiteralPattern = regexp.MustCompile(`^\[(?:[iI][pP][vV]4:)?` + ipv4AddressLiteralRegex + `\]$`)
	ipv6AddressLiteralPattern = regexp.MustCompile(`^\[[iI][pP][vV]6:` + ipv6AddrRegex + `\]$`)
)

func ValidateAddressLiteral(addressLiteral string) bool {
	if ipv4AddressLiteralPattern.MatchString(addressLiteral) {
		return true
	}
	if ipv6AddressLiteralPattern.MatchString(addressLiteral) {
		return true
	}
	return false
}
