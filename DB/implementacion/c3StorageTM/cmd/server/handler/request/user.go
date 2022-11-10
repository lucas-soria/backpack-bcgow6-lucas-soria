package request

import "github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"

type UserPOSTRequest struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IP         string `json:"ip"`
	MacAddress string `json:"macAddress"`
	Website    string `json:"website"`
	Image      string `json:"image"`
}

func (userPOSTRequest *UserPOSTRequest) MapToDomain() domain.User {
	return domain.User{
		Firstname:  userPOSTRequest.Firstname,
		Lastname:   userPOSTRequest.Lastname,
		Username:   userPOSTRequest.Username,
		Password:   userPOSTRequest.Password,
		Email:      userPOSTRequest.Email,
		IP:         userPOSTRequest.IP,
		MacAddress: userPOSTRequest.MacAddress,
		Website:    userPOSTRequest.Website,
		Image:      userPOSTRequest.Image,
	}
}
