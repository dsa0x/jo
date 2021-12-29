server {
  enabled          = true
}

client {
  enabled       = true

  host_volume "pgdata" {
    path = "/Users/sam/Dev/docman/data"
    read_only = false
  }
}

plugin "raw_exec" {
  config {
    enabled = true
  }
}
