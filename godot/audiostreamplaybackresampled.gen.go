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

//func NewAudioStreamPlaybackResampledFromPointer(ptr gdnative.Pointer) AudioStreamPlaybackResampled {
func newAudioStreamPlaybackResampledFromPointer(ptr gdnative.Pointer) AudioStreamPlaybackResampled {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStreamPlaybackResampled{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type AudioStreamPlaybackResampled struct {
	AudioStreamPlayback
	owner gdnative.Object
}

func (o *AudioStreamPlaybackResampled) BaseClass() string {
	return "AudioStreamPlaybackResampled"
}

// AudioStreamPlaybackResampledImplementer is an interface that implements the methods
// of the AudioStreamPlaybackResampled class.
type AudioStreamPlaybackResampledImplementer interface {
	AudioStreamPlaybackImplementer
}
