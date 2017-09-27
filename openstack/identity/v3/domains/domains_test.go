package domains

import (
	"fmt"
	"testing"
	"encoding/json"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/pagination"
	"github.com/rackspace/gophercloud/openstack"
)

var OpenstackUrl = ""
var Username     = ""
var Password     = ""
var Project      = ""
var DomainName   = ""

func TestList(*testing.T) {

	auth := gophercloud.AuthOptions{}
	auth.IdentityEndpoint = OpenstackUrl
	auth.Username = Username
	auth.Password = Password
	auth.TenantName = Project
	auth.AllowReauth = true
	auth.DomainName = DomainName
	client, err := openstack.AuthenticatedClient(auth)
	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}
	identity := openstack.NewIdentityV3(client)

	opts := ListOpts{}
	pager := List(identity, &opts)

	err = pager.EachPage(func(page pagination.Page) (bool, error) {
		domainList, err := ExtractDomains(page)
		for _, dom := range domainList {
			// "dom" will be a Domain
			//fmt.Printf("%+v\n\n", dom)
			pprint(dom)
		}
		return true, err
	})

	if err != nil {
		fmt.Printf("ERROR: %s\n", err.Error())
		return
	}
}

func pprint(data interface{}) {
	str, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("pprint() error: %s\n", err.Error())
		return
	}
	fmt.Printf("%s\n",str)
}
