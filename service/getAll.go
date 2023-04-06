package service

import "exchange/domain"

type GetAllUseCase interface {
	Execute([]domain.Exchange) ([]domain.Exchange, error)
}

type GetAllRepository interface {
	GetAll([]domain.Exchange) ([]domain.Exchange, error)
}

type getAllRepository struct{
	getAllRepo GetAllRepository
}

func NewGetAllUseCase (getAllRepo GetAllRepository) *getAllRepository{
	return &getAllRepository{
		getAllRepo: getAllRepo,
	}
}

func (ga getAllRepository) Execute(e []domain.Exchange) ([]domain.Exchange, error){
	exchanges, err := ga.getAllRepo.GetAll(e)
	if err != nil{
		return []domain.Exchange{}, err
	}
	return exchanges, nil
}