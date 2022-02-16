package entity

type HomePageResponse struct {
	Posts []Post `json:"posts,omitempty"`
}
