package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns a terraform.ResourceProvider.
func Provider() *schema.Provider {
	provider := &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREEIPA_HOST", ""),
				Description: descriptions["host"],
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREEIPA_USERNAME", ""),
				Description: descriptions["username"],
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				DefaultFunc: schema.EnvDefaultFunc("FREEIPA_PASSWORD", ""),
				Description: descriptions["password"],
			},
			"insecure": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: descriptions["insecure"],
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"freeipa_dns_record":                      resourceFreeIPADNSRecord(),
			"freeipa_dns_zone":                        resourceFreeIPADNSZone(),
			"freeipa_user":                            resourceFreeIPAUser(),
			"freeipa_group":                           resourceFreeIPAGroup(),
			"freeipa_user_group_membership":           resourceFreeIPAUserGroupMembership(),
			"freeipa_host":                            resourceFreeIPAHost(),
			"freeipa_hostgroup":                       resourceFreeIPAHostGroup(),
			"freeipa_host_hostgroup_membership":       resourceFreeIPAHostHostGroupMembership(),
			"freeipa_sudo_cmd":                        resourceFreeIPASudocmd(),
			"freeipa_sudo_cmdgroup":                   resourceFreeIPASudocmdgroup(),
			"freeipa_sudo_cmdgroup_membership":        resourceFreeIPASudocmdgroupMembership(),
			"freeipa_sudo_rule":                       resourceFreeIPASudoRule(),
			"freeipa_sudo_rule_option":                resourceFreeIPASudoRuleOption(),
			"freeipa_sudo_rule_allowcmd_membership":   resourceFreeIPASudoRuleAllowCommandMembership(),
			"freeipa_sudo_rule_denycmd_membership":    resourceFreeIPASudoRuleDenyCommandMembership(),
			"freeipa_sudo_rule_host_membership":       resourceFreeIPASudoRuleHostMembership(),
			"freeipa_sudo_rule_user_membership":       resourceFreeIPASudoRuleUserMembership(),
			"freeipa_sudo_rule_runasuser_membership":  resourceFreeIPASudoRuleRunAsUserMembership(),
			"freeipa_sudo_rule_runasgroup_membership": resourceFreeIPASudoRuleRunAsGroupMembership(),
			"freeipa_hbac_policy":                     resourceFreeIPAHBACPolicy(),
			"freeipa_hbac_policy_host_membership":     resourceFreeIPAHBACPolicyHostMembership(),
			"freeipa_hbac_policy_user_membership":     resourceFreeIPAHBACPolicyUserMembership(),
			"freeipa_hbac_policy_service_membership":  resourceFreeIPAHBACPolicyServiceMembership(),
		},

		DataSourcesMap: map[string]*schema.Resource{},

		ConfigureFunc: providerConfigure,
	}
	return provider
}

var descriptions map[string]string

func init() {
	descriptions = map[string]string{
		"host": "The FreeIPA host",

		"username": "Username to use for connection",

		"password": "Password to use for connection",

		"insecure": "Whether to verify the server's SSL certificate",
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return &Config{
		Host:               d.Get("host").(string),
		Username:           d.Get("username").(string),
		Password:           d.Get("password").(string),
		InsecureSkipVerify: d.Get("insecure").(bool),
	}, nil
}
