package main

import (
	"context"
	"fmt"
	"net/http"
	"testing"
)

// HTTPClient là interface để mock HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// MyService là struct chứa logic gọi API
type MyService struct {
	client HTTPClient
}

// CallAPI là phương thức gọi API sử dụng HTTPClient
func (s *MyService) CallAPI(ctx context.Context) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://example.com", nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	resp, err := s.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to perform request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return nil
}

// MockHTTPClient là struct để mock HTTPClient
type MockHTTPClient struct {
	DoFunc func(*http.Request) (*http.Response, error)
}

// Do triển khai phương thức Do của HTTPClient interface
func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return m.DoFunc(req)
}

// Unit test cho MyService
func TestMyService_CallAPI(t *testing.T) {
	tests := []struct {
		name          string
		mockResponse  *http.Response
		mockError     error
		expectedError bool
	}{
		{
			name:          "Success",
			mockResponse:  &http.Response{StatusCode: http.StatusOK},
			mockError:     nil,
			expectedError: false,
		},
		{
			name:          "Failed with bad status",
			mockResponse:  &http.Response{StatusCode: http.StatusBadRequest},
			mockError:     nil,
			expectedError: true,
		},
		{
			name:          "Failed with error",
			mockResponse:  nil,
			mockError:     fmt.Errorf("network error"),
			expectedError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Tạo mock client
			mockClient := &MockHTTPClient{
				DoFunc: func(req *http.Request) (*http.Response, error) {
					return tt.mockResponse, tt.mockError
				},
			}

			// Khởi tạo service với mock client
			service := &MyService{client: mockClient}

			// Gọi CallAPI với context
			ctx := context.Background()
			err := service.CallAPI(ctx)

			// Kiểm tra kết quả
			if (err != nil) != tt.expectedError {
				t.Errorf("CallAPI() error = %v, expectedError %v", err, tt.expectedError)
			}
		})
	}
}
