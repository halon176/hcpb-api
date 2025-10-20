package schemas

type Call struct {
	ServiceID   int	`json:"service_id"`
	TypeID	  	int `json:"type_id"`
	ChatID      string `json:"chat_id"`
	Coin        string `json:"coin"`
	CreatedAt   string `json:"created_at"`
}

type PaginationParams struct {
	Page     int
	PageSize int
	Offset   int
}

type PaginatedCallsResponse struct {
	Data       any   `json:"data"`
	Page       int   `json:"page"`
	PageSize   int   `json:"page_size"`
	TotalItems int64 `json:"total_items"`
	TotalPages int   `json:"total_pages"`
}