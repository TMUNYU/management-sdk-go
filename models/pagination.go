package models

type Pagination struct {
	ContinuationToken string `json:"continuation_token"`
	NextPage          string `json:"next_page"`
}