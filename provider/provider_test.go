package provider

import (
	"testing"

	"github.com/hashicorp/terraform/config"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var (
	testAccProviders map[string]terraform.ResourceProvider
	testAccProvider  *schema.Provider
)

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"jenkins": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := testAccProvider.InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProviderConfigure(t *testing.T) {
	raw := map[string]interface{}{
		"address":  "#address",
		"username": "#username",
		"token":    "#token",
	}
	rawConfig, configErr := config.NewRawConfig(raw)
	if configErr != nil {
		t.Fatal(configErr)
	}
	c := terraform.NewResourceConfig(rawConfig)
	err := testAccProvider.Configure(c)
	if err != nil {
		t.Fatal(err)
	}

	config := testAccProvider.Meta().(*Services).Config
	if config.Address != raw["address"] {
		t.Fatalf("address should be %#v, not %#v", raw["address"], config.Address)
	}
	if config.Username != raw["username"] {
		t.Fatalf("username should be %#v, not %#v", raw["username"], config.Username)
	}
	if config.Token != raw["token"] {
		t.Fatalf("token should be %#v, not %#v", raw["token"], config.Token)
	}
}

func testAccPreCheck(t *testing.T) {
	c := terraform.NewResourceConfig(nil)
	err := testAccProvider.Configure(c)
	if err != nil {
		t.Fatal(err)
	}
}
