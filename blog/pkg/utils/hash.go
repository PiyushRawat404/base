package utils

import "golang.org/x/crypto/bcrypt"


func HashPassword(password){
	bytes,err:=bcrypt.GenerateFromPassword([]byte(password),
	bcrypt.DefaultCost,
	)
	return string(bytes),err	
}

