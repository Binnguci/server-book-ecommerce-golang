package exception

const (
	// Success codes
	SuccessCode       = 200
	CreateSuccessCode = 201
	AcceptedCode      = 202

	// Client errors
	BadRequestCode      = 400
	UnauthorizedCode    = 401
	ForbiddenCode       = 403
	NotFoundCode        = 404
	ConflictCode        = 409
	UnprocessableEntity = 422

	// Server errors
	InternalServerErrorCode = 500
	ServiceUnavailableCode  = 503
	GatewayTimeoutCode      = 504

	// Custom application errors
	UserExistsCode   = 1001
	ErrorInvalidOTP  = 1002
	ErrorSendEmail   = 1003
	CreateFailedCode = 1004
	ErrorUpdateCode  = 3001
)

var message = map[int]string{
	// Success messages
	SuccessCode:       "Success",
	CreateSuccessCode: "Create Success",
	AcceptedCode:      "Request accepted for processing",

	// Client error messages
	BadRequestCode:      "Bad request: Please check the input",
	UnauthorizedCode:    "Unauthorized: Invalid credentials",
	ForbiddenCode:       "Forbidden: Access denied",
	NotFoundCode:        "Resource not found",
	ConflictCode:        "Conflict: Resource already exists",
	UnprocessableEntity: "Unprocessable entity: Validation failed",

	// Server error messages
	InternalServerErrorCode: "Internal server error: Something went wrong",
	ServiceUnavailableCode:  "Service unavailable: Please try again later",
	GatewayTimeoutCode:      "Gateway timeout: Request timed out",

	// Custom error messages
	UserExistsCode:   "User already exists",
	ErrorInvalidOTP:  "Invalid OTP",
	ErrorSendEmail:   "Error when sending email",
	CreateFailedCode: "Error when creating resource",
	ErrorUpdateCode:  "Error when updating database",
}

func GetMessage(code int) string {
	if msg, exists := message[code]; exists {
		return msg
	}
	return "Unknown error"
}
