package main

import (
	"github.com/remijouannet/terraform-provider-ovh/ovh"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ovh.Provider,
	})
}
