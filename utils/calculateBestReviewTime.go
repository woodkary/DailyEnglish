package utils

import (
	"math"
	"strconv"
	"strings"
)

const (
	iterations = 200000
	max_index  = 122
	min_index  = -30
	base       = 1.05
)

func calStartHalflife(difficulty int) float64 {
	p := math.Max(0.925-0.05*float64(difficulty), 0.025)
	return -1 / math.Log2(p)
}

func calNextRecallHalflife(h, p float64, d, recall int) float64 {
	if recall == 1 {
		return h * (1 + math.Exp(3.81140723)*math.Pow(float64(d), -0.5345194)*math.Pow(h, -0.12641492)*math.Pow(1-p, 0.97043354))
	} else {
		return math.Exp(-0.04141891) * math.Pow(float64(d), -0.04074844) * math.Pow(h, 0.37749318) * math.Pow(1-p, -0.22722912)
	}
}

func calHalflifeIndex(h float64) int {
	return int(math.Max(math.Round(math.Log(h)/math.Log(base))-float64(min_index), 0))
}

func CalculateBestInterval(d int, interval_history, feedback_history string) int {
	if interval_history == "" && feedback_history == "" {
		return calHalflifeIndex(calStartHalflife(d))
	}
	interval_history_array := strings.Split(interval_history, ",")
	feedback_history_array := strings.Split(feedback_history, ",")
	best_interval := 1
	for i := 0; i < iterations; i++ {
		for j := 0; j < len(interval_history_array); j++ {
			interval, _ := strconv.Atoi(interval_history_array[j])
			feedback, _ := strconv.Atoi(feedback_history_array[j])
			h := calStartHalflife(d)
			p_recall := math.Exp2(-float64(interval) / h)
			recall_h := calNextRecallHalflife(h, p_recall, d, feedback)
			h_index := calHalflifeIndex(recall_h)
			if h_index > best_interval {
				best_interval = h_index
			}
		}
	}
	return best_interval
}
