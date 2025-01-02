terraform {
  required_providers {
    confluence = {
      source = "test-tf-registry.com/akroshchenko/confluence"
    }
  }
}

provider "confluence" {
  site  = "test.atlassian.net"
  token = "<changeme>"
  service_deployment_model = "datacenter"
}

resource confluence_content "default" {
  space  = "IC"
  title  = "akros-test-provider"
  body   = "This page was built with Terraform"
}
