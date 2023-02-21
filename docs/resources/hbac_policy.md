---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "freeipa_hbac_policy Resource - freeipa"
subcategory: ""
description: |-
  
---

# freeipa_hbac_policy (Resource)





<!-- schema generated by tfplugindocs -->
## Example Usage
```hcl
resource "freeipa_hbac_policy" "hbac-0" {
  name            = "test-hbac"
  description     = "Test HBAC policy"
  enabled         = true
  hostcategory    = "all"
  servicecategory = "all"
}
```

### Required

- `name` (String) HBAC policy name

### Optional

- `description` (String) HBAC policy description
- `enabled` (Boolean) Enable this policy (Defaults to `true`)
- `usercategory` (String) User category the policy is applied to (allowed value: `all`)
- `hostcategory` (String) Host category the policy is applied to (allowed value: `all`)
- `servicecategory` (String) Service category the policy is applied to (allowed value: `all`)

### Read-Only

- `id` (String) The ID of this resource.

