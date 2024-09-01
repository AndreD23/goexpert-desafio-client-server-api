package quotation

type Quotation struct {
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
}

type QuotationResponse struct {
	USD_BRL Quotation `json:"USDBRL"`
}

func ConvertToQuotationDB(qr *QuotationResponse) *Quotation {
	return &Quotation{
		Code:       qr.USD_BRL.Code,
		Codein:     qr.USD_BRL.Codein,
		Name:       qr.USD_BRL.Name,
		High:       qr.USD_BRL.High,
		Low:        qr.USD_BRL.Low,
		VarBid:     qr.USD_BRL.VarBid,
		PctChange:  qr.USD_BRL.PctChange,
		Bid:        qr.USD_BRL.Bid,
		Ask:        qr.USD_BRL.Ask,
		Timestamp:  qr.USD_BRL.Timestamp,
		CreateDate: qr.USD_BRL.CreateDate,
	}
}
