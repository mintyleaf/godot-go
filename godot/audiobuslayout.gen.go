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

//func NewAudioBusLayoutFromPointer(ptr gdnative.Pointer) AudioBusLayout {
func newAudioBusLayoutFromPointer(ptr gdnative.Pointer) AudioBusLayout {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioBusLayout{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Stores position, muting, solo, bypass, effects, effect position, volume, and the connections between buses. See [AudioServer] for usage.
*/
type AudioBusLayout struct {
	Resource
	owner gdnative.Object
}

func (o *AudioBusLayout) BaseClass() string {
	return "AudioBusLayout"
}

// AudioBusLayoutImplementer is an interface that implements the methods
// of the AudioBusLayout class.
type AudioBusLayoutImplementer interface {
	ResourceImplementer
}
