package main

import (
	"context"
)

type Customer struct {
	ID    string `redis:"id"`
	Name  string `redis:"name"`
	Email string `redis:"email"`
	Phone string `redis:"phone"`
	Age   int    `redis:"age"`
}

type CustomerRepo struct {
	Cli *RedisManager
	Db  int
}

func NewCustomerRepo() (*CustomerRepo, error) {
	cli, err := NewRedisClient(CustomerDb)
	if err != nil {
		return nil, err
	}
	return &CustomerRepo{
		Cli: cli,
	}, nil
}

func (c *CustomerRepo) Save(customer *Customer) error {
	return c.Cli.Client.HSet(context.TODO(), customer.ID, customer).Err()
}

func (c *CustomerRepo) Get(id string) (*Customer, error) {
	customer := &Customer{}
	resMap := c.Cli.Client.HGetAll(context.TODO(), id)
	if resMap.Err() != nil {
		return nil, resMap.Err()
	}
	if len(resMap.Val()) == 0 {
		return nil, nil
	}
	err := resMap.Scan(customer)
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (c *CustomerRepo) Delete(id string) error {
	return c.Cli.Client.Del(context.TODO(), id).Err()
}

func (c *CustomerRepo) Update(customer *Customer) error {
	return c.Cli.Client.HSet(context.TODO(), customer.ID, customer).Err()
}
