package http_handler

import (
	"fmt"
	"net/http"

	"github.com/giicoo/GiicooAuth/internal/models"
	"github.com/giicoo/GiicooAuth/pkg/data"
	"github.com/giicoo/GiicooAuth/pkg/errTools"
)

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()

	user := models.UserRequest{}

	if err := data.FromJSON(&user, body); err != nil {
		h.log.Error(fmt.Errorf("error with json decoder: %s", err))
		JSONHandleError(w, errTools.ErrInvalidJSON, err)
		return
	}

	err := data.ValidateStructure(user)
	if err != nil {
		h.log.Error(fmt.Errorf("error with validate struct: %s", err))
		JSONHandleError(w, err, err)
		return
	}

	accessToken, refreshToken, err := h.services.AuthService.Login(user.Email, user.Password)
	if err != nil {
		h.log.Error(fmt.Errorf("error with generate tokens: %s", err))
		JSONHandleError(w, err, err)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteStrictMode,
		Path:     "/",
	})

	jwtResponse := models.AccessTokenResponse{AccessToken: accessToken}

	if err := SendResponse(w, jwtResponse); err != nil {
		h.log.Error(fmt.Errorf("err with send response: %s", err))
		JSONHandleError(w, err, nil)
		return
	}
}

func (h *Handler) Validate(w http.ResponseWriter, r *http.Request) {
	body := r.Body
	defer body.Close()

	accessTokenReq := models.AccessTokenRequest{}

	if err := data.FromJSON(&accessTokenReq, body); err != nil {
		h.log.Error(fmt.Errorf("error with json decoder: %s", err))
		JSONHandleError(w, errTools.ErrInvalidJSON, err)
		return
	}

	err := data.ValidateStructure(accessTokenReq)
	if err != nil {
		h.log.Error(fmt.Errorf("error with validate struct: %s", err))
		JSONHandleError(w, err, err)
		return
	}

	// authHeader := r.Header.Get("Authorization")
	// if authHeader == "" {
	// 	h.log.Error(fmt.Errorf("error with get access token: %s", err))
	// 	JSONHandleError(w, err, err)
	// 	return
	// }

	// parts := strings.Split(authHeader, " ")
	// if len(parts) != 2 || parts[0] != "Bearer" {
	// 	h.log.Error(fmt.Errorf("error with get access token: %s", err))
	// 	JSONHandleError(w, err, err)
	// 	return
	// }

	accessToken := accessTokenReq.AccessToken
	userID, err := h.services.AuthService.ValidateAccessToken(accessToken)
	// if access token is valid
	if err == nil {
		validResponse := models.ValidAccessTokenResponse{Valid: true, UserID: userID}
		if err := SendResponse(w, validResponse); err != nil {
			h.log.Error(fmt.Errorf("err with send response: %s", err))
			JSONHandleError(w, err, err)
			return
		}
		return
	}

	refreshCookie, err := r.Cookie("refresh_token")
	if err != nil {
		h.log.Error(fmt.Errorf("err get refresh cookie: %s", err))
		JSONHandleError(w, errTools.ErrUnauthorized, err)
		return
	}

	userID, err = h.services.AuthService.ValidateRefreshToken(refreshCookie.Value)
	if err != nil {
		h.log.Error(fmt.Errorf("err validate refresh token: %s", err))
		JSONHandleError(w, errTools.ErrUnauthorized, err)
		return
	}

	// Генерируем новый Access Token
	access, err := h.services.AuthService.GenerateNewAccessToken(userID)
	if err != nil {
		h.log.Error(fmt.Errorf("err get generate new access token: %s", err))
		JSONHandleError(w, err, err)
		return
	}

	invalidResponse := models.InvalidAccessTokenResponse{
		Valid:       false,
		AccessToken: access,
	}

	if err := SendResponse(w, invalidResponse); err != nil {
		h.log.Error(fmt.Errorf("err with send response: %s", err))
		JSONHandleError(w, err, nil)
		return
	}
}
