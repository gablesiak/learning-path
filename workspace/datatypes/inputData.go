package datatypes

type InputUserData struct {
	FirstName    string `validate:"required"gte=2,lte=15`
	LastName     string `validate:"required"gte=2,lte=30`
	Age          int    `validate:"required,gte=18,lte=80"`
	City         string `validate:"required",gte=2,lte=30`
	Organization string `validate:"required,contains=/"`
	
}