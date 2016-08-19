package paypal

type NotificationChannel string

const (
	SMS   NotificationChannel = "SMS"
	EMAIL NotificationChannel = "EMAIL"
)

type Language string

const (
	DA_DK Language = "da_DK"
	DE_DE Language = "de_DE"
	EN_AU Language = "en_AU"
	EN_GB Language = "en_GB"
	EN_US Language = "en_US"
	ES_ES Language = "es_ES"
	ES_XC Language = "es_XC"
	FR_CA Language = "fr_CA"
	FR_FR Language = "fr_FR"
	FR_XC Language = "fr_XC"
	HE_IL Language = "he_IL"
	ID_ID Language = "id_ID"
	IT_IT Language = "it_IT"
	JA_JP Language = "ja_JP"
	NL_NL Language = "nl_NL"
	NO_NO Language = "no_NO"
	PL_PL Language = "pl_PL"
	PT_BR Language = "pt_BR"
	PT_PT Language = "pt_PT"
	RU_RU Language = "ru_RU"
	SV_SE Language = "sv_SE"
	TH_TH Language = "th_TH"
	ZH_CN Language = "zh_CN"
	ZH_HK Language = "zh_HK"
	ZH_TW Language = "zh_TW"
	ZH_XC Language = "zh_XC"
)

type TermType string

const (
	DUE_ON_RECEIPT        TermType = "DUE_ON_RECEIPT"
	DUE_ON_DATE_SPECIFIED TermType = "DUE_ON_DATE_SPECIFIED"
	NET_10                TermType = "NET_10"
	NET_15                TermType = "NET_15"
	NET_30                TermType = "NET_30"
	NET_45                TermType = "NET_45"
	NET_60                TermType = "NET_60"
	NET_90                TermType = "NET_90"
	NO_DUE_DATE           TermType = "NO_DUE_DATE"
)

type (
	Phone struct {
		CountryCode    string `json:"country_code"`
		NationalNumber string `json:"national_number"`
	}

	MerchantInfo struct {
		Email               string   `json:"email"`
		FirstName           string   `json:"first_name"`
		LastName            string   `json:"last_name"`
		BusinessName        string   `json:"business_name"`
		Phone               *Phone   `json:"phone"`
		Fax                 *Phone   `json:"fax"`
		Website             string   `json:"website"`
		Address             *Address `json:"address"`
		TaxID               string   `json:"tax_id"`
		AdditionalInfoLabel string   `json:"additional_info_label"`
		AdditionalInfo      string   `json:"additional_info"`
	}

	BillingInfo struct {
		Email               string              `json:"email"`
		FirstName           string              `json:"first_name"`
		LastName            string              `json:"last_name"`
		BusinessName        string              `json:"business_name"`
		Address             *Address            `json:"address"`
		Language            Language            `json:"language"`
		AdditionalInfoLabel string              `json:"additional_info_label"`
		AdditionalInfo      string              `json:"additional_info"`
		NotificationChannel NotificationChannel `json:"notification_channel"`
		Phone               *Phone              `json:"phone"`
	}

	Participant struct {
		Email        string   `json:"email"`
		FirstName    string   `json:"first_name"`
		LastName     string   `json:"last_name"`
		BusinessName string   `json:"business_name"`
		Phone        *Phone   `json:"phone"`
		Fax          *Phone   `json:"fax"`
		Website      string   `json:"website"`
		Address      *Address `json:"address"`
	}

	ShippingInfo struct {
		FirstName    string   `json:"first_name"`
		LastName     string   `json:"last_name"`
		BusinessName string   `json:"business_name"`
		Address      *Address `json:"address"`
	}

	PaymentTerm struct {
		DueDate  string   `json:"due_date"` // The date when the invoice was enabled. The date format is yyyy-MM-dd z as defined in Internet Date/Time Format. (https://tools.ietf.org/html/rfc3339#section-5.6)
		TermType TermType `json:"term_type"`
	}

	Discount struct {
		Percent float64 `json:"percent"`
		Amount  *Amount `json:"amount"`
	}

	Tax struct {
		ID      string  `json:"id"`
		Name    string  `json:"name"`
		Percent float64 `json:"percent"`
		Amount  *Amount `json:"amount"`
	}

	ShippingCost struct {
		Amount *Amount `json:"amount"`
		Tax    *Tax    `json:"tax"`
	}

	CustomAmount struct {
		Amount *Amount `json:"amount"`
		Label  string  `json:"label"`
	}

	Currency struct {
		Currency string `json:"currency"`
		Value    string `json:"value"`
	}

	FileAttachments struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	Invoice struct {
		Number                     int                `json:"number"`
		TemplateID                 string             `json:"template_id"`
		MerchantInfo               *MerchantInfo      `json:"merchant_info"`
		BillingInfo                *BillingInfo       `json:"billing_info"`
		CCInfo                     []*Participant     `json:"cc_info"`
		ShippingInfo               *ShippingInfo      `json:"shipping_info"`
		Items                      []*Item            `json:"items"`
		InvoiceDate                string             `json:"invoice_date"` // The date when the invoice was enabled. The date format is yyyy-MM-dd z as defined in Internet Date/Time Format. (https://tools.ietf.org/html/rfc3339#section-5.6)
		PaymentTerm                *PaymentTerm       `json:"payment_term"`
		Reference                  string             `json:"reference"`
		Discount                   *Discount          `json:"discount"`
		ShippingCost               *ShippingCost      `json:"shipping_cost"`
		Custom                     *CustomAmount      `json:"custom"`
		AllowPartialPayment        bool               `json:"allow_partial_payment"`
		MinimumAmountDue           *Currency          `json:"minimum_amount_due"`
		TaxCalculatedAfterDiscount bool               `json:"tax_calculated_after_discount"`
		TaxInclusive               bool               `json:"tax_inclusive"`
		Terms                      string             `json:"terms"`
		Note                       string             `json:"note"`
		MerchantMemo               string             `json:"merchant_memo"`
		LogoURL                    string             `json:"logo_url"`
		Attachments                []*FileAttachments `json:"attachments"`
	}
)
