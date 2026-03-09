package userHandler_test

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"userHandler"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
		code        int
		response    string
		contentType string
	}

	tests := []struct {
		name   string
		preset string
		want   want
	}{
		{
			name:   "OK 200",
			preset: "u1",
			want: want{
				code:        200,
				response:    `{"ID":"1","FirstName":"Albert","LastName":"Taygibov"}`,
				contentType: "application/json",
			},
		},
		{
			name: "BadRequestException 400",
			want: want{
				code:        400,
				response:    "user_id is empty",
				contentType: "text/plain; charset=utf-8",
			},
		},
		{
			name:   "NotFoundException 404",
			preset: "u3",
			want: want{
				code:        404,
				response:    "user not found",
				contentType: "text/plain; charset=utf-8",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users := map[string]userHandler.User{
				"u1": {ID: "1", FirstName: "Albert", LastName: "Taygibov"},
				"u2": {ID: "2", FirstName: "Ruslan", LastName: "Taygibov"},
			}

			url := "/users"
			if tt.preset != "" {
				url = "/users?user_id=" + tt.preset
			}

			request := httptest.NewRequest(http.MethodGet, url, nil)
			w := httptest.NewRecorder()
			w.Header().Set("Content-Type", "application/json")
			handler := userHandler.UserViewHandler(users)
			handler(w, request)

			result := w.Result()
			defer result.Body.Close()

			assert.Equal(t, tt.want.code, result.StatusCode)

			resBody, err := io.ReadAll(result.Body)

			require.NoError(t, err)

			assert.Equal(t, tt.want.contentType, result.Header.Get("Content-Type"))

			assert.Equal(t, strings.TrimSpace(string(resBody)), tt.want.response)
		})
	}
}
