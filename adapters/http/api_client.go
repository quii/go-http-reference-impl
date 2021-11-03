package http

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/quii/go-http-reference-impl/application/recipe"
)

type APIClientLogger interface {
	Log(...interface{})
	Logf(string, ...interface{})
}

type APIClient struct {
	baseURL    string
	httpClient *http.Client
	logger     APIClientLogger
}

func NewAPIClient(baseURL string, logger APIClientLogger) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 5 * time.Second},
		logger:     logger,
	}
}

func (a *APIClient) CheckIfHealthy() error {
	url := a.baseURL + "/internal/healthcheck"
	a.logger.Log("GET", url)

	res, err := a.httpClient.Get(url)
	if err != nil {
		return fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status %d from POST %q", res.StatusCode, url)
	}

	return nil
}

func (a *APIClient) WaitForAPIToBeHealthy(retries int) error {
	var (
		err   error
		start = time.Now()
	)

	for retries > 0 {
		if err = a.CheckIfHealthy(); err != nil {
			retries--
			time.Sleep(1 * time.Second)
		} else {
			return nil
		}
	}
	return fmt.Errorf("given up checking healthcheck after %dms, %v", time.Since(start).Milliseconds(), err)
}

func (a *APIClient) Greet(name string) (string, error) {
	url := a.baseURL + "/greet/" + name
	a.logger.Log("GET", url)

	res, err := a.httpClient.Get(url)
	if err != nil {
		return "", fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return "", fmt.Errorf("unexpected status %d from GET %q", res.StatusCode, url)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func (a *APIClient) Save(r recipe.Recipe) (id string, err error) {
	url := a.baseURL + "/recipes"
	a.logger.Log("POST", url)

	recipeAsJSON, _ := json.Marshal(r)
	res, err := a.httpClient.Post(url, "application/json", bytes.NewReader(recipeAsJSON))
	if err != nil {
		return "", fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusCreated {
		return "", fmt.Errorf("unexpected status %d from %q", res.StatusCode, url)
	}

	type createdResponse struct {
		ID string `json:"id"`
	}

	var createdRes createdResponse

	if err = json.NewDecoder(res.Body).Decode(&createdRes); err != nil {
		return "", fmt.Errorf("could not parse created response from %s, %w", url, err)
	}

	return createdRes.ID, nil
}

func (a *APIClient) Get(id string) (recipe.Recipe, error) {
	url := a.baseURL + "/recipes/" + id
	res, err := a.httpClient.Get(url)
	if err != nil {
		return recipe.Recipe{}, fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return recipe.Recipe{}, fmt.Errorf("unexpected status %d from %q", res.StatusCode, url)
	}

	var r recipe.Recipe
	if err = json.NewDecoder(res.Body).Decode(&r); err != nil {
		return recipe.Recipe{}, fmt.Errorf("could not parse created response from %s, %w", url, err)
	}

	return r, nil
}