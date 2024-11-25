package response

type ResponseWithData struct {
	Code     int         `json:"code"`
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Paginate *Paginate   `json:"paginate,omitempty"`
	Data     interface{} `json:"data"`
}

type ResponseWithoutData struct {
	Code    int    `json:"code"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

func Response(params ResponseParam) interface{} {
	var status string
	if params.StatusCode >= 200 && params.StatusCode < 300 {
		status = "success"
	} else {
		status = "failed"
	}

	if params.Data != nil {
		return &ResponseWithData{
			Code:     params.StatusCode,
			Status:   status,
			Message:  params.Message,
			Paginate: params.Paginate,
			Data:     params.Data,
		}
	}
	return &ResponseWithoutData{
		Code:    params.StatusCode,
		Status:  status,
		Message: params.Message,
	}
}
