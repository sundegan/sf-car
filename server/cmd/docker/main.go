package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"time"
)

func main() {
	// Create a docker client.
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Create a mongo container and Exposes its 27017 port and maps to the host post 27018.
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "mongo",
		ExposedPorts: nat.PortSet{
			"27017:tcp": {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			"27017/tcp": []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0", // 0 indicates that a free port is mapped randomly
				},
			},
		},
	}, nil, nil, "mongo_test")
	if err != nil {
		panic(err)
	}

	// Start the mongodb container in the background。
	fmt.Println("mongo_test container start...")
	err = cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	time.Sleep(30 * time.Second)

	// Get the port that mongodb container maps to the host.
	inspectRes, err := cli.ContainerInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}
	hostPort := inspectRes.NetworkSettings.Ports["27017/tcp"][0].HostPort
	fmt.Printf("listening at host port: %+v\n", hostPort)

	// Force removal of the mongodb container.
	fmt.Println("mongo_test container removed...")
	err = cli.ContainerRemove(ctx, resp.ID, types.ContainerRemoveOptions{
		Force: true,
	})
	if err != nil {
		panic(err)
	}
}
