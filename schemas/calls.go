package schemas

type Call struct {
	ServiceID   int	`json:"service_id"`
	TypeID	  	int `json:"type_id"`
	ChatID      string `json:"chat_id"`
	Coin        string `json:"coin"`
	CreatedAt   string `json:"created_at"`
}