package main

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func moveServo(servo *gpio.ServoDriver, angle uint8) {
	if angle < 11 {
		angle = 11
	}
	if angle > 24 {
		angle = 24
	}
	servo.Move(angle)
}

func NewCCTVBot() *gobot.Robot {
	adapter := raspi.NewAdaptor()

	panServo := gpio.NewServoDriver(adapter, "11")
	tiltServo := gpio.NewServoDriver(adapter, "12")

	moveServo(panServo, 0)
	moveServo(tiltServo, 0)

	robot := gobot.NewRobot("CCTVBot",
		[]gobot.Connection{adapter},
		[]gobot.Device{panServo, tiltServo},
		func() {},
	)

	robot.AddCommand("pan", func(params map[string]interface{}) interface{} {
		switch params["direction"] {
		case "left":
			panServo.Move(panServo.CurrentAngle + 1)
		case "right":
			panServo.Move(panServo.CurrentAngle - 1)
		case "up":
			moveServo(tiltServo, tiltServo.CurrentAngle-1)
		case "down":
			moveServo(tiltServo, tiltServo.CurrentAngle+1)
		}
		return fmt.Sprintf("This command pans the cctv %+v", params)
	})

	return robot
}
