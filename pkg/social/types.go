package social

type RetweetRes struct {
	Data   []RetweetResData `json:"data"`
	Meta   RetweetResMeta   `json:"meta"`
	Status int              `json:"status"`
}

type RetweetResData struct {
	Username string `json:"username"`
}

type RetweetResMeta struct {
	ResultCount int    `json:"result_count"`
	NextToken   string `json:"next_token"`
}
