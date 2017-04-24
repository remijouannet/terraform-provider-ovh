package ovh

import (
	"log"

	"github.com/ovh/go-ovh/ovh"
	"github.com/hashicorp/terraform/terraform"
)

type Config struct {
	Endpoint string
	AppKey string
	AppSecret string
	ConsumerKey string
}

// Client represents the DNSimple provider client.
// This is a convenient container for the configuration and the underlying API client.
type Client struct {
	client *ovh.Client
	config *Config
}

// Client() returns a new client for accessing dnsimple.
func (c *Config) Client() (*Client, error) {

    client, err := ovh.NewClient(
        c.endpoint, 
        c.AppKey, 
        c.AppSecret, 
        c.consumerKey
    )
    
    if err != nil {
		fmt.Printf("Error Client : %q\n", err)
		return
	}

	provider := &Client{
		client: client,
		config: c,
	}

    err := provider.client.Ping(); err != nil  {
        log.Printf("[INFO] failed ping API")
        return
    } 

	log.Printf("[INFO] OVH Client configured for account: %s", c.AppKey)

	return provider, nil
}
