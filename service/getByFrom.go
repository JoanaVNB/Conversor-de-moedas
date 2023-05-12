package service

type GetByFromUseCase interface {
	Execute([]Exchange) ([]Exchange, error)
}

type GetByFromRepository interface {
	GetByFrom(string, []Exchange) ([]Exchange, error)
}

type getByFromRepository struct{
	getByFromRepo GetByFromRepository
}

func NewGetByFromUseCase (getByFromRepo GetByFromRepository) *getByFromRepository{
	return &getByFromRepository{
		getByFromRepo: getByFromRepo,
	}
}

func (ga getByFromRepository) Execute(from string, e []Exchange) ([]Exchange, error){
	exchanges, err := ga.getByFromRepo.GetByFrom(from, e)
	if err != nil{
		return []Exchange{}, err
	}
	return exchanges, nil
}