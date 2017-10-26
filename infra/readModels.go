package infra

import (
	"net/http"
	"strings"
	"log"
	"github.com/jdextraze/go-gesclient/client"
	"encoding/json"
	"errors"
)

type EventData struct {
	Event interface{}
}

type TrxRepository struct {
	Data []interface{}
}

type EventHandler func (*TrxRepository, *EventData) error
type EventHandlerProvider func (string) EventHandler

func (r *TrxRepository) Create(record interface{}) error {
	r.Data = append(r.Data, record)
	return nil
}

func NewTrxRepository() *TrxRepository {
	return &TrxRepository{
		Data: make([]interface{}, 0),
	}
}

func (i *Infra) FindAll(modelName string) ([]interface{}, error) {
	pos := client.Position_Start

	handler := i.readModelHandlerProvider(modelName)
	if handler == nil {
		return nil, errors.New("no read model handler for " + modelName)
	}

	trxRepo := NewTrxRepository()

	for {
		var result *client.AllEventsSlice
		if t, err := i.conn.ReadAllEventsForwardAsync(pos, 1024, false, i.userCredentials); err != nil {
			return nil, err
		} else if err := t.Wait(); err != nil {
			return nil, err
		} else {
			result = t.Result().(*client.AllEventsSlice)
		}

		for _, event := range result.GetEvents() {
			if event.OriginalStreamId()[0:1] == "$" {
				continue
			}
			payload := i.eventFactory(event.OriginalEvent().EventType())
			if err := json.Unmarshal(event.OriginalEvent().Data(), payload); err != nil {
				return nil, err
			}
			eventData := &EventData{
				Event: payload,
			}
			if err := handler(trxRepo, eventData); err != nil {
				return nil, err
			}
		}

		if result.IsEndOfStream() {
			break
		}
		pos = result.GetNextPosition()
	}

	return trxRepo.Data, nil
}

type GenericReadModelHandler struct {
	infra *Infra
}

func (h *GenericReadModelHandler) ServeHTTP(w http.ResponseWriter, r* http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		return
	}

	parts := strings.Split(r.URL.Path, "/")
	modelName := parts[len(parts)-1]

	log.Printf("%s %s\n", modelName, r.URL.Query().Get("filter"))

	if results, err := h.infra.FindAll(modelName); err != nil {
		WriteError(w, err)
		return
	} else {
		WriteJson(w, results)
	}
	//TODO findOne
	//TODO filter
}

func (i *Infra) NewGenericReadModelsHandler() http.Handler {
	return &GenericReadModelHandler{infra:i}
}