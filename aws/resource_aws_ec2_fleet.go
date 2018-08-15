package aws

import (
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceAwsEC2Fleet() *schema.Resource {
	return &schema.Resource{
		Create: nil,
		Read:   nil,
		Delete: nil,
		Update: nil,

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"default_target_capacity_type": {
				Type: schema.TypeString,
			},
			"excess_capacity_termination_policy": {
				Type: schema.TypeString,
			},

			"on_demand_target_capacity": {
				Type: schema.TypeInt,
			},
			"replace_unhealthy_instances": {
				Type: schema.TypeBool,
			},
			"request_type": {
				Type: schema.TypeString,
			},
			"spot_target_capacity": {
				Type: schema.TypeInt,
			},
			"terminate-instances-with-expiration": {
				Type: schema.TypeBool,
			},
			"total_target_capacity": {
				Type: schema.TypeInt,
			},
			"valid_from": {
				Type: schema.TypeString,
			},
			"valid_until": {
				Type: schema.TypeString,
			},
			"tags": tagsSchema(),
			"launch_template": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:         schema.TypeString,
							Required:     true,
							ValidateFunc: validateLaunchTemplateId,
						},
						"version": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"launch_template_overrides": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"instance_type": {
							Type: schema.TypeString,
						},
						"max_price": {
							Type: schema.TypeString,
						},
						"subnet_id": {
							Type: schema.TypeString,
						},
						"availability_zone": {
							Type: schema.TypeString,
						},
						"weighted_capacity": {
							Type: schema.TypeFloat,
						},
						"priority": {
							Type: schema.TypeFloat,
						},
					},
				},
			},
			"on_demand_options": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allocation_strategy": {
							Type: schema.TypeString,
						},
					},
				},
			},
			"spot_options": {
				Type: schema.TypeSet,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"allocation_strategy": {
							Type: schema.TypeString,
						},
						"instance_interruption_behaviour": {
							Type: schema.TypeString,
						},
						"instance_pools_to_use_count": {
							Type: schema.TypeInt,
						},
					},
				},
			},
		},
	}
}
