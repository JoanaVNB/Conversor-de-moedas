package handlers
//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"exchange/service"
	"exchange/presenter"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GetHandler interface {
	Execute(int, service.Exchange) (service.Exchange, error)
}

type getHandler struct{
	repository GetHandler
}

func NewGetHandler (repository GetHandler) *getHandler{
	return &getHandler{repository: repository}
}

func (g getHandler) Get(c *gin.Context) {
	var u service.Exchange

	id, _ := strconv.Atoi(c.Param("id"))
	ex, err := g.repository.Execute(id, u)
	if err != nil{
		c.JSON(404, "not found")
	}

	presenter := *presenter.PresenterFX(ex)
	c.JSON(200, presenter)
}