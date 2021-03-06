package dto

//UserUpdateDTO is used by client when PUT update profile
type UserUpdateDTO struct {
	ID       uint64 `json:"id" form:"id"`
	Name     string `json:"name" form:"name" binding:"required"`
	Email    string `json:"email" form:"email" binding:"required,email"`
	Password string `json:"password,omitempty" form:"password,omitempty"`
}

//UserCreateDTO is used when register a user
// type UserCreateDTO struct {
// 	ID       uint64 `json:"id" form:"id" binding:"required"`
// 	Name     string `json:"name" form:"name" binding:"required"`
// 	Email    string `json:"email" form:"email" binding:"required" validate:"email"`                           //khusus untuk go-playground v10 ada validate email
// 	Password string `json:"password,omitempty" form:"password,omitempty" validate:"min:6" binding:"required"` //validate password minimal 6 char
// }
