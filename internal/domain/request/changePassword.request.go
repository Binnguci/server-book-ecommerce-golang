package request

type ChangePasswordRequest struct {
	OldPassword string `json:"old_password" binding:"required,min=8,max=20"`
	NewPassword string `json:"new_password" binding:"required,min=8,max=20"`
}
