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

//func NewInputDefaultFromPointer(ptr gdnative.Pointer) InputDefault {
func newInputDefaultFromPointer(ptr gdnative.Pointer) InputDefault {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputDefault{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type InputDefault struct {
	input
	owner gdnative.Object
}

func (o *InputDefault) BaseClass() string {
	return "InputDefault"
}

// InputDefaultImplementer is an interface that implements the methods
// of the InputDefault class.
type InputDefaultImplementer interface {
	InputImplementer
}
