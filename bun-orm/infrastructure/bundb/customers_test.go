package bundb

import (
	"fmt"
	"go.code-bucket/bun-orm/models"
	"testing"
)

func TestCustomersRepo_GetByEmail(t *testing.T) {
	cusRepo, _ := NewCustomerRepo()
	email := "john.down@test.com"
	byEmail, err := cusRepo.GetByEmail(email)
	fmt.Println(byEmail)
	fmt.Println(err)
}

func TestCustomersRepo_Save(t *testing.T) {
	customer := models.Customers{
		FirstName: "john",
		LastName:  "Down",
		Email:     "john.down@test.com",
		Age:       18,
	}
	cusRepo, _ := NewCustomerRepo()
	save, err := cusRepo.Save(&customer)
	fmt.Println(save)
	fmt.Println(err)
}

func TestCustomersRepo_GetCustomers(t *testing.T) {
	cusRepo, _ := NewCustomerRepo()
	save, err := cusRepo.GetCustomers()
	fmt.Println(save)
	fmt.Println(err)
}

func TestCustomersRepo_Delete(t *testing.T) {
	customer := models.Customers{
		ID:        2,
		FirstName: "john",
		LastName:  "Down",
		Email:     "john.down@test.com",
		Age:       18,
	}
	cusRepo, _ := NewCustomerRepo()
	err := cusRepo.Delete(&customer)
	fmt.Println(err)
}

func TestCustomersRepo_Update(t *testing.T) {
	customer := models.Customers{
		ID:        2,
		FirstName: "kakashi",
		LastName:  "hatake",
		Email:     "k.hatake@test.com",
		Age:       38,
	}
	cusRepo, _ := NewCustomerRepo()
	save, err := cusRepo.Update(&customer)
	fmt.Println(save)
	fmt.Println(err)
}
