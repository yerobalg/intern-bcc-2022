package user

type LoginUserFormat struct {
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

type RegisterUserFormat struct {
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type ProfilUserFormat struct {
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Email    string `json:"email"`
	NomorHp  string `json:"nomorHP"`
}

func LoginFormat(user *Users, token string) LoginUserFormat {
	return LoginUserFormat{
		Nama:     user.Nama,
		Username: user.Username,
		Token:    token,
	}
}

func RegisterFormat(user *Users) RegisterUserFormat {
	return RegisterUserFormat{
		Nama:     user.Nama,
		Username: user.Username,
		Email:    user.Email,
	}
}

func ProfilUserFormatter (user *Users) ProfilUserFormat{
	return ProfilUserFormat{
		Nama:     user.Nama,
		Username: user.Username,
		Email:    user.Email,
		NomorHp:  user.NomorHp,
	}
}
