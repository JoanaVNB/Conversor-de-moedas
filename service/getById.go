package service

import "exchange/domain"

type GetUseCase interface {
	Execute(int, domain.Exchange) (domain.Exchange, error)
}

type GetRepository interface {
	Get(int, domain.Exchange) (domain.Exchange, error)
}

type getRepository struct{
	getRepo GetRepository
}

func NewGetUseCase (getRepo GetRepository) *getRepository{
	return &getRepository{
		getRepo: getRepo,
	}
}

func (g getRepository) Execute(id int, u domain.Exchange) (domain.Exchange, error){
	user, err := g.getRepo.Get(id, u)
	if err != nil{
		return domain.Exchange{}, err
	}
	user.ID = u.ID
	return user, nil
}