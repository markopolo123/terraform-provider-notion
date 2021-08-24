package user

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
)

var dataSchema = map[string]*schema.Schema{
	"database": {
		Type:     schema.TypeString,
		Required: true,
	},
	"title": {
		Type:     schema.TypeString,
		Computed: true,
	},
	"title_column_title": {
		Type:     schema.TypeString,
		Computed: true,
	},
}

func Data() *schema.Resource {

	return &schema.Resource{
		ReadContext: read,
		Schema:      dataSchema,
	}
}

func read(ctx context.Context, data *schema.ResourceData, m interface{}) (diags diag.Diagnostics) {
	client := m.(*api.Client)

	database, err := client.GetDatabase(ctx, data.Id())

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Failed to Get Database from API",
			Detail:   err.Error(),
		})
		return
	}

	err = data.Set("title", database.Title[0].Get())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Failed to set title from API response",
			Detail:   err.Error(),
		})
	}

	err = data.Set("parent", database.Parent.PageId)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Failed to set parent from API response",
			Detail:   err.Error(),
		})
	}

	foundTitle := false
	for _, property := range database.Properties {
		if property.Title != nil {
			foundTitle = true
			err = data.Set("title_column_title", property.Name)
			if err != nil {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Failed to set title_column_title from API response",
					Detail:   err.Error(),
				})
			}
		}
	}
	if foundTitle == false {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Warning,
			Summary:  "Failed to find title_column_title in API response",
			Detail:   err.Error(),
		})
	}

	return diags
}