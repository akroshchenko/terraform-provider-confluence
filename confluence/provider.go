package confluence

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider returns the ResourceProvider for Confluence
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"site": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confluence hostname (<name>.atlassian.net if using Cloud Confluence, otherwise hostname)",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_SITE", nil),
			},
			"site_scheme": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Optional https or http scheme to use for API calls",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_SITE_SCHEME", "https"),
			},
			"public_site": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Optional public Confluence Server hostname if different than API hostname",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_PUBLIC_SITE", ""),
			},
			"public_site_scheme": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Optional https or http scheme to use for public site URLs",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_PUBLIC_SITE_SCHEME", "https"),
			},
			"context": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confluence path context (Will default to /wiki if using an atlassian.net hostname)",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_CONTEXT", ""),
			},
			"user": {
				Type: schema.TypeString,
				// We make the 'user' field optional as there is a difference between authentication for Cloud-based and Datacenter versions of Confluence:
				// Datacenter version requires only Personal Accees Token (without username)
				// https://confluence.atlassian.com/enterprise/using-personal-access-tokens-1026032365.html
				// whereas Cloud-based uses both username and API token
				// https://developer.atlassian.com/cloud/confluence/basic-auth-for-rest-apis/
				Optional:    true,
				Description: "User's email address for Cloud Confluence or username for Confluence Server",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_USER", nil),
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Confluence API Token for Cloud Confluence or password for Confluence Server",
				DefaultFunc: schema.EnvDefaultFunc("CONFLUENCE_TOKEN", nil),
			},
		},
		// TODO: implement corresponding data sources
		DataSourcesMap: map[string]*schema.Resource{},

		ResourcesMap: map[string]*schema.Resource{
			"confluence_content":    resourceContent(),
			"confluence_attachment": resourceAttachment(),
		},
		// TODO: switch to ConfigureContextFunc
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	return NewClient(&NewClientInput{
		context:                d.Get("context").(string),
		publicSite:             d.Get("public_site").(string),
		publicSiteScheme:       d.Get("public_site_scheme").(string),
		site:                   d.Get("site").(string),
		siteScheme:             d.Get("site_scheme").(string),
		token:                  d.Get("token").(string),
		user:                   d.Get("user").(string),
	}), nil
}
