package aws

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-provider-aws/internal/provider"
)

func Provider() *schema.Provider {
	return provider.Provider()
}
