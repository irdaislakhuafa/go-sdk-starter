package pagination

type (
	Pagination struct {
		Limit int `json:"limit" query:"limit"`
		Page  int `json:"page" query:"page"`
	}
)

func (p *Pagination) Parse() error {
	if p.Limit == 0 {
		p.Page = 0
		p.Limit = 15
	} else {
		p.Page = (p.Page * p.Limit)
	}

	return nil
}
