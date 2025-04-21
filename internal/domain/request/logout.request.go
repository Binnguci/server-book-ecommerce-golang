package request

type LogoutData struct {
	Token string `json:"token" binding:"required"`
}
