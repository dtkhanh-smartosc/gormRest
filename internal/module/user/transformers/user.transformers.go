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
	transformer.Validator.RegisterStructValidation(ValidatePhoneNumber, dto.CreateUserRequest{})
	if errValid := transformer.Validator.Struct(data); errValid != nil {
		return errValid
	}
	return nil
}

//check by default
func ValidatePhoneNumber(sl validator.StructLevel) {
	user := sl.Current().Interface().(dto.CreateUserRequest)
	if len(user.PhoneNumber) == 10 {
		for _, c := range user.PhoneNumber {
			if c < '0' || c > '9' {
				sl.ReportError(user.PhoneNumber, "phone_number", "PhoneNumber", "vnpnumber", "")
			}
		}
	} else {
		sl.ReportError(user.PhoneNumber, "phone_number", "PhoneNumber", "vnpnumber", "")
	}
}

//func ValidatePhoneNumber(sl validator.StructLevel) {
//	user := sl.Current().Interface().(dto.CreateUserRequest)
//	if match, _ := regexp.MatchString("/(((\\+|)84)|0)(3|5|7|8|9)+([0-9]{8})\\b/", user.PhoneNumber); match != true {
//		sl.ReportError(user.PhoneNumber, "phone_number", "PhoneNumber", "vnpnumber", "")
//	}
//}
