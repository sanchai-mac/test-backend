package error_wrapper

var (
	ErrorCodeStatus = map[ErrorCode]string{
		"200": "SUCCESS",
		"201": "CREATED",
		"400": "BAD_REQUEST",
		"401": "UNAUTHORIZED",
		"403": "FORBIDEN",
		"404": "NOT FOUND",
		"410": "GONE",
		"500": "INTERNAL ERROR",
		"503": "SERVER UNAVAILABLE",
	}
)

const (
	SUCCESS ErrorCode = "200"
	CREATED ErrorCode = "201"
	// server error
	BAD_REQUEST        ErrorCode = "400"
	UNAUTHORIZED       ErrorCode = "401"
	FORBIDEN           ErrorCode = "403"
	NOT_FOUND          ErrorCode = "404"
	GONE               ErrorCode = "410"
	INTERNAL_ERROR     ErrorCode = "500"
	SERVER_UNAVAILABLE ErrorCode = "503"
)
