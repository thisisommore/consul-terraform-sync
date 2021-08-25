schema = "1"

project "consul-terraform-sync" {
  team = "consul team" #what should go here?
  slack {
    notification_channel = "#consul channel" #what should go here?
  }
  github {
    organization = "hashicorp"
    repository = "consul-terraform-sync"
    release_branches = ["main"]
  }
}

event "merge" {}

event "build" {
  depends = ["merge"]
  action "build" {
    organization = "hashicorp"
    repository = "consul-terraform-sync"
    workflow = "build"
  }
}

event "upload-dev" {
  depends = ["build"]
  action "upload-dev" {
    organization = "hashicorp"
    repository = "consul-terraform-sync"
    workflow = "upload-dev"
    depends = ["build"]
  }

  notification {
    on = "fail"
    message_template = "{{stage_name}} failed with {{stage_output}}"
  }
}

event "notarize-darwin-amd64" {
  action "notarize-darwin-amd64" {
    organization = "hashicorp"
    repository = "consul-terraform-sync"
    workflow = "notarize-darwin-amd64"

  }

  notification {
    on = "fail"
    message_template = "{{stage_name}} {{version}} failed with {{stage_output}}"
  }
}

event "notarize-windows-386" {
  depends = ["notarize-darwin-amd64"]
  action "notarize-windows-386" {
    organization = "hashicorp"
    repository = "consul-terraform-sync"
    workflow = "notarize-windows-386"

  }

  notification {
    on = "fail"
    message_template = "{{stage_name}} {{version}} failed with {{stage_output}}"
  }
}

event "notarize-windows-amd64" {
  depends = ["notarize-windows-386"]
  action "notarize-windows-amd64" {
    organization = "hashicorp"
    repository = "consul-terraform-sync"
    workflow = "notarize-windows-amd64"
  }

  notification {
    on = "fail"
    message_template = "{{stage_name}} {{version}} failed with {{stage_output}}"
  }
}

event "sign" {
  depends = ["notarize-windows-amd64"]
  action "sign" {
    organization = "hashicorp"
    repository = "crt-workflows-common"
    workflow = "sign"

  }

  notification {
    on = "fail"
    message_template = "{{stage_name}} {{version}} failed with {{stage_output}}"
  }
}