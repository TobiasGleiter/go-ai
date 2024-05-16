package types

type OllamaOptions struct {
	NumCtx int `json:"num_ctx,omitempty"`
}

type OllamaModel struct {
    Model    string `json:"model"`
    Prompt   string `json:"prompt"`
    Options  OllamaOptions
    Stream bool `json:"stream"`
	Format   string   `json:"format,omitempty"`
}

type OllamaStreamResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type OllamaChatResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Message  Message `json:"message"`
	Done      bool   `json:"done"`
}

type OllamaResponse struct {
	Model              string        `json:"model"`
	CreatedAt          string        `json:"created_at"`
	Response           string        `json:"response"`
	Done               bool          `json:"done"`
	Context            []interface{} `json:"context"`
	TotalDuration      int64         `json:"total_duration"`
	LoadDuration       int64         `json:"load_duration"`
	PromptEvalCount    int           `json:"prompt_eval_count"`
	PromptEvalDuration int64         `json:"prompt_eval_duration"`
	EvalCount          int           `json:"eval_count"`
	EvalDuration       int64         `json:"eval_duration"`
}