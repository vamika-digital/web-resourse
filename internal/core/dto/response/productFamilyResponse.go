package response

type ProductFamilyResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Unit        string `json:"unit"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
