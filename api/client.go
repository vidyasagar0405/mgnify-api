package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
    baseURL    = "https://www.ebi.ac.uk/metagenomics/api/v1/"
    // userAgent  = "MGnify-CLI/1.1"
)

type APIClient struct {
    baseURL    *url.URL
    httpClient *http.Client
}

func NewAPIClient() *APIClient {
    parsedURL, _ := url.Parse(baseURL)
    return &APIClient{
        baseURL: parsedURL,
        httpClient: &http.Client{
            Timeout: 30 * time.Second,
            Transport: &http.Transport{
                MaxIdleConnsPerHost: 10,
            },
        },
    }
}

func (c *APIClient) FetchPaginatedData(path string, queryParams map[string]string) ([]interface{}, error) {
    var results []interface{}

    // Build initial URL with proper query parameters
    endpointURL := c.baseURL.ResolveReference(&url.URL{Path: path})
    query := endpointURL.Query()

    // Add query parameters
    for k, v := range queryParams {
        query.Add(k, v)
    }
    endpointURL.RawQuery = query.Encode()

    // Start with the first URL
    nextURL := endpointURL
    fmt.Println(nextURL)

    for {
        req, err := http.NewRequest("GET", nextURL.String(), nil)
        if err != nil {
            return nil, fmt.Errorf("error creating request: %w", err)
        }

        req.Header.Set("Accept", "application/json")

        resp, err := c.httpClient.Do(req)
        if err != nil {
            return nil, fmt.Errorf("API request failed: %w", err)
        }
        defer resp.Body.Close()

        if resp.StatusCode != http.StatusOK {
            return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
        }

        var response struct {
            Data  []interface{} `json:"data"`
            Links struct {
                Next string `json:"next"`
            } `json:"links"`
        }

        if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
            return nil, fmt.Errorf("error decoding response: %w", err)
        }

        results = append(results, response.Data...)

        if response.Links.Next == "" {
            break
        }

        nextURL, err = url.Parse(response.Links.Next)
        if err != nil {
            return nil, fmt.Errorf("invalid next URL: %w", err)
        }
    }

    return results, nil
}
