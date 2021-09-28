package service

import (
	"support-bot/internal/infrastructure/auth"
	"support-bot/internal/models"
)

type repo interface {
	UpsertBot(bot *models.Bot) error
	GetBotByID(id string) (*models.Bot, error)

	UpsertUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)

	UpsertTenant(tenant *models.Tenant) error
	GetTenantByID(id string) (*models.Tenant, error)
	GetTenantByName(name string) (*models.Tenant, error)
	TenantHasSuperuser(tenantID string) (bool, error)
}

type botsPlatform interface {
	ConnectNewBot(token, path string) error
	SendMessage(msg *models.Message, token, receiver string) error
}

type Service struct {
	repo repo
	tg   botsPlatform
	vb   botsPlatform
	auth auth.Authenticator
}

func New(repo repo, tg botsPlatform, vb botsPlatform, auth auth.Authenticator) *Service {
	return &Service{repo: repo, tg: tg, vb: vb, auth: auth}
}
