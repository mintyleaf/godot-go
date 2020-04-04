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

//func NewAnimationTrackEditPluginFromPointer(ptr gdnative.Pointer) AnimationTrackEditPlugin {
func newAnimationTrackEditPluginFromPointer(ptr gdnative.Pointer) AnimationTrackEditPlugin {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationTrackEditPlugin{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AnimationTrackEditPlugin struct {
	Reference
	owner gdnative.Object
}

func (o *AnimationTrackEditPlugin) BaseClass() string {
	return "AnimationTrackEditPlugin"
}

// AnimationTrackEditPluginImplementer is an interface that implements the methods
// of the AnimationTrackEditPlugin class.
type AnimationTrackEditPluginImplementer interface {
	ReferenceImplementer
}
