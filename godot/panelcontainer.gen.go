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

//func NewPanelContainerFromPointer(ptr gdnative.Pointer) PanelContainer {
func newPanelContainerFromPointer(ptr gdnative.Pointer) PanelContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PanelContainer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Panel container type. This container fits controls inside of the delimited area of a stylebox. It's useful for giving controls an outline.
*/
type PanelContainer struct {
	Container
	owner gdnative.Object
}

func (o *PanelContainer) BaseClass() string {
	return "PanelContainer"
}

// PanelContainerImplementer is an interface that implements the methods
// of the PanelContainer class.
type PanelContainerImplementer interface {
	ContainerImplementer
}
