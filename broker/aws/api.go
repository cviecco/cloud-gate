package aws

import (
	"errors"
	"time"

	"github.com/Symantec/Dominator/lib/log"
	"github.com/Symantec/cloud-gate/broker"
	"github.com/Symantec/cloud-gate/broker/configuration"
	"github.com/Symantec/cloud-gate/lib/userinfo"
)

type userAllowedCredentialsCacheEntry struct {
	PermittedAccounts []broker.PermittedAccount
	Expiration        time.Time
}

type Broker struct {
	config                      *configuration.Configuration
	userInfo                    userinfo.UserInfo
	credentialsFilename         string
	logger                      log.DebugLogger
	userAllowedCredentialsCache map[string]userAllowedCredentialsCacheEntry
}

func New(userInfo userinfo.UserInfo,
	credentialsFilename string,
	logger log.DebugLogger) *Broker {
	return &Broker{userInfo: userInfo,
		credentialsFilename: credentialsFilename,
		logger:              logger,
		userAllowedCredentialsCache: make(map[string]userAllowedCredentialsCacheEntry),
	}
}

func (b *Broker) UpdateConfiguration(
	config *configuration.Configuration) error {
	if config == nil {
		return errors.New("nill config passed")
	}
	b.logger.Debugf(1, "config=%+v", *config)
	b.config = config
	return nil
}

func (b *Broker) GetUserAllowedAccounts(username string) ([]broker.PermittedAccount, error) {
	return b.getUserAllowedAccounts(username)
}

func (b *Broker) UserAllowedToAssumeRole(username string, accountName string, roleName string) (bool, error) {
	return b.userAllowedToAssumeRole(username, accountName, roleName)
}

func (b *Broker) GetConsoleURLForAccountRole(accountName string, roleName string, userName string) (string, error) {
	return b.getConsoleURLForAccountRole(accountName, roleName, userName)
}

func (b *Broker) GenerateTokenCredentials(accountName string, roleName string, userName string) (*broker.AWSCredentialsJSON, error) {
	return b.generateTokenCredentials(accountName, roleName, userName)
}
