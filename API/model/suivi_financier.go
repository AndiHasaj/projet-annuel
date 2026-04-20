package model

type SuiviFinancier struct {
	ChiffreAffaire float32 `json:"chiffre_affaire"`
	Charges        float32 `json:"charges"`
	Benefice       float32 `json:"benefice"`
}
