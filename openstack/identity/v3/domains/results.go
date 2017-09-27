package domains

import (
	"github.com/mitchellh/mapstructure"
	"github.com/rackspace/gophercloud/pagination"
)

// Domain
type Domain struct {
	ID           string                  `mapstructure:"id"          json:"id"`
	Name         string                  `mapstructure:"name"        json:"name"`
	Description  string                  `mapstructure:"description" json:"description"`
	Enabled      bool                    `mapstructure:"enabled"     json:"enabled"`
	Links        map[string]interface{}  `mapstructure:"links"       json:"links"`
}

// DomainPage is a single page of Domain results.
type DomainPage struct {
	pagination.LinkedPageBase
}

// IsEmpty returns true if no Domain were returned.
func (p DomainPage) IsEmpty() (bool, error) {
	es, err := ExtractDomains(p)
	if err != nil {
		return true, err
	}
	return len(es) == 0, nil
}

// ExtractDomains extracts a list of Domains from a Page.
func ExtractDomains(page pagination.Page) ([]Domain, error) {

	var response struct {
		Domains []Domain `mapstructure:"domains"`
	}

	err := mapstructure.Decode(page.(DomainPage).Body, &response)

	return response.Domains, err
}
