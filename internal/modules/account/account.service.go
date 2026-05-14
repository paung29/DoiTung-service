package account

type AccountService interface {
	CreateAccount(form AccountCreateForm) (AccountCreateResponse, error)
	UpdateAccount(form AccountUpdateForm) (AccountUpdateResponse, error)
}

type service struct {
	accountRepo AccountRepository
}

func NewAuthService(accountRepo AccountRepository) AccountService {
	return &service{
		accountRepo: accountRepo,
	}
}
