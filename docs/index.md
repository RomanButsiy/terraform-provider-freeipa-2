---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "FreeIPA Provider"
subcategory: ""
description: |-
  
---

# FreeIPA Provider

## Example Usage
```hcl
terraform {
  required_providers {
    freeipa = {
      version = "[Version]"
      source  = "[Terraform registry provider path]"
    }
  }
}

provider "freeipa" {
  host = "ipa.example.test"
  username = "admin"
  password = "123456789"
  insecure = true
}
```

### Optional

- `host` (String) The FreeIPA host
- `insecure` (Boolean) Whether to verify the server's SSL certificate
- `password` (String) Password to use for connection
- `username` (String) Username to use for connection
