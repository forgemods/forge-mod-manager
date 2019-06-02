package install

type Mod struct {
	name string
	Dependencies []string `json:"dependencies"`
}