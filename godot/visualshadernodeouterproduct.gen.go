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

//func NewVisualShaderNodeOuterProductFromPointer(ptr gdnative.Pointer) VisualShaderNodeOuterProduct {
func newVisualShaderNodeOuterProductFromPointer(ptr gdnative.Pointer) VisualShaderNodeOuterProduct {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeOuterProduct{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualShaderNodeOuterProduct struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeOuterProduct) BaseClass() string {
	return "VisualShaderNodeOuterProduct"
}

// VisualShaderNodeOuterProductImplementer is an interface that implements the methods
// of the VisualShaderNodeOuterProduct class.
type VisualShaderNodeOuterProductImplementer interface {
	VisualShaderNodeImplementer
}
