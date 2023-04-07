package handlers
//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type DeleteHandler interface {
	Execute(int) (error)
}

type deleteHandler struct{
	repository DeleteHandler
}

func NewDeleteHandler (repository DeleteHandler) *deleteHandler{
	return &deleteHandler{repository: repository}
}

func (d deleteHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := d.repository.Execute(id)
	if err != nil{
		c.JSON(400, "not found")
		return
	}
	c.JSON(200, "deleted")
}