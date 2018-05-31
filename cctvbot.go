package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func moveServo(servo *gpio.ServoDriver, angle uint8) {
	if angle < 18 {
		angle = 19
	}
	if angle > 30 {
		angle = 30
	}
	servo.Move(angle)
}

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
	servo.Move(0)

	robot := gobot.NewRobot("CCTVBot",
		[]gobot.Connection{adapter},
		[]gobot.Device{baseStepper, servo},
		func() {},
	)

	robot.AddCommand("pan", func(params map[string]interface{}) interface{} {
		switch params["direction"] {
		case "left":
			baseStepper.Move(100)
		case "right":
			baseStepper.Move(-100)
		case "up":
			moveServo(servo, servo.CurrentAngle+1)
		case "down":
			moveServo(servo, servo.CurrentAngle-1)
		}
		return fmt.Sprintf("This command pans the cctv %+v", params)
	})

	return robot
}
