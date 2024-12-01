package api

import (
	"github.com/brentellingson/entra-playground/internal/config"
	"github.com/gin-gonic/gin"
)

// Server represents the API server.
//
//	@Title									Entra Playground API
//	@Version								1.0
//	@Accept									json
//	@Produce								json
//
//	@SecurityDefinitions.OAuth2.AccessCode	OAuth2Entra
//	@AuthorizationUrl						http://localhost:8080/api/authorize
//	@TokenUrl								http://localhost:8080/api/token
type Server struct {
	Cfg *config.Config
}

// NewServer creates a new API server.
func NewServer(cfg *config.Config) *Server {
	return &Server{cfg}
}

// Register registers the API routes.
func (s *Server) Register(r *gin.Engine) {
	r.POST("/api/token", s.Token)
	r.GET("/api/authorize", s.Authorize)
	r.GET("/api/validate", s.Validate)
}

// Token represents the OAuth 2.0 Token Mediation Endpoint.
//
//	@Summary		Get a token
//	@Description	OAuth 2.0 Token Mediation Endpoint.
//
//	@Description	See https://datatracker.ietf.org/doc/html/draft-ietf-oauth-browser-based-apps#name-token-mediating-backend
//
//	@Accept			x-www-form-urlencoded
//	@Produce		json
//	@Param			code			formData	string	true	"OAuth 2.0 Authorization Code"
//	@Param			code_verifier	formData	string	false	"PKCE Code Verifier"
//	@Param			redirect_uri	formData	string	true	"OAuth 2.0 Redirect URI"
//	@Success		200				{object}	any
//	@Router			/token [post]
func (s *Server) Token(g *gin.Context) {
	g.String(200, "Hello, World!")
}

// Authorize represents the OAuth 2.0 Authorization Mediation Endpoint.
//
//	@Summary		Get a token
//	@Description	OAuth 2.0 Authorization Mediation Endpoint.
//	@Description	See https://datatracker.ietf.org/doc/html/draft-ietf-oauth-browser-based-apps#name-token-mediating-backend
//	@Param			redirect_uri			query	string	true	"OAuth 2.0 Redirect URI"
//	@Param			state					query	string	false	"OAuth 2.0 State"
//	@Param			code_challenge			query	string	false	"PKCE Code Challenge"
//	@Param			code_challenge_method	query	string	false	"PKCD Code Challenge Method"
//	@Success		302
//	@Router			/authorize [get]
func (s *Server) Authorize(g *gin.Context) {
	g.Redirect(302, "/")
}

// Validate represents the OAuth 2.0 Token Validation Endpoint.
//
//	@Summary	Get a token
//	@Router		/validate [get]
//	@Security	OAuth2Entra
func (s *Server) Validate(g *gin.Context) {
	g.String(200, "Hello, World!")
}
