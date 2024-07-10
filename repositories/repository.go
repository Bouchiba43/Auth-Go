package repositories

type Repository interface {
	FindAll() ([]interface{}, error)
	FindById(id int) (interface{}, error)
	Create(data interface{}) (interface{}, error)
	UpdateById(id int, data interface{}) (interface{}, error)
	DeleteById(id int) error
	Save() error
}
