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

//func NewAnimationNodeTimeScaleFromPointer(ptr gdnative.Pointer) AnimationNodeTimeScale {
func newAnimationNodeTimeScaleFromPointer(ptr gdnative.Pointer) AnimationNodeTimeScale {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeTimeScale{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Allows scaling the speed of the animation (or reversing it) in any children nodes. Setting it to 0 will pause the animation.
*/
type AnimationNodeTimeScale struct {
	AnimationNode
	owner gdnative.Object
}

func (o *AnimationNodeTimeScale) BaseClass() string {
	return "AnimationNodeTimeScale"
}

// AnimationNodeTimeScaleImplementer is an interface that implements the methods
// of the AnimationNodeTimeScale class.
type AnimationNodeTimeScaleImplementer interface {
	AnimationNodeImplementer
}
