package service

type CreateUseCase interface{
	Execute (Exchange) (Exchange, error)
}

type CreateRepository interface {
	Create(Exchange) (Exchange, error)
}

type createRepository struct{
	createRepo CreateRepository
}

func NewCreateUseCase (createRepo CreateRepository) *createRepository{
	return &createRepository{
		createRepo: createRepo,
	}
}

func (c createRepository) Execute(amount float64, from string, to string, rate float64) (ex Exchange, err error){
	ex = GetConverted(amount, from, to, rate) //converte para a struct Exchange do service
	err = Validate(ex); if err != nil{
		return Exchange{}, err
	}
	exCreated, err := c.createRepo.Create(ex)
	if err != nil{
		return Exchange{}, err
	}
	return exCreated, nil
}