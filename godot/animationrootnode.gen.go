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

//func NewAnimationRootNodeFromPointer(ptr gdnative.Pointer) AnimationRootNode {
func newAnimationRootNodeFromPointer(ptr gdnative.Pointer) AnimationRootNode {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationRootNode{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AnimationRootNode struct {
	AnimationNode
	owner gdnative.Object
}

func (o *AnimationRootNode) BaseClass() string {
	return "AnimationRootNode"
}

// AnimationRootNodeImplementer is an interface that implements the methods
// of the AnimationRootNode class.
type AnimationRootNodeImplementer interface {
	AnimationNodeImplementer
}
