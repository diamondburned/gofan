package main

import (
	"fmt"
	"strings"
	"testing"

	"github.com/diamondburned/gofan/easings"
)

func TestCalculatePWM(t *testing.T) {
	inputs := []struct {
		p    threshold
		t    threshold
		ease func(float64) float64
	}{
		{p: threshold{10, 255}, t: threshold{58, 85}, ease: easings.EaseInOutCubic},
		{p: threshold{10, 255}, t: threshold{58, 85}, ease: easings.EaseInOutElastic},
	}

	const (
		start = 50
		end   = 90
	)

	for i, input := range inputs {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			lines := make([]string, 1, end-start+1)
			lines[0] = "Temperature,PWM"

			for temp := start; temp <= end; temp++ {
				pwm := calculatePWM(input.p, input.t, temp, input.ease)
				lines = append(lines, fmt.Sprintf("%d,%d", temp, pwm))
			}

			t.Log("CSV output:\n" + strings.Join(lines, "\n"))
		})
	}
}
