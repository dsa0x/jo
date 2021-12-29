job "fabio" {
  datacenters = ["dc1"]


  group "lb" {

    count = 1
    
    service {
      name = "fabio-lb"
      tags = ["lb","reverse-proxy"]
    }

    task "fabio" {
      driver = "raw_exec"

      config {
        command = "/Users/sam/Dev/docman/bin/fabio"
      }
    }
    // task "fabio" {
    //   driver = "raw_exec"
    //   config {
    //     command = "fabio"
    //     args = ["-proxy.strategy=rr"]
    //   }

    //   artifact {
    //     source      = "https://github.com/fabiolb/fabio/releases/download/v1.5.15/fabio-1.5.15-go1.15.5-darwin_amd64"
    //     destination = "local/fabio"
    //     mode        = "file"
    //   }
    // }
  }

}