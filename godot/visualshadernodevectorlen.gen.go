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

//func NewVisualShaderNodeVectorLenFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorLen {
func newVisualShaderNodeVectorLenFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorLen {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeVectorLen{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeVectorLen struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeVectorLen) BaseClass() string {
	return "VisualShaderNodeVectorLen"
}

// VisualShaderNodeVectorLenImplementer is an interface that implements the methods
// of the VisualShaderNodeVectorLen class.
type VisualShaderNodeVectorLenImplementer interface {
	VisualShaderNodeImplementer
}
