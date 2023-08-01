package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)


func Test_GetDrones(t *testing.T) {
	tests := []struct{
		name  string
		wantStatus int
	} {
		{
			name: "Test get end point status code",
			wantStatus: http.StatusOK,
		},
    }

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &droneAPI{}
			request, _ := http.NewRequest(http.MethodGet, "/api/dron/", nil)
			response := httptest.NewRecorder()
			api.Get(response, request)
			if status := response.Code; status != tt.wantStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusCreated)
			}
		})
	}
}