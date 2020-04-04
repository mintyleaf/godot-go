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

// DirectionalLightShadowDepthRange is an enum for ShadowDepthRange values.
type DirectionalLightShadowDepthRange int

const (
	DirectionalLightShadowDepthRangeOptimized DirectionalLightShadowDepthRange = 1
	DirectionalLightShadowDepthRangeStable    DirectionalLightShadowDepthRange = 0
)

// DirectionalLightShadowMode is an enum for ShadowMode values.
type DirectionalLightShadowMode int

const (
	DirectionalLightShadowOrthogonal      DirectionalLightShadowMode = 0
	DirectionalLightShadowParallel2Splits DirectionalLightShadowMode = 1
	DirectionalLightShadowParallel4Splits DirectionalLightShadowMode = 2
)

//func NewDirectionalLightFromPointer(ptr gdnative.Pointer) DirectionalLight {
func newDirectionalLightFromPointer(ptr gdnative.Pointer) DirectionalLight {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := DirectionalLight{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type DirectionalLight struct {
	Light
	owner gdnative.Object
}

func (o *DirectionalLight) BaseClass() string {
	return "DirectionalLight"
}

/*
        Undocumented
	Args: [], Returns: enum.DirectionalLight::ShadowDepthRange
*/
func (o *DirectionalLight) GetShadowDepthRange() DirectionalLightShadowDepthRange {
	//log.Println("Calling DirectionalLight.GetShadowDepthRange()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DirectionalLight", "get_shadow_depth_range")

	// Call the parent method.
	// enum.DirectionalLight::ShadowDepthRange
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return DirectionalLightShadowDepthRange(ret)
}

/*
        Undocumented
	Args: [], Returns: enum.DirectionalLight::ShadowMode
*/
func (o *DirectionalLight) GetShadowMode() DirectionalLightShadowMode {
	//log.Println("Calling DirectionalLight.GetShadowMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DirectionalLight", "get_shadow_mode")

	// Call the parent method.
	// enum.DirectionalLight::ShadowMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return DirectionalLightShadowMode(ret)
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *DirectionalLight) IsBlendSplitsEnabled() gdnative.Bool {
	//log.Println("Calling DirectionalLight.IsBlendSplitsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DirectionalLight", "is_blend_splits_enabled")

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
	Args: [{ false enabled bool}], Returns: void
*/
func (o *DirectionalLight) SetBlendSplits(enabled gdnative.Bool) {
	//log.Println("Calling DirectionalLight.SetBlendSplits()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DirectionalLight", "set_blend_splits")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *DirectionalLight) SetShadowDepthRange(mode gdnative.Int) {
	//log.Println("Calling DirectionalLight.SetShadowDepthRange()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DirectionalLight", "set_shadow_depth_range")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *DirectionalLight) SetShadowMode(mode gdnative.Int) {
	//log.Println("Calling DirectionalLight.SetShadowMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DirectionalLight", "set_shadow_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// DirectionalLightImplementer is an interface that implements the methods
// of the DirectionalLight class.
type DirectionalLightImplementer interface {
	LightImplementer
	IsBlendSplitsEnabled() gdnative.Bool
	SetBlendSplits(enabled gdnative.Bool)
	SetShadowDepthRange(mode gdnative.Int)
	SetShadowMode(mode gdnative.Int)
}
