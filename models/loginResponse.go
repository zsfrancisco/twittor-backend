package models

/* LoginResponse has the token returned with login */
type LoginResponse struct {
	Token string `json:"token,omitempty"`
}
