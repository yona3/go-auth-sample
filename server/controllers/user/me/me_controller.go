package controllersMe

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/yona3/go-auth-sample/database"
	"github.com/yona3/go-auth-sample/ent/user"
	"github.com/yona3/go-auth-sample/utils"
)

type MeController struct{}

type TokenResponse struct {
	Ok          bool   `json:"ok"`
	AccessToken string `json:"access_token"`
}

func NewMeController() *MeController {
	return &MeController{}
}

func (c *MeController) Index(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		c.get(w, r)
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

// GET: /users/me
func (c *MeController) get(w http.ResponseWriter, r *http.Request) {
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
	db := database.GetClient()

	// get user
	u, err := db.User.Query().Where(user.UUID(id)).Only(ctx)
	if err != nil {
		msg := "User not found"
		log.Println(msg)

		opts := utils.HandleServerErrorOptions{
			Code:    http.StatusNotFound,
			Message: msg,
		}
		utils.HandleServerError(w, nil, opts)
		return
	}

	// return user
	res, err := json.Marshal(u)
	if err != nil {
		utils.HandleServerError(w, err)
		return
	}

	log.Println("GET: /users/me")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
