package enums

var UserTypes = newUserTypeRegistry()

type userTypeRegistry struct {
	Mentor string
	Mentee string
	Admin  string
}

func newUserTypeRegistry() *userTypeRegistry {
	return &userTypeRegistry{
		Mentor: "MENTOR",
		Mentee: "MENTEE",
		Admin:  "ADMIN",
	}
}
