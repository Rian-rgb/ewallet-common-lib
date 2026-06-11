package response

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

type ErrorResponse struct {
	ErrorCode string `json:"error_code"`
	Message   string `json:"message"`
}

type ValidationErrorField struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

type BadRequestResponse struct {
	ErrorCode string                 `json:"error_code"`
	Message   string                 `json:"message"`
	Errors    []ValidationErrorField `json:"errors"`
}

type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	PageSize    int `json:"page_size"`
	TotalPage   int `json:"total_page"`
	TotalData   int `json:"total_data"`
}
