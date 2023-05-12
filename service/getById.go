package service

type GetUseCase interface {
	Execute(int, Exchange) (Exchange, error)
}

type GetRepository interface {
	Get(int, Exchange) (Exchange, error)
}

type getRepository struct{
	getRepo GetRepository
}

func NewGetUseCase (getRepo GetRepository) *getRepository{
	return &getRepository{
		getRepo: getRepo,
	}
}

func (g getRepository) Execute(id int, u Exchange) (Exchange, error){
	user, err := g.getRepo.Get(id, u)
	if err != nil{
		return Exchange{}, err
	}
	user.ID = u.ID
	return user, nil
}