package wg

import (
	"os/exec"
	"strings"
)

func ActiveInterfaces() ([]InterfaceStatus, error) {
	cmd := exec.Command("wg", "show", "interfaces")
	out, err := cmd.Output()
	if err != nil {
		return []InterfaceStatus{}, err
	}

	interfaceNames := strings.TrimSpace(string(out))

	if interfaceNames == "" {
		return []InterfaceStatus{}, nil
	}

	return parseActiveInterfaces(interfaceNames), nil
}

func parseActiveInterfaces(output string) []InterfaceStatus {
	interfaceStatuses := []InterfaceStatus{}
	for name := range strings.FieldsSeq(output) {
		interfaceStatuses = append(interfaceStatuses, InterfaceStatus{
			Name: string(name),
		})
	}
	return interfaceStatuses
}
