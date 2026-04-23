package wg

import "os/exec"

func Show() (string, error) {
	cmd := exec.Command("wg", "show")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(out), nil
}
