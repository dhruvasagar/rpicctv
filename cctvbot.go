package main

import (
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func tiltServoMove(servo *gpio.ServoDriver, angle uint8) {
	if angle < 11 {
		angle = 11
	}
	if angle > 24 {
		angle = 24
	}
	servo.Move(angle)
}

func panServoMove(servo *gpio.ServoDriver, angle uint8) {
	if angle < 11 {
		angle = 11
	}
	if angle > 42 {
		angle = 42
	}
	servo.Move(angle)
}

func autoPan(panServo, tiltServo *gpio.ServoDriver) {
	tiltServo.Move(20)
	angle := 10
	angleDiff := 1
	gobot.Every(200*time.Millisecond, func() {
		if angle == 42 {
			angleDiff = -1
		} else if angle == 10 {
			angleDiff = 1
		}
		angle += angleDiff
		fmt.Println("Angle: ", angle)
		panServo.Move(uint8(angle))
	})
}

func NewCCTVBot() *gobot.Robot {
	adapter := raspi.NewAdaptor()

	panServo := gpio.NewServoDriver(adapter, "11")
	tiltServo := gpio.NewServoDriver(adapter, "12")

	panServoMove(panServo, 0)
	tiltServoMove(tiltServo, 0)

	robot := gobot.NewRobot("CCTVBot",
		[]gobot.Connection{adapter},
		[]gobot.Device{panServo, tiltServo},
		func() {},
		// func() { autoPan(panServo, tiltServo) },
	)

	robot.AddCommand("pan", func(params map[string]interface{}) interface{} {
		switch params["direction"] {
		case "left":
			panServoMove(panServo, panServo.CurrentAngle+1)
		case "right":
			panServoMove(panServo, panServo.CurrentAngle-1)
		case "up":
			tiltServoMove(tiltServo, tiltServo.CurrentAngle-1)
		case "down":
			tiltServoMove(tiltServo, tiltServo.CurrentAngle+1)
		}
		return fmt.Sprintf("This command pans the cctv %+v", params)
	})

	return robot
}
