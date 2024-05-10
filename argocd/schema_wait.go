package argocd

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func waitSchema() *schema.Schema {
	return &schema.Schema{
		Type:        schema.TypeSet,
		Description: "Configuration for waiting for a resource to reach a desired state based on create, update, or delete.  This takes priority over the wait attribute.",
		Optional:    true,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				"create": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"update": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
				"delete": {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  false,
				},
			},
		},
	}
}
