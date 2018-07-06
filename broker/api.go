package broker

import (
	"github.com/Symantec/cloud-gate/broker/configuration"
)

type PermittedAccount struct {
	Name              string
	HumanName         string
	PermittedRoleName []string
}

type AWSCredentialsJSON struct {
	SessionId    string `json:"sessionId"`
	SessionKey   string `json:"sessionKey"`
	SessionToken string `json:"sessionToken"`
}

type Broker interface {
	UpdateConfiguration(config *configuration.Configuration) error
	GetUserAllowedAccounts(username string) ([]PermittedAccount, error)
	UserAllowedToAssumeRole(username string, accountName string, roleName string) (bool, error)
	GetConsoleURLForAccountRole(accountName string, roleName string, username string) (string, error)
	GenerateTokenCredentials(accountName string, roleName string, username string) (*AWSCredentialsJSON, error)
}
