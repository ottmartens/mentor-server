package enums

var RequestTypes = newRequestTypeRegistry()

type requestTypeRegistry struct {
	CreateGroup string
	JoinGroup   string
}

func newRequestTypeRegistry() *requestTypeRegistry {
	return &requestTypeRegistry{
		CreateGroup: "CREATE_GROUP",
		JoinGroup:   "JOIN_GROUP",
	}
}
