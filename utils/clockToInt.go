package utils

import (
	"strconv"
	"strings"
)

func TimeRangeToMinutes(timeRange string) int {
	parts := strings.Split(timeRange, "~")
	if len(parts) != 2 {
		return 0
	}

	startTime := strings.TrimSpace(parts[0])
	endTime := strings.TrimSpace(parts[1])

	startParts := strings.Split(startTime, ":")
	startHour, _ := strconv.Atoi(startParts[0])
	startMinute, _ := strconv.Atoi(startParts[1])

	endParts := strings.Split(endTime, ":")
	endHour, _ := strconv.Atoi(endParts[0])
	endMinute, _ := strconv.Atoi(endParts[1])

	startTotalMinutes := startHour*60 + startMinute
	endTotalMinutes := endHour*60 + endMinute

	return endTotalMinutes - startTotalMinutes
}
