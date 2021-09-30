package controllersRefreshToken

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/yona3/go-auth-sample/database"
	"github.com/yona3/go-auth-sample/ent/user"
	"github.com/yona3/go-auth-sample/utils"
)

type RefreshTokenController struct{}

type RefreshTokenResponse struct {
	Ok          bool      `json:"ok"`
	Message     string    `json:"message"`
	LoggedOutAt time.Time `json:"logged_out_at"`
}

func NewRefreshTokenController() *RefreshTokenController {
	return &RefreshTokenController{}
}

func (c *RefreshTokenController) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "DELETE":
		c.delete(w, r)
	default:
		msg := "Method not allowed"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusMethodNotAllowed,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
	}
}

// DELETE: /refresh_token
func (c *RefreshTokenController) delete(w http.ResponseWriter, r *http.Request) {
	// (pass: auth token middleware)

	ctx := r.Context()

	// get jwt claims from context
	userUUID, err := utils.GetContextUserUUID(ctx)
	if err != nil {
		utils.HandleServerError(w, err)
	}

	// parse uuid from string
	id, err := uuid.Parse(userUUID)
	if err != nil {
		msg := "UUID is invalid"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusBadRequest,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	// get user from database
	db := database.GetClient()
	u, err := db.User.Query().Where(user.UUID(id)).Only(ctx)
	if err != nil {
		msg := "User not found"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusNotFound,
			Message: msg,
		}
		utils.HandleServerError(w, err, opts)
		return
	}

	// update logged out at in the database
	now := time.Now()
	if _, err := db.User.UpdateOne(u).SetLoggedOutAt(now).Save(ctx); err != nil {
		msg := "Failed to update logged out at"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		utils.HandleServerError(w, err, opts)
		return
	}

	// revoke the cookie refresh token
	if err := utils.RevokeRefreshToken(w, r); err != nil {
		msg := "Failed to revoke the cookie refresh token"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		utils.HandleServerError(w, err, opts)
	}

	// return response
	data := RefreshTokenResponse{
		Ok:          true,
		Message:     "Logged out successfully",
		LoggedOutAt: now,
	}

	res, err := json.Marshal(data)
	if err != nil {
		msg := "Failed to marshal response"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Message: msg,
		}
		utils.HandleServerError(w, err, opts)
		return
	}

	log.Println("DELETE: /refresh_token")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
