package infra

import (
	"reflect"
	"github.com/jdextraze/go-gesclient/client"
	"encoding/json"
	"github.com/satori/go.uuid"
	"errors"
	"fmt"
)

type EventFactory func (string) interface{}

type Aggregate interface {
	Hydrate(interface{})
	Execute(interface{}) ([]interface{}, error)
}

type CommandHandler func (aggregateId string, aggregate Aggregate, command interface{}) error

func (i *Infra) HandleCommand(aggregateId string, aggregate Aggregate, command interface{}) error {
	aggType := reflect.TypeOf(aggregate).Elem()
	stream := aggType.Name() + "-" + aggregateId

	var result *client.StreamEventsSlice
	if t, err := i.conn.ReadStreamEventsForwardAsync(stream, 0, 4096, false, nil); err != nil {
		return err
	} else if err := t.Wait(); err != nil {
		return err
	} else {
		result = t.Result().(*client.StreamEventsSlice)
	}
	if !result.IsEndOfStream() {
		return errors.New("stream is longer than 4096 events")
	}

	for _, event := range result.Events() {
		originalEvent := event.OriginalEvent()
		data := originalEvent.Data()
		if ev := i.eventFactory(originalEvent.EventType()); ev == nil {
			return errors.New(fmt.Sprintf("Unknow eventType %s", originalEvent.EventType()))
		} else if err := json.Unmarshal(data, &ev); err != nil {
			return err
		} else {
			aggregate.Hydrate(ev)
		}
	}

	uncommitedEvents, err := aggregate.Execute(command)
	if err != nil {
		return err
	}
	eventsToCommit := make([]*client.EventData, len(uncommitedEvents))
	for i, ev := range uncommitedEvents {
		if data, err := json.Marshal(ev); err != nil {
			return err
		} else {
			evType := reflect.TypeOf(ev).Elem()
			eventsToCommit[i] = client.NewEventData(uuid.NewV4(), evType.Name(), true, data, nil)
		}
	}
	if t, err := i.conn.AppendToStreamAsync(stream, result.LastEventNumber(), eventsToCommit, nil); err != nil {
		return err
	} else if err := t.Wait(); err != nil {
		return err
	}

	return nil
}