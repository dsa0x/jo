### Project using hashistack

Creating an app that uses the [Hashistack](https://hashicorp.com) Tools.

- Nomad for container ochestration
- Consul for service discovery
- Fabio\* for load balancing

- Go application is running on 5 different instances, each with a different port.
- Nomad is used to schedule the instances.
- Fabio balances the load between the instances.
- Consul is used to register the instances with the load balancer.

#### Run nomad jobs

`cd nomad`
`nomad job init`
`nomad job run consul.nomad`
`nomad job run postgres.nomad`
`nomad job run web.nomad`
`nomad job run fabio.nomad`

If you are using a non-M1 machine, replace `fabio.nomad` with the right artifact.
