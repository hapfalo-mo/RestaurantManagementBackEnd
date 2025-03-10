package models

type PagingRequest struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}
