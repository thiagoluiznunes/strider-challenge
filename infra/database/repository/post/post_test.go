package post

import (
	"context"
	"database/sql"
	"flag"
	"strider-challenge/domain/contract"
	"strider-challenge/domain/entity"
	"strider-challenge/infra/config"
	"strider-challenge/infra/database"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var startContainer bool
var conn *sql.DB

func init() {
	flag.BoolVar(&startContainer, "startcontainer", true, "test container started")
}

func TestPostRepo(t *testing.T) {

	var mysqlContainer *database.MySQLContainer
	var err error

	ctx := context.Background()
	cfg := config.Config{
		DBName: "strider",
		DBPass: "secret",
		DBHost: "mysql",
		DBPort: "3306",
		DBUser: "root",
	}

	if testing.Short() {
		t.Skip("skipping integration test")
	}

	if startContainer {
		t.Run("create mysql container instance", func(t *testing.T) {
			mysqlContainer, err = database.SetupMySQLContainer(ctx, cfg, startContainer)
			if err != nil {
				t.Errorf("database.() error = %v", err)
			}
		})
		cfg.DBHost = mysqlContainer.HostIP
		cfg.DBPort = mysqlContainer.HostPort

		defer mysqlContainer.Terminate(ctx)
	}

	conn, err = database.ConnectDataBase(&cfg)
	if err != nil {
		t.Errorf("ConnectDataBase() error = %v", err)
		return
	}

	repo := NewPostRepository(conn)
	assert.NotEmpty(t, repo)

	t.Run("create strider database", func(t *testing.T) {
		CreateDatabase(t, ctx, conn)
	})

	t.Run("create post table", func(t *testing.T) {
		CreatePostTable(t, ctx, conn)
	})

	t.Run("add post", func(t *testing.T) {
		AddPost(t, ctx, repo)
	})
}

func CreateDatabase(t *testing.T, ctx context.Context, conn *sql.DB) {
	result, err := conn.Exec(CreateDatabaseQuery)
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func CreatePostTable(t *testing.T, ctx context.Context, conn *sql.DB) {
	result, err := conn.Exec(CreatePostTableQuery)
	assert.Nil(t, err)
	assert.NotEmpty(t, result)
}

func AddPost(t *testing.T, ctx context.Context, repo contract.PostRepo) {

	var postsIDs []int64

	type args struct {
		post entity.Post
	}

	tests := []struct {
		name      string
		args      args
		wantError bool
	}{
		{
			name: "insert post without uuid",
			args: args{
				entity.Post{
					Type:      "original",
					Text:      NewString("My original twitter test without uuid"),
					UserID:    1,
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: true,
		},
		{
			name: "insert post with empty type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "",
					Text:      NewString("My twitter test with empty type"),
					UserID:    1,
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: true,
		},
		{
			name: "insert post without type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Text:      NewString("My twitter test without type"),
					UserID:    1,
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: true,
		},
		{
			name: "insert post with wrongly type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "wrong",
					Text:      NewString("My twitter test with wrongly type"),
					UserID:    1,
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: true,
		},
		{
			name: "insert post with empty text",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "original",
					UserID:    1,
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: true,
		},
		{
			name: "insert post/userid=1 with original type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "original",
					Text:      NewString("My twitter test with original type"),
					UserID:    1,
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: false,
		},
		{
			name: "insert post/userid=2 with original type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "original",
					Text:      NewString("My twitter test with original type"),
					UserID:    2,
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: false,
		},
		{
			name: "insert post/userid=3 with original type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "original",
					Text:      NewString("My twitter test with original type"),
					UserID:    3,
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			postID, err := repo.Add(ctx, tt.args.post)
			if (err != nil) && !tt.wantError {
				t.Errorf("Add() error = %v, wantError %v", err, tt.wantError)
			}
			if postID != 0 {
				postsIDs = append(postsIDs, postID)
			}
		})
	}

	tests = []struct {
		name      string
		args      args
		wantError bool
	}{
		{
			name: "insert post with repost type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "repost",
					Text:      NewString("My twitter test with repost type"),
					UserID:    1,
					PostID:    &postsIDs[0],
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: false,
		},
		{
			name: "insert post with repost type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "quote",
					Text:      NewString("My twitter test with quote type"),
					UserID:    1,
					PostID:    &postsIDs[1],
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: false,
		},
		{
			name: "insert post with repost type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "quote",
					Text:      NewString("My twitter test with quote type"),
					UserID:    2,
					PostID:    &postsIDs[2],
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: false,
		},
		{
			name: "insert post with repost type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "quote",
					Text:      NewString("My twitter test with quote type"),
					UserID:    3,
					PostID:    &postsIDs[0],
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: false,
		},
		{
			name: "insert post with repost type",
			args: args{
				entity.Post{
					UUID:      NewUUID(),
					Type:      "quote",
					Text:      NewString("My twitter test with quote type"),
					UserID:    1,
					PostID:    &postsIDs[1],
					UpdatedAt: time.Now(),
					CreatedAt: time.Now(),
				},
			},
			wantError: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.Add(ctx, tt.args.post)
			if (err != nil) && !tt.wantError {
				t.Errorf("Add() error = %v, wantError %v", err, tt.wantError)
			}
		})
	}
}
