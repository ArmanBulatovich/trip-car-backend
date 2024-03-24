package responses

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	Code    int         `json:"code"`
}

func CreateSuccessResponse(result interface{}, message string) ApiResponse {
	return ApiResponse{Result: result, Message: message, Status: true}
}

func CreateErrorResponse(result interface{}, message string, code int) ApiResponse {
	return ApiResponse{Result: result, Message: message, Status: false, Code: code}
}

// Codes
const (
	OK = iota
	Error
	InvalidRequestBody
	InvalidName
	InvalidSlug
	InvalidMetadata
)
