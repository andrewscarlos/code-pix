package model

import (
	"github.com/asaskevish/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)
type Account struct {
	Base `valid:"required"`
	OwnerName string `json:"owner_name" valid"notnull"`
	Bank *Bank `valid:"-"`
	Number string `json:"number" valid:"notnull"`	
}

func (account *Account) isValid() error{
	_, erro := govalidator.ValidateStruct(account)
	if err != nil{
		return err
	}
	return nil
}

func NewAccount(bank *Bank, number string, OwnerName string)(*Account, error){
	account := Account{
		OwnerName: ownerName,
		Bank : bank,
		Number:     number
	}

	account.ID = uuid.NewV4().String()
	account.CreatedAt = time.Now()

	err := account.IsValid()
	if err != nil{
		return nil, err
	}
	return &account, nil
}