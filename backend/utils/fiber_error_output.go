package utils

// ErrorOutput Error struct that must be returned then a request failed
type ErrorOutput struct {
	Message string `json:"message"`
	Data    string `json:"data"`
	Reason  string `json:"reason,omitempty"`
}
