package controllers

type ResponseDTO struct {
	Status string      `json:"status"`
	Error  interface{} `json:"error"`
	Data   interface{} `json:"data"`
}

func Ok(d interface{}) ResponseDTO {
	return ResponseDTO{"ok", nil, d}
}

func Error(e error) ResponseDTO {
	return ResponseDTO{"error", e.Error(), nil}
}
