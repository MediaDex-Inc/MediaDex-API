package authentication

import (
	"mediadex/config"
	"mediadex/database/dbmodel"
	"mediadex/pkg/user"
	"net/http"

	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

type AuthConfig struct {
	*config.Config
}

func New(configuration *config.Config) *AuthConfig {
	return &AuthConfig{configuration}
}

// @Summary		User login
// @Description	Authenticate user and return JWT token
// @Tags			authentication
// @Accept			json
// @Produce		json
// @Param			request	body		user.UserRequest	true	"Login credentials"
// @Success		200		{object}	TokenResponse
// @Failure 400 {object} map[string]string
// @Router			/auth/login [post]
func (config *AuthConfig) LoginHandler(w http.ResponseWriter, r *http.Request) {
	req := &user.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request req"})
		return
	}

	user, err := config.UserRepository.FindByEmail(req.Email)
	if err != nil {
		user, err = config.UserRepository.FindByUsername(req.Username)
		if err != nil {
			render.JSON(w, r, map[string]string{"error": "Invalid email or password"})
			return
		}
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)) != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid email or password"})
		return
	}

	accessToken, err := GenerateToken(config.JWTSecret, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(config.JWTSecret, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}

	tokens := &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}

// @Summary		User register
// @Description	Create a new user and return JWT tokens
// @Tags			authentication
// @Accept			json
// @Produce		json
// @Param			request	body		user.UserRequest	true	"Register credentials"
// @Success		200		{object}	TokenResponse
// @Failure 400 {object} map[string]string
// @Router			/auth/register [post]
func (config *AuthConfig) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	req := &user.UserRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request payload"})
		return
	}

	_, err := config.UserRepository.FindByEmail(req.Email)
	if err == nil {
		render.JSON(w, r, map[string]string{"error": " email or pseudo already in use"})
		return
	}
	_, err = config.UserRepository.FindByUsername(req.Username)
	if err == nil {
		render.JSON(w, r, map[string]string{"error": " email or pseudo already in use"})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	req.Password = string(hashedPassword)

	userEntry := &dbmodel.User{Email: req.Email, Password: req.Password, Username: req.Username}
	res, err := config.UserRepository.Create(userEntry)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to create user"})
		return
	}
	user := &user.UserResponse{ID: res.ID, Email: res.Email, Username: res.Username}

	accessToken, err := GenerateToken(config.JWTSecret, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(config.JWTSecret, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}
	tokens := &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}

// @Summary		Refresh token
// @Description	Generate a new JWT token from an existing valid refresh token
// @Tags			authentication
// @Accept			json
// @Produce		json
// @Param			request	body		TokenRequest	true	"Refresh token"
// @Success		200		{object}	TokenResponse
// @Failure 400 {object} map[string]string
// @Router			/auth/refresh [post]
func (config *AuthConfig) RefreshHandler(w http.ResponseWriter, r *http.Request) {
	req := &TokenRequest{}
	if err := render.Bind(r, req); err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid request req"})
		return
	}

	email, err := ParseToken(config.JWTSecret, req.RefreshToken)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Invalid refresh token"})
		return
	}

	user, err := config.UserRepository.FindByEmail(email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "User not found"})
		return
	}
	accessToken, err := GenerateToken(config.JWTSecret, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate token"})
		return
	}
	refreshToken, err := GenerateRefreshToken(config.JWTSecret, user.Email)
	if err != nil {
		render.JSON(w, r, map[string]string{"error": "Failed to generate refresh token"})
		return
	}

	tokens := &TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		TokenType:    "bearer",
	}

	render.JSON(w, r, tokens)
}
