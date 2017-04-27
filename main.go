package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/remijouannet/terraform-provider-ovh/ovh"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ovh.Provider,
	})
}
