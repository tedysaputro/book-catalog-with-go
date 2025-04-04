package hello

// HelloService defines the interface for hello operations
type HelloService interface {
	GetHelloMessage() HelloResponseDTO
}

// helloServiceImpl implements HelloService
type helloServiceImpl struct{}

// NewHelloService creates a new instance of HelloService
func NewHelloService() HelloService {
	return &helloServiceImpl{}
}

// GetHelloMessage returns a hello message
func (s *helloServiceImpl) GetHelloMessage() HelloResponseDTO {
	return HelloResponseDTO{
		Message: "Hello, World!",
	}
}
