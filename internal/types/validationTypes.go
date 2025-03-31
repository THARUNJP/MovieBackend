package types

type LoginRequest struct {
	MobileNo string `json:"mobileNo" validate:"min=10,max=10"`
	Password string `json:"password" validate:"min=8,containsany=!@#$%^&*+,containsany=0123456789"`
}
