package subdomain

import (
	"testing"
)

func TestSubdomainVisits(t *testing.T) {
	var te = []string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}
	// fmt.Println(strings.SplitN("discuss.leetcode.com", ".", 3))
	SubdomainVisits(te)
}
