# This file is generated by Consul Terraform Sync.
#
# The HCL blocks, arguments, variables, and values are derived from the
# operator configuration for Sync. Any manual changes to this file
# may not be preserved and could be overwritten by a subsequent update.
#
# Task: test
# Description: user description for task named 'test'

catalog_services = {
  "api" = ["tag"]
  "consul" = []
  "web" = ["tag_a", "tag_b"]
}

services = {
  "api.worker-01.dc1" = {
    id              = "api"
    name            = "api"
    kind            = ""
    address         = "1.2.3.4"
    port            = 8080
    meta            = {}
    tags            = ["tag"]
    namespace       = ""
    status          = "passing"
    node            = "worker-01"
    node_id         = "39e5a7f5-2834-e16d-6925-78167c9f50d8"
    node_address    = "127.0.0.1"
    node_datacenter = "dc1"
    node_tagged_addresses = {
      lan      = "127.0.0.1"
      lan_ipv4 = "127.0.0.1"
      wan      = "127.0.0.1"
      wan_ipv4 = "127.0.0.1"
    }
    node_meta = {
      consul-network-segment = ""
    }
    cts_user_defined_meta = {}
  },
  "api-2.worker-01.dc1" = {
    id              = "api-2"
    name            = "api"
    kind            = ""
    address         = "5.6.7.8"
    port            = 8080
    meta            = {}
    tags            = ["tag"]
    namespace       = ""
    status          = "passing"
    node            = "worker-01"
    node_id         = "39e5a7f5-2834-e16d-6925-78167c9f50d8"
    node_address    = "127.0.0.1"
    node_datacenter = "dc1"
    node_tagged_addresses = {
      lan      = "127.0.0.1"
      lan_ipv4 = "127.0.0.1"
      wan      = "127.0.0.1"
      wan_ipv4 = "127.0.0.1"
    }
    node_meta = {
      consul-network-segment = ""
    }
    cts_user_defined_meta = {}
  },
  "web.worker-01.dc1" = {
    id              = "web"
    name            = "web"
    kind            = ""
    address         = "1.1.1.1"
    port            = 8000
    meta            = {}
    tags            = ["tag_a", "tag_b"]
    namespace       = ""
    status          = "passing"
    node            = "worker-01"
    node_id         = "39e5a7f5-2834-e16d-6925-78167c9f50d8"
    node_address    = "127.0.0.1"
    node_datacenter = "dc1"
    node_tagged_addresses = {
      lan      = "127.0.0.1"
      lan_ipv4 = "127.0.0.1"
      wan      = "127.0.0.1"
      wan_ipv4 = "127.0.0.1"
    }
    node_meta = {
      consul-network-segment = ""
    }
    cts_user_defined_meta = {}
  },
}
