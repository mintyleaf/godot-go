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

//func NewVisualShaderNodeVectorScalarMixFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorScalarMix {
func newVisualShaderNodeVectorScalarMixFromPointer(ptr gdnative.Pointer) VisualShaderNodeVectorScalarMix {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeVectorScalarMix{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeVectorScalarMix struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeVectorScalarMix) BaseClass() string {
	return "VisualShaderNodeVectorScalarMix"
}

// VisualShaderNodeVectorScalarMixImplementer is an interface that implements the methods
// of the VisualShaderNodeVectorScalarMix class.
type VisualShaderNodeVectorScalarMixImplementer interface {
	VisualShaderNodeImplementer
}
