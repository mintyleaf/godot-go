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

//func NewVisualShaderNodeDotProductFromPointer(ptr gdnative.Pointer) VisualShaderNodeDotProduct {
func newVisualShaderNodeDotProductFromPointer(ptr gdnative.Pointer) VisualShaderNodeDotProduct {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeDotProduct{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Translates to [code]dot(a, b)[/code] in the shader language.
*/
type VisualShaderNodeDotProduct struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeDotProduct) BaseClass() string {
	return "VisualShaderNodeDotProduct"
}

// VisualShaderNodeDotProductImplementer is an interface that implements the methods
// of the VisualShaderNodeDotProduct class.
type VisualShaderNodeDotProductImplementer interface {
	VisualShaderNodeImplementer
}
