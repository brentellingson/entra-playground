package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/brentellingson/entra-playground/internal/config"
	"github.com/brentellingson/entra-playground/internal/errhandler"
	"github.com/golang-jwt/jwt/v5"
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
func (s *Server) Register() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/token", s.Token)
	mux.HandleFunc("GET /api/authorize", s.Authorize)
	mux.HandleFunc("GET /api/validate", s.Validate)
	return mux
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
//	@Router			/api/token [post]
func (s *Server) Token(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		errhandler.Abort(w, http.StatusBadRequest, fmt.Errorf("failed to parse form: %w", err))
		return
	}

	form := req.Form
	form.Set("client_id", s.Cfg.WebAPIA.ClientID)
	form.Set("client_secret", s.Cfg.WebAPIA.ClientSecret)
	form.Set("grant_type", "authorization_code")
	form.Set("scope", s.Cfg.WebAPIA.Scope)

	resp, err := http.PostForm(s.Cfg.WebAPIA.TokenEndpoint, form)
	if err != nil {
		errhandler.Abort(w, http.StatusInternalServerError, fmt.Errorf("failed to forward request: %w", err))
		return
	}
	defer resp.Body.Close()
	w.WriteHeader(resp.StatusCode)
	_, _ = io.Copy(w, resp.Body)
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
//	@Router			/api/authorize [get]
func (s *Server) Authorize(w http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		errhandler.Abort(w, http.StatusBadRequest, fmt.Errorf("failed to parse form: %w", err))
		return
	}

	form := req.Form
	form.Set("client_id", s.Cfg.WebAPIA.ClientID)
	form.Set("response_type", "code")
	form.Set("scope", s.Cfg.WebAPIA.Scope)

	redirect, err := url.Parse(s.Cfg.WebAPIA.AuthorizationEndpoint)
	if err != nil {
		errhandler.Abort(w, http.StatusInternalServerError, fmt.Errorf("failed to parse URL: %w", err))
		return
	}
	redirect.RawQuery = form.Encode()
	http.Redirect(w, req, redirect.String(), http.StatusTemporaryRedirect)
}

// Validate represents the OAuth 2.0 Token Validation Endpoint.
//
//	@Summary	Get a token
//	@Router		/api/validate [get]
//	@Security	OAuth2Entra
func (s *Server) Validate(w http.ResponseWriter, req *http.Request) {
	auth := req.Header.Get("Authorization")
	if !strings.HasPrefix(auth, "Bearer ") {
		errhandler.Abort(w, http.StatusUnauthorized, fmt.Errorf("missing bearer token"))
		return
	}

	tokenString := strings.TrimPrefix(auth, "Bearer ")
	parser := jwt.NewParser(jwt.WithAudience(s.Cfg.WebAPIA.Audience), jwt.WithIssuer(s.Cfg.WebAPIA.Issuer))
	token, _, _ := parser.ParseUnverified(tokenString, &jwt.MapClaims{})
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(token.Claims)
}
