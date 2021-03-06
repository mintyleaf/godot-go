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

//func NewVideoStreamFromPointer(ptr gdnative.Pointer) VideoStream {
func newVideoStreamFromPointer(ptr gdnative.Pointer) VideoStream {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VideoStream{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base resource type for all video streams. Classes that derive from [VideoStream] can all be used as resource types to play back videos in [VideoPlayer].
*/
type VideoStream struct {
	Resource
	owner gdnative.Object
}

func (o *VideoStream) BaseClass() string {
	return "VideoStream"
}

// VideoStreamImplementer is an interface that implements the methods
// of the VideoStream class.
type VideoStreamImplementer interface {
	ResourceImplementer
}
