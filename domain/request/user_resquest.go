package request

type UpdateTierByUserIDRequest struct {
	ID int `json:"id"`
	Tier int `json:"tier_rank"`
}
