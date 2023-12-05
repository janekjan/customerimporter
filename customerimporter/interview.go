// Package customerimporter reads from a CSV file and returns a sorted (data
// structure of your choice) of email domains along with the number of customers
// with e-mail addresses for each domain. This should be able to be ran from the
// CLI and output the sorted domains to the terminal or to a file. Any errors
// should be logged (or handled). Performance matters (this is only ~3k lines,
// but could be 1m lines or run on a small machine).

package customerimporter

// DomainCount is a pair of values representing a domain and a number
// of customers with an email address in the domain.
type DomainCount struct {
	domain string
	count int
}

// CustomerCSVToDomainCount accepts a filepath pointing to a CSV file.
// It returns a slice of (domain, count) pairs, containing email domains
// along with the number of customers with email address in that domain.
// The slice is sorted alphabetically by domains.
func CustomerCSVToDomainCount(filepath string) ([]DomainCount, error) {

}

// loadCustomerCSV loads CSV file from filepath to a 2D slice of values.
func loadCustomerCSV(filepath string) ([][]string, error) {

}

// domainsFromCustomers produces a slice of email domains, given a 2D slice
// of imported customer data.
// Size of the resulting slice is equal to the number of customers, which means
// that the repetitions of domains are preserved.
// Errors dealing with CSV syntax are returned from here.
func domainsFromCustomers(rawCustomers [][]string) ([]string, error) {

}

// domainCount returns a slice of pairs of (domain, count), containing
// email domains along with the number of customers with email address in
// that domain. The slice is sorted alphabetically by domain.
func domainCount(rawDomains []string) ([]DomainCount, error) {

}

