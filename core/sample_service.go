package core

func SampleServiceCode() string {
	return `package service

/// TODO: this code example, you can use the following structure:

type Repository interface {
	GetByID(ctx context.Context, userID int) (*entity.User, error)
}

type Service struct {
	repo Repository
}

func New(repository Repository) *Service {
	return &Service{repo: repository}
}`
}

func SampleMethodServiceCode() string {
	return `package service

/// TODO: this code example, you can use the following structure:

func (s Service) GetByID(ctx context.Context, req param.GetUserReq) (param.GetUserRes, error) {
	user, err := s.repo.GetByID(ctx, req.ID)
	if err != nil {
		return param.GetUserRes{}, err
	}
	return params.GetUserRes{
		FirstName:  user.FirstName,
		FamilyName: user.FamilyName,
	}, nil
}`
}
