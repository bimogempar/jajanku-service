package user

type Service interface {
	GetAllUser() ([]AllUsers, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAllUser() ([]AllUsers, error) {
	users, err := s.repository.GetAllUser()
	return users, err
}
