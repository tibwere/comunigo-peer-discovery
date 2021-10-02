package main

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func main() {

	var ports []uint16

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	containers, err := cli.ContainerList(ctx, types.ContainerListOptions{})
	if err != nil {
		panic(err)
	}

	for _, container := range containers {
		if strings.Contains(container.Image, "peer_ms") {
			ports = append(ports, container.Ports[0].PublicPort)
		}
	}

	if len(ports) == 0 {
		fmt.Println("No peers found, bye bye :)")
	}

	sort.Slice(ports, func(i, j int) bool {
		return ports[i] < ports[j]
	})

	for i, port := range ports {
		fmt.Printf("%v-th peer is available at http://localhost:%v/\n", i+1, port)
	}
}
