package oxygen

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const APIBaseURL = "https://app.oxyapi.uk"

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type OxygenClient struct {
	Httpclient HTTPClient
	Token      string
	Baseurl    string
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewClient(token string) *OxygenClient {
	return &OxygenClient{
		Httpclient: http.DefaultClient,
		Token:      token,
		Baseurl:    APIBaseURL,
	}
}

func (oc *OxygenClient) post(ctx context.Context, task string, payload any) ([]byte, error) {
	url := oc.resolveURL(ctx, task)
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	if oc.Token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", oc.Token))
	}
	res, err := oc.Httpclient.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		errResp := ErrorResponse{}
		if err := json.Unmarshal(resBody, &errResp); err != nil {
			return nil, fmt.Errorf("oxygenAI error: %s", resBody)
		}

		return nil, fmt.Errorf("oxygenAI error: %s", errResp.Error)
	}

	return resBody, nil
}

func (oc *OxygenClient) resolveURL(ctx context.Context, task string) string {

	return fmt.Sprintf("%s/%s", oc.Baseurl, task)
}
