package routehomepage

import (
	"flag"
	"net/http"
	"net/http/httptest"
	"strider-challenge/mock"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var startContainer bool

func init() {
	flag.BoolVar(&startContainer, "startcontainer", true, "test container started")
}

func TestHomePageController(t *testing.T) {

	homePageService := mock.NewMockHomePageService(*mock.NewMockService())
	ctrl := &Controller{
		homeService: homePageService,
	}

	t.Run("crontoller get all posts", func(t *testing.T) {
		controllerGetAllPosts(t, ctrl)
	})
}

func controllerGetAllPosts(t *testing.T, ctrl *Controller) {

	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()

	type args struct {
		ctx        echo.Context
		path       string
		paramName  string
		paramValue string
		queryParam string
	}

	tests := []struct {
		name           string
		args           args
		wantError      bool
		wantStatusCode int64
	}{
		{
			name: "get all posts with wrong param name",
			args: args{
				ctx:        e.NewContext(req, rec),
				path:       "/homepage/:test",
				paramName:  "test",
				paramValue: "test",
			},
			wantError:      true,
			wantStatusCode: 422,
		},
		{
			name: "get all posts with wrong param value",
			args: args{
				ctx:        e.NewContext(req, rec),
				path:       "/homepage/:switch",
				paramName:  "switch",
				paramValue: "test",
			},
			wantError:      true,
			wantStatusCode: 422,
		},
		{
			name: "get all posts with empty param value",
			args: args{
				ctx:        e.NewContext(req, rec),
				path:       "/homepage/:switch",
				paramName:  "switch",
				paramValue: "",
			},
			wantError:      true,
			wantStatusCode: 422,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.args.ctx.SetPath(tt.args.path)
			tt.args.ctx.SetParamNames(tt.args.paramName)
			tt.args.ctx.SetParamValues(tt.args.paramValue)

			if assert.NoError(t, ctrl.GetAllPosts(tt.args.ctx)) {
				assert.Equal(t, tt.wantStatusCode, int64(rec.Code))
			}
		})
	}
}
