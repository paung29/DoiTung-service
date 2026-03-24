package account

import (
	"github.com/doitung/DoiTung-service/internal/models"
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"
)


func (repo *repository) FindByEmail(email string) (*models.Account, error) {
	var account models.Account
	if err := repo.db.Where("email", email).First(&account).Error; err != nil {
		return  nil, err
	}
	return &account, nil
}

func (repo *repository) FindByID(id uint) (*models.Account, error) {
	return commonrepo.FindByID[models.Account](repo.db, id)
}

func (repo *repository) Create(account *models.Account) error {
	return commonrepo.Create(repo.db, account)
}