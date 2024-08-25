package services

type ConvertService struct{}

func NewConvertService() *ConvertService {
	service := ConvertService{}
	return &service
}
