package argocd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func waitSchema() *schema.Schema {
	return &schema.Schema{
		Type:         schema.TypeMap,
		Description:  "Configuration for waiting for a resource to reach a desired state based on create, update, or delete.",
		Optional:     true,
		Elem:         &schema.Schema{Type: schema.TypeBool},
		ValidateFunc: validateWait,
	}
}
