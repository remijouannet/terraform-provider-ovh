package ovh

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"ovh": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("OVH_APPLICATION_KEY"); v != "" {
		t.Fatal("OVH_APPLICATION_KEY must be set for acceptance tests")
	}

	if v := os.Getenv("OVH_APPLICATION_SECRET"); v == "" {
		t.Fatal("OVH_APPLICATION_SECRET must be set for acceptance tests")
	}

	if v := os.Getenv("OVH_APPLICATION_SECRET"); v == "" {
		t.Fatal("OVH_APPLICATION_SECRET must be set for acceptance tests")
	}

	if v := os.Getenv("OVH_ZONE"); v == "" {
		t.Fatal("OVH_ZONE must be set for acceptance tests. The domain is used to create and destroy record against.")
	}
}
