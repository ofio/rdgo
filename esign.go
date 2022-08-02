package rdgo

import (
	"github.com/ofio/esign"
	"github.com/ofio/esign/v2/model"
)

type Metadata struct {
	Userid string
	Email  string
}

type Urls struct {
	SenderViewURL string `json:"senderviewurl,omitempty"`
	ConsentURL    string `json:"consentURL,omitempty"`
	TemplateURL   string `json:"templateURL,omitempty"`
}

type EnvSummary struct {
	EnvelopeID     string `json:"envelopeid"`
	EnvelopeStatus string `json:"envelopestatus,omitempty"`
	URLS           Urls   `json:"urls,omitempty"`
}
type TemplateSummary struct {
	TemplateID string `json:"templateid"`
	URLS       Urls   `json:"urls,omitempty"`
}

type EnvelopesStatuses struct {
	ImpersonatedUserGUID string                   `json:"impersonated_user_guid"`
	EnvelopeIdsRequest   model.EnvelopeIdsRequest `json:"envelopestatusrequest"`
	BusinessID           int                      `json:"business_id"`
}

type Emailname struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RequestFile struct {
	Secret     string `json:"secret"`
	UUID       string `json:"uuid"`
	Generation int64  `json:"generation,omitempty"`
	Name       string `json:"name"`
}

type RequestBody struct {
	ApprovalRequest *int   `json:"request,omitempty"`
	Code            string `json:"code,omitempty"`
	Api             string `json:"api,omitempty"`
}

type RedirectEmbed struct {
	Redirect string `json:"redirect"`
	Embed    string `json:"embed"`
	Login    bool   `json:"login"`
	Error    string `json:"error"`
}

type UnisuiteRequestBody struct {
	EmailTemplate int                               `json:"email_template,omitempty"`
	ObjectUuid    string                            `json:"object_uuid,omitempty"`
	WhereExp      map[string]map[string]interface{} `json:"where_exp,omitempty"`
}

type ApprovalRequestUpdate struct {
	ExternalId string `json:"external_id,omitempty"`
	Status     string `json:"status,omitempty"`
}

type MyError struct {
	Msg string
}

func (error *MyError) Error() string {
	return error.Msg
}

type Urlredirect struct {
	Redirect string `json:"redirect"`
}

type ApprovalRequestInsertExternalLogInput struct {
	Log               esign.EnvelopeStatusXML `json:"log,omitempty"`
	ApprovalRequestId int                     `json:"approval_request_id,omitempty"`
	CreatedBy         string                  `json:"created_by,omitempty"`
	UpdatedBy         string                  `json:"updated_by,omitempty"`
	InstanceId        int                     `json:"instance_id,omitempty"`
}

type InternalStorage struct {
	Id   int    `json:"id"`
	UUID string `json:"uuid"`
	Gen  int64  `json:"gen"`
}

type DocusignToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	APIBaseURL  string `json:"api_base_url"`
}

type ContractAttachmentMutation struct {
	ContractID   int    `json:"contract_id"`
	AttachmentID int    `json:"attachment_id"`
	InstanceID   int    `json:"instance_id"`
	CreatedBy    string `json:"created_by"`
	UpdatedBy    string `json:"updated_by"`
	Generation   int64  `json:"generation"`
	Mime         string `json:"mime_type"`
	Name         string `json:"name"`
}

type Integrationinsertinput struct {
	UUID            string  `json:"object_uuid"`
	CreatedBy       string  `json:"created_by"`
	UpdatedBy       string  `json:"updated_by"`
	IntegrationType string  `json:"type"`
	ModuleName      string  `json:"module_name"`
	ExternalID      string  `json:"external_id"`
	Status          string  `json:"status"`
	InstanceID      int64   `json:"instance_id"`
	Options         Options `json:"options"`
}
type Integrationupdateinput struct {
	ExternalID string  `json:"external_id,omitempty"`
	Status     string  `json:"status,omitempty"`
	Options    Options `json:"options,omitempty"`
}

type Options struct {
	EnvelopeInfo  esign.EnvelopeStatusXML `xml:"EnvelopeStatus" json:"envelopeStatus,omitempty"`
	AgreementInfo WebhookEvent            `json:"agreementinfo,omitempty"`
}

type IntegrationQuery struct {
	BusinessID int    `json:"business_id"`
	Data       string `json:"data,omitempty"`
	Type       string `json:"type,omitempty"`
}
type UserInfo struct {
	FirstName  string `json:"firstName,omitempty"`
	LastName   string `json:"lastName,omitempty"`
	Email      string `json:"email,omitempty"`
	JobTitle   string `json:"jobTitle,omitempty"`
	MiddleName string `json:"middleName,omitempty"`
}

type NewUsersInfo struct {
	NewUsers []UserInfo `json:"userinfo"`
}

type BusinessIntegrationQuery struct {
	BusinessID int    `json:"business_id"`
	Data       string `json:"data,omitempty"`
	Type       string `json:"type,omitempty"`
}

