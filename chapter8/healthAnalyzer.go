package main

import "fmt"

type Status int

const (
	TooBad   Status = 1
	Bad      Status = 2
	Normal   Status = 3
	Good     Status = 4
	VeryGood Status = 5
)

type HealthAnalyzer struct {
	User User
}

func NewHealthAnalyzer(user User) *HealthAnalyzer {
	healthAnalyzer := new(HealthAnalyzer)
	healthAnalyzer.User = user
	return healthAnalyzer
}

func (h *HealthAnalyzer) getDisplayedStatus(status Status) string {
	switch status {
	case VeryGood:
		return "매우 좋음"
	case Good:
		return "좋음"
	case Normal:
		return "보통"
	case Bad:
		return "나쁨"
	case TooBad:
		return "매우나쁨"
	default:
		return ""
	}
}

func (h *HealthAnalyzer) GetRecommendedCalories() (
	cal int, err error) {
	if h.User.Age < 10 {
		return 1500, nil
	}

	if h.User.Gender == Male {
		if h.User.Age < 15 {
			return 2300, nil
		} else if h.User.Age < 20 {
			return 2700, nil
		} else if h.User.Age < 50 {
			return 2500, nil
		} else if h.User.Age < 65 {
			return 2300, nil
		} else if h.User.Age < 75 {
			return 2000, nil
		}
		return 1800, nil
	} else if h.User.Gender == Female {
		if h.User.Age < 15 {
			return 2000, nil
		} else if h.User.Age < 20 {
			return 2100, nil
		} else if h.User.Age < 65 {
			return 2000, nil
		} else if h.User.Age < 75 {
			return 1700, nil
		}
		return 1600, nil
	}
	return 0, fmt.Errorf("%d는 유효한 성별이 아닙니다.", h.User.Gender)
}

func (h *HealthAnalyzer) GetCaloryDiff(
	cal int) (diffCal int, err error) {
	recommendedCal, err := h.GetRecommendedCalories()
	if err != nil {
		return 0, err
	}
	return recommendedCal - cal, nil
}

func (h *HealthAnalyzer) GetCaloryStatus(
	cal int) (status string, err error) {
	recommendedCal, err := h.GetRecommendedCalories()
	if err != nil {
		return "", err
	}
	diffCal, err := h.GetCaloryDiff(cal)
	if err != nil {
		return "", err
	}

	diffPercent := float32(diffCal) / float32(recommendedCal)
	if diffPercent < 0.05 {
		return h.getDisplayedStatus(VeryGood), nil
	} else if diffPercent < 0.1 {
		return h.getDisplayedStatus(Good), nil
	} else if diffPercent < 0.2 {
		return h.getDisplayedStatus(Normal), nil
	} else if diffPercent < 0.3 {
		return h.getDisplayedStatus(Bad), nil
	}
	return h.getDisplayedStatus(TooBad), nil
}
