package commonModel

// Notification 是否发送通知
type Notification struct {
	SendNotification bool   `json:"sendNotification"`
	NotifyWay        string `json:"notifyWay,omitempty"`
	NotifyAddress    string `json:"notifyAddress,omitempty"`
}
