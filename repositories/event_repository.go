package repositories

type EventRepository struct{}

func NewEventRepository() *EventRepository {
	return &EventRepository{}
}
