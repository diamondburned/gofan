package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/diamondburned/gofan/easings"
	"github.com/diamondburned/gofan/internal/fileutil"
)

var (
	function = "EaseInOutCubic"
	interval = 2 * time.Second

	p     threshold
	pPath string
	t     threshold
	tPath string
)

type threshold struct {
	min, max int
}

func needInt(v int, flag string) {
	if v == 0 {
		log.Fatalf("missing flag %s", flag)
	}
}

func needStr(v, flag string) {
	if v == "" {
		log.Fatalf("missing flag %s", flag)
	}
}

func main() {
	flag.StringVar(&function, "function", function, "function to use for curve")
	flag.DurationVar(&interval, "interval", interval, "polling interval (delay)")
	flag.IntVar(&p.min, "pmin", 0, "PWM minimum")
	flag.IntVar(&p.max, "pmax", 0, "PWM maximum")
	flag.StringVar(&pPath, "p", "", "PWM path")
	flag.IntVar(&t.min, "tmin", 0, "temperature minimum")
	flag.IntVar(&t.max, "tmax", 0, "temperature maximum")
	flag.StringVar(&pPath, "t", "", "temperature path")
	flag.Parse()

	needStr(tPath, "-t")
	needStr(pPath, "-p")
	needInt(p.min, "-pmin")
	needInt(p.max, "-pmax")
	needInt(t.min, "-tmin")
	needInt(t.max, "-tmax")

	ease, ok := easings.CurveFunctions[function]
	if !ok {
		log.Fatalf("unknown easing function %q", function)
	}

	if err := setManualPWM(true); err != nil {
		log.Fatalln("cannot set manual PWM:", err)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// TODO: reduce polling frequency if the temperature is too idly. Mostly an
	// energy saving measure, but it doesn't really matter.
	ticker := time.NewTicker(interval)
	temp := newTempObserver(tPath)
	pwm := newPWMObserver(ease)

pollLoop:
	for {
		select {
		case <-ctx.Done():
			break pollLoop
		case <-ticker.C:
			// do
		}

		tchanged, err := temp.check()
		if err != nil {
			log.Fatalln("cannot check temperature:", err)
		}
		if !tchanged {
			continue
		}

		pchanged := pwm.recalculate(temp.val)
		if !pchanged {
			continue
		}

		if err := changePWM(pwm.val); err != nil {
			log.Fatalln("cannot change PWM:", err)
		}
	}

	if err := setManualPWM(false); err != nil {
		log.Fatalln("cannot set automatic PWM:", err)
	}
}

type tempObserver struct {
	file *fileutil.Scanner
	val  int
}

func newTempObserver(path string) tempObserver {
	return tempObserver{
		file: fileutil.NewScanner(path),
	}
}

func (o *tempObserver) check() (changed bool, err error) {
	t, err := o.file.ScanInt()
	if err != nil {
		return false, err
	}
	if t == o.val {
		return false, nil
	}
	o.val = t
	return true, nil
}

type pwmObserver struct {
	val  int
	ease func(float64) float64
}

func newPWMObserver(ease func(float64) float64) pwmObserver {
	return pwmObserver{
		ease: ease,
	}
}

func (o *pwmObserver) recalculate(temp int) (changed bool) {
	newPWM := calculatePWM(p, t, temp, o.ease)
	if newPWM == o.val {
		return false
	}
	o.val = newPWM
	return true
}

func calculatePWM(p, t threshold, temp int, ease func(float64) float64) int {
	switch {
	case temp <= t.min:
		return p.min
	case temp >= t.max:
		return p.max
	}

	// n within [0, 1]
	n := float64(temp-t.min) / float64(t.max-t.min)

	v := ease(n)
	if v < 0 {
		v = 0
	} else if v > 1 {
		v = 1
	}

	return p.min + roundInt(v*float64(p.max-p.min))
}

func roundInt(f float64) int {
	// f always positive, so we can do this
	return int(f + 0.5)
}

func changePWM(pwm int) error {
	return writeIntFile(pPath, pwm)
}

func setManualPWM(manual bool) error {
	if manual {
		return writeIntFile(pPath+"_enable", 1)
	}

	err := writeIntFile(pPath+"_enable", 0)
	if err == nil {
		return nil // worked
	}

	// Uh oh, auto mode doesn't work. Ramp the user's fan all the way up!
	changePWM(p.max)
	return err
}

func writeIntFile(file string, v int) error {
	return os.WriteFile(file, []byte(strconv.Itoa(v)), 0)
}
