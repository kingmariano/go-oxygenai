package oxygen

import (
	"context"
	"encoding/json"
	"errors"
)

type ChatRequest struct {
	//Required
	Messages []ChatMessage `json:"messages"`

	// ID of the model to use. See the model endpoint compatibility (https://app.oxyapi.uk/v1/models) table for details on which models work with the Chat API.
	Model string `json:"model"`

	// Number between -2.0 and 2.0. Positive values penalize new tokens based on their existing frequency in the text so far, decreasing the model's likelihood to repeat the same line verbatim.
	FrequencyPenalty *float64 `json:"frequency_penalty,omitempty"`

	// The maximum number of tokens that can be generated in the chat completion.
	MaxTokens *int `json:"max_tokens,omitempty"`

	// How many chat completion choices to generate for each input message. Note that you will be charged based on the number of generated tokens across all of the choices. Keep n as 1 to minimize costs. nmax = 2
	N *int `json:"n,omitempty"`

	// if set, partial message deltas will be sent, like in ChatGPT. Tokens will be sent as data-only server-sent events as they become available, with the stream terminated by a data: [DONE] message. Example Python code.
	Stream *bool `json:"stream,omitempty"`

	//What sampling temperature to use, between 0 and 2. Higher values like 0.8 will make the output more random, while lower values like 0.2 will make it more focused and deterministic.
	Temperature *int `json:"temperature,omitempty"`

	//An alternative to sampling with temperature, called nucleus sampling, where the model considers the results of the tokens with top_p probability mass. So 0.1 mea
	TopP *int `json:"top,omitempty"`
}
type ChatResponse struct {
	ID                string        `json:"id"`                 // Unique identifier for the chat completion
	Choices           []interface{} `json:"choices"`            // List of chat completion choices
	Created           int64         `json:"created"`            // Unix timestamp (in seconds) of when the chat completion was created
	Model             string        `json:"model"`              // Model used for the chat completion
	SystemFingerprint string        `json:"system_fingerprint"` // Fingerprint representing the backend configuration that the model runs with
	Object            string        `json:"object"`             // Object type, always set to "chat.completion"
}

// ChatMessage represents a message in the conversation
type ChatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

func (oc *OxygenClient) ChatCompletion(ctx context.Context, req *ChatRequest) (*ChatResponse, error) {
	if len(req.Messages) == 0 {
		return nil, errors.New("input text is required")
	}
	if req.Model == "" {
		req.Model = "gpt-3.5-turbo-1106"
	}
	body, err := oc.post(ctx, "/v1/chat/completions", req)
	if err != nil {
		return nil, err
	}
	chatResponse := ChatResponse{}
	if err := json.Unmarshal(body, &chatResponse); err != nil {
		return nil, err
	}

	return &chatResponse, nil
}
