package common

// Store - this interface defines all the methods
// that our service needs in order to operate 
type Store interface {}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new
//  service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}