package healthcheck

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

func GetGpuTemp() float64 {
	cmd := exec.Command("vcgencmd", "measure_temp")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error checking temperature:", err)
	}

	return parseVcgencmd(string(output))
}

func parseVcgencmd(cmdOutput string) float64 {
	re := regexp.MustCompile("=|'")
	split := re.Split(cmdOutput, -1)

	parsed, _ := strconv.ParseFloat(split[1], 64)

	return parsed
}
