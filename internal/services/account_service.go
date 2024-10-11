package services

type AccountService interface {
}

type AccountServiceImpl struct {
}

func NewAccountService() AccountService {
	return &AccountServiceImpl{}
}
