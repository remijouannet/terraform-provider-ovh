package ovh

import (
	"fmt"
	"github.com/ovh/go-ovh/ovh"
	"log"
	//"github.com/hashicorp/terraform/terraform"
)

type Config struct {
	Endpoint    string
	AppKey      string
	AppSecret   string
	ConsumerKey string
}

// Client represents the OVH provider client.
// This is a convenient container for the configuration and the underlying API client.
type Client struct {
	client *ovh.Client
	config *Config
}

// Client() returns a new client for accessing ovh API.
func (c *Config) Client() (*Client, error) {
	log.Printf("[INFO] client init")

	client, err := ovh.NewClient(
		c.Endpoint,
		c.AppKey,
		c.AppSecret,
		c.ConsumerKey,
	)

	if err != nil {
		fmt.Printf("Error Client : %q\n", err)
		return nil, nil
	}

	provider := &Client{
		client: client,
		config: c,
	}

	if err := provider.client.Ping(); err != nil {
		log.Printf("[INFO] failed ping API %s", err)
		return nil, nil
	}

	log.Printf("[INFO] OVH Client configured")

	return provider, nil
}
