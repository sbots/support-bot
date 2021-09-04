package server

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"support-bot/infrastructure/auth"
	"support-bot/models"
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
	GetBot(id string) (*models.Bot, error)
	UpsertBot(bot *models.Bot) error

	UpsertTenant(tenant *models.Tenant) error
	UpsertUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id string) (*models.User, error)
	GetTenantByID(id string) (*models.Tenant, error)

	TenantHasSuperuser(tenantID string) (bool, error)
}

type botsPlatform interface {
	ConnectNewBot(token, path string) error
	SendMessage(msg *models.Message, token, receiver string) error
}

type authenticator interface {
	GetToken(user *models.User) (string, error)
	ParseToken(token string) (auth.JWTToken, error)

	SetServiceTokenToContext(ctx context.Context, serviceToken auth.JWTToken) context.Context
	GetServiceTokenFromContext(ctx context.Context) auth.JWTToken
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

func prepareResponse(w http.ResponseWriter, rsp interface{}) error {
	bytes, err := json.Marshal(rsp)
	if err != nil {
		return err
	}

	if _, err := w.Write(bytes); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	return nil
}

func closeBody(b io.ReadCloser) {
	if err := b.Close(); err != nil {
		log.Println(err)
	}
}
