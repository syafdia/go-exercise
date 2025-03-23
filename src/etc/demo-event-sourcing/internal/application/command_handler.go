package application

import (
	"context"
	"fmt"

	"github.com/syafdia/demo-es/internal/domain"
	"github.com/syafdia/demo-es/internal/domain/communication"
	"github.com/syafdia/demo-es/internal/domain/task"
)

type CreateTaskCommand struct {
	TaskTitle       string
	TaskDescription string
}

type UpdateTaskStatusCommand struct {
	TaskID     domain.ID
	TaskStatus task.TaskStatus
}

type UpvoteTaskCommand struct {
	TaskID domain.ID
}

type CommandHandler struct {
	taskRepository task.TaskRepository
	emailService   communication.EmailService
}

func NewCommandHandler(
	taskRepository task.TaskRepository,
	emailService communication.EmailService,
) *CommandHandler {
	return &CommandHandler{
		taskRepository: taskRepository,
	}
}

func (c *CommandHandler) CreateTask(ctx context.Context, cmd CreateTaskCommand) error {
	t, err := task.NewTask(cmd.TaskTitle, cmd.TaskDescription)
	if err != nil {
		return err
	}

	err = c.taskRepository.Store(ctx, t)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommandHandler) UpdateTaskStatus(ctx context.Context, cmd UpdateTaskStatusCommand) error {
	t, err := c.taskRepository.Find(ctx, cmd.TaskID)
	if err != nil {
		return err
	}

	err = t.ChangeStatus(cmd.TaskStatus)
	if err != nil {
		return err
	}

	err = c.taskRepository.Store(ctx, t)
	if err != nil {
		return err
	}

	if t.Status == task.TaskStatusCompleted {
		c.emailService.SendEmail(ctx, fmt.Sprintf("Task %s has been completed", t.Title))
	}

	return nil
}

func (c *CommandHandler) UpvoteTask(ctx context.Context, cmd UpvoteTaskCommand) error {
	t, err := c.taskRepository.Find(ctx, cmd.TaskID)
	if err != nil {
		return err
	}

	err = t.Upvote()
	if err != nil {
		return err
	}

	err = c.taskRepository.Store(ctx, t)
	if err != nil {
		return err
	}

	return nil
}

func (c *CommandHandler) DownvoteTask(ctx context.Context, cmd UpvoteTaskCommand) error {
	t, err := c.taskRepository.Find(ctx, cmd.TaskID)
	if err != nil {
		return err
	}

	err = t.Downvote()
	if err != nil {
		return err
	}

	err = c.taskRepository.Store(ctx, t)
	if err != nil {
		return err
	}

	return nil
}
