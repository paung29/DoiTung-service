package account

type AccountService interface {
	CreateAccount(form AccountCreateForm) (AccountCreateResponse, error)
	UpdateAccountInfo(form AccountUpdateInfoForm) (AccountUpdateInfoResponse, error)
	UpdatePassword(form AccountPasswordUpdateForm) (AccountPasswordUpdateResponse, error)
	GetAllAccounts() (AccountLists, error)
	GetAccountById(userId uint) (AccountDetails, error)
}

type service struct {
	accountRepo AccountRepository
}

func NewAuthService(accountRepo AccountRepository) AccountService {
	return &service{
		accountRepo: accountRepo,
	}
}
