// models/urls.go

package models

type URLCollection struct {
	ID	uint   `json:"id" gorm:"primary_key"`
	ActualURL string
	ShortURL string
}

type URLCollectionInput struct {
	ActualURL string
	ShortURL string
}

type SuccessResponse struct {
	Code     int
	Message  string
	Response URLCollection
}
