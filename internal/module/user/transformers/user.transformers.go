package transformers

import (
	"github.com/HiBang15/sample-gorm.git/internal/module/user/dto"
	"github.com/HiBang15/sample-gorm.git/internal/module/user/entities"
	"github.com/go-playground/validator/v10"
)

type UserTransformer struct {
	Validator *validator.Validate
}

func NewUserTransformer() *UserTransformer {
	return &UserTransformer{Validator: validator.New()}
}

func (transformer *UserTransformer) UserEntityToDto(data *entities.User) *dto.User {
	return &dto.User{
		Id:          data.Id,
		FirstName:   data.FirstName,
		LastName:    data.LastName,
		Email:       data.Email,
		PhoneNumber: data.PhoneNumber,
		IsVerify:    data.IsVerify,
		CreatedAt:   data.CreatedAt,
		UpdatedAt:   data.UpdatedAt,
	}
}

func (transformer *UserTransformer) VerifyCreateUserRequest(data *dto.CreateUserRequest) error {
	if errValid := transformer.Validator.Struct(data); errValid != nil {
		return errValid
	}
	return nil
}
