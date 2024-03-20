package entity

func NewApiDolarQuotation() *apiDolarQuotation {
	return &apiDolarQuotation{}
}

type apiDolarQuotation struct {
	Usdbrl struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func (e *apiDolarQuotation) GetUrl() string {
	return "https://economia.awesomeapi.com.br/json/last/USD-BRL"
}

func (e *apiDolarQuotation) GetBid() string {
	return e.Usdbrl.Bid
}
