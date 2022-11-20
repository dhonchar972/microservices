package author

/*
* Classic Data Transfer Object
*
* Use for tagging and remove redundant fealds.
 */

type CreateAuthorDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
