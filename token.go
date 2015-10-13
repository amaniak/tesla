package tesla

import (
	"encoding/json"
	"fmt"
)

//
type Token struct {
	AccessToken *string `json:"access_token,omitempty"`
	Type        *string `json:"token_type,omitempty"`
	ExpiresIn   *int    `json:"expires_in,omitempty"`
}

func (t *Token) String() string {
	return Stringify(t)
}

func (t *Token) ToJSON() string {
	b, err := json.Marshal(t)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	return string(b)
}
