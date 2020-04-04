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

//func NewStyleBoxEmptyFromPointer(ptr gdnative.Pointer) StyleBoxEmpty {
func newStyleBoxEmptyFromPointer(ptr gdnative.Pointer) StyleBoxEmpty {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := StyleBoxEmpty{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Empty stylebox (really does not display anything).
*/
type StyleBoxEmpty struct {
	StyleBox
	owner gdnative.Object
}

func (o *StyleBoxEmpty) BaseClass() string {
	return "StyleBoxEmpty"
}

// StyleBoxEmptyImplementer is an interface that implements the methods
// of the StyleBoxEmpty class.
type StyleBoxEmptyImplementer interface {
	StyleBoxImplementer
}
