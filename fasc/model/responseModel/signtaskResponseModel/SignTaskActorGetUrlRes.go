package signtaskResponseModel

type SignTaskActorGetUrlRes struct {
	RequestId string                  `json:"requestId"`
	Code      string                  `json:"code"`
	Msg       string                  `json:"msg"`
	Data      SignTaskActorGetUrlData `json:"data"`
}

type SignTaskActorGetUrlData struct {
	ActorSignTaskUrl         string       `json:"actorSignTaskUrl,omitempty"`
	ActorSignTaskEmbedUrl    string       `json:"actorSignTaskEmbedUrl,omitempty"`
	ActorSignTaskMiniAppInfo *MiniAppInfo `json:"actorSignTaskMiniAppInfo,omitempty"`
}

type MiniAppInfo struct {
	WxOriginalId string `json:"wxOriginalId,omitempty"`
	Path         string `json:"path,omitempty"`
}
