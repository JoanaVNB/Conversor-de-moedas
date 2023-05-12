package handlers
//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"exchange/service"
	"log"
	"strconv"
	"exchange/presenter"
	"github.com/gin-gonic/gin"
)

type CreateRHandler interface {
	Execute(float64, string, string) (service.Exchange, error)
}

type createRHandler struct{
	repository CreateRHandler
}

func NewCreateRHandler (repository CreateRHandler) *createRHandler{
	return &createRHandler{repository: repository}
}

func (cr createRHandler) Create(c *gin.Context) {
	amount, _ := strconv.ParseFloat(c.Param("amount"), 64)
	from := c.Param("from")
	to := c.Param("to")

	ex, err := cr.repository.Execute(amount, from, to)
	if err != nil{
		c.JSON(400, err)
	}

	presenter := *presenter.PresenterFX(ex)
	c.JSON(201, presenter)
}