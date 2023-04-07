package rate

import (
	"exchange/domain"
	"exchange/service"
)

type CreateWithRateUseCase interface{
	Execute (domain.Exchange) (domain.Exchange, error)
}

type CreateWithRateRepository interface {
	Create(domain.Exchange) (domain.Exchange, error)
}

type createWithRateRepository struct{
	createWithRateRepo CreateWithRateRepository
}

func NewCreateWithRateUseCase (createWithRateRepo CreateWithRateRepository) *createWithRateRepository{
	return &createWithRateRepository{
		createWithRateRepo: createWithRateRepo,
	}
}

func (c createWithRateRepository) Execute(amount float64, from string, to string) (domain.Exchange, error){
	result, rate := ChoiseFX(amount, from, to)
	ex := GetConverted(amount, from, to, rate, result)
	err := service.Validate(ex); if err != nil{
		return domain.Exchange{}, err
	}
	exCreated, err := c.createWithRateRepo.Create(ex)
	if err != nil{
		return domain.Exchange{}, err
	}
	return exCreated, nil

}