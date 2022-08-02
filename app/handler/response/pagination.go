package response

type Pagination struct {
	Data  interface{} `json:"data"`
	Limit int         `json:"limit"`
	Page  int         `json:"page"`
	Total int         `json:"total"`
}
