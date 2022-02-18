package database

import (
	"context"
	"flag"
	"strider-challenge/infra/config"
	"strider-challenge/mock"
	"testing"

	"github.com/stretchr/testify/assert"
)

var startContainer bool

func init() {
	flag.BoolVar(&startContainer, "startcontainer", true, "test container started")
}

func TestConnection(t *testing.T) {

	ctx := context.Background()
	mysqlContainer, cfg, err := mock.InitDatabaseInstance(t, ctx, startContainer)
	assert.NoError(t, err)

	defer mysqlContainer.Terminate(ctx)

	t.Run("test mysql connection instance", func(t *testing.T) {
		type args struct {
			ctx    context.Context
			config *config.Config
		}
		tests := []struct {
			name      string
			args      *args
			wantError bool
		}{
			{
				name: "init mysql connection with background context and container",
				args: &args{
					ctx:    nil,
					config: &cfg,
				},
				wantError: false,
			},
			{
				name: "init mysql connection without config",
				args: &args{
					ctx: nil,
					config: &config.Config{
						DBHost: "",
						DBPort: "",
						DBName: "",
						DBUser: "",
						DBPass: "",
					},
				},
				wantError: false,
			},
		}
		for _, tt := range tests {
			_, err := ConnectDataBase(tt.args.config)
			if (err != nil) != tt.wantError {
				t.Errorf("ConnectDataBase() error = %v, wantErr %v", err, tt.wantError)
				return
			}
		}
	})
}
