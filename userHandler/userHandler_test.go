package userHandler_test

import (
	"testing"
)

func TestUserViewHandler(t *testing.T) {
	type want struct {
    code        int
    response    string
    contentType string
  }

	tests := []struct {
		name string
		want want
	}{
		{
			name: "OK 200",
			want: want{
				code: 200,
				response: "",
				contentType: "application/json",
			},
		},
		{
			name: "BadRequestException 400",
			want: want{
				code: 400,
				response: "",
				contentType: "application/json",
			},
		},
		{
			name: "NotFoundException 404",
			want: want{
				code: 404,
				response: "",
				contentType: "application/json",
			},
		},
		{
			name: "InternalServerException 500",
			want: want{
				code: 500,
				response: "",
				contentType: "application/json",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