type EnvBusiness struct {
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

type DocumentResponse struct {
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

type WorkflowRequet struct {
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
type FolderCreationResponse struct {
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

type RelatedDocumentResponse []struct {
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

type Returning struct {
	Id int `json:"id"`
}

type QueryReq struct {
	ExternalID string `json:"external_id"`
}

type AccessCode struct {
	Code string `json:"code"`
}

type ApiErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type JsonArray struct {
	Headers [][]string
}

type AccessToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token,omitempty"`
	TokenType    string `json:"token_type"`
	Expires      int    `json:"expires_in"`
}
type Ccs struct {
	Email string
}
type FilesInfo struct {
	FileType                  string `json:"fileType"`
	TransientDocumentID       string `json:"transientDocumentId"`
	ChildAgreementsInfoFileID string `json:"childAgreementsInfoFileId"`
}

type ChildAgreementsInfo struct {
	FileInfo FilesInfo `json:"fileInfo"`
}
type MergeFieldInfo struct {
	DefaultValue string `json:"defaultValue"`
	FieldName    string `json:"fieldName"`
}

type AutomationInfo struct {
	AuditTrailUUID string `json:"audit_trail"`
}

type ApprovalRequestQuery struct {
	ID int `json:"id"`
}

type UserPrefQuery struct {
	Uid string `json:"user_id"`
}

type ApprovalRequestInsertExternalLogInputAdobesign struct {
	Log               *WebhookEvent `json:"log,omitempty"`
	ApprovalRequestId int           `json:"approval_request_id,omitempty"`
	CreatedBy         string        `json:"created_by,omitempty"`
	UpdatedBy         string        `json:"updated_by,omitempty"`
	InstanceId        int           `json:"instance_id,omitempty"`
}

type WebhookEvent struct {
	WebhookID             string `json:"webhookId"`
	WebhookName           string `json:"webhookName"`
	WebhookNotificationID string `json:"webhookNotificationId"`
	WebhookURLInfo        struct {
		URL string `json:"url"`
	} `json:"webhookUrlInfo"`
	WebhookScope                       string `json:"webhookScope"`
	WebhookNotificationApplicableUsers []struct {
		ID                string `json:"id"`
		Email             string `json:"email"`
		Role              string `json:"role"`
		PayloadApplicable bool   `json:"payloadApplicable"`
	} `json:"webhookNotificationApplicableUsers"`
	Event                   string `json:"event"`
	SubEvent                string `json:"subEvent"`
	EventDate               string `json:"eventDate"`
	EventResourceType       string `json:"eventResourceType"`
	EventResourceParentType string `json:"eventResourceParentType"`
	EventResourceParentID   string `json:"eventResourceParentId"`
	ParticipantRole         string `json:"participantRole"`
	ActionType              string `json:"actionType"`
	ParticipantUserID       string `json:"participantUserId"`
	ParticipantUserEmail    string `json:"participantUserEmail"`
	ActingUserID            string `json:"actingUserId"`
	ActingUserEmail         string `json:"actingUserEmail"`
	ActingUserIPAddress     string `json:"actingUserIpAddress"`
	InitiatingUserID        string `json:"initiatingUserId"`
	InitiatingUserEmail     string `json:"initiatingUserEmail"`
	Agreement               struct {
		ID            string `json:"id"`
		Name          string `json:"name"`
		SignatureType string `json:"signatureType"`
		Status        string `json:"status"`
		Ccs           []struct {
			Email        string   `json:"email"`
			Label        string   `json:"label"`
			VisiblePages []string `json:"visiblePages"`
		} `json:"ccs"`
		DeviceInfo struct {
			ApplicationDescription string `json:"applicationDescription"`
			DeviceDescription      string `json:"deviceDescription"`
			Location               struct {
				Latitude  string `json:"latitude"`
				Longitude string `json:"longitude"`
			} `json:"location"`
			DeviceTime string `json:"deviceTime"`
		} `json:"deviceInfo"`
		DocumentVisibilityEnabled bool   `json:"documentVisibilityEnabled"`
		CreatedDate               string `json:"createdDate"`
		ExpirationTime            string `json:"expirationTime"`
		ExternalID                struct {
			ID string `json:"id"`
		} `json:"externalId"`
		PostSignOption struct {
			RedirectDelay int    `json:"redirectDelay"`
			RedirectURL   string `json:"redirectUrl"`
		} `json:"postSignOption"`
		FirstReminderDelay int    `json:"firstReminderDelay"`
		Locale             string `json:"locale"`
		Message            string `json:"message"`
		ReminderFrequency  string `json:"reminderFrequency"`
		SenderEmail        string `json:"senderEmail"`
		VaultingInfo       struct {
			Enabled string `json:"enabled"`
		} `json:"vaultingInfo"`
		WorkflowID          string `json:"workflowId"`
		ParticipantSetsInfo struct {
			ParticipantSets []struct {
				MemberInfos []struct {
					ID             string `json:"id"`
					Email          string `json:"email"`
					Company        string `json:"company"`
					Name           string `json:"name"`
					PrivateMessage string `json:"privateMessage"`
					Status         string `json:"status"`
				} `json:"memberInfos"`
				Order          int    `json:"order"`
				Role           string `json:"role"`
				Status         string `json:"status"`
				ID             string `json:"id"`
				Name           string `json:"name"`
				PrivateMessage string `json:"privateMessage"`
			} `json:"participantSets"`
		} `json:"participantSetsInfo"`
		DocumentsInfo struct {
			Documents []struct {
				ID       string `json:"id"`
				Label    string `json:"label"`
				NumPages int    `json:"numPages"`
				MimeType string `json:"mimeType"`
				Name     string `json:"name"`
			} `json:"documents"`
			SupportingDocuments []struct {
				DisplayLabel string `json:"displayLabel"`
				FieldName    string `json:"fieldName"`
				ID           string `json:"id"`
				MimeType     string `json:"mimeType"`
				NumPages     int    `json:"numPages"`
			} `json:"supportingDocuments"`
		} `json:"documentsInfo"`
	} `json:"agreement"`
}

type ApprovalRequestInsertInput struct {
	ObjectType  string `json:"object_type,omitempty"`
	ObjectUUID  string `json:"object_uuid,omitempty"`
	Status      string `json:"status,omitempty"`
	DueDate     string `json:"due_date,omitempty"`
	WorkflowId  string `json:"workflow_id,omitempty"`
	Message     string `json:"message,omitempty"`
	ServiceName string `json:"service_name,omitempty"`
}
