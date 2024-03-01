package oxygen
import (
	"context"
	"errors"
	"encoding/json"
)

type ImageRequest struct {

	//Required
	Prompt string `json:"prompt"`

	// ID of the model to use. See the model endpoint compatibility (https://app.oxyapi.uk/v1/models) table for details on which models work with the Chat API.
	Model string `json:"model"`

	//If set, partial images with progression deltas will be sent.
	Stream *string `json:"stream,omitempty"`
}
type ImageResponse map[string]interface{}
func (oc *OxygenClient) ImageGeneration(ctx context.Context, req *ImageRequest) (*ImageResponse, error) {
	if len(req.Prompt) == 0 {
		return nil, errors.New("input text is required")
	}
	if req.Model == "" {
		req.Model = "dall-e-3"
	}
	body, err := oc.post(ctx, "/v1/images/generations", req)
	if err != nil {
		return nil, err
	}
	imageResponse := ImageResponse{}
	if err := json.Unmarshal(body, &imageResponse); err != nil {
		return nil, err
	}
	return &imageResponse, nil
}

