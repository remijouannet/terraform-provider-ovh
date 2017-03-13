package ovh

import (
	"errors"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"endpoint": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVH_ENDPOINT", nil),
				Description: "the OVH endpoint, should be ovh-eu or ovh-us",
			},
			"application_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVH_APPLICATION_KEY", nil),
				Description: "Application Key",
			},
			"application_secret": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVH_APPLICATION_SECRET", nil),
				Description: "Application secret key.",
			},
			"consumer_key": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("OVH_CONSUMER_KEY", nil),
				Description: "Consumer Key.",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ovh_dns_record": resourceDNSimpleRecord(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	if email := d.Get("email").(string); email != "" {
		return nil, errors.New(
			"DNSimple API v2 requires an account identifier and the new OAuth token. " +
				"Please upgrade your configuration.")
	}

	config := Config{
		Token:   d.Get("token").(string),
		Account: d.Get("account").(string),
	}

	return config.Client()
}
