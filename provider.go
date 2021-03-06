package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		DataSourcesMap: map[string]*schema.Resource{
			"download_file": downloadFile(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"download_file": schema.DataSourceResourceShim(
				"download_file",
				downloadFile(),
			),
		},
	}
}
