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

//func NewAudioStreamPlayer2DFromPointer(ptr gdnative.Pointer) AudioStreamPlayer2D {
func newAudioStreamPlayer2DFromPointer(ptr gdnative.Pointer) AudioStreamPlayer2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioStreamPlayer2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Plays audio that dampens with distance from screen center.
*/
type AudioStreamPlayer2D struct {
	Node2D
	owner gdnative.Object
}

func (o *AudioStreamPlayer2D) BaseClass() string {
	return "AudioStreamPlayer2D"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *AudioStreamPlayer2D) X_BusLayoutChanged() {
	//log.Println("Calling AudioStreamPlayer2D.X_BusLayoutChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "_bus_layout_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AudioStreamPlayer2D) X_IsActive() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer2D.X_IsActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "_is_active")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *AudioStreamPlayer2D) X_SetPlaying(enable gdnative.Bool) {
	//log.Println("Calling AudioStreamPlayer2D.X_SetPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "_set_playing")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *AudioStreamPlayer2D) GetAreaMask() gdnative.Int {
	//log.Println("Calling AudioStreamPlayer2D.GetAreaMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_area_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer2D) GetAttenuation() gdnative.Real {
	//log.Println("Calling AudioStreamPlayer2D.GetAttenuation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_attenuation")

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
	Args: [], Returns: String
*/
func (o *AudioStreamPlayer2D) GetBus() gdnative.String {
	//log.Println("Calling AudioStreamPlayer2D.GetBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_bus")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer2D) GetMaxDistance() gdnative.Real {
	//log.Println("Calling AudioStreamPlayer2D.GetMaxDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_max_distance")

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
func (o *AudioStreamPlayer2D) GetPitchScale() gdnative.Real {
	//log.Println("Calling AudioStreamPlayer2D.GetPitchScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_pitch_scale")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the position in the [AudioStream].
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer2D) GetPlaybackPosition() gdnative.Real {
	//log.Println("Calling AudioStreamPlayer2D.GetPlaybackPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_playback_position")

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
	Args: [], Returns: AudioStream
*/
func (o *AudioStreamPlayer2D) GetStream() AudioStreamImplementer {
	//log.Println("Calling AudioStreamPlayer2D.GetStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_stream")

	// Call the parent method.
	// AudioStream
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAudioStreamFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AudioStreamImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AudioStream" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AudioStreamImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AudioStreamPlayer2D) GetStreamPaused() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer2D.GetStreamPaused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_stream_paused")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns the [AudioStreamPlayback] object associated with this [AudioStreamPlayer2D].
	Args: [], Returns: AudioStreamPlayback
*/
func (o *AudioStreamPlayer2D) GetStreamPlayback() AudioStreamPlaybackImplementer {
	//log.Println("Calling AudioStreamPlayer2D.GetStreamPlayback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_stream_playback")

	// Call the parent method.
	// AudioStreamPlayback
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newAudioStreamPlaybackFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(AudioStreamPlaybackImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "AudioStreamPlayback" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(AudioStreamPlaybackImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *AudioStreamPlayer2D) GetVolumeDb() gdnative.Real {
	//log.Println("Calling AudioStreamPlayer2D.GetVolumeDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "get_volume_db")

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
	Args: [], Returns: bool
*/
func (o *AudioStreamPlayer2D) IsAutoplayEnabled() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer2D.IsAutoplayEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "is_autoplay_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AudioStreamPlayer2D) IsPlaying() gdnative.Bool {
	//log.Println("Calling AudioStreamPlayer2D.IsPlaying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "is_playing")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Plays the audio from the given position [code]from_position[/code], in seconds.
	Args: [{0 true from_position float}], Returns: void
*/
func (o *AudioStreamPlayer2D) Play(fromPosition gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer2D.Play()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(fromPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "play")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the position from which audio will be played, in seconds.
	Args: [{ false to_position float}], Returns: void
*/
func (o *AudioStreamPlayer2D) Seek(toPosition gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer2D.Seek()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(toPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "seek")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *AudioStreamPlayer2D) SetAreaMask(mask gdnative.Int) {
	//log.Println("Calling AudioStreamPlayer2D.SetAreaMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "set_area_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false curve float}], Returns: void
*/
func (o *AudioStreamPlayer2D) SetAttenuation(curve gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer2D.SetAttenuation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(curve)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "set_attenuation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *AudioStreamPlayer2D) SetAutoplay(enable gdnative.Bool) {
	//log.Println("Calling AudioStreamPlayer2D.SetAutoplay()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "set_autoplay")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bus String}], Returns: void
*/
func (o *AudioStreamPlayer2D) SetBus(bus gdnative.String) {
	//log.Println("Calling AudioStreamPlayer2D.SetBus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(bus)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "set_bus")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pixels float}], Returns: void
*/
func (o *AudioStreamPlayer2D) SetMaxDistance(pixels gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer2D.SetMaxDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(pixels)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "set_max_distance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pitch_scale float}], Returns: void
*/
func (o *AudioStreamPlayer2D) SetPitchScale(pitchScale gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer2D.SetPitchScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(pitchScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "set_pitch_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false stream AudioStream}], Returns: void
*/
func (o *AudioStreamPlayer2D) SetStream(stream AudioStreamImplementer) {
	//log.Println("Calling AudioStreamPlayer2D.SetStream()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(stream.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "set_stream")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pause bool}], Returns: void
*/
func (o *AudioStreamPlayer2D) SetStreamPaused(pause gdnative.Bool) {
	//log.Println("Calling AudioStreamPlayer2D.SetStreamPaused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(pause)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "set_stream_paused")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false volume_db float}], Returns: void
*/
func (o *AudioStreamPlayer2D) SetVolumeDb(volumeDb gdnative.Real) {
	//log.Println("Calling AudioStreamPlayer2D.SetVolumeDb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(volumeDb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "set_volume_db")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Stops the audio.
	Args: [], Returns: void
*/
func (o *AudioStreamPlayer2D) Stop() {
	//log.Println("Calling AudioStreamPlayer2D.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AudioStreamPlayer2D", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AudioStreamPlayer2DImplementer is an interface that implements the methods
// of the AudioStreamPlayer2D class.
type AudioStreamPlayer2DImplementer interface {
	Node2DImplementer
	X_BusLayoutChanged()
	X_IsActive() gdnative.Bool
	X_SetPlaying(enable gdnative.Bool)
	GetAreaMask() gdnative.Int
	GetAttenuation() gdnative.Real
	GetBus() gdnative.String
	GetMaxDistance() gdnative.Real
	GetPitchScale() gdnative.Real
	GetPlaybackPosition() gdnative.Real
	GetStream() AudioStreamImplementer
	GetStreamPaused() gdnative.Bool
	GetStreamPlayback() AudioStreamPlaybackImplementer
	GetVolumeDb() gdnative.Real
	IsAutoplayEnabled() gdnative.Bool
	IsPlaying() gdnative.Bool
	Play(fromPosition gdnative.Real)
	Seek(toPosition gdnative.Real)
	SetAreaMask(mask gdnative.Int)
	SetAttenuation(curve gdnative.Real)
	SetAutoplay(enable gdnative.Bool)
	SetBus(bus gdnative.String)
	SetMaxDistance(pixels gdnative.Real)
	SetPitchScale(pitchScale gdnative.Real)
	SetStream(stream AudioStreamImplementer)
	SetStreamPaused(pause gdnative.Bool)
	SetVolumeDb(volumeDb gdnative.Real)
	Stop()
}
