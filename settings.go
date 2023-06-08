package gorf

// GlobalSettings Global Project settings
type GlobalSettings struct {
	SecretKey  string
	UserObjKey string
	UserObjId  string
	DbBackends DbBackend
}

var Settings = &GlobalSettings{
	UserObjKey: "user",
	UserObjId:  "id",
	DbBackends: nil,
}
