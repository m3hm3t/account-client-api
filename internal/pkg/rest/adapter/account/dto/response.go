package dto

type ResponseDto struct {
	Data DataResponseDto `json:"data"`
}

type DataResponseDto struct {
	ID             string                   `json:"id"`
	Type           string                   `json:"type"`
	OrganisationID string                   `json:"organisation_id"`
	Version        int                      `json:"version"`
	ModifiedOn     string                   `json:"modified_on"`
	CreatedOn      string                   `json:"created_on"`
	Attributes     AttributesResponseDto    `json:"attributes"`
	Relationships  RelationshipsResponseDto `json:"relationships"`
}

type AttributesResponseDto struct {
	Country                 string   `json:"country"`
	BaseCurrency            string   `json:"base_currency"`
	BankID                  string   `json:"bank_id"`
	BankIDCode              string   `json:"bank_id_code"`
	AccountNumber           string   `json:"account_number"`
	IBAN                    string   `json:"iban"`
	Name                    []string `json:"name"`
	AlternativeName         []string `json:"alternative_names"`
	AccountClassification   string   `json:"account_classification"`
	JointAccount            bool     `json:"joint_account"`
	AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
	SecondaryIdentification string   `json:"secondary_identification"`
	Status                  string   `json:"status"`
}

type RelationshipsResponseDto struct {
	AccountEvents AccountEventsResponseDto `json:"account_events"`
}

type AccountEventsResponseDto struct {
	Data []AccountEventsDataResponseDto `json:"data"`
}

type AccountEventsDataResponseDto struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}
