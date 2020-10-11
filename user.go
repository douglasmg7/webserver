package main

type UserGroups uint8

type Session struct {
	UserName          string
	UserGroups        UserGroups
	CartProductsCount uint
	Categories        []string
}

const (
	GroupRoot UserGroups = 1 << iota
	GroupAdmin
	GorupSeller
)

// Set.
func (g *UserGroups) Set(flag UserGroups) {
	// r := *g | flag
	// g = &r
	*g = *g | flag
}

// Clear.
func (g *UserGroups) Clear(flag UserGroups) {
	*g = *g &^ flag
}

// Toggle.
func (g *UserGroups) Toggle(flag UserGroups) {
	*g = *g ^ flag
}

// Has.
func (g UserGroups) Has(flag UserGroups) bool {
	return (g & flag) != 0
}

// Has.
func (g UserGroups) Hav() bool {
	// *g & flag
	return false
}

// Is Admin.
func (g UserGroups) IsAdmin() bool {
	return g&GroupAdmin != 0
}
