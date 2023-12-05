package customerimporter

import (
	"testing"
	"reflect"
)

type domainFromEmailTest struct {
	email, expected string
}

var domainFromEmailTests = []domainFromEmailTest{
	domainFromEmailTest{"name@domain.co", "domain.co"},
	domainFromEmailTest{"a.b.b.b@toplevel", "toplevel"},
	domainFromEmailTest{"emptydomain@", ""},
	domainFromEmailTest{"@no.user.name", "no.user.name"},
	domainFromEmailTest{"normal.user@email.domain", "email.domain"},
	domainFromEmailTest{"@", ""},
	domainFromEmailTest{"trailing.punctuation@dom.ain.", "dom.ain."},
}

func TestDomainFromEmail(t *testing.T) {
	for _, test := range domainFromEmailTests {
		out, _ := domainFromEmail(test.email)
		if out != test.expected {
			t.Errorf("domainFromEmail: got %s expected %s", out, test.expected)
		}
	}
}

var domainFromEmailErrorTests = []domainFromEmailTest{
	domainFromEmailTest{"justthe.domain", ""},
	domainFromEmailTest{"justtoplevel", ""},
	domainFromEmailTest{"2.very.deep.domain.name", ""},
}

func TestDomainFromEmailErrors(t *testing.T) {
	for _, test := range domainFromEmailErrorTests {
		out, err := domainFromEmail(test.email)
		if err == nil {
			t.Errorf("domainFromEmail: got %s expected an error", out)
		}
	}
}

type domainCountTest struct {
	domains []string
	result []DomainCount
}

var domainCountTests = []domainCountTest{
	domainCountTest{
		[]string{"a", "a", "aa", "aa", "aa", "ab"},
		[]DomainCount{
			DomainCount{"a", 2},
			DomainCount{"aa", 3},
			DomainCount{"ab", 1},
		},
	},
	domainCountTest{
		[]string{"", "", "", "", "empty"},
		[]DomainCount{
			DomainCount{"", 4},
			DomainCount{"empty", 1},
		},
	},
}

func TestDomainCount(t *testing.T) {
	for _, test := range domainCountTests {
		out, _ := domainCount(test.domains)
		if !reflect.DeepEqual(out, test.result) {
			t.Error("domainCount: got ", out, " expected ", test.result)
		}
	}
}


