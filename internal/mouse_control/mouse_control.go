package mouse_control

import (
	"computer-control/internal/handle"
	"github.com/go-vgo/robotgo"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func SetMousePosition(r *http.Request) handle.Response {

	vars := mux.Vars(r)
	x, errorX := vars["x"]
	if errorX == false {
		return handle.Response{HttpCode: 400, Message: "Invalid x position param"}
	}

	y, errorY := vars["y"]
	if errorY == false {
		return handle.Response{HttpCode: 400, Message: "Invalid y position param"}
	}

	positionX, _ := strconv.Atoi(x)
	positionY, _ := strconv.Atoi(y)

	robotgo.MoveMouse(positionX, positionY)

	return handle.Response{HttpCode: 200, Message: "Position of mouse updated"}
}

func ClickLeftMouse(r *http.Request) handle.Response {

	vars := mux.Vars(r)
	doubleClick, _ := strconv.ParseBool(vars["doubleClick"])

	robotgo.MouseClick("left", doubleClick)

	return handle.Response{HttpCode: 200, Message: "Left mouse clicked"}
}

func ClickRightMouse(r *http.Request) handle.Response {

	robotgo.MouseClick("right", false)

	return handle.Response{HttpCode: 200, Message: "Right mouse clicked"}
}