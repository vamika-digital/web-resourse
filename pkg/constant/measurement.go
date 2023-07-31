package constant

type Measurement string

const (
	Kg  Measurement = "KG"
	Gm  Measurement = "GM"
	Pcs Measurement = "PCS"
	Ltr Measurement = "LTR"
)

func (m Measurement) AsText() string {
	switch m {
	case Kg:
		return "Kg."
	case Gm:
		return "Gram"
	case Pcs:
		return "PCs"
	case Ltr:
		return "Ltr."
	default:
		return "Unknown"
	}
}

func (m Measurement) IsValid() bool {
	return m == Kg || m == Gm || m == Pcs || m == Ltr
}
