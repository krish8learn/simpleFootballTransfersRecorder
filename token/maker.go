package token

import "time"

type Maker interface{
	//creating token 
	CreateToken(username string, duration time.Duration)(string, error)
	//verifying token
	VerifyToken(token string)(*Payload, error)
}
