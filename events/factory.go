package events

func CreateEvent(eventType string) interface{} {
	switch eventType {
	case "GameUpdated":
		return &GameUpdated {}
	
	}
	return nil
}
