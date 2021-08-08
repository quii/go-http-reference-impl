package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/quii/go-http-reference-impl/domain"
	"io"
	"net/http"
	"time"
)

type APIClient struct {
	baseURL    string
	httpClient *http.Client
}

func NewAPIClient(baseURL string) *APIClient {
	return &APIClient{
		baseURL:    baseURL,
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

func (a *APIClient) CheckIfHealthy() error {
	url := a.baseURL + "/internal/healthcheck"
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
	var err error

	for retries > 0 {
		if err = a.CheckIfHealthy(); err != nil {
			retries -= 1
			time.Sleep(1 * time.Second)
		} else {
			return nil
		}
	}
	return fmt.Errorf("given up checking healthcheck, %v", err)
}

func (a *APIClient) Greet(name string) (string, error) {
	url := a.baseURL + "/greet/" + name

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

func (a *APIClient) Save(recipe domain.Recipe) (id string, err error) {
	url := a.baseURL + "/recipes"

	recipeAsJSON, _ := json.Marshal(recipe)
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

func (a *APIClient) Get(id string) (domain.Recipe, error) {
	url := a.baseURL + "/recipes/" + id
	res, err := a.httpClient.Get(url)
	if err != nil {
		return domain.Recipe{}, fmt.Errorf("problem reaching %s, %w", url, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return domain.Recipe{}, fmt.Errorf("unexpected status %d from %q", res.StatusCode, url)
	}

	var recipe domain.Recipe
	if err = json.NewDecoder(res.Body).Decode(&recipe); err != nil {
		return domain.Recipe{}, fmt.Errorf("could not parse created response from %s, %w", url, err)
	}

	return recipe, nil
}
