package models

type ListOptions struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type GetOptions struct {
	Id int `json:"id"`
}
