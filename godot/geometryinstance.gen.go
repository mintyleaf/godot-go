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

// GeometryInstanceFlags is an enum for Flags values.
type GeometryInstanceFlags int

const (
	GeometryInstanceFlagDrawNextFrameIfVisible GeometryInstanceFlags = 1
	GeometryInstanceFlagMax                    GeometryInstanceFlags = 2
	GeometryInstanceFlagUseBakedLight          GeometryInstanceFlags = 0
)

// GeometryInstanceShadowCastingSetting is an enum for ShadowCastingSetting values.
type GeometryInstanceShadowCastingSetting int

const (
	GeometryInstanceShadowCastingSettingDoubleSided GeometryInstanceShadowCastingSetting = 2
	GeometryInstanceShadowCastingSettingOff         GeometryInstanceShadowCastingSetting = 0
	GeometryInstanceShadowCastingSettingOn          GeometryInstanceShadowCastingSetting = 1
	GeometryInstanceShadowCastingSettingShadowsOnly GeometryInstanceShadowCastingSetting = 3
)

//func NewGeometryInstanceFromPointer(ptr gdnative.Pointer) GeometryInstance {
func newGeometryInstanceFromPointer(ptr gdnative.Pointer) GeometryInstance {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GeometryInstance{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base node for geometry based visual instances. Shares some common functionality like visibility and custom materials.
*/
type GeometryInstance struct {
	VisualInstance
	owner gdnative.Object
}

func (o *GeometryInstance) BaseClass() string {
	return "GeometryInstance"
}

/*
        Undocumented
	Args: [], Returns: enum.GeometryInstance::ShadowCastingSetting
*/
func (o *GeometryInstance) GetCastShadowsSetting() GeometryInstanceShadowCastingSetting {
	//log.Println("Calling GeometryInstance.GetCastShadowsSetting()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_cast_shadows_setting")

	// Call the parent method.
	// enum.GeometryInstance::ShadowCastingSetting
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return GeometryInstanceShadowCastingSetting(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *GeometryInstance) GetExtraCullMargin() gdnative.Real {
	//log.Println("Calling GeometryInstance.GetExtraCullMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_extra_cull_margin")

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
	Args: [{ false flag int}], Returns: bool
*/
func (o *GeometryInstance) GetFlag(flag gdnative.Int) gdnative.Bool {
	//log.Println("Calling GeometryInstance.GetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_flag")

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
	Args: [], Returns: float
*/
func (o *GeometryInstance) GetLodMaxDistance() gdnative.Real {
	//log.Println("Calling GeometryInstance.GetLodMaxDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_lod_max_distance")

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
func (o *GeometryInstance) GetLodMaxHysteresis() gdnative.Real {
	//log.Println("Calling GeometryInstance.GetLodMaxHysteresis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_lod_max_hysteresis")

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
func (o *GeometryInstance) GetLodMinDistance() gdnative.Real {
	//log.Println("Calling GeometryInstance.GetLodMinDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_lod_min_distance")

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
func (o *GeometryInstance) GetLodMinHysteresis() gdnative.Real {
	//log.Println("Calling GeometryInstance.GetLodMinHysteresis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_lod_min_hysteresis")

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
	Args: [], Returns: Material
*/
func (o *GeometryInstance) GetMaterialOverride() MaterialImplementer {
	//log.Println("Calling GeometryInstance.GetMaterialOverride()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "get_material_override")

	// Call the parent method.
	// Material
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMaterialFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MaterialImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Material" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(MaterialImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false shadow_casting_setting int}], Returns: void
*/
func (o *GeometryInstance) SetCastShadowsSetting(shadowCastingSetting gdnative.Int) {
	//log.Println("Calling GeometryInstance.SetCastShadowsSetting()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(shadowCastingSetting)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_cast_shadows_setting")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false aabb AABB}], Returns: void
*/
func (o *GeometryInstance) SetCustomAabb(aabb gdnative.Aabb) {
	//log.Println("Calling GeometryInstance.SetCustomAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromAabb(aabb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_custom_aabb")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false margin float}], Returns: void
*/
func (o *GeometryInstance) SetExtraCullMargin(margin gdnative.Real) {
	//log.Println("Calling GeometryInstance.SetExtraCullMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_extra_cull_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false flag int} { false value bool}], Returns: void
*/
func (o *GeometryInstance) SetFlag(flag gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling GeometryInstance.SetFlag()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_flag")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode float}], Returns: void
*/
func (o *GeometryInstance) SetLodMaxDistance(mode gdnative.Real) {
	//log.Println("Calling GeometryInstance.SetLodMaxDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_lod_max_distance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode float}], Returns: void
*/
func (o *GeometryInstance) SetLodMaxHysteresis(mode gdnative.Real) {
	//log.Println("Calling GeometryInstance.SetLodMaxHysteresis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_lod_max_hysteresis")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode float}], Returns: void
*/
func (o *GeometryInstance) SetLodMinDistance(mode gdnative.Real) {
	//log.Println("Calling GeometryInstance.SetLodMinDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_lod_min_distance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode float}], Returns: void
*/
func (o *GeometryInstance) SetLodMinHysteresis(mode gdnative.Real) {
	//log.Println("Calling GeometryInstance.SetLodMinHysteresis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_lod_min_hysteresis")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false material Material}], Returns: void
*/
func (o *GeometryInstance) SetMaterialOverride(material MaterialImplementer) {
	//log.Println("Calling GeometryInstance.SetMaterialOverride()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GeometryInstance", "set_material_override")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// GeometryInstanceImplementer is an interface that implements the methods
// of the GeometryInstance class.
type GeometryInstanceImplementer interface {
	VisualInstanceImplementer
	GetExtraCullMargin() gdnative.Real
	GetFlag(flag gdnative.Int) gdnative.Bool
	GetLodMaxDistance() gdnative.Real
	GetLodMaxHysteresis() gdnative.Real
	GetLodMinDistance() gdnative.Real
	GetLodMinHysteresis() gdnative.Real
	GetMaterialOverride() MaterialImplementer
	SetCastShadowsSetting(shadowCastingSetting gdnative.Int)
	SetCustomAabb(aabb gdnative.Aabb)
	SetExtraCullMargin(margin gdnative.Real)
	SetFlag(flag gdnative.Int, value gdnative.Bool)
	SetLodMaxDistance(mode gdnative.Real)
	SetLodMaxHysteresis(mode gdnative.Real)
	SetLodMinDistance(mode gdnative.Real)
	SetLodMinHysteresis(mode gdnative.Real)
	SetMaterialOverride(material MaterialImplementer)
}
