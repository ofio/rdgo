package main

type metadata struct {
	userid string
	email  string
}

type urls struct {
	SenderViewURL string `json:"senderviewurl,omitempty"`
	ConsentURL    string `json:"consentURL,omitempty"`
	TemplateURL   string `json:"templateURL,omitempty"`
}

type envSummary struct {
	EnvelopeID     string `json:"envelopeid"`
	EnvelopeStatus string `json:"envelopestatus,omitempty"`
	URLS           urls   `json:"urls,omitempty"`
}
type templateSummary struct {
	TemplateID string `json:"templateid"`
	URLS       urls   `json:"urls,omitempty"`
}

type envelopesStatuses struct {
	ImpersonatedUserGUID string                   `json:"impersonated_user_guid"`
	EnvelopeIdsRequest   model.EnvelopeIdsRequest `json:"envelopestatusrequest"`
	BusinessID           int                      `json:"business_id"`
}

type emailname struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type requestFile struct {
	Secret     string `json:"secret"`
	UUID       string `json:"uuid"`
	Generation int64  `json:"generation,omitempty"`
	Name       string `json:"name"`
}

type requestBody struct {
	ApprovalRequest *int   `json:"request,omitempty"`
	Code            string `json:"code,omitempty"`
}

type redirectEmbed struct {
	Redirect string `json:"redirect"`
	Embed    string `json:"embed"`
	Login    bool   `json:"login"`
	Error    string `json:"error"`
}

type unisuiteRequestBody struct {
	EmailTemplate int                               `json:"email_template,omitempty"`
	ObjectUuid    string                            `json:"object_uuid,omitempty"`
	WhereExp      map[string]map[string]interface{} `json:"where_exp,omitempty"`
}

type approvalRequestUpdate struct {
	ExternalId string `json:"external_id,omitempty"`
	Status     string `json:"status,omitempty"`
}

type MyError struct {
	msg string
}

type urlredirect struct {
	Redirect string `json:"redirect"`
}

type approvalRequestInsertExternalLogInput struct {
	Log               esign.EnvelopeStatusXML `json:"log,omitempty"`
	ApprovalRequestId int                     `json:"approval_request_id,omitempty"`
	CreatedBy         string                  `json:"created_by,omitempty"`
	UpdatedBy         string                  `json:"updated_by,omitempty"`
	InstanceId        int                     `json:"instance_id,omitempty"`
}

type internalStorage struct {
	Id   int    `json:"id"`
	UUID string `json:"uuid"`
	Gen  int64  `json:"gen"`
}

