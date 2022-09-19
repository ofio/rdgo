package rdgo

import (
	"time"

	"github.com/ofio/esign"
	"gopkg.in/guregu/null.v4"
)

func NewError() error {
	return nil
}

type ContractQuery struct {
	ExternalID string `json:"external_id"`
}

type PrimaryContact struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Contract struct {
	BoardItemContracts      *[]BoardItemContract   `json:"board_item_contracts"`
	Note                    string                 `json:"note"`
	AnnualizedValue         float64                `json:"annualized_value"`
	Business                Business               `json:"business"`
	FundingDepartment       Department             `json:"funding_department"`
	IncreasePercent         float64                `json:"increase_percent"`
	RenewalType             string                 `json:"renewal_type"`
	RenegotiationAlertDate  string                 `json:"renegotiation_alert_date"`
	RenewalNotificationDays int                    `json:"renewal_notification_days"`
	PaymentTerms            string                 `json:"payment_terms"`
	EffectiveDate           string                 `json:"effective_date"`
	EndDate                 string                 `json:"end_date"`
	Owner                   Owner                  `json:"owner"`
	PrimaryContact          PrimaryContact         `json:"primary_contact"`
	ManagingDepartment      Department             `json:"managing_department"`
	ContractDiscountTerms   []ContractDiscountTerm `json:"contract_discount_terms"`
	TotalValue              float64                `json:"total_value"`
	ContractStatus          ContractStatus         `json:"contract_status"`
	ContractCommodities     []ContractCommodities  `json:"contract_commodities"`
	ContractAttachments     []ContractAttachment   `json:"contract_attachments"`
	ID                      int                    `json:"id"`
	Name                    string                 `json:"name"`
	UUID                    string                 `json:"uuid"`
	CurrencyCode            string                 `json:"currency_code"`
	ApprovalRequests        []ApprovalRequest      `json:"approval_requests"`
	PaymentSchedule         string                 `json:"payment_schedule"`
	SignedDate              time.Time              `json:"signed_date"`
}

type Invoice struct {
	AccountNumber       string         `json:"account_number"`
	Amount              float64        `json:"amount"`
	ApprovalStatus      string         `json:"approval_status"`
	ApprovedAt          time.Time      `json:"approved_at"`
	BankAccountName     string         `json:"bank_account_name"`
	BankName            string         `json:"bank_name"`
	BusinessAddress     string         `json:"business_address"`
	BusinessID          int            `json:"business_id"`
	BusinessName        string         `json:"business_name"`
	BusinessPhone       string         `json:"business_phone"`
	ChecksPayableTo     string         `json:"checks_payable_to"`
	CreatedAt           time.Time      `json:"created_at"`
	CreatedBy           string         `json:"created_by"`
	CreatedByInstanceID int            `json:"created_by_instance_id"`
	CreatedByUserJsonb  UserJsonb      `json:"created_by_user_jsonb"`
	CurrencyCode        string         `json:"currency_code"`
	ID                  int            `json:"id"`
	ImportData          ImportData     `json:"import_data"`
	ImportStatus        string         `json:"import_status"`
	InstanceID          int            `json:"instance_id"`
	InvoiceNumber       string         `json:"invoice_number"`
	PaidStatus          string         `json:"paid_status"`
	PoNumber            string         `json:"po_number"`
	RoutingNumber       string         `json:"routing_number"`
	TermsAndConditions  string         `json:"terms_and_conditions"`
	Type                string         `json:"type"`
	UpdatedAt           time.Time      `json:"updated_at"`
	UpdatedBy           string         `json:"updated_by"`
	UpdatedByInstanceID int            `json:"updated_by_instance_id"`
	UpdatedByUserJsonb  UserJsonb      `json:"updated_by_user_jsonb"`
	UserEmail           string         `json:"user_email"`
	UUID                string         `json:"uuid"`
	ValidationStatus    string         `json:"validation_status"`
	Instance            Instance       `json:"instance"`
	InvoiceLines        []InvoiceLines `json:"invoice_lines"`
	Business            Business       `json:"business"`
}
type ImportData struct {
}

type UserJsonb struct {
	Email      string `json:"email"`
	UUID       string `json:"uuid"`
	InstanceID int    `json:"instance_id"`
	Name       string `json:"name"`
	Typename   string `json:"__typename"`
	ID         string `json:"id"`
}
type Instance struct {
	ID               int                `json:"id"`
	InstanceSettings []InstanceSettings `json:"instance_settings"`
	Business         Business           `json:"business"`
}
type InstanceSettings struct {
	BrandingLogoUUID string `json:"branding_logo_uuid"`
}

