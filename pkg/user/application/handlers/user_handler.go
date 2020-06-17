package user

import (
	logs "github.com/ernesto2108/AP_CreatyHelp/internal/logs"
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
	id, _ := c.Params.Get("id")
	idN, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusBadRequest)
		return
	}

	u, err := parseUpdateRequest(c,idN)
	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusBadRequest)
		return
	}
	res := h.gtw.Update(u)

	c.JSON(http.StatusOK, gin.H{"user": &res})
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
	id, _ := c.Params.Get("id")
	idN, err := strconv.ParseInt(id, 10, 64)

	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusNotFound)
		return
	}

	u := h.gtw.Delete(idN)
	if err != nil {
		log.Fatal(err)
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{"delete user": u})
}

func (h *CreateUsersHandler) GetAllUsersEndPoint(c *gin.Context) {
	u := h.gtw.GetAll()
	if u == nil || len(u) == 0 {
		u = []*model.User{}
	}
	c.JSON(http.StatusOK, gin.H{"all user": u})
}

type CreateUsersHandler struct {
	gtw user.UsersCreateGateway
}

func parseUpdateRequest(c *gin.Context, id int64) (*model.UpdateUserCmd, error)  {
	var u model.UpdateUserCmd
	err := c.BindJSON(&u)

	if err != nil {
		logs.Log().Error(err.Error())
		return nil, err
	}
	u.ID = id
	return &u, nil
}

func NewCreateUsersHandler(cl *internal.PostSqlClient) *CreateUsersHandler {
	return &CreateUsersHandler{user.NewUsersCreateGateway(cl)}
}
