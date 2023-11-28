package models

type User struct {
	ID       int    `json:"userID" db:"userid"`
	Nick     string `json:"nick" db:"nick"`
	Email 	 string `json:"email" db:"email"`
	Password string `json:"-" db:"password"`
}

func (u User) Validate() {

}
