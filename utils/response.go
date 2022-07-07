package utils

type OkWithPage struct {
	List       interface{} `json:"list"`
	Pagination *Pagination `json:"pagination"`
}
