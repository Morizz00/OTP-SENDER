package data

type OTPData struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type VerifyData struct {
	User User   `json:"user" validate:"required"`
	Code string `json:"code" validate:"required"`
}

type User struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
}

type SendMSGData struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Message     string `json:"message" validate:"required"`
}
