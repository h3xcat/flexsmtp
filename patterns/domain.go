package patterns

import "regexp"

var domainPattern = regexp.MustCompile(`^[A-Za-z\d](?:[A-Za-z\d\-]*[A-Za-z\d])?(?:\.[A-Za-z\d](?:[A-Za-z\d\-]*[A-Za-z\d])?)*$`)

func ValidateDomain(domain string) bool {
	return domainPattern.MatchString(domain)
}
