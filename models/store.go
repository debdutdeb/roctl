package models

type Store interface {
	create()
	insert()
	delete()
	update()
}
