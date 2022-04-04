package book

type LoginService interface {
	LoginUser(userRequest LoginCredentialsRequest) bool
}

type loginInformation struct {
	repository Repository
}

func NewLoginService(repository Repository) *loginInformation {
	return &loginInformation{repository}
}

/*func StaticLoginService() *loginInformation {
	return &loginInformation{
		email:    "praditio.nugraha@gmail.com",
		password: "test123",
	}
}*/

func (info *loginInformation) LoginUser(userRequest LoginCredentialsRequest) bool {
	var user LoginCredentials

	user.Email = userRequest.Email
	user.Password = userRequest.Password

	userResponse, _ := info.repository.Authenticated(user)

	return userResponse.Email != "" && userResponse.Password != ""
}
