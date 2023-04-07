package service

import (
	"exchange/domain"
)

type CreateUseCase interface{
	Execute (domain.Exchange) (domain.Exchange, error)
}

type CreateRepository interface {
	Create(domain.Exchange) (domain.Exchange, error)
}

type createRepository struct{
	createRepo CreateRepository
}

func NewCreateUseCase (createRepo CreateRepository) *createRepository{
	return &createRepository{
		createRepo: createRepo,
	}
}

func (c createRepository) Execute(amount float64, from string, to string, rate float64) (ex domain.Exchange, err error){
	ex = GetConverted(amount, from, to, rate)
	err = Validate(ex); if err != nil{
		return domain.Exchange{}, err
	}
	exCreated, err := c.createRepo.Create(ex)
	if err != nil{
		return domain.Exchange{}, err
	}
	return exCreated, nil
}