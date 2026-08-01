package internal

type BrasilAPIError struct {
	Message string `json:"message"`
	Type    string `json:"type"`
}
