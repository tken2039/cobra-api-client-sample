package apicaller

import (
	"encoding/json"
	"fmt"

	"github.com/tken2039/cobra-api-client-sample/domain/user"
)

type APICaller struct {
	APIClient     UserAPI
	outputOptions OutputOptions
}

type OutputOptions struct {
	Format string
	Full   bool
}

func NewAPICaller(apiClient UserAPI, opts OutputOptions) *APICaller {
	return &APICaller{
		APIClient:     apiClient,
		outputOptions: opts,
	}
}

type UserAPI interface {
	Get() ([]*user.User, error)
}

func (c *APICaller) Run() error {
	users, err := c.APIClient.Get()
	if err != nil {
		return err
	}

	if c.outputOptions.Format == "json" {
		return c.outputJson(users)
	} else {
		return c.outputTable(users)
	}
}

func (c *APICaller) outputJson(users []*user.User) error {
	if c.outputOptions.Full {
		json, err := json.Marshal(users)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
	} else {
		type ommitedUser struct {
			UserID     string `json:"user_id"`
			Email      string `json:"email"`
			Name       string `json:"name"`
			GivenName  string `json:"given_name"`
			FamilyName string `json:"family_name"`
			Nickname   string `json:"nickname"`
		}

		ommitedUsers := []*ommitedUser{}
		for _, item := range users {
			ommitedUsers = append(ommitedUsers, &ommitedUser{
				UserID:     item.UserID,
				Email:      item.Email,
				Name:       item.Name,
				GivenName:  item.GivenName,
				FamilyName: item.FamilyName,
				Nickname:   item.Nickname,
			})
		}

		json, err := json.Marshal(ommitedUsers)
		if err != nil {
			return err
		}

		fmt.Println(string(json))
	}

	return nil
}

func (c *APICaller) outputTable(users []*user.User) error {
	if c.outputOptions.Full {
		// Output header
		fmt.Printf("%-24s %-18s %-20s %-10s %-10s %-10s %-15s %-15s %-30s %-30s %-30s %-14s\n",
			"User ID", "Email", "Name", "Given Name", "Family Name", "Nickname", "Last IP", "Logins Count",
			"Created At", "Updated At", "Last Login", "Email Verified")

		// Output data
		for _, item := range users {
			fmt.Printf("%-24s %-18s %-20s %-10s %-10s %-10s %-15s %-15d %-30s %-30s %-30s %-14v\n",
				item.UserID, item.Email, item.Name, item.GivenName, item.FamilyName, item.Nickname, item.LastIP,
				item.LoginsCount, item.CreatedAt, item.UpdatedAt, item.LastLogin, item.EmailVerified)
		}
	} else {
		// Output header
		fmt.Printf("%-24s %-18s %-20s %-10s %-10s %-10s\n",
			"User ID", "Email", "Name", "Given Name", "Family Name", "Nickname")

		// Output data
		for _, item := range users {
			fmt.Printf("%-24s %-18s %-20s %-10s %-10s %-10s\n",
				item.UserID, item.Email, item.Name, item.GivenName, item.FamilyName, item.Nickname)
		}
	}

	return nil
}
