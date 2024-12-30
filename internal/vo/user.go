// Value Object for User
package vo

type UserRegisterVO struct {
	Email   string `json:"email" binding:"required,email"`
	Purpose string `json:"purpose" binding:"required"`
}
