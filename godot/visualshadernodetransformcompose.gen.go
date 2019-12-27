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

//func NewVisualShaderNodeTransformComposeFromPointer(ptr gdnative.Pointer) VisualShaderNodeTransformCompose {
func newVisualShaderNodeTransformComposeFromPointer(ptr gdnative.Pointer) VisualShaderNodeTransformCompose {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeTransformCompose{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeTransformCompose struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeTransformCompose) BaseClass() string {
	return "VisualShaderNodeTransformCompose"
}

// VisualShaderNodeTransformComposeImplementer is an interface that implements the methods
// of the VisualShaderNodeTransformCompose class.
type VisualShaderNodeTransformComposeImplementer interface {
	VisualShaderNodeImplementer
}
