package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

var testFmtPtr = flag.Bool("t", false, "Comma separated available ports")

func main() {
	var ports []uint16

	flag.Parse()

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
		if strings.Contains(container.Image, "comunigo/peer") {
			ports = append(ports, container.Ports[0].PublicPort)
		}
	}

	if len(ports) == 0 {
		if *testFmtPtr {
			os.Exit(1)
		} else {
			fmt.Println("No peers found, bye bye :)")
			os.Exit(1)
		}
	}

	sort.Slice(ports, func(i, j int) bool {
		return ports[i] < ports[j]
	})

	if *testFmtPtr {
		output := ""
		for i, port := range ports {
			if i == len(ports)-1 {
				output += fmt.Sprintf("%v", port)
			} else {
				output += fmt.Sprintf("%v,", port)
			}
		}
		fmt.Println(output)
	} else {
		for i, port := range ports {
			fmt.Printf("%v-th peer is available at http://localhost:%v/\n", i+1, port)
		}
	}

	os.Exit(0)
}
