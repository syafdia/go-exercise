package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"reflect"
	"time"

	"github.com/syafdia/demo-es/internal/domain"
	"github.com/syafdia/demo-es/internal/domain/task"
)

type TaskRepository struct {
	idWithEvents map[domain.ID][]domain.Event
	events       []domain.Event
}

func NewTaskRepository() *TaskRepository {
	return &TaskRepository{
		idWithEvents: map[domain.ID][]domain.Event{},
	}
}

func (tr *TaskRepository) Store(ctx context.Context, task task.Task) error {
	tr.events = append(tr.events, task.GetEvents()...)

	// f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	// if err != nil {
	// 	panic(err)
	// }

	// defer f.Close()

	// tr.allEvents(ctx, func(eventName, aggregateID string, payload any) {
	// 	row := fmt.Sprintf("event_name: %s, aggregate_id: %s, payload: %+v\n", eventName, aggregateID, payload)
	// 	if _, err := f.WriteString(row); err != nil {
	// 		panic(err)
	// 	}
	// })

	return nil
}

var fileName = fmt.Sprintf("./tmp/storage-%d.txt", time.Now().Unix())

func (tr *TaskRepository) Find(ctx context.Context, id domain.ID) (task.Task, error) {
	task := task.Task{}

	for _, event := range tr.events {
		if event.AggregateID() != id {
			continue
		}

		task.ApplyEvent(event)
	}

	return task, nil
}

func (tr *TaskRepository) Peek(ctx context.Context) {
	fmt.Println("Stored events:\n")
	tr.allEvents(ctx, func(eventName, aggregateID string, at time.Time, payload any) {
		jsonPayload, err := json.Marshal(payload)
		if err != nil {
			panic(err)
		}

		fmt.Printf(
			"event_name: %s, aggregate_id: %s, occured_at: %s, payload: %+v\n\n",
			eventName, aggregateID, at.Format(time.RFC3339), string(jsonPayload),
		)
	})

}

func (tr *TaskRepository) allEvents(ctx context.Context, fn func(string, string, time.Time, any)) {
	for _, event := range tr.events {
		tEvent := reflect.TypeOf(event)
		fn(tEvent.Name(), string(event.AggregateID()), event.OccuredAt(), event)
	}
}
