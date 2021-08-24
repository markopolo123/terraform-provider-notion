---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "notion_database_entries Data Source - terraform-provider-notion"
subcategory: ""
description: |-
  
---

# notion_database_entries (Data Source)



## Example Usage

```terraform
resource "notion_database" "example" {
  title              = "Some title"
  parent             = var.parent_page_id
  title_column_title = "Name"
}

data "notion_database_entries" "example" {
  database = notion_database.example.id
}

output "entries" {
  value = data.notion_database_entries.example.entries
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **database** (String)

### Optional

- **id** (String) The ID of this resource.

### Read-Only

- **entries** (List of Map of String)

