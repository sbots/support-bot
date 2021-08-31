package server

import (
	"context"
	"log"
	"net/http"
	models2 "support-bot/service/models"
)

type Server struct {
	http          *http.Server
	tg            botsPlatform
	vb            botsPlatform
	domain        string
	repo          repo
	authenticator authenticator
}

type repo interface {
	GetBot(id string) (*models2.Bot, error)
	UpsertBot(bot *models2.Bot) (*models2.Bot, error)

	UpsertTenant(tenant *models2.Tenant) (*models2.Tenant, error)
	UpsertUser(user *models2.User) (*models2.User, error)
	GetUserByEmail(email, tenant string) (*models2.User, error)
	GetUserByID(id string) (*models2.User, error)
	GetTenantByID(id string) (*models2.Tenant, error)

	TenantHasSuperuser(tenantID string) (bool, error)
}

type botsPlatform interface {
	ConnectNewBot(token, path string) error
	SendMessage(msg *models2.Message, token, receiver string) error
}

type authenticator interface {
	GetJWT(user *models2.User) (string, error)
}

func New(addr, domain string, tg, vb botsPlatform, r repo, auth authenticator) *Server {
	s := &Server{
		http:          &http.Server{Addr: addr},
		tg:            tg,
		vb:            vb,
		domain:        domain,
		repo:          r,
		authenticator: auth,
	}
	s.http.Handler = s.router()
	return s
}

func (s Server) Run(ctx context.Context) error {
	go func() {
		<-ctx.Done()
		err := s.http.Shutdown(ctx)
		if err != nil {
			log.Println("http service shutdown (", err, ")")
		}
		log.Println("service gracefully stopped")
	}()
	log.Print("starting server")
	return s.http.ListenAndServe()
}

func (s *Server) getEndpointForTgBot(id string) string {
	return telegramEndpoint + id
}

func (s *Server) getEndpointForVbBot(id string) string {
	return viberEndpoint + id
}
