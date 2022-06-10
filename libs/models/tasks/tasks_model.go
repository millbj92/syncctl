package tasks

import (
	"github.com/google/uuid"
)

type TASK_TYPE string

const (
	TASK_TYPE_DELETE = "delete"
	TASK_TYPE_COPY = "copy"
	TASK_TYPE_MOVE = "move"
	TASK_TYPE_RENAME = "rename"
	TASK_TYPE_SYNC = "sync"
)

type Task struct {
	Id          uuid.UUID  `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Type        TASK_TYPE  `json:"type"`
	Args		IArgable   `json:"args"`
}

func NewTask(id uuid.UUID, name string, description string, t TASK_TYPE) *Task {
	switch t {
	case TASK_TYPE_COPY:
		return &Task{Id: id, Name: name, Description: description, Type: t, Args: &CopyArgs{}}
	case TASK_TYPE_MOVE:
		return &Task{Id: id, Name: name, Description: description, Type: t, Args: &MoveArgs{}}
	case TASK_TYPE_RENAME:
		return &Task{Id: id, Name: name, Description: description, Type: t, Args: &RenameArgs{}}
	case TASK_TYPE_SYNC:
		return &Task{Id: id, Name: name, Description: description, Type: t, Args: &SyncArgs{}}
	}
	return nil
}

type TaskArgs struct {
	Source string `json:"source"`
	Include string `json:"include"`
	Exclude string `json:"exclude"`
	Recursive bool `json:"recursive"`
	Force bool `json:"force"`
}

type IArgable interface {
	GetArgs() interface{}
}

type CopyArgs struct {
	TaskArgs
	Destination   string     `json:"destination"`
}

type MoveArgs struct {
	TaskArgs
	Destination   string     `json:"destination"`
}
type SyncArgs 	struct {
	TaskArgs
	Destination string     `json:"destination"`
	Delete 		bool       `json:"delete"`
	Destructive bool       `json:"destructive"`
}
type RenameArgs struct {
	TaskArgs
	Prefix 	    string    `json:"prefix"`
	Extension 	string    `json:"extension"`
}


func (c *CopyArgs) GetArgs() interface{} {
	return c
}

func (m *MoveArgs) GetArgs() interface{} {
	return m
}

func (s *SyncArgs) GetArgs() interface{} {
	return s
}

func (r *RenameArgs) GetArgs() interface{} {
	return r
}

func (t *Task) GetArgs() interface{} {
	switch t.Type {
	case TASK_TYPE_COPY:
		return t.Args.(*CopyArgs)
	case TASK_TYPE_MOVE:
		return t.Args.(*MoveArgs)
	case TASK_TYPE_RENAME:
		return t.Args.(*RenameArgs)
	case TASK_TYPE_SYNC:
		return t.Args.(*SyncArgs)
	}
	return nil
}

func (t *Task) GetType() TASK_TYPE {
	return t.Type
}

func (t *Task) GetId() uuid.UUID {
	return t.Id
}

func (t *Task) GetName() string {
	return t.Name
}

func (t *Task) GetDescription() string {
	return t.Description
}




