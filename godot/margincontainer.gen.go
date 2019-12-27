package godot

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewMarginContainerFromPointer(ptr gdnative.Pointer) MarginContainer {
func newMarginContainerFromPointer(ptr gdnative.Pointer) MarginContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MarginContainer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Adds a top, left, bottom, and right margin to all [Control] nodes that are direct children of the container. To control the [MarginContainer]'s margin, use the [code]margin_*[/code] theme properties listed below. [b]Note:[/b] Be careful, [Control] margin values are different than the constant margin values. If you want to change the custom margin values of the [MarginContainer] by code you should use the following examples: [codeblock] var margin_value = 100 set("custom_constants/margin_top", margin_value) set("custom_constants/margin_left", margin_value) set("custom_constants/margin_bottom", margin_value) set("custom_constants/margin_right", margin_value) [/codeblock]
*/
type MarginContainer struct {
	Container
	owner gdnative.Object
}

func (o *MarginContainer) BaseClass() string {
	return "MarginContainer"
}

// MarginContainerImplementer is an interface that implements the methods
// of the MarginContainer class.
type MarginContainerImplementer interface {
	ContainerImplementer
}
