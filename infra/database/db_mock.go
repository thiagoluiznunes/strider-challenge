package database

import (
	"context"
	"fmt"
	"strider-challenge/infra/config"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MySQLContainer struct {
	Terminate func(context.Context) error
	HostIP    string
	HostPort  string
}

func SetupMySQLContainer(ctx context.Context, cfg config.Config, startContainer bool) (*MySQLContainer, error) {

	tcpPort := fmt.Sprintf("%s/tcp", cfg.DBPort)

	// Create the container
	req := testcontainers.ContainerRequest{
		Image: "mysql:latest",
		Env: map[string]string{
			"MYSQL_DATABASE":      cfg.DBName,
			"MYSQL_ROOT_USER":     cfg.DBUser,
			"MYSQL_ROOT_PASSWORD": cfg.DBPass,
		},
		ExposedPorts: []string{tcpPort},
		WaitingFor:   wait.ForListeningPort(nat.Port(tcpPort)),
	}
	container, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          startContainer,
	})
	if err != nil {
		return nil, err
	}

	// Find ports assigned to the new container
	ports, err := container.Ports(ctx)
	if err != nil {
		return nil, err
	}

	return &MySQLContainer{
		Terminate: container.Terminate,
		HostPort:  ports[nat.Port(tcpPort)][0].HostPort,
		HostIP:    ports[nat.Port(tcpPort)][0].HostIP,
	}, nil
}
