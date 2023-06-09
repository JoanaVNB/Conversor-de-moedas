package handlers
//go:generate mockgen -source=$GOFILE -destination=mock_$GOFILE -package=$GOPACKAGE
import (
	"exchange/domain"
	"exchange/presenter"

	"github.com/gin-gonic/gin"
)

type GetAllHandler interface {
	Execute([]domain.Exchange) ([]domain.Exchange, error)
}

type getAllHandler struct{
	repository GetAllHandler
}

func NewGetAllHandler (repository GetAllHandler) *getAllHandler{
	return &getAllHandler{repository: repository}
}

func (ga getAllHandler) GetAll(c *gin.Context){
	var u []domain.Exchange

	Exchanges, err := ga.repository.Execute(u)
	if err != nil{
		c.JSON(400, err)
		return
	}

	for _, eachExchange := range Exchanges{
		Exchange := eachExchange
		presenterExchange := *presenter.PresenterFX(Exchange)
		c.JSON(200, presenterExchange)
	}
}