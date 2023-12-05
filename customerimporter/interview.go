// Package customerimporter reads from a CSV file and returns a sorted (data
// structure of your choice) of email domains along with the number of customers
// with e-mail addresses for each domain. This should be able to be ran from the
// CLI and output the sorted domains to the terminal or to a file. Any errors
// should be logged (or handled). Performance matters (this is only ~3k lines,
// but could be 1m lines or run on a small machine).

package customerimporter

import (
	"encoding/csv"
	"errors"
	"os"
	"slices"
	"strings"
	"log"
)

// DomainCount is a pair of values representing a domain and a number
// of customers with an email address in the domain.
type DomainCount struct {
	domain string
	count int
}

// InvalidAddressError is a email address domain parsing error
var ErrInvalidAddress = errors.New("Invalid email address")

// CustomerCSVToDomainCount accepts a filepath pointing to a CSV file
// and a logger to inform of non-fatal errors.
// It returns a slice of (domain, count) pairs, containing email domains
// along with the number of customers with email address in that domain.
// The slice is sorted alphabetically by domains.
func CustomerCSVToDomainCount(filepath string, logger *log.Logger) ([]DomainCount, error) {
	rawCustomers, err := loadCustomerCSV(filepath)
	if err != nil {
		return nil, err
	}

	rawDomains, err := domainsFromCustomers(rawCustomers, logger)
	if err != nil {
		return nil, err
	}

	sortedCountedDomains, err := domainCount(rawDomains)
	if err != nil {
		return nil, err
	}

	return sortedCountedDomains, nil
}

// loadCustomerCSV loads CSV file from filepath to a 2D slice of values.
func loadCustomerCSV(filepath string) ([][]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// Discard the title line
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	// Read and return the parsed CSV lines, will also happily return 0 lines.
	customers, err := reader.ReadAll()
	return customers, err
}

// domainsFromCustomers produces a slice of email domains, given a 2D slice
// of imported customer data.
// Size of the resulting slice is equal to the number of customers, which means
// that the repetitions of domains are preserved.
// Errors dealing with CSV syntax are returned from here,
// or those not considered fatal are logged via logger.
func domainsFromCustomers(rawCustomers [][]string, logger *log.Logger) ([]string, error) {
	domains := make([]string, 1, len(rawCustomers)) // Allocate all the memory at once
	for _, customer := range rawCustomers {
		if len(customer) < 3 { // Catch wrong number of columns
			return nil, errors.New("Faulty CSV syntax")
		}
		domain, err := domainFromEmail(customer[2])
		if err != nil {
			if err == ErrInvalidAddress {
				log.Print(err, domain)
			} else {
				return nil, err
			}
		}
		domains = append(domains, domain)
	}
	return domains, nil
}

// domainCount returns a slice of pairs of (domain, count), containing
// email domains along with the number of customers with email address in
// that domain. The slice is sorted alphabetically by domain.
func domainCount(rawDomains []string) ([]DomainCount, error) {
	countedDomains := make([]DomainCount, 1)
	var previousDomain string
	index := 0

	// sort domains, repetitions will be consecutive
	slices.Sort(rawDomains)

	for i, domain := range rawDomains {
		// Prepare initial data (no do-while in Go)
		if i == 0 {
			countedDomains = append(countedDomains, DomainCount{domain, 1})
			previousDomain = domain
			continue
		}
		
		if domain == previousDomain {
			countedDomains[index].count++
		} else {
			// Start counting the next domain
			countedDomains = append(countedDomains, DomainCount{domain, 1})
			// Advance pointer
			index++
			// Check against the new domain
			previousDomain = domain
		}
	}
	return countedDomains, nil
}

// domainFromEmail parses the email string and returns just a domain.
func domainFromEmail(email string) (string, error) {
	whereIsAt := strings.Index(email, "@")
	if whereIsAt == -1 {
		return "", ErrInvalidAddress
	}

	return email[whereIsAt+1:], nil
}
