package enums

var GlobalSettingsTypes = newGlobalSettingsTypeRegistry()

type globalSettingsTypeRegistry struct {
	MentorsCanRegister string
	MenteesCanRegister string
}

func newGlobalSettingsTypeRegistry() *globalSettingsTypeRegistry {
	return &globalSettingsTypeRegistry{
		MentorsCanRegister: "MENTORS_CAN_REGISTER",
		MenteesCanRegister: "MENTEES_CAN_REGISTER",
	}
}
