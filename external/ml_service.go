package external

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/Capstone-Bangkit-C242-PS001/Backend-Service/config"
	dto "github.com/Capstone-Bangkit-C242-PS001/Backend-Service/dto/external"
	"io"
	"net/http"
	"time"
)

type MLService interface {
	Predict(body dto.PredictRequest) (*dto.PredictResponse, error)
}

type mlService struct {
	BaseURL    string
	HTTPClient *http.Client
	APIKey     string
}

func NewMLService() *mlService {
	baseURL := config.GetConfig().ML_SERVICE_BASE_URL

	return &mlService{
		BaseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

func (c *mlService) Predict(body dto.PredictRequest) (*dto.PredictResponse, error) {
	// Build the URL
	url := fmt.Sprintf("%s/predict", c.BaseURL)

	// Marshal the request body to JSON
	data, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf("error marshalling input: %w", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Execute the HTTP request
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	// Check for non-200 HTTP response
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("non-200 response: %d, body: %s", resp.StatusCode, string(body))
	}

	// Parse the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Unmarshal response into PredictResponse
	var predictResponse dto.PredictResponse
	if err := json.Unmarshal(responseBody, &predictResponse); err != nil {
		return nil, fmt.Errorf("error unmarshalling response: %w", err)
	}

	return &predictResponse, nil
}