type ApprovalRequestCreator struct {
	Id              string   `json:"id"`
	InstanceID      int      `json:"instance_id"`
	Instance        Instance `json:"instance"`
	UserPreferences struct {
		DocusignRefreshToken  string          `json:"docusign_refresh_token"`
		DocusignUserInfo      *esign.UserInfo `json:"docusign_user_info"`
		AdobeSignRefreshToken string          `json:"adobe_sign_refresh_token"`
		AdobeSignApi          string          `json:"adobe_sign_api"`
	} `json:"user_preference"`
}

type Department struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type ApprovalRequestContract struct {
	FundingDepartment   Department           `json:"funding_department"`
	UUID                string               `json:"uuid"`
	ID                  int                  `json:"id"`
	Name                string               `json:"name"`
	ContractAttachments []ContractAttachment `json:"contract_attachments"`
	IncreasePercent     float64              `json:"increase_percent"`
	Business            struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"business"`
	RenewalType             string      `json:"renewal_type"`
	RenegotiationAlertDate  interface{} `json:"renegotiation_alert_date"`
	RenewalNotificationDays int         `json:"renewal_notification_days"`
	PaymentTerms            string      `json:"payment_terms"`
	EffectiveDate           string      `json:"effective_date"`
	EndDate                 string      `json:"end_date"`
	Owner                   struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"owner"`
	PrimaryContact struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"primary_contact"`
	ManagingDepartment    Department             `json:"managing_department"`
	Note                  string                 `json:"note"`
	ContractDiscountTerms []ContractDiscountTerm `json:"contract_discount_terms"`
	TotalValue            float64                `json:"total_value"`
	AnnualizedValue       float64                `json:"annualized_value"`
	ContractStatus        struct {
		ID int `json:"id"`
	} `json:"contract_status"`
	ContractCommodities []struct {
		ID int `json:"id"`
	} `json:"contract_commodities"`

	BoardItemContracts *[]BoardItemContract `json:"board_item_contracts"`
}

type BoardItemContract struct {
	BoardItem BoardItem `json:"board_item"`
}

type BoardItem struct {
	Board Board                  `json:"board"`
	Data  map[string]interface{} `json:"data"`
	ID    string                 `json:"id"`
	UUID  string                 `json:"uuid"`
}

type Board struct {
	BoardDef BoardDef `json:"board_def"`
}

type BoardDef struct {
	DndField     string       `json:"dndField"`
	ColumnDefs   []ColumnDef  `json:"columnDefs"`
	StatusBarDef StatusBarDef `json:"statusBarDef"`
}

type Aggregation struct {
	Aggr          string `json:"aggr"`
	Type          string `json:"type"`
	Field         string `json:"field"`
	Label         string `json:"label"`
	FormatOptions struct {
		Style    string `json:"style"`
		Currency string `json:"currency"`
		Notation string `json:"notation"`
	} `json:"formatOptions"`
}

type StatusBarDef struct {
	Aggregations []Aggregation `json:"aggregations"`
}

type ContractStatus struct {
	ID int `json:"id"`
}
type ContractCommodities struct {
	ID int `json:"id"`
}

type ContractDiscountTerm struct {
	DiscountDays       int `json:"discount_days"`
	DiscountPercentage int `json:"discount_percentage"`
	ID                 int `json:"id"`
}

type ColumnDef struct {
	Type          string `json:"type"`
	Field         string `json:"field"`
	AggFunc       string `json:"aggFunc,omitempty"`
	HeaderName    string `json:"headerName"`
	ChartDataType string `json:"chartDataType,omitempty"`
	Width         int    `json:"width,omitempty"`
	Choices       []struct {
		Text       string `json:"text"`
		Background string `json:"background"`
	} `json:"choices,omitempty"`
	MaxWidth        int     `json:"maxWidth,omitempty"`
	MinWidth        int     `json:"minWidth,omitempty"`
	ColumnGroupShow string  `json:"columnGroupShow,omitempty"`
	Children        []Child `json:"children,omitempty"`
}

type Child struct {
	Type            string `json:"type"`
	Field           string `json:"field"`
	Width           int    `json:"width"`
	MaxWidth        int    `json:"maxWidth,omitempty"`
	HeaderName      string `json:"headerName"`
	ChartDataType   string `json:"chartDataType"`
	ColumnGroupShow string `json:"columnGroupShow,omitempty"`
}

type ContractAttachment struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	IsDeleted bool   `json:"is_deleted"`
}

type InstanceSetting struct {
	BrandingLogoUUID string `json:"branding_logo_uuid"`
}
type ApprovalRequest struct {
	UUID       string `json:"uuid"`
	ObjectUUID string `json:"object_uuid"`
	Id         int    `json:"id"`
	Instance   struct {
		ID               int               `json:"id"`
		InstanceSettings []InstanceSetting `json:"instance_settings"`
		Business         Business          `json:"business"`
	} `json:"instance"`
	Message                    string                      `json:"message"`
	Status                     string                      `json:"status"`
	ExternalID                 string                      `json:"external_id"`
	ServiceName                string                      `json:"service_name"`
	ApprovalRequestAttachments []ApprovalRequestAttachment `json:"approval_request_attachments"`
	Creator                    ApprovalRequestCreator      `json:"creator"`
	Contract                   Contract                    `json:"contract"`
	CreatedBy                  string                      `json:"created_by"`
	InstanceID                 int                         `json:"instance_id"`
	AttachmentRank             []string                    `json:"attachment_rank"`
	AdminAutomation            struct {
		AuditTrailUUID string `json:"audit_trail"`
	} `json:"admin_automation"`
	CoverPage bool `json:"cover_page"`
}
type Business struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Address       string `json:"address"`
	City          string `json:"city"`
	StateProvince string `json:"state_province"`
	PostalCode    string `json:"postal_code"`
	Country       string `json:"country"`
}
type Approver struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	ID    string `json:"id"`
}

type ApprovalRequestApprover struct {
	Sequence  int       `json:"sequence"`
	IsSigner  bool      `json:"is_signer"`
	Approver  Approver  `json:"approver"`
	Status    string    `json:"status"`
	UpdatedAt time.Time `json:"updated_at"`
	Message   string    `json:"message"`
}

type ApprovalRequestAttachment struct {
	Attachment          Attachment                `json:"contract_attachment"`
	AttachmentApprovers []ApprovalRequestApprover `json:"approval_request_attachment_approvers"`
}

type Owner struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Contractupdateinput struct {
	ContractStatus string    `json:"status"`
	SignedDate     null.Time `json:"signed_date"`
}

type Data struct {
	InsertAttachment struct {
		Returning []struct {
			Id   int    `json:"id"`
			Uuid string `json:"uuid"`
		} `json:"returning"`
	} `json:"insert_attachment"`
	InsertContractAttachment struct {
		Returning []struct {
			Id   int    `json:"id"`
			UUID string `json:"uuid"`
		} `json:"returning"`
	} `json:"insert_contract_attachment"`
	InsertPoHeaderAttachment struct {
		Returning []struct {
			Id   int    `json:"id"`
			UUID string `json:"uuid"`
		} `json:"returning"`
	} `json:"insert_po_header_attachment"`
	InsertInvoiceAttachment struct {
		Returning []struct {
			Id   int    `json:"id"`
			UUID string `json:"uuid"`
		} `json:"returning"`
	} `json:"insert_invoice_attachment"`
	Integration []struct {
		ModuleName string `json:"module_name"`
		Type       string `json:"type"`
		Contract   struct {
			ID         int    `json:"id"`
			UUID       string `json:"uuid"`
			Name       string `json:"name"`
			BusinessID int    `json:"business_id"`
		} `json:"contract"`
	} `json:"integration"`
	BusinessIntegration []struct {
		Type string            `json:"type"`
		Data map[string]string `json:"data"`
	} `json:"business_integration"`
	ApprovalRequest            []ApprovalRequest            `json:"approval_request"`
	EmailTemplate              []EmailTemplate              `json:"email_template"`
	Instance                   []Instance                   `json:"instance"`
	User                       []User                       `json:"user"`
	DocumentGeneratorCondition []DocumentGeneratorCondition `json:"document_generator_condition"`
	PoHeader                   []PoHeader                   `json:"po_header"`
	Invoice                    []Invoice                    `json:"invoice"`
	InvoiceAttachment          []Attachment                 `json:"invoice_attachment"`
	ContractAttachment         []Attachment                 `json:"contract_attachment"`
	PoHeaderAttachment         []Attachment                 `json:"po_header_attachment"`
}

type PoHeader struct {
	CurrencyCode          string           `json:"currency_code"`
	PoNumber              string           `json:"po_number"`
	PaymentTerms          string           `json:"payment_terms"`
	Status                string           `json:"status"`
	InvoicingInstructions string           `json:"invoicing_instructions"`
	TermsAndConditions    string           `json:"terms_and_conditions"`
	Notes                 string           `json:"notes"`
	SoldToEntity          string           `json:"sold_to_entity"`
	Instance              Instance         `json:"instance"`
	DepartmentID          int              `json:"department_id"`
	Rev                   int              `json:"rev"`
	BuyerJsonb            BuyerJsonb       `json:"buyer_jsonb"`
	RequesterJsonb        RequesterJsonb   `json:"requester_jsonb"`
	SupplierContact       SupplierContact  `json:"supplier_contact"`
	UpdatedAt             time.Time        `json:"updated_at"`
	BusinessBillTo        BusinessBillTo   `json:"businessBillTo"`
	BusinessShipTo        BusinessShipTo   `json:"businessShipTo"`
	BusinessSupplier      BusinessSupplier `json:"businessSupplier"`
	PoLines               []PoLines        `json:"po_lines"`
	Contract              Contract         `json:"contract"`
	Department            Department       `json:"department"`
	ID                    int              `json:"id"`
	CreatedBy             string           `json:"created_by"`
}

type BusinessBillTo struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	City          string `json:"city"`
	StateProvince string `json:"state_province"`
	PostalCode    string `json:"postal_code"`
	Country       string `json:"country"`
}
type BusinessShipTo struct {
	Name                  string `json:"name"`
	ShippingAddress       string `json:"shipping_address"`
	ShippingCity          string `json:"shipping_city"`
	ShippingStateProvince string `json:"shipping_state_province"`
	ShippingPostalCode    string `json:"shipping_postal_code"`
	ShippingCountry       string `json:"shipping_country"`
}
type BusinessSupplier struct {
	Name          string `json:"name"`
	Address       string `json:"address"`
	City          string `json:"city"`
	StateProvince string `json:"state_province"`
	PostalCode    string `json:"postal_code"`
	Country       string `json:"country"`
}

type PoLines struct {
	ID              int       `json:"id"`
	LineNumber      int       `json:"line_number"`
	Commodity       Commodity `json:"commodity"`
	ItemDescription string    `json:"item_description"`
	Quantity        float64   `json:"quantity"`
	NetPricePerUnit float64   `json:"net_price_per_unit"`
	CommodityID     int       `json:"commodity_id"`
}

type InvoiceLines struct {
	ID              int     `json:"id"`
	LineNumber      int     `json:"line_number"`
	ItemDescription string  `json:"item_description"`
	Quantity        float64 `json:"quantity"`
	UomCode         string  `json:"uom_code"`
	ItemCode        string  `json:"item_code"`
	UnitPrice       float64 `json:"unit_price"`
	LineAmount      float64 `json:"line_amount"`
}

type Commodity struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type SupplierContact struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type BuyerJsonb struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	UUID       string `json:"uuid"`
	Email      string `json:"email"`
	Typename   string `json:"__typename"`
	InstanceID int    `json:"instance_id"`
}

type RequesterJsonb struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	UUID       string `json:"uuid"`
	Email      string `json:"email"`
	Typename   string `json:"__typename"`
	InstanceID int    `json:"instance_id"`
}

type Builder struct {
	Fields    []interface{} `json:"fields"`
	Operation string        `json:"operation"`
}
type DocumentGeneratorCondition struct {
	ObjectType     string   `json:"object_type"`
	Builder        Builder  `json:"builder"`
	ConditionGql   string   `json:"condition_gql"`
	QueryVariables []string `json:"query_variables"`
	CreatedBy      string   `json:"created_by"`
	UpdatedBy      string   `json:"updated_by"`
}

type Role struct {
	Name string `json:"name"`
}
type User struct {
	Role Role `json:"role"`
}

type Hasuraerror struct {
	Extentions Hasuraerrorext `json:"extentions"`
	Message    string         `json:"message"`
}

type Hasuraerrorext struct {
	Path string `json:"path"`
	Code string `json:"code"`
}

type Responsedata struct {
	Data   Data          `json:"data"`
	Errors []Hasuraerror `json:"errors"`
}
type Attachment struct {
	ID         int    `json:"id"`
	UUID       string `json:"uuid"`
	ReadSecret string `json:"read_secret"`
	Generation int64  `json:"generation"`
	Name       string `json:"name"`
	Mime       string `json:"mime_type"`
}

type EmailTemplate struct {
	ID       int    `json:"id"`
	Template string `json:"template"`
}

type EmailTemplateRequestBody struct {
	EmailTemplate  int                               `json:"email_template,omitempty"`
	QueryVariables map[string]interface{}            `json:"query_variables,omitempty"`
	WhereExp       map[string]map[string]interface{} `json:"where_exp,omitempty"`
}
