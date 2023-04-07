package handlers
//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"exchange/domain"
	"log"
	"strconv"
	"exchange/presenter"
	"github.com/gin-gonic/gin"
)

type CreateHandler interface {
	Execute(float64, string, string, float64) (domain.Exchange, error)
}

type createHandler struct{
	repository CreateHandler
}

func NewCreateHandler (repository CreateHandler) *createHandler{
	return &createHandler{repository: repository}
}

func (cr createHandler) Create(c *gin.Context) {
	amount, _ := strconv.ParseFloat(c.Param("amount"), 64)
	from := c.Param("from")
	to := c.Param("to")
	rate, _ := strconv.ParseFloat(c.Param("rate"), 64)

	ex, err := cr.repository.Execute(amount, from, to, rate)
	if err != nil{
		log.Fatal("error to create")
	}

	presenter := *presenter.PresenterFX(ex)
	c.JSON(201, presenter)
}








