package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/syafdia/demo-es/internal/application"
	"github.com/syafdia/demo-es/internal/domain"
	"github.com/syafdia/demo-es/internal/domain/task"
	"github.com/syafdia/demo-es/internal/infrastructure/repository"
	"github.com/syafdia/demo-es/internal/infrastructure/service"
)

func main() {

	taskRepository := repository.NewTaskRepository()
	emailService := service.NewEmailService()
	cmdHandler := application.NewCommandHandler(taskRepository, emailService)

	ctx := context.Background()

	for {
		opt := p(`
Available option
	[ct] Create task
	[ut] Update task status
	[dv] Downvote task
	[uv] Upvote task
	[vt] View task
	[ve] View all events
Enter option: `)

		switch strings.TrimSpace(opt) {
		case "ct":
			title := p("[ct] Enter task title:")
			description := p("[ct] Enter task description:")

			err := cmdHandler.CreateTask(ctx, application.CreateTaskCommand{
				TaskTitle:       title,
				TaskDescription: description,
			})
			handleError(err)

		case "ut":
			taskID := p("[ut] Enter task ID:")
			taskStatus := p("[ut] Enter status (NOT_STARTED, IN_PROGRESS, IN_REVIEW, COMPLETED):")

			err := cmdHandler.UpdateTaskStatus(ctx, application.UpdateTaskStatusCommand{
				TaskID:     domain.ID(taskID),
				TaskStatus: task.TaskStatus(taskStatus),
			})
			handleError(err)

		case "uv":
			taskID := p("[uv] Enter task ID:")

			err := cmdHandler.UpvoteTask(ctx, application.UpvoteTaskCommand{
				TaskID: domain.ID(taskID),
			})
			handleError(err)

		case "dv":
			taskID := p("[dv] Enter task ID:")

			err := cmdHandler.DownvoteTask(ctx, application.UpvoteTaskCommand{
				TaskID: domain.ID(taskID),
			})
			handleError(err)

		case "vt":
			taskID := p("[vt] Enter task ID:")

			t, err := taskRepository.Find(ctx, domain.ID(taskID))
			handleError(err)

			fmt.Printf("Current state of task: %+v\n", t)

		case "ve":
			taskRepository.Peek(ctx)
		}

		// fmt.Print("Enter your city: ")
		// city, _ := reader.ReadString('\n')
		// fmt.Print("You live in " + city)
	}
}

func p(message string) string {
	fmt.Print(message + " ")
	result, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(result)
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Got error,", err)
	}
}
