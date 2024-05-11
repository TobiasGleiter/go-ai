package types

type Llama3StreamResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Response  string `json:"response"`
	Done      bool   `json:"done"`
}

type Message struct {
	Role string `json:"role"`
	Content string `json:"content"`
}

type Llama3ChatResponse struct {
	Model     string `json:"model"`
	CreatedAt string `json:"created_at"`
	Message  Message `json:"message"`
	Done      bool   `json:"done"`
}

type Llama3Response struct {
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