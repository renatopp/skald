package inputs

import "github.com/hajimehoshi/ebiten/v2"

type GamepadButton struct {
	id   ebiten.GamepadID
	code ebiten.GamepadButton

	prevValue float64
	curValue  float64
}

func (k *GamepadButton) Id() int {
	return int(k.id)
}

func (k *GamepadButton) Code() int {
	return int(k.code)
}

func (k *GamepadButton) Name() string {
	return string(k.code)
}

func (k *GamepadButton) Device() Device {
	return DeviceGamepad
}

func (k *GamepadButton) IsDown() bool {
	return k.curValue > 0
}

func (k *GamepadButton) IsUp() bool {
	return k.curValue == 0
}

func (k *GamepadButton) IsPressed() bool {
	return k.curValue > 0 && k.prevValue == 0
}

func (k *GamepadButton) IsReleased() bool {
	return k.curValue == 0 && k.prevValue > 0
}

func (k *GamepadButton) GetBool() bool {
	return k.curValue > 0
}

func (k *GamepadButton) GetFloat() float64 {
	return k.curValue
}

func (k *GamepadButton) update() {
	k.prevValue = k.curValue
	if ebiten.IsGamepadButtonPressed(k.id, k.code) {
		k.curValue = 1
	} else {
		k.curValue = 0
	}
}
