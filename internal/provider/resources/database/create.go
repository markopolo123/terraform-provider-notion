package database

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/theostanton/terraform-provider-notion/internal/api"
	"github.com/theostanton/terraform-provider-notion/internal/model"
)

func create(ctx context.Context, data *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*api.Client)
	var diags diag.Diagnostics

	abort := false
	database := model.Database{}

	title, ok := data.GetOk("title")
	if ok {
		database.Title = []model.RichText{
			model.NewRichText(title.(string)),
		}
	} else {
		abort = true
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No title provided",
			Detail:   "title is required for creating a database",
		})
	}

	titleColumnTitle, ok := data.GetOk("title_column_title")
	if ok {
		database.Properties = map[string]model.DatabaseProperty{
			titleColumnTitle.(string): model.NewTitleDatabaseProperty(nil),
		}
	} else {
		abort = true
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No title column title provided",
			Detail:   "title column title is required for creating a database",
		})
	}

	parent, ok := data.GetOk("parent")
	if ok {
		parent := model.NewParentFromPageId(parent.(string))
		database.Parent = &parent
	} else {
		abort = true
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "No parent provided",
			Detail:   "parent is required for creating a database",
		})
	}

	if abort {
		return diags
	}

	databaseId, err := client.CreateDatabase(ctx, database)

	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Create Database API call failed",
			Detail:   err.Error(),
		})
		return diags
	}

	data.SetId(databaseId)

	return diags
}
