package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name          string
		headers       http.Header
		want          string
		wantErr       bool
		expectedError string
	}{
		{
			name: "Success - Valid ApiKey",
			headers: http.Header{
				"Authorization": []string{"ApiKey some-api-key"},
			},
			want:    "some-api-key",
			wantErr: false,
		},
		{
			name:          "Failure - Missing Authorization Header",
			headers:       http.Header{},
			want:          "",
			wantErr:       true,
			expectedError: "no authorization header included",
		},
		{
			name: "Failure - Malformed Header (No Key)",
			headers: http.Header{
				"Authorization": []string{"ApiKey"},
			},
			want:          "",
			wantErr:       true,
			expectedError: "malformed authorization header",
		},
		{
			name: "Failure - Malformed Header (Wrong Prefix)",
			headers: http.Header{
				"Authorization": []string{"Bearer some-api-key"},
			},
			want:          "",
			wantErr:       true,
			expectedError: "malformed authorization header",
		},
		{
			name: "Failure - Malformed Header (Empty)",
			headers: http.Header{
				"Authorization": []string{""},
			},
			want:          "",
			wantErr:       true,
			expectedError: "no authorization header included",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetAPIKey(tt.headers)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAPIKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr && err != nil && tt.expectedError != "" && err.Error() != tt.expectedError {
				t.Errorf("GetAPIKey() error = %v, want %v", err, tt.expectedError)
			}
			if got != tt.want {
				t.Errorf("GetAPIKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
