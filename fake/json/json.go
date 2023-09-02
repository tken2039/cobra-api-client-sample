package json

import (
	"encoding/json"
	"os"

	"github.com/tken2039/cobra-api-client-sample/domain/user"
)

type APIClientFakeJSON struct {
	filepath string
}

func NewAPIClientFakeJSON(filepath string) *APIClientFakeJSON {
	return &APIClientFakeJSON{
		filepath: filepath,
	}
}

func (c *APIClientFakeJSON) Get() ([]*user.User, error) {
	bytes, err := os.ReadFile(c.filepath)
	if err != nil {
		return nil, err
	}

	// json bytes to []*user.User
	users := []*user.User{}
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}

	return users, nil
}
