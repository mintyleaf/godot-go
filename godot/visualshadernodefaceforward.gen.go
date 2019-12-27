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

//func NewVisualShaderNodeFaceForwardFromPointer(ptr gdnative.Pointer) VisualShaderNodeFaceForward {
func newVisualShaderNodeFaceForwardFromPointer(ptr gdnative.Pointer) VisualShaderNodeFaceForward {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeFaceForward{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualShaderNodeFaceForward struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeFaceForward) BaseClass() string {
	return "VisualShaderNodeFaceForward"
}

// VisualShaderNodeFaceForwardImplementer is an interface that implements the methods
// of the VisualShaderNodeFaceForward class.
type VisualShaderNodeFaceForwardImplementer interface {
	VisualShaderNodeImplementer
}
