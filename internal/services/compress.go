package services

type CompressService struct{}

func NewCompressService() *CompressService {
	service := CompressService{}
	return &service
}
