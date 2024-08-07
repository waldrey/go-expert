package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-chi/jwtauth"
	"github.com/waldrey/go-expert/apis/internal/dto"
	"github.com/waldrey/go-expert/apis/internal/entity"
	"github.com/waldrey/go-expert/apis/internal/infra/database"
)

type Error struct {
	Message string `json:"message"`
}

type UserHandler struct {
	userDB       database.UserInterface
	Jwt          *jwtauth.JWTAuth
	JwtExpiresIn int
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{
		userDB: userDB}
}

// GetJWT godoc
// @Summary 	Get a JWT
// @Description Get a JWT
// @Tags 		users
// @Accept  	json
// @Produce  	json
// @Param 		request		user 	body 		dto.GetJWTInput	true	"user credentials"
// @Success 	200 	{object} 	dto.GetJWTOuput
// @Failure 	400
// @Failure 	500 	{object} 	Error
// @Router 		/users/generate_token [post]
func (h *UserHandler) GetJWT(w http.ResponseWriter, r *http.Request) {
	jwt := r.Context().Value("jwt").(*jwtauth.JWTAuth)
	jwtExpiresIn := r.Context().Value("jwtExpiresIn").(int)
	var user dto.GetJWTInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	u, err := h.userDB.FindByEmail(user.Email)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	if !u.ValidatePassword(user.Password) {
		w.WriteHeader(http.StatusUnauthorized)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}
	_, tokenString, err := jwt.Encode(map[string]interface{}{
		"sub":   u.ID.String(),
		"email": user.Email,
		"exp":   time.Now().Add(time.Second * time.Duration(jwtExpiresIn)).Unix(),
	})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	accessToken := dto.GetJWTOuput{AccessToken: tokenString}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accessToken)
}

// Create user godoc
// @Summary 	Create a user
// @Description Create a user
// @Tags 		users
// @Accept  	json
// @Produce  	json
// @Param 		user 	body 		dto.CreateUserInput 	true 	"user request"
// @Success 	201
// @Failure 	500 	{object} 	Error
// @Router 		/users [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user dto.CreateUserInput
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	u, err := entity.NewUser(user.Name, user.Email, user.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = h.userDB.Create(u)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
