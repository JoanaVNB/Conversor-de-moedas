package rate

import (
	"exchange/service"
)

type CreateWithRateUseCase interface{
	Execute (service.Exchange) (service.Exchange, error)
}

type CreateWithRateRepository interface {
	Create(service.Exchange) (service.Exchange, error)
}

type createWithRateRepository struct{
	createWithRateRepo CreateWithRateRepository
}

func NewCreateWithRateUseCase (createWithRateRepo CreateWithRateRepository) *createWithRateRepository{
	return &createWithRateRepository{
		createWithRateRepo: createWithRateRepo,
	}
}

func (c createWithRateRepository) Execute(amount float64, from string, to string) (service.Exchange, error){
	result, rate := ChoiseFX(amount, from, to)
	ex := GetConverted(amount, from, to, rate, result)
	err := service.Validate(ex); if err != nil{
		return service.Exchange{}, err
	}
	exCreated, err := c.createWithRateRepo.Create(ex)
	if err != nil{
		return service.Exchange{}, err
	}
	return exCreated, nil

}