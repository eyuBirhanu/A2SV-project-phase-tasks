package concurrency

import (
	"fmt"
	"library_management/services"
)

type ReservationRequest struct {
	BookID   int
	MemberID int
}

type ReservationHandler struct {
	LibraryService services.LibraryManager
	RequestQueue   chan ReservationRequest
}

func NewReservationHandler(service services.LibraryManager, bufferSize int) *ReservationHandler {
	return &ReservationHandler{
		LibraryService: service,
		RequestQueue:   make(chan ReservationRequest, bufferSize),
	}
}

func (rh *ReservationHandler) StartWorkers(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go func(workerID int) {
			for req := range rh.RequestQueue {
				err := rh.LibraryService.ReserveBook(req.BookID, req.MemberID)
				if err != nil {
					fmt.Printf("[Worker %d] Error: %v\n", workerID, err)
				}
			}
		}(i)
	}
}

func (rh *ReservationHandler) AddRequest(bookID, memberID int) {
	rh.RequestQueue <- ReservationRequest{BookID: bookID, MemberID: memberID}
}
