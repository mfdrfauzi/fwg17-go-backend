package services

type PageInfo struct {
	Page      int `json:"page"`
	TotalPage int `json:"totalPage"`
	NextPage  int `json:"nextPage,omitempty"`
	PrevPage  int `json:"prevPage,omitempty"`
}

type ResponseList struct {
	Success  bool        `json:"success"`
	Message  string      `json:"message"`
	PageInfo PageInfo    `json:"pageInfo"`
	Results  interface{} `json:"results"`
}
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Results interface{} `json:"results"`
}

type ResponseOnly struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
