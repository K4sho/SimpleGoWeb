package model

type Model struct {
	db
}

// Метод создания экземпляра Model
func New(db db) *Model {
	return &Model {
		db: db,
	}
}

func (m *Model) People() ([]*Person, error)  {
	return m.SelectPeople()
}