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

//func NewAudioEffectNotchFilterFromPointer(ptr gdnative.Pointer) AudioEffectNotchFilter {
func newAudioEffectNotchFilterFromPointer(ptr gdnative.Pointer) AudioEffectNotchFilter {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectNotchFilter{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Attenuates frequencies in a narrow band around the [member AudioEffectFilter.cutoff_hz] and cuts frequencies outside of this range.
*/
type AudioEffectNotchFilter struct {
	AudioEffectFilter
	owner gdnative.Object
}

func (o *AudioEffectNotchFilter) BaseClass() string {
	return "AudioEffectNotchFilter"
}

// AudioEffectNotchFilterImplementer is an interface that implements the methods
// of the AudioEffectNotchFilter class.
type AudioEffectNotchFilterImplementer interface {
	AudioEffectFilterImplementer
}
