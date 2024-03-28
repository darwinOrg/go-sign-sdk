package commonModel

type Doc struct {
	DocId         string  `json:"docId,omitempty"`
	DocName       string  `json:"docName,omitempty"`
	DocFileId     string  `json:"docFileId,omitempty"`
	DocTemplateId string  `json:"docTemplateId,omitempty"`
	DocFields     []Field `json:"docFields,omitempty"`
}
