package entities

type Profile struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	UserID    int    `json:"user_id"`
}

func NewProfile(firstName, lastName, email, phone string, userId int) *Profile {
	return &Profile{
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Phone:     phone,
		UserID:    userId,
	}
}
func NewFakeProfile() Profile {
	return Profile{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "finanzas@gmail.com",
		Phone:     "9123456789",
		UserID:    1,
	}
}
