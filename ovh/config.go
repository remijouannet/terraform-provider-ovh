package ovh

import (
	"fmt"
	"log"

	"github.com/ovh/go-ovh/ovh"
)

type Config struct {
	Endpoint          string
	ApplicationKey    string
	ApplicationSecret string
	ConsumerKey       string
	OVHClient         *ovh.Client
}

func clientDefault(c *Config) (*ovh.Client, error) {
	client, err := ovh.NewClient(
		c.Endpoint,
		c.ApplicationKey,
		c.ApplicationSecret,
		c.ConsumerKey,
	)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Config) loadAndValidate() error {
	validEndpoint := false

	ovhEndpoints := [2]string{ovh.OvhEU, ovh.OvhCA}

	for _, e := range ovhEndpoints {
		if ovh.Endpoints[c.Endpoint] == e {
			validEndpoint = true
		}
	}

	if !validEndpoint {
		return fmt.Errorf("%s must be one of %#v endpoints\n", c.Endpoint, ovh.Endpoints)
	}

	targetClient, err := clientDefault(c)
	if err != nil {
		return fmt.Errorf("Error getting ovh client: %q\n", err)
	}

	log.Printf("[DEBUG] Logged in on OVH API!")
	c.OVHClient = targetClient

	return nil
}
