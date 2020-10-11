package models

type Session struct {
	UserName          string
	UserGroups        UserGroups
	CartProductsCount uint
	Categories        []string
}
