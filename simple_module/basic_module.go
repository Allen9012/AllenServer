package simple_module

import (
	"fmt"
	"github.com/Allen9012/AllenGame/service"
)

func init() {

}

//module其实就是serivce

type Module1 struct {
	service.Module
}

func (slf *Module1) OnInit() error {
	fmt.Printf("Module1 OnInit.\n")
	return nil
}
func (slf *Module1) OnRelease() {
	fmt.Printf("Module1 Release.\n")
}

type Module2 struct {
	service.Module
}

func (slf *Module2) OnInit() error {
	fmt.Printf("Module2 OnInit.\n")
	return nil
}

func (slf *Module2) OnRelease() {
	fmt.Printf("Module2 Release.\n")
}

type Module3 struct {
	service.Module
}

func (slf *Module3) OnInit() error {
	fmt.Printf("Module3 OnInit.\n")
	return nil
}

func (slf *Module3) OnRelease() {
	fmt.Printf("Module3 Release.\n")
}
