package admin

type Permission struct {
	Id          int
	Name        string
	Description string
	Roles       []*Role `orm:"reverse(many)"`
}
