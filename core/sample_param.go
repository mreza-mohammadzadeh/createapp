package core

func SampleParamsCode() string {
	return `package param

/// TODO: this code example, you can use the following structure:

type GetUserReq struct {
	ID string ` + "`param:\"id\"`" + `
}

type GetUserRes struct {
	FirstName  string ` + "`json:\"first_name\"`" + `
	FamilyName string ` + "`json:\"family_name\"`" + `
}`
}
