package responses

type ApiResponse struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
}

func CreateSuccessResponse(result interface{}, message string) ApiResponse {
	return ApiResponse{Result: result, Message: message, Status: true}
}

func CreateErrorResponse(result interface{}, message string) ApiResponse {
	return ApiResponse{Result: result, Message: message, Status: false}
}
