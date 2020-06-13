package user

import (
	internal "github.com/ernesto2108/AP_CreatyHelp/internal/storage/psql"
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/infrestructure/psql/gateway"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h *CreateUsersHandler) SaveUsersEndPoint(c *gin.Context) {

		var us model.CreateUserCmd
		err := c.BindJSON(&us)

		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusBadRequest)
			return
		}

		u, err := h.Create(&us)
		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user create": u})

}

type CreateUsersHandler struct {
	user.UsersCreateGateway
}

func NewCreateUsersHandler(cl *internal.PostSqlClient) *CreateUsersHandler {
	return &CreateUsersHandler{user.NewUsersCreateGateway(cl)}
}
