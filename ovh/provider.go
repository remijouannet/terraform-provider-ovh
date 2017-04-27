package ovh

import (
//	"errors"
	"log"
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
				DefaultFunc: schema.EnvDefaultFunc("OVH_APPLICATION_SECRET", nil),
				Description: "Consumer key",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"ovh_domain_zone_record": resourceOVHDomainZoneRecord(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
    log.Printf("[INFO] provider init")
	config := Config{
		Endpoint:       d.Get("endpoint").(string),
		AppKey:         d.Get("application_key").(string),
		AppSecret :     d.Get("application_secret").(string),
		ConsumerKey :     d.Get("consumer_key").(string),
	}

	return config.Client()
}
