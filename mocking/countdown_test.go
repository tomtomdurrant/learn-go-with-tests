package main

import (
	"time"
	"reflect"
	"bytes"
	"testing"
)
const write = "write"
const sleep = "sleep"
type SpySleeper struct {
	Calls int
}
func (s *SpySleeper) Sleep() {
	s.Calls++
}
type SpyTime struct {
	durationSlept time.Duration
}
func (s *SpyTime) Sleep(duration time.Duration)  {
	s.durationSlept = duration
}
type CountdownOperationsSpy struct {
	Calls []string
}
func (s *CountdownOperationsSpy) Sleep()  {
	s.Calls = append(s.Calls, sleep)
}
func (s *CountdownOperationsSpy) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("Get correct output", func (t *testing.T) {
	buffer := bytes.Buffer{}
	spySleeper := CountdownOperationsSpy{}
	Countdown(&buffer, &spySleeper)
	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
	})
	t.Run("sleep before every print", func (t *testing.T) {
		spySleepPrinter := &CountdownOperationsSpy{}
		Countdown(spySleepPrinter, spySleepPrinter)
		want := []string{
			sleep,
			write, 
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	
	spyTime := SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()
	
	if spyTime.durationSlept != sleeper.duration {
		t.Errorf("Should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}