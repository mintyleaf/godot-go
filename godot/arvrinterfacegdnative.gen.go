package godot

import (
	"github.com/gabstv/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewARVRInterfaceGDNativeFromPointer(ptr gdnative.Pointer) ARVRInterfaceGDNative {
func newARVRInterfaceGDNativeFromPointer(ptr gdnative.Pointer) ARVRInterfaceGDNative {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ARVRInterfaceGDNative{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type ARVRInterfaceGDNative struct {
	ARVRInterface
	owner gdnative.Object
}

func (o *ARVRInterfaceGDNative) BaseClass() string {
	return "ARVRInterfaceGDNative"
}

// ARVRInterfaceGDNativeImplementer is an interface that implements the methods
// of the ARVRInterfaceGDNative class.
type ARVRInterfaceGDNativeImplementer interface {
	ARVRInterfaceImplementer
}