type contractAttachmentMutation struct {
	ContractID   int    `json:"contract_id"`
	AttachmentID int    `json:"attachment_id"`
	InstanceID   int    `json:"instance_id"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
}

type options struct {
	EnvelopeInfo esign.EnvelopeStatusXML `xml:"EnvelopeStatus" json:"envelopeStatus,omitempty"`
}

type integrationinsertinput struct {
	UUID            string  `json:"object_uuid"`
	CreatedBy       string  `json:"created_by"`
	UpdatedBy       string  `json:"updated_by"`
	IntegrationType string  `json:"type"`
	ModuleName      string  `json:"module_name"`
	ExternalID      string  `json:"external_id"`
	Status          string  `json:"status"`
	InstanceID      int64   `json:"instance_id"`
	Options         options `json:"options"`
}
type integrationupdateinput struct {
	ExternalID string  `json:"external_id,omitempty"`
	Status     string  `json:"status,omitempty"`
	Options    options `json:"options,omitempty"`
}

type contractupdateinput struct {
	ContractStatus string `json:"status"`
}

type Data struct {
	Attachment struct {
		Returning []struct {
			Id   int    `json:"id"`
			Uuid string `json:"uuid"`
		} `json:"returning"`
	} `json:"insert_attachment"`
	ContractAttachment struct {
		Returning []struct {
			Id int `json:"id"`
		} `json:"returning"`
	} `json:"insert_contract_attachment"`
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
	ApprovalRequestQueryResult []ApprovalRequest `json:"approval_request"`
	EmailTemplate              []struct {
		ID int `json:"id"`
	} `json:"email_template"`
}

type hasuraerror struct {
	Extentions hasuraerrorext `json:"extentions"`
	Message    string         `json:"message"`
}

type hasuraerrorext struct {
	Path string `json:"path"`
	Code string `json:"code"`
}

type responsedata struct {
	Data   Data          `json:"data"`
	Errors []hasuraerror `json:"errors"`
}

type integrationQuery struct {
	ExternalID string `json:"external_id"`
}

type UserInfo struct {
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	Email      string `json:"email,omitempty"`
	JobTitle   string `json:"jobTitle,omitempty"`
	MiddleName string `json:"middleName,omitempty"`
}

type newUsersInfo struct {
	NewUsers []UserInfo `json:"userinfo"`
}

type businessIntegrationQuery struct {
	BusinessID int    `json:"business_id"`
	Data       string `json:"data,omitempty"`
	Type       string `json:"type,omitempty"`
}

type envBusiness struct {
	ExternalID string `json:"external_id"`
}

type Items struct {
	UserEmail       string `json:"UserEmail"`
	Action          string `json:"Action"`
	MoreInfo        string `json:"MoreInfo"`
	Comment         string `json:"Comment"`
	User            string `json:"User"`
	RelatedDocument string `json:"RelatedDocument"`
	CreatedDate     string `json:"CreatedDate"`
}

type HistoryItems struct {
	Items    []Items `json:"Items"`
	Href     string  `json:"Href"`
	Offset   int     `json:"Offset"`
	Limit    int     `json:"Limit"`
	First    string  `json:"First"`
	Previous string  `json:"Previous"`
	Next     string  `json:"Next"`
	Last     string  `json:"Last"`
	Total    int     `json:"Total"`
}
type AccessLevel struct {
	See       bool `json:"See"`
	Read      bool `json:"Read"`
	Write     bool `json:"Write"`
	Move      bool `json:"Move"`
	Create    bool `json:"Create"`
	SetAccess bool `json:"SetAccess"`
}

type EosParentInfo struct {
	Name       string `json:"Name"`
	Path       string `json:"Path"`
	ObjectID   string `json:"ObjectId"`
	ObjectType string `json:"ObjectType"`
	Folder     string `json:"Folder"`
}
type Lock struct {
	IsLocked      bool   `json:"IsLocked"`
	LockDate      string `json:"LockDate"`
	Type          string `json:"Type"`
	Comment       string `json:"Comment"`
	LockOwner     string `json:"LockOwner"`
	CheckInHref   string `json:"CheckInHref"`
	SignatureHref string `json:"SignatureHref"`
	Href          string `json:"Href"`
}

type Version struct {
	Items    []string `json:"Items"`
	Href     string   `json:"Href"`
	Offset   int      `json:"Offset"`
	Limit    int      `json:"Limit"`
	First    string   `json:"First"`
	Previous string   `json:"Previous"`
	Next     string   `json:"Next"`
	Last     string   `json:"Last"`
	Total    int      `json:"Total"`
}

type ShareLinks struct {
	Items    []string `json:"Items"`
	Href     string   `json:"Href"`
	Offset   int      `json:"Offset"`
	Limit    int      `json:"Limit"`
	First    string   `json:"First"`
	Previous string   `json:"Previous"`
	Next     string   `json:"Next"`
	Last     string   `json:"Last"`
	Total    int      `json:"Total"`
}
type ProcessUID struct {
}

type AssignedTo struct {
	Email    string `json:"Email"`
	FullName string `json:"FullName"`
	User     string `json:"User"`
}

type AssignedBy struct {
	Email    string `json:"Email"`
	FullName string `json:"FullName"`
	User     string `json:"User"`
}
type ActionBy struct {
	Email    string `json:"Email"`
	FullName string `json:"FullName"`
	User     string `json:"User"`
}

type UserItems struct {
	AssignedTo    AssignedTo `json:"AssignedTo"`
	WorkflowQueue string     `json:"WorkflowQueue"`
	Action        string     `json:"Action"`
	Output        string     `json:"Output"`
	Comments      string     `json:"Comments"`
	AssignedBy    AssignedBy `json:"AssignedBy"`
	ActionBy      ActionBy   `json:"ActionBy"`
	StartDate     string     `json:"StartDate"`
	UpdatedDate   string     `json:"UpdatedDate"`
	EndDate       string     `json:"EndDate"`
}

type UserActions struct {
	Items    []UserItems `json:"Items"`
	Href     string      `json:"Href"`
	Offset   int         `json:"Offset"`
	Limit    int         `json:"Limit"`
	First    string      `json:"First"`
	Previous string      `json:"Previous"`
	Next     string      `json:"Next"`
	Last     string      `json:"Last"`
	Total    int         `json:"Total"`
}

type DocumentTrackingItems struct {
	Name        string      `json:"Name"`
	TypeName    string      `json:"TypeName"`
	Status      string      `json:"Status"`
	Output      string      `json:"Output"`
	StageName   string      `json:"StageName"`
	StartDate   string      `json:"StartDate"`
	EndDate     string      `json:"EndDate"`
	DueDate     string      `json:"DueDate"`
	ProcessUID  ProcessUID  `json:"ProcessUid"`
	ProcessName string      `json:"ProcessName"`
	Document    string      `json:"Document"`
	UserActions UserActions `json:"UserActions"`
	Href        string      `json:"Href"`
}

type DocumentProcessTrackingActivities struct {
	Items    DocumentTrackingItems `json:"Items"`
	Href     string                `json:"Href"`
	Offset   int                   `json:"Offset"`
	Limit    int                   `json:"Limit"`
	First    string                `json:"First"`
	Previous string                `json:"Previous"`
	Next     string                `json:"Next"`
	Last     string                `json:"Last"`
	Total    int                   `json:"Total"`
}

type DocumentReminders struct {
	Items    []string `json:"Items"`
	Href     string   `json:"Href"`
	Offset   int      `json:"Offset"`
	Limit    int      `json:"Limit"`
	First    string   `json:"First"`
	Previous string   `json:"Previous"`
	Next     string   `json:"Next"`
	Last     string   `json:"Last"`
	Total    int      `json:"Total"`
}

type RelatedDocuments struct {
	Items    []string `json:"Items"`
	Href     string   `json:"Href"`
	Offset   int      `json:"Offset"`
	Limit    int      `json:"Limit"`
	First    string   `json:"First"`
	Previous string   `json:"Previous"`
	Next     string   `json:"Next"`
	Last     string   `json:"Last"`
	Total    int      `json:"Total"`
}

type WorkItems struct {
	Items    []string `json:"Items"`
	Href     string   `json:"Href"`
	Offset   int      `json:"Offset"`
	Limit    int      `json:"Limit"`
	First    string   `json:"First"`
	Previous string   `json:"Previous"`
	Next     string   `json:"Next"`
	Last     string   `json:"Last"`
	Total    int      `json:"Total"`
}

type documentResponse struct {
	Name                              string                            `json:"Name"`
	CreatedDate                       string                            `json:"CreatedDate"`
	CreatedBy                         string                            `json:"CreatedBy"`
	UpdatedDate                       string                            `json:"UpdatedDate"`
	UpdatedBy                         string                            `json:"UpdatedBy"`
	Description                       string                            `json:"Description"`
	ParentFolder                      string                            `json:"ParentFolder"`
	Path                              string                            `json:"Path"`
	HistoryItems                      HistoryItems                      `json:"HistoryItems"`
	AttributeGroups                   string                            `json:"AttributeGroups"`
	AccessLevel                       AccessLevel                       `json:"AccessLevel"`
	PageCount                         int                               `json:"PageCount"`
	EosParentInfo                     EosParentInfo                     `json:"EosParentInfo"`
	Lock                              Lock                              `json:"Lock"`
	Version                           Version                           `json:"Version"`
	PreviewURL                        string                            `json:"PreviewUrl"`
	Versions                          Versions                          `json:"Versions"`
	ShareLinks                        ShareLinks                        `json:"ShareLinks"`
	DocumentProcessTrackingActivities DocumentProcessTrackingActivities `json:"DocumentProcessTrackingActivities"`
	DocumentReminders                 DocumentReminders                 `json:"DocumentReminders"`
	RelatedDocuments                  RelatedDocuments                  `json:"RelatedDocuments"`
	WorkItems                         WorkItems                         `json:"WorkItems"`
	DownloadDocumentHref              string                            `json:"DownloadDocumentHref"`
	NativeFileSize                    int                               `json:"NativeFileSize"`
	PdfFileSize                       int                               `json:"PdfFileSize"`
	ContentCreatedDate                string                            `json:"ContentCreatedDate"`
	Href                              string                            `json:"Href"`
}

type workflowRequet struct {
	Name              string            `json:"Name"`
	StartDate         string            `json:"StartDate"`
	EndDate           string            `json:"EndDate"`
	Status            string            `json:"Status"`
	Info              string            `json:"Info"`
	Params            string            `json:"Params"`
	WorkflowDocuments WorkflowDocuments `json:"WorkflowDocuments"`
	Href              string            `json:"Href"`
}

type WorkflowDocuments struct {
	Items    []string `json:"Items"`
	Href     string   `json:"Href"`
	Offset   int      `json:"Offset"`
	Limit    int      `json:"Limit"`
	First    string   `json:"First"`
	Previous string   `json:"Previous"`
	Next     string   `json:"Next"`
	Last     string   `json:"Last"`
	Total    int      `json:"Total"`
}

//To create an EOS folder, supply the EosInfo parameters. To create a regular folder, supply the Name and ParentFolder properties.
type folderCreationResponse struct {
	Name               string        `json:"Name"`
	CreatedDate        string        `json:"CreatedDate"`
	CreatedBy          string        `json:"CreatedBy"`
	UpdatedDate        string        `json:"UpdatedDate"`
	UpdatedBy          string        `json:"UpdatedBy"`
	Description        string        `json:"Description"`
	ParentFolder       string        `json:"ParentFolder"`
	BrowseDocumentsURL string        `json:"BrowseDocumentsUrl"`
	AccessLevel        AccessLevel   `json:"AccessLevel"`
	Documents          Documents     `json:"Documents"`
	Folders            Folders       `json:"Folders"`
	Path               string        `json:"Path"`
	AttributeGroups    string        `json:"AttributeGroups"`
	EosInfo            EosInfo       `json:"EosInfo"`
	EosParentInfo      EosParentInfo `json:"EosParentInfo"`
	ShareLinks         ShareLinks    `json:"ShareLinks"`
	Security           Security      `json:"Security"`
	CreateDocumentHref string        `json:"CreateDocumentHref"`
	Href               string        `json:"Href"`
}

type Folders struct {
	Items    []string `json:"Items"`
	Href     string   `json:"Href"`
	Offset   int      `json:"Offset"`
	Limit    int      `json:"Limit"`
	First    string   `json:"First"`
	Previous string   `json:"Previous"`
	Next     string   `json:"Next"`
	Last     string   `json:"Last"`
	Total    int      `json:"Total"`
}

type EosInfo struct {
	Name       string `json:"Name"`
	Path       string `json:"Path"`
	ObjectID   string `json:"ObjectId"`
	ObjectType string `json:"ObjectType"`
	Folder     string `json:"Folder"`
}

type Versions struct {
	Items    []string `json:"Items"`
	Href     string   `json:"Href"`
	Offset   int      `json:"Offset"`
	Limit    int      `json:"Limit"`
	First    string   `json:"First"`
	Previous string   `json:"Previous"`
	Next     string   `json:"Next"`
	Last     string   `json:"Last"`
	Total    int      `json:"Total"`
}
type Roles struct {
	Item       string `json:"Item"`
	AccessType string `json:"AccessType"`
}
type Groups struct {
	Item       string `json:"Item"`
	AccessType string `json:"AccessType"`
}

type Users struct {
	Item       string `json:"Item"`
	AccessType string `json:"AccessType"`
}
type Security struct {
	Roles  []Roles  `json:"Roles"`
	Groups []Groups `json:"Groups"`
	Users  []Users  `json:"Users"`
}

type Documents struct {
	Items    []string `json:"Items"`
	Href     string   `json:"Href"`
	Offset   int      `json:"Offset"`
	Limit    int      `json:"Limit"`
	First    string   `json:"First"`
	Previous string   `json:"Previous"`
	Next     string   `json:"Next"`
	Last     string   `json:"Last"`
	Total    int      `json:"Total"`
}

type relatedDocumentResponse []struct {
	Name                              string                            `json:"Name"`
	CreatedDate                       string                            `json:"CreatedDate"`
	CreatedBy                         string                            `json:"CreatedBy"`
	UpdatedDate                       string                            `json:"UpdatedDate"`
	UpdatedBy                         string                            `json:"UpdatedBy"`
	Description                       string                            `json:"Description"`
	ParentFolder                      string                            `json:"ParentFolder"`
	Path                              string                            `json:"Path"`
	HistoryItems                      HistoryItems                      `json:"HistoryItems"`
	AttributeGroups                   string                            `json:"AttributeGroups"`
	AccessLevel                       AccessLevel                       `json:"AccessLevel"`
	PageCount                         int                               `json:"PageCount"`
	EosParentInfo                     EosParentInfo                     `json:"EosParentInfo"`
	Lock                              Lock                              `json:"Lock"`
	Version                           string                            `json:"Version"`
	PreviewURL                        string                            `json:"PreviewUrl"`
	Versions                          Versions                          `json:"Versions"`
	ShareLinks                        ShareLinks                        `json:"ShareLinks"`
	DocumentProcessTrackingActivities DocumentProcessTrackingActivities `json:"DocumentProcessTrackingActivities"`
	DocumentReminders                 struct {
		Items    []string `json:"Items"`
		Href     string   `json:"Href"`
		Offset   int      `json:"Offset"`
		Limit    int      `json:"Limit"`
		First    string   `json:"First"`
		Previous string   `json:"Previous"`
		Next     string   `json:"Next"`
		Last     string   `json:"Last"`
		Total    int      `json:"Total"`
	} `json:"DocumentReminders"`
	RelatedDocuments struct {
		Items    []string `json:"Items"`
		Href     string   `json:"Href"`
		Offset   int      `json:"Offset"`
		Limit    int      `json:"Limit"`
		First    string   `json:"First"`
		Previous string   `json:"Previous"`
		Next     string   `json:"Next"`
		Last     string   `json:"Last"`
		Total    int      `json:"Total"`
	} `json:"RelatedDocuments"`
	WorkItems struct {
		Items    []string `json:"Items"`
		Href     string   `json:"Href"`
		Offset   int      `json:"Offset"`
		Limit    int      `json:"Limit"`
		First    string   `json:"First"`
		Previous string   `json:"Previous"`
		Next     string   `json:"Next"`
		Last     string   `json:"Last"`
		Total    int      `json:"Total"`
	} `json:"WorkItems"`
	DownloadDocumentHref string `json:"DownloadDocumentHref"`
	NativeFileSize       int    `json:"NativeFileSize"`
	PdfFileSize          int    `json:"PdfFileSize"`
	ContentCreatedDate   string `json:"ContentCreatedDate"`
	Href                 string `json:"Href"`
}

type returning struct {
	Id int `json:"id"`
}

type Approver struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

type ApprovalRequestApprover struct {
	Sequence int      `json:"sequence"`
	IsSigner bool     `json:"is_signer"`
	Approver Approver `json:"approver"`
}

type Attachment struct {
	UUID       string `json:"uuid"`
	ReadSecret string `json:"read_secret"`
	Generation int64  `json:"generation"`
	Name       string `json:"name"`
	Mime       string `json:"mime_type"`
}
type ApprovalRequestAttachment struct {
	Attachment          Attachment                `json:"contract_attachment"`
	AttachmentApprovers []ApprovalRequestApprover `json:"approval_request_attachment_approvers"`
}

type instance struct {
	InstanceSettings []instanceSettings `json:"instance_settings"`
}

type instanceSettings struct {
	BrandingLogoUUID string `json:"branding_logo_uuid"`
}

type ApprovalRequestCreator struct {
	Id              string   `json:"id"`
	InstanceID      int      `json:"instance_id"`
	Instance        instance `json:"instance"`
	UserPreferences struct {
		DocusignRefreshToken string          `json:"docusign_refresh_token"`
		DocusignUserInfo     *esign.UserInfo `json:"docusign_user_info"`
	} `json:"user_preference"`
}

type ApprovalRequestContract struct {
	FundingDepartment   string               `json:"funding_department"`
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
	ManagingDepartment    string                 `json:"managing_department"`
	Note                  string                 `json:"note"`
	ContractDiscountTerms []ContractDiscountTerm `json:"contract_discount_terms"`
	TotalValue            int                    `json:"total_value"`
	AnnualizedValue       int                    `json:"annualized_value"`
	ContractStatus        struct {
		ID int `json:"id"`
	} `json:"contract_status"`
	ContractCommodities []struct {
		ID int `json:"id"`
	} `json:"contract_commodities"`

	BoardItemContracts *[]boardItemContract `json:"board_item_contracts"`
}

type boardItemContract struct {
	BoardItem boardItem `json:"board_item"`
}
type boardItem struct {
	Board board                  `json:"board"`
	Data  map[string]interface{} `json:"data"`
	ID    string                 `json:"id"`
	UUID  string                 `json:"uuid"`
}
type board struct {
	BoardDef boardDef `json:"board_def"`
}
type boardDef struct {
	DndField     string       `json:"dndField"`
	ColumnDefs   []columnDef  `json:"columnDefs"`
	StatusBarDef statusBarDef `json:"statusBarDef"`
}

type statusBarDef struct {
	Aggregations []Aggregation `json:"aggregations"`
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

type columnDef struct {
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
	Children        []child `json:"children,omitempty"`
}
type child struct {
	Type            string `json:"type"`
	Field           string `json:"field"`
	Width           int    `json:"width"`
	MaxWidth        int    `json:"maxWidth,omitempty"`
	HeaderName      string `json:"headerName"`
	ChartDataType   string `json:"chartDataType"`
	ColumnGroupShow string `json:"columnGroupShow,omitempty"`
}

// type boardItemData struct {
// 	UUID       string                 `json:"uuid"`
// 	ItemData   map[string]interface{} `json:"data"`
// 	InstanceID int                    `json:"instance_id"`
// 	IsArchived bool                   `json:"is_archived"`
// 	CreatedBy  string                 `json:"created_by"`
// 	UpdatedAt  time.Time              `json:"updated_at"`
// 	CreatedAt  time.Time              `json:"created_at"`
// 	ID         string                 `json:"id"`
// 	BoardID    *int                   `json:"board_id"`
// 	UpdatedBy  string                 `json:"updated_by"`
// }
type ContractAttachment struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	IsDeleted bool   `json:"is_deleted"`
}

type instanceSetting struct {
	BrandingLogoUUID string `json:"branding_logo_uuid"`
}
type ApprovalRequest struct {
	Id       int `json:"id"`
	Instance struct {
		ID               int               `json:"id"`
		InstanceSettings []instanceSetting `json:"instance_settings"`
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
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Owner struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
type PrimaryContact struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Contract struct {
	BoardItemContracts      *[]boardItemContract   `json:"board_item_contracts"`
	Note                    string                 `json:"note"`
	AnnualizedValue         int                    `json:"annualized_value"`
	Business                Business               `json:"business"`
	FundingDepartment       string                 `json:"funding_department"`
	IncreasePercent         float64                `json:"increase_percent"`
	RenewalType             string                 `json:"renewal_type"`
	RenegotiationAlertDate  string                 `json:"renegotiation_alert_date"`
	RenewalNotificationDays int                    `json:"renewal_notification_days"`
	PaymentTerms            string                 `json:"payment_terms"`
	EffectiveDate           string                 `json:"effective_date"`
	EndDate                 string                 `json:"end_date"`
	Owner                   Owner                  `json:"owner"`
	PrimaryContact          PrimaryContact         `json:"primary_contact"`
	ManagingDepartment      string                 `json:"managing_department"`
	ContractDiscountTerms   []ContractDiscountTerm `json:"contract_discount_terms"`
	TotalValue              int                    `json:"total_value"`
	ContractStatus          ContractStatus         `json:"contract_status"`
	ContractCommodities     []ContractCommodities  `json:"contract_commodities"`
	ContractAttachments     []ContractAttachment   `json:"contract_attachments"`
	ID                      int                    `json:"id"`
	Name                    string                 `json:"name"`
	UUID                    string                 `json:"uuid"`
	CurrencyCode            string                 `json:"currency_code"`
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

type queryReq struct {
	ExternalID string `json:"external_id"`
}
