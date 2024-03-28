package dgsign

// Actor 签署任务参与方
type Actor struct {
	ActorType   ActorType         `desc:"参与方类型" json:"actorType" binding:"required,mustIn=corp#person"`
	ActorId     string            `desc:"参与方id" json:"actorId" binding:"required"`
	ActorName   string            `desc:"参与方名称" json:"actorName" binding:"required"`
	Permissions []ActorPermission `desc:"参与方权限" json:"permissions" binding:"required"`
}

var (
	ACTOR_PLATFORM = &Actor{
		ActorType:   ACTOR_TYPE_CORP,
		ActorId:     "平台",
		ActorName:   "平台",
		Permissions: []ActorPermission{ACTOR_PERMISSION_SIGN},
	}
	ACTOR_CUSTOMER = &Actor{
		ActorType:   ACTOR_TYPE_CORP,
		ActorId:     "客户",
		ActorName:   "客户",
		Permissions: []ActorPermission{ACTOR_PERMISSION_SIGN},
	}
	ACTOR_EMPLOYEE = &Actor{
		ActorType:   ACTOR_TYPE_PERSON,
		ActorId:     "雇员",
		ActorName:   "雇员",
		Permissions: []ActorPermission{ACTOR_PERMISSION_SIGN},
	}
)

type ActorType string

func (at ActorType) String() string {
	return string(at)
}

const (
	ACTOR_TYPE_CORP   ActorType = "corp"
	ACTOR_TYPE_PERSON ActorType = "person"
)

type ActorPermission string

func (ap ActorPermission) String() string {
	return string(ap)
}

const (
	ACTOR_PERMISSION_FILL ActorPermission = "fill"
	ACTOR_PERMISSION_SIGN ActorPermission = "sign"
	ACTOR_PERMISSION_CC   ActorPermission = "cc"
)

type FileType string

func (ft FileType) String() string {
	return string(ft)
}

const (
	FILE_TYPE_DOC    FileType = "doc"
	FILE_TYPE_ATTACH FileType = "attach"
)

type NotifyWay string

func (nw NotifyWay) String() string {
	return string(nw)
}

const (
	NOTIFY_WAY_MOBILE NotifyWay = "mobile"
	NOTIFY_WAY_EMAIL  NotifyWay = "email"
)
