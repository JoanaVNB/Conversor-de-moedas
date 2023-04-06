package service

type DeleteUseCase interface {
	Execute(int) (error)
}

type DeleteRepository interface {
	Delete(int) (error)
}

type deleteRepository struct{
	deleteRepo DeleteRepository
}

func NewDeleteUseCase (deleteRepo DeleteRepository) *deleteRepository{
	return &deleteRepository{
		deleteRepo: deleteRepo,
	}
}

func (d deleteRepository) Execute(id int) (error){
	err := d.deleteRepo.Delete(id)
	if err != nil{
		return  err
	}
	return nil
}