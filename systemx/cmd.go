package systemx

import (
	"os/exec"
)

func RunCommand(output func([]byte) error, name string, arg ...string) error {
	cmd := exec.Command(name, arg...)

	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	for {
		buf := make([]byte, 1024)
		_, err := stdout.Read(buf)
		if err != nil {
			break
		}
		if err := output(buf); err != nil {
			return err
		}
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}
