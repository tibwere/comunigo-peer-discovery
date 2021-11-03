# Discovery service for comuniGO peers
This simple go program uses the Docker SDK to identify the ports exposed by each peer of the [comuniGO](https://gitlab.com/tibwere/comunigo) application.

You can use these applications in the two modes shown below:
- To find out which port the peers' frontends are listening on use `go run comunigo-peer-discovery.go`
- To get the same result but as comma separated values insteady a pretty format use `-t` flag