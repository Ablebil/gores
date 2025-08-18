package gores

// http status codes
const (
	StatusOK                  = 200
	StatusCreated             = 201
	StatusAccepted            = 202
	StatusNoContent           = 204
	StatusBadRequest          = 400
	StatusUnauthorized        = 401
	StatusForbidden           = 403
	StatusNotFound            = 404
	StatusMethodNotAllowed    = 405
	StatusConflict            = 409
	StatusUnprocessableEntity = 422
	StatusTooManyRequests     = 429
	StatusInternalServerError = 500
	StatusBadGateway          = 502
	StatusServiceUnavailable  = 503
)

// error codes
const (
	ErrCodeBadRequest         = "BAD_REQUEST"
	ErrCodeUnauthorized       = "UNAUTHORIZED"
	ErrCodeForbidden          = "FORBIDDEN"
	ErrCodeNotFound           = "NOT_FOUND"
	ErrCodeMethodNotAllowed   = "METHOD_NOT_ALLOWED"
	ErrCodeConflict           = "CONFLICT"
	ErrCodeValidation         = "VALIDATION_ERROR"
	ErrCodeTooManyRequests    = "TOO_MANY_REQUESTS"
	ErrCodeInternalServer     = "INTERNAL_SERVER_ERROR"
	ErrCodeBadGateway         = "BAD_GATEWAY"
	ErrCodeServiceUnavailable = "SERVICE_UNAVAILABLE"
)

// default messages
const (
	MsgSuccess            = "Success"
	MsgCreated            = "Resource created successfully"
	MsgUpdated            = "Resource updated successfully"
	MsgDeleted            = "Resource deleted successfully"
	MsgBadRequest         = "Bad request"
	MsgUnauthorized       = "Unauthorized"
	MsgForbidden          = "Forbidden"
	MsgNotFound           = "Resource not found"
	MsgMethodNotAllowed   = "Method not allowed"
	MsgConflict           = "Resource conflict"
	MsgValidation         = "Validation error"
	MsgTooManyRequests    = "Too many requests"
	MsgInternalServer     = "Internal server error"
	MsgBadGateway         = "Bad gateway"
	MsgServiceUnavailable = "Service unavailable"
)
