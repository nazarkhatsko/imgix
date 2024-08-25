package services

type ResizeService struct{}

func NewResizeService() *ResizeService {
	service := ResizeService{}
	return &service
}
