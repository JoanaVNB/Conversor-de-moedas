package service

type GetAllUseCase interface {
	Execute([]Exchange) ([]Exchange, error)
}

type GetAllRepository interface {
	GetAll([]Exchange) ([]Exchange, error)
}

type getAllRepository struct{
	getAllRepo GetAllRepository
}

func NewGetAllUseCase (getAllRepo GetAllRepository) *getAllRepository{
	return &getAllRepository{
		getAllRepo: getAllRepo,
	}
}

func (ga getAllRepository) Execute(e []Exchange) ([]Exchange, error){
	exchanges, err := ga.getAllRepo.GetAll(e)
	if err != nil{
		return []Exchange{}, err
	}
	return exchanges, nil
}