package mongotest

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
	"testing"
)

const (
	image         = "mongo"
	containerPort = "27017/tcp"
	containerName = "mongo_test"
)

// RunWithMongoInDocker runs the tests with
// a mongodb instance in a docker container.
func RunWithMongoInDocker(m *testing.M, mongoURI *string) int {
	// Create a docker client.
	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	// Create a mongo container and Exposes its 27017 port and maps to the host post 27018.
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: image,
		ExposedPorts: nat.PortSet{
			containerPort: {},
		},
	}, &container.HostConfig{
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					HostIP:   "127.0.0.1",
					HostPort: "0", // 0 indicates that a free port is mapped randomly
				},
			},
		},
	}, nil, nil, containerName)
	if err != nil {
		panic(err)
	}
	containerID := resp.ID
	// Forcibly remove the mongodb container when the test is complete.
	defer func() {
		fmt.Println("mongo_test container removed...")
		err = cli.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
			Force: true,
		})
		if err != nil {
			panic(err)
		}
	}()

	// Start the mongodb container in the background。
	fmt.Println("mongo_test container start...")
	err = cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}

	// Get the hostIP and host port, then concatenate the mongodb connection URI.
	inspectRes, err := cli.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	hostIP := inspectRes.NetworkSettings.Ports[containerPort][0].HostIP
	hostPort := inspectRes.NetworkSettings.Ports[containerPort][0].HostPort
	*mongoURI = fmt.Sprintf("mongodb://%s:%s", hostIP, hostPort)

	return m.Run()
}

// NewObjID produces the mongodb ObjectID.
func NewObjID(hex string) primitive.ObjectID {
	objID, err := primitive.ObjectIDFromHex(hex)
	if err != nil {
		log.Fatalf("create ObjectID failed: %v", err)
	}
	return objID
}
