package service

import "exchange/domain"

type GetByFromUseCase interface {
	Execute([]domain.Exchange) ([]domain.Exchange, error)
}

type GetByFromRepository interface {
	GetByFrom(string, []domain.Exchange) ([]domain.Exchange, error)
}

type getByFromRepository struct{
	getByFromRepo GetByFromRepository
}

func NewGetByFromUseCase (getByFromRepo GetByFromRepository) *getByFromRepository{
	return &getByFromRepository{
		getByFromRepo: getByFromRepo,
	}
}

func (ga getByFromRepository) Execute(from string, e []domain.Exchange) ([]domain.Exchange, error){
	exchanges, err := ga.getByFromRepo.GetByFrom(from, e)
	if err != nil{
		return []domain.Exchange{}, err
	}
	return exchanges, nil
}