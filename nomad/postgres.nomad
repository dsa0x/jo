job "postgres" {
  datacenters = ["dc1"]



  group "db" {

    count = 1

    volume "pgdata" {
      type      = "host"
      read_only = false
      source    = "pgdata"
    }

    network {
      port "db" {
        static = 5431
        to = 5432
      }
    }
    
    service {
      name = "postgres-db"
      tags = ["db","postgres"]
      port = "db"

      check {
          name     = "alive"
          type     = "tcp"
          interval = "10s"
          timeout  = "2s"
        }
    }

    task "postgres" {
      driver = "docker"

      volume_mount {
        volume      = "pgdata"
        destination = "/var/lib/postgresql/data"
        read_only   = false
      }

      config {
        image = "postgres:latest"
        ports = ["db"]        
      }

      env {
        POSTGRES_PASSWORD = "somePassword"
      }
    }
  }

}