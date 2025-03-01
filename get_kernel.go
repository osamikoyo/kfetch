package main

import (
	"fmt"

	"golang.org/x/sys/unix"
)

func GetKernel() string {
	var utsname unix.Utsname

	err := unix.Uname(&utsname)
	if err != nil {
		fmt.Println("Ошибка при получении информации о ядре:", err)
		return ""
	}

	kernelVersion := string(utsname.Release[:])
	return kernelVersion
}