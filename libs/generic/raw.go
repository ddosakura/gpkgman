package generic

// Name is the name of the mod
func (*Mod) Name() string {
	return "generic"
}

// Version is the version of the mod
func (*Mod) Version() (int, int, int) {
	return 0, 0, 1
}

// Description is the description of the mod
func (*Mod) Description() string {
	return "Generic Mod committed to solving most dependency problems"
}

// Dependences is the dependences of the mod
func (*Mod) Dependences() [][]string {
	return [][]string{
		[]string{
			"git",
		},
	}
}
