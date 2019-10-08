package cli

import (
	"fmt"
	"github.com/fishi0x01/vsh/client"
	"io"
)

// CdCommand container for all 'cd' parameters
type CdCommand struct {
	name string

	client *client.Client
	stderr io.Writer
	stdout io.Writer
	Path   string
}

// NewCdCommand creates a new CdCommand parameter container
func NewCdCommand(c *client.Client, stdout io.Writer, stderr io.Writer) *CdCommand {
	return &CdCommand{
		name:   "cd",
		client: c,
		stdout: stdout,
		stderr: stderr,
	}
}

// GetName returns the CdCommand's name identifier
func (cmd *CdCommand) GetName() string {
	return cmd.name
}

// IsSane returns true if command is sane
func (cmd *CdCommand) IsSane() bool {
	return cmd.Path != ""
}

// Parse given arguments and return status
func (cmd *CdCommand) Parse(args []string) (success bool) {
	if len(args) == 2 {
		cmd.Path = args[1]
		success = true
	} else {
		fmt.Println("Usage:\ncd <path>")
	}
	return success
}

// Run executes 'cd' with given CdCommand's parameters
func (cmd *CdCommand) Run() error {
	newPwd := cmdPath(cmd.client.Pwd, cmd.Path)

	t, err := cmd.client.GetType(newPwd)
	if err != nil {
		return err
	}

	if t == client.LEAF {
		fmt.Fprintln(cmd.stderr, "cannot cd to '"+newPwd+"' because it is a file")
		return nil
	}
	newPwd = newPwd + "/"
	if newPwd == "//" {
		newPwd = "/"
	}
	cmd.client.Pwd = newPwd
	return err
}
