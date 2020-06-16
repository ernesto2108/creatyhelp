package user

import (
	internal "github.com/ernesto2108/AP_CreatyHelp/internal/storage/psql"
	model "github.com/ernesto2108/AP_CreatyHelp/pkg/user/domain/models"
	"github.com/ernesto2108/AP_CreatyHelp/pkg/user/infrestructure/psql/gateway"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h *CreateUsersHandler) SaveUsersEndPoint(c *gin.Context) {

		var us model.CreateUserCmd
		err := c.BindJSON(&us)

		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusBadRequest)
			return
		}

		u, err := h.gtw.Create(&us)
		if err != nil {
			log.Fatal(err)
			c.Status(http.StatusBadRequest)
			return
		}

		c.JSON(http.StatusOK, gin.H{"user create": u})

}

func (h *CreateUsersHandler) UpdateUsersEndPoint(c *gin.Context) {
}

func (h *CreateUsersHandler) GetIdUsersEndPoint(c *gin.Context) {
	id, _ := c.Params.Get("id")
	idN, _ := strconv.ParseInt(id, 10, 64)

	user, err := h.gtw.GetId(idN)
	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})

}

func (h *CreateUsersHandler) DeleteUsersEndPoint(c *gin.Context) {
}

func (h *CreateUsersHandler) GetAllUsersEndPoint(c *gin.Context) {
}

type CreateUsersHandler struct {
	gtw user.UsersCreateGateway
}

func NewCreateUsersHandler(cl *internal.PostSqlClient) *CreateUsersHandler {
	return &CreateUsersHandler{user.NewUsersCreateGateway(cl)}
}
