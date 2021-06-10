package dto

type RequestDto struct {
	Data DataRequestDto `json:"data"`
}

type DataRequestDto struct {
	ID             string               `json:"id"`
	Type           string               `json:"type"`
	OrganisationID string               `json:"organisation_id"`
	Attributes     AttributesRequestDto `json:"attributes"`
}

type AttributesRequestDto struct {
	Country                 string   `json:"country"`
	BaseCurrency            string   `json:"base_currency"`
	BankID                  string   `json:"bank_id"`
	BankIDCode              string   `json:"bank_id_code"`
	BIC                     string   `json:"bic"`
	Name                    []string `json:"name"`
	AlternativeName         []string `json:"alternative_names"`
	AccountClassification   string   `json:"account_classification"`
	JointAccount            bool     `json:"joint_account"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
	SecondaryIdentification string   `json:"secondary_identification"`
}
