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

//func NewAudioEffectPhaserFromPointer(ptr gdnative.Pointer) AudioEffectPhaser {
func newAudioEffectPhaserFromPointer(ptr gdnative.Pointer) AudioEffectPhaser {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffectPhaser{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Combines phase-shifted signals with the original signal. The movement of the phase-shifted signals is controlled using a low-frequency oscillator.
*/
type AudioEffectPhaser struct {
	AudioEffect
	owner gdnative.Object
}

func (o *AudioEffectPhaser) BaseClass() string {
	return "AudioEffectPhaser"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectPhaser) GetDepth() gdnative.Real {
	//log.Println("Calling AudioEffectPhaser.GetDepth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "get_depth")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectPhaser) GetFeedback() gdnative.Real {
	//log.Println("Calling AudioEffectPhaser.GetFeedback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "get_feedback")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectPhaser) GetRangeMaxHz() gdnative.Real {
	//log.Println("Calling AudioEffectPhaser.GetRangeMaxHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "get_range_max_hz")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectPhaser) GetRangeMinHz() gdnative.Real {
	//log.Println("Calling AudioEffectPhaser.GetRangeMinHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "get_range_min_hz")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioEffectPhaser) GetRateHz() gdnative.Real {
	//log.Println("Calling AudioEffectPhaser.GetRateHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "get_rate_hz")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false depth float}], Returns: void
*/
func (o *AudioEffectPhaser) SetDepth(depth gdnative.Real) {
	//log.Println("Calling AudioEffectPhaser.SetDepth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(depth)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "set_depth")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false fbk float}], Returns: void
*/
func (o *AudioEffectPhaser) SetFeedback(fbk gdnative.Real) {
	//log.Println("Calling AudioEffectPhaser.SetFeedback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(fbk)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "set_feedback")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false hz float}], Returns: void
*/
func (o *AudioEffectPhaser) SetRangeMaxHz(hz gdnative.Real) {
	//log.Println("Calling AudioEffectPhaser.SetRangeMaxHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(hz)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "set_range_max_hz")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false hz float}], Returns: void
*/
func (o *AudioEffectPhaser) SetRangeMinHz(hz gdnative.Real) {
	//log.Println("Calling AudioEffectPhaser.SetRangeMinHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(hz)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "set_range_min_hz")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false hz float}], Returns: void
*/
func (o *AudioEffectPhaser) SetRateHz(hz gdnative.Real) {
	//log.Println("Calling AudioEffectPhaser.SetRateHz()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(hz)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioEffectPhaser", "set_rate_hz")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioEffectPhaserImplementer is an interface that implements the methods
// of the AudioEffectPhaser class.
type AudioEffectPhaserImplementer interface {
	AudioEffectImplementer
	GetDepth() gdnative.Real
	GetFeedback() gdnative.Real
	GetRangeMaxHz() gdnative.Real
	GetRangeMinHz() gdnative.Real
	GetRateHz() gdnative.Real
	SetDepth(depth gdnative.Real)
	SetFeedback(fbk gdnative.Real)
	SetRangeMaxHz(hz gdnative.Real)
	SetRangeMinHz(hz gdnative.Real)
	SetRateHz(hz gdnative.Real)
}
