package mock

import (
	"context"
	"fmt"
	"strider-challenge/infra/config"
	"testing"

	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

type MySQLContainer struct {
	Terminate func(context.Context) error
	HostIP    string
	HostPort  string
}

func setupMySQLContainer(ctx context.Context, cfg config.Config, startContainer bool) (*MySQLContainer, error) {

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

func InitDatabaseInstance(t *testing.T, ctx context.Context, startContainer bool) (mysqlContainer *MySQLContainer, cfg config.Config, err error) {

	cfg = config.Config{
		DBName: "strider",
		DBPass: "secret",
		DBUser: "root",
	}

	if testing.Short() {
		t.Skip("skipping integration test")
	}

	if startContainer {
		t.Run("create mysql container instance", func(t *testing.T) {
			mysqlContainer, err = setupMySQLContainer(ctx, cfg, startContainer)
			if err != nil {
				t.Errorf("setupMySQLContainer() error = %v", err)
			}
		})
		cfg.DBHost = mysqlContainer.HostIP
		cfg.DBPort = mysqlContainer.HostPort
	}

	return mysqlContainer, cfg, err
}
