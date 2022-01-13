package datatypes

type InputUserData struct {
	FirstName    string `validate:"required"`
	LastName     string `validate:"required"`
	Age          int    `validate:"required,gte=18,lte=80"`
	City         string `validate:"required"`
	Organization string `validate:"required,contains=/"`
}