package models

type SignUpCredentials struct {
	Nick     string
	Email    string
	Password string
}

type SignInCredentials struct {
	Login string
	Password string
}