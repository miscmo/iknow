package model

import "encoding/json"

type RspCode int32

const (
	Success        RspCode = 0
	ParamInvild    RspCode = 1
	UpdateDBFailed RspCode = 2
)

type OptionalString struct {
	Defined bool
	Value *string
}

func (os *OptionalString)UnmarshalJSON(data []byte) error {
	os.Defined = true
	return json.Unmarshal(data, &os.Value)
}

type OptionalInt struct {
	Defined bool
	Value *int
}
func (oi *OptionalInt)UnmarshalJSON(data []byte) error {
	oi.Defined = true
	return json.Unmarshal(data, &oi.Value)
}

type OptionalBool struct {
	Defined bool
	Value *bool
}
func (ob *OptionalBool)UnmarshalJSON(data []byte) error {
	ob.Defined = true
	return json.Unmarshal(data, &ob.Value)
}

type NoteDiff struct {
	Version int `json:"version,omitempty"`
	Diff string `json:"diff,omitempty"`
	Desc string `json:"desc,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

type Note struct {
	Id int `json:"id,omitempty"`
	Name OptionalString `json:"name,omitempty"`
	Desc OptionalString `json:"desc,omitempty"`
	Content OptionalString `json:"content,omitempty"`
	Type OptionalString	`bson:"type,omitempty"`	// "text", "audio", "markdown", "video", "pdf"
	Diff []NoteDiff `json:"diff,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

type Node struct {
	Id int `json:"id,omitempty"`
	Name OptionalString 	`json:"name,omitempty"`
	Desc OptionalString `json:"desc,omitempty"`
	Type OptionalInt	`json:"type,omitempty"`
	Notes []int	`json:"notes,omitempty"`
	SubNodes []int `json:"subNodes,omitempty"`
	Collapse OptionalBool `json:"collapse,omitempty"`
	Score      OptionalInt `json:"score,omitempty"`
	Permission OptionalInt `json:"permission,omitempty"`
	CreateTime string `json:"createTime,omitempty"`
	UpdateTime string `json:"updateTime,omitempty"`
}

type SearchNodesReq struct {
	Id       int64 `json:"id,omitempty"`
	Page     int32 `json:"page,omitempty"`
	PageSize int32 `json:"pageSize,omitempty"`
}
type SearchNodesRsp struct {
	Code  RspCode `json:"code"`
	Msg   string  `json:"msg"`
	Nodes []*Node `json:"nodes,omitempty"`
}

type AddNodeReq struct {
	Node   *Node `json:"node,omitempty"`
}
type AddNodeRsp struct {
	Code  RspCode `json:"code"`
	Msg   string  `json:"msg"`
}

type AddSubNodeReq struct {
	Parent int64 `json:"parent,omitempty"`
	Node *Node `json:"node,omitempty"`
}

type AddSubNodeRsp struct {
	Code RspCode `json:"code"`
	Msg string `json:"msg"`
}

type AddNoteReq struct {
	Note *Note `json:"note,omitempty"`
}
type AddNoteRsp struct {
	Code RspCode `json:"code"`
	Msg string `json:"msg"`
}

type UpdateNodeReq struct {
	Id int64 `json:"id,omitempty"`
	Node *Node `json:"node,omitempty"`
}
type UpdateNodeRsp struct {
	Code RspCode `json:"code"`
	Msg string `json:"msg"`
}

type UpdateNoteReq struct {
	Id int64 `json:"id,omitempty"`
	Note *Note `json:"note,omitempty"`
}
type UpdateNoteRsp struct {
	Code RspCode `json:"code"`
	Msg string `json:"msg"`
}