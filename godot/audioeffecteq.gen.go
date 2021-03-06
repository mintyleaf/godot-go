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

//func NewAudioEffectEQFromPointer(ptr gdnative.Pointer) AudioEffectEQ {
func newAudioEffectEQFromPointer(ptr gdnative.Pointer) AudioEffectEQ {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectEQ{}
	obj.SetBaseObject(owner)

	return obj
}

/*
AudioEffectEQ gives you control over frequencies. Use it to compensate for existing deficiencies in audio. AudioEffectEQs are useful on the Master bus to completely master a mix and give it more character. They are also useful when a game is run on a mobile device, to adjust the mix to that kind of speakers (it can be added but disabled when headphones are plugged).
*/
type AudioEffectEQ struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectEQ) BaseClass() string {
	return "AudioEffectEQ"
}

/*
        Returns the number of bands of the equalizer.
	Args: [], Returns: int
*/
func (o *AudioEffectEQ) GetBandCount() gdnative.Int {
	//log.Println("Calling AudioEffectEQ.GetBandCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectEQ", "get_band_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the band's gain at the specified index, in dB.
	Args: [{ false band_idx int}], Returns: float
*/
func (o *AudioEffectEQ) GetBandGainDb(bandIdx gdnative.Int) gdnative.Real {
	//log.Println("Calling AudioEffectEQ.GetBandGainDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bandIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectEQ", "get_band_gain_db")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Sets band's gain at the specified index, in dB.
	Args: [{ false band_idx int} { false volume_db float}], Returns: void
*/
func (o *AudioEffectEQ) SetBandGainDb(bandIdx gdnative.Int, volumeDb gdnative.Real) {
	//log.Println("Calling AudioEffectEQ.SetBandGainDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bandIdx)
	ptrArguments[1] = gdnative.NewPointerFromReal(volumeDb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectEQ", "set_band_gain_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioEffectEQImplementer is an interface that implements the methods
// of the AudioEffectEQ class.
type AudioEffectEQImplementer interface {
	AudioEffectImplementer
	GetBandCount() gdnative.Int
	GetBandGainDb(bandIdx gdnative.Int) gdnative.Real
	SetBandGainDb(bandIdx gdnative.Int, volumeDb gdnative.Real)
}
