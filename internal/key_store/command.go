package keystore

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Operation string

const (
	GET           Operation = "GET"
	SET           Operation = "SET"
	DELETE        Operation = "DELETE"
	GET_OR_CREATE Operation = "GET_OR_CREATE"
)

type Command struct {
	Operation Operation `json:"command"`
	Key       string    `json:"key"`
	Value     string    `json:"value,omitempty"`
}

type ErrUnknownCommand struct {
	operation Operation
}

func (err *ErrUnknownCommand) Error() string {
	return fmt.Sprintf("unknown command: %s", err.operation)
}

func GetCommand(data []byte) (*Command, error) {
	var cmd Command
	err := json.Unmarshal(data, &cmd)

	if err != nil {
		logger.Error("error decoding the command")
		return nil, err
	}

	switch cmd.Operation {
	case GET, SET, DELETE, GET_OR_CREATE:
		return &cmd, nil
	default:
		logger.Debug("unknown command recieved")
		return nil, &ErrUnknownCommand{operation: cmd.Operation}
	}
}

func getRawCommand(cmd Command) ([]byte, error) {
	raw_cmd, err := json.Marshal(cmd)

	if err != nil {
		logger.Error("unable to marshal the delete request command!")
		return nil, errors.New("unable to marshal the delete request command")
	}

	return raw_cmd, nil
}
