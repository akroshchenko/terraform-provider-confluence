terraform {
  required_providers {
    confluence = {
      source = "registry.terraform.io/akroshchenko/confluence"
    }
  }
}

provider "confluence" {
#   site  = "selfhosted.confluence.com" # expose it via env var
#   token = "<changeme>"                # expose it via env var
}

resource confluence_content "default" {
  space  = "IC"
  title  = "akros-test-provider"
  body   = "This page was built with Terraform"
}
