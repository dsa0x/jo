job "web" {
  datacenters = ["dc1"]


  group "api" {

    count = 5

    network {
      port "http" {}
    }
    
    service {
      name = "api"
      tags = ["api","http","urlprefix-/"]
      port = "http"

      check {
        type = "http"
        path = "/health"
        interval = "10s"
        timeout  = "2s"
      }
    }

    task "api" {
      driver = "raw_exec"

      config {
        command = "/Users/sam/Dev/docman/docman"
        args = ["-port", "${NOMAD_PORT_http}", "-dbhost", "localhost", "-dbport", "5431", "-dbname", "postgres", "-dbuser", "postgres", "-dbpassword", "somePassword"]
      }
    }
  }

}