package signtaskResponseModel

type GetSignTaskSlicingTicketIdRes struct {
	RequestId string                         `json:"requestId"`
	Code      string                         `json:"code"`
	Msg       string                         `json:"msg"`
	Data      GetSignTaskSlicingTicketIdData `json:"data"`
}

type GetSignTaskSlicingTicketIdData struct {
	SlicingTicketId string `json:"slicingTicketId,omitempty"`
}
