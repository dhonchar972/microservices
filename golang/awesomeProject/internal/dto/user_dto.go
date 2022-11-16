package dto

type UserDto struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func CreateNewUserDto(name string, age int) *UserDto {
	return &UserDto{
		Name: name,
		Age:  age,
	}
}
