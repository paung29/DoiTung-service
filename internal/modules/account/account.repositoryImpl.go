package account

import (
	commonrepo "github.com/doitung/DoiTung-service/internal/common/repository"
	"github.com/doitung/DoiTung-service/internal/models"
)

func (repo *repository) FindByEmail(email string) (*models.Account, error) {
	var account models.Account
	if err := repo.db.Where("email", email).First(&account).Error; err != nil {
		return nil, err
	}
	return &account, nil
}

func (repo *repository) FindByID(id uint) (*models.Account, error) {
	return commonrepo.FindByID[models.Account](repo.db, id)
}

func (repo *repository) Create(account *models.Account) error {
	return commonrepo.Create(repo.db, account)
}

func (repo *repository) Update(account *models.Account) error {
	return commonrepo.Save(repo.db, account)
}

func (repo *repository) GetAll() ([]models.Account, error) {
	var accounts []models.Account
	if err := repo.db.Find(&accounts).Error; err != nil {
		return nil, err
	}
	return accounts, nil
}
