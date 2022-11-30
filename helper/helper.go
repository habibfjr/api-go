package helper

type JsonResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EmptyData struct{}

func WriteResponse(status int, message string, data interface{}) JsonResponse {
	return JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func WriteErrResponse(status int, message string, data interface{}) JsonResponse {
	// split := strings.Split(err, "\n")
	return JsonResponse{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
