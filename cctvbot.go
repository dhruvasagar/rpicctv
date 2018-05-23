package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func NewCCTVBot() *gobot.Robot {
	adapter := raspi.NewAdaptor()
	motor_pins := [4]string{
		"7",
		"11",
		"16",
		"18",
	}
	baseStepper := gpio.NewStepperDriver(adapter, motor_pins, gpio.StepperModes.HalfStepping, 2048)

	servo := gpio.NewServoDriver(adapter, "12")

	robot := gobot.NewRobot("CCTVBot",
		[]gobot.Connection{adapter},
		[]gobot.Device{baseStepper, servo},
		func() {},
	)

	robot.AddCommand("pan", func(params map[string]interface{}) interface{} {
		switch params["direction"] {
		case "left":
			baseStepper.Move(1000)
		case "right":
			baseStepper.Move(-1000)
		case "up":
			servo.Move(servo.CurrentAngle + 5)
		case "down":
			servo.Move(servo.CurrentAngle - 5)
		}
		return fmt.Sprintf("This command pans the cctv %+v", params)
	})

	return robot
}
