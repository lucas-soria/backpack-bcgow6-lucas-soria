package request

import "github.com/lucas-soria/backpack-bcgow6-lucas-soria/internal/domain"

type UserRequest struct {
	Firstname  string `json:"first_name"`
	Lastname   string `json:"last_name"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	IP         string `json:"ip"`
	MacAddress string `json:"mac_address"`
	Website    string `json:"website"`
	Image      string `json:"image"`
}

func (userRequest *UserRequest) MapToDomain() domain.User {
	return domain.User{
		Firstname:  userRequest.Firstname,
		Lastname:   userRequest.Lastname,
		Username:   userRequest.Username,
		Password:   userRequest.Password,
		Email:      userRequest.Email,
		IP:         userRequest.IP,
		MacAddress: userRequest.MacAddress,
		Website:    userRequest.Website,
		Image:      userRequest.Image,
	}
}
