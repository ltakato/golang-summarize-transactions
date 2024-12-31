package config

import "os"

type ImapConfig struct {
	Email    string
	Password string
}

type StorageConfig struct {
	BucketName     string
	EmailCsvFolder string
}

type EmailEngineConfig struct {
	ImapConfig
	LocalFolder string
	StorageConfig
}

func NewEmailEngineConfig() *EmailEngineConfig {
	email := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASSWORD")
	bucketName := "summary-transactions"
	emailCsvFolder := "email-csv"
	localFolder := "/tmp"

	return &EmailEngineConfig{
		ImapConfig: ImapConfig{
			Email:    email,
			Password: password,
		},
		LocalFolder: localFolder,
		StorageConfig: StorageConfig{
			BucketName:     bucketName,
			EmailCsvFolder: emailCsvFolder,
		},
	}
}
