package core

func SampleDBRepositoryCode() string {
	return `package repository

/// TODO: this code example, you can use the following structure:

type Repository struct {
	mysqlAdapter *mysqladapter.DB
}

func NewRepository(mysqlAdapter *mysqladapter.DB) *Repository {
	return &Repository{
		mysqlAdapter: mysqlAdapter,
	}
}`
}

func SampleRepositoryCode() string {
	return `package repository

/// TODO: this code example, you can use the following structure:

func (repo *Repository) GetByID(ctx context.Context, id int) (*entity.User, error) {
	return repo.mysqlAdapter.GetByID(id)
}`
}
