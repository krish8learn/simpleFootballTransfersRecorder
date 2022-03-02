package Util

import "fmt"

type CreateusersParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

/*Need --> Generate Random username*/
func Randomusername() string {
	return RandomStringGenerator(10)
}

/*Need --> generate Random fullname*/
func Randomfullname() string {
	return RandomStringGenerator(20)
}

/*Need --> generate random Email*/
func Randomemail() string {
	return fmt.Sprintf(RandomStringGenerator(10) + "@gmail.com")
}
