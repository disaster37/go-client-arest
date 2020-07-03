package gpio

import (
	"errors"

	"github.com/disaster37/go-client-arest/v1"
)

const (
	NO       int = 0
	NC       int = 1
	ON       int = 1
	OFF      int = 0
	SateOn       = "on"
	StateOff     = "off"
)

type Relay struct {
	GPIO         *GPIO
	Level        int
	Output       int
	DefaultState string
	state        string
}

func NewRelay(c client.Client, pin int, level int, output int, defaultState string) (relay relay, err error) {

	if level != client.HIGH && level != client.LOW {
		errors.New("Level must be HIGH or LOW")
	}
	if output != NO && output != NC {
		errors.New("Output must be NO or NC")
	}
	if defaultState != SateOn && defaultState != StateOff {
		errors.New("DefaultState must be StateOn or StateOff")
	}

	gpio := &GPIO{
		Client: c,
		Pin:    pin,
		Mode:   client.OUTPUT,
	}

	relay = &Relay{
		GPIO:         gpio,
		Level:        level,
		Output:       output,
		DefaultState: defaultState,
	}

	// Set pin mode
	err = c.SetPinMode(pin, gpio.Mode)
	if err != nil {
		return nil, err
	}

	// Set default state
	switch defaultState {
	case SateOn:
		err = relay.On()
	case StateOff:
		err = relay.Off()
	}

	return relay, err

}

func (r *Relay) On() (err error) {

	switch r.Output {
	case NO:
		// Normaly Open
		switch r.Level {
		case client.HIGH:
			// High signal
			err = r.GPIO.Client.DigitalWrite(r.GPIO.Pin, client.HIGH)
		case client.LOW:
			// Low signal
			err = r.GPIO.Client.DigitalWrite(r.GPIO.Pin, client.LOW)
		}
	case NC:
		// Normaly Close
		switch r.Level {
		case client.HIGH:
			// High signal
			err = r.GPIO.Client.DigitalWrite(r.GPIO.Pin, client.LOW)
		case client.LOW:
			// Low signal
			err = r.GPIO.Client.DigitalWrite(r.GPIO.Pin, client.HIGH)
		}
	}

	if err == nil {
		r.state = SateOn
	}

	return err
}

func (r *Relay) Off() (err error) {
	switch r.Output {
	case NO:
		// Normaly Open
		switch r.Level {
		case client.HIGH:
			// High signal
			err = r.GPIO.Client.DigitalWrite(r.GPIO.Pin, client.LOW)
		case client.LOW:
			// Low signal
			err = r.GPIO.Client.DigitalWrite(r.GPIO.Pin, client.HIGH)
		}
	case NC:
		// Normaly Close
		switch r.Level {
		case client.HIGH:
			// High signal
			err = r.GPIO.Client.DigitalWrite(r.GPIO.Pin, client.HIGH)
		case client.LOW:
			// Low signal
			err = r.GPIO.Client.DigitalWrite(r.GPIO.Pin, client.LOW)
		}
	}

	if err == nil {
		r.state = StateOff
	}

	return err
}

func (r *Relay) State() string {
	return r.State()
}
