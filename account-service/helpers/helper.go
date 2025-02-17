package helpers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"os"
	"account-service/models"
	"github.com/supabase-community/supabase-go"
)

type SupabaseUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type SupabaseResponse struct {
	Users []SupabaseUser `json:"users"`
}

func GetUserUUIDByEmail(email string) (string, error) {
	baseURL := os.Getenv("DB_URL")
	url := fmt.Sprintf("%s/auth/v1/admin/users?email=%s", baseURL, email)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	req.Header.Set("apikey", os.Getenv("SERVICE_ROLE_KEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SERVICE_ROLE_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response SupabaseResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return "", err
	}

	if len(response.Users) == 0 {
		return "", fmt.Errorf("user not found")
	}

	return response.Users[0].ID, nil
}

func DeleteUserByUUID(uuid string) error {
	baseURL := os.Getenv("DB_URL")
	url := fmt.Sprintf("%s/auth/v1/admin/users/%s", baseURL, uuid)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}

	req.Header.Set("apikey", os.Getenv("SERVICE_ROLE_KEY"))
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SERVICE_ROLE_KEY"))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("failed to delete user, status code: %d", resp.StatusCode)
	}

	return nil
}

func GenerateAccountNumber(client *supabase.Client) (string, error) {
	for {
		num, err := rand.Int(rand.Reader, big.NewInt(9999999999999999))
		if err != nil {
			return "", err
		}

		accountNum := num.String()

		for len(accountNum) < 16 {
			accountNum = "0" + accountNum
		}

		_, count, err := client.From("account").Select("account_number", "exact", false).Eq("account_number", accountNum).Execute()
		if err != nil {
			return "", err
		}

		if count == 0 {
			return accountNum, nil
		}
	}
}

func DataCheck(account models.Account) bool {
	return account.Email != "" &&
		account.Full_Name != "" &&
		account.Phone != "" &&
		account.Password != "" &&
		account.AccountType != ""
}
