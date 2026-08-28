package db

import "log"

type PostingService interface {
	NewOrder(b Book) error
}

type StubbedPostingService struct {
}

// NewPostingService initializes the PostingService.
func NewPostingService() PostingService {
	return &StubbedPostingService{}
}

// NewOrder
func (sps *StubbedPostingService) NewOrder(b Book) error {
	log.Printf("STUBBED POSTING SERVICE: book %s posted: %v", b.ID, b)
	return nil
}
