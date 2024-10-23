// // api/handlers/auth_handler.go
package handlers

// import (
// 	models "cleanAchitech/entities"
// 	"cleanAchitech/infrastucture/usecases"
// 	"encoding/json"
// 	"net/http"
// )

// type AuthHandler struct {
// 	authUseCase *usecases.AuthUseCase
// }

// func NewAuthHandler(authUseCase *usecases.AuthUseCase) *AuthHandler {
// 	return &AuthHandler{authUseCase: authUseCase}
// }

// func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
// 	var user models.User
// 	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	if err := h.authUseCase.Register(r.Context(), &user); err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.WriteHeader(http.StatusCreated)
// 	json.NewEncoder(w).Encode(map[string]string{"message": "User registered successfully"})
// }

// func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
// 	var loginRequest struct {
// 		Email    string `json:"email"`
// 		Password string `json:"password"`
// 	}
// 	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	user, err := h.authUseCase.Login(r.Context(), loginRequest.Email, loginRequest.Password)
// 	if err != nil {
// 		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(user)
// }
