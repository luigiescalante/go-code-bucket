package bundb

import (
	"context"
	"github.com/uptrace/bun"
	"go.code-bucket/bun-orm/models"
)

type CustomersRepo struct {
	Repo *bun.DB
}

func NewCustomerRepo() (*CustomersRepo, error) {
	db := Db()
	return &CustomersRepo{
		Repo: db,
	}, nil
}

func (cs CustomersRepo) Save(customer *models.Customers) (*models.Customers, error) {
	err := cs.Repo.NewInsert().Model(customer).Scan(context.TODO(), customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (cs CustomersRepo) Update(customer *models.Customers) (*models.Customers, error) {
	_, err := cs.Repo.NewUpdate().Model(customer).Where("id = ? ", customer.ID).Exec(context.TODO())
	return customer, err
}

func (cs CustomersRepo) GetByEmail(email string) (*models.Customers, error) {
	var customer models.Customers
	err := cs.Repo.NewSelect().Model(&customer).Where("email = ?", email).Scan(context.TODO(), &customer)
	return &customer, err
}

func (cs CustomersRepo) GetCustomers() ([]*models.Customers, error) {
	customers := make([]*models.Customers, 0)
	err := cs.Repo.NewSelect().Model(&models.Customers{}).Scan(context.TODO(), &customers)
	return customers, err
}

func (cs CustomersRepo) Delete(customer *models.Customers) error {
	_, err := cs.Repo.NewDelete().Model(customer).Where("id = ?", customer.ID).Exec(context.TODO())
	return err
}
