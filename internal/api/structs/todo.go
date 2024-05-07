package structs

type TodoResponse struct {
	Todos []Todo `json:"todos"`
	Total int    `json:"total"`
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
}
type Todo struct {
	ID        int    `json:"id"`
	Todo      string `json:"todo"`
	Completed bool   `json:"completed"`
	Userid    int    `json:"userid"`
}
