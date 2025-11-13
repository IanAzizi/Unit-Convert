package service

type Unit struct {
	From  string  `json:"from"`
	To    string  `json:"to"`
	Value float64 `json:"value"`
}

const (
	Gram      = 1.0
	Kilogram  = 1000.0
	Meter     = 1.0
	Kilometer = 1000.0
)

func Convert(u Unit) float64 {
	switch u.From {
	case "Gram":
		if u.To == "Kilogram" {
			return u.Value / 1000
		}
	case "Kilogram":
		if u.To == "Gram" {
			return u.Value * 1000
		}
	case "Meter":
		if u.To == "Kilometer" {
			return u.Value / 1000
		}
	case "Kilometer":
		if u.To == "Meter" {
			return u.Value * 1000
		}
	}
	return 0
}
