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

// ParticlesDrawOrder is an enum for DrawOrder values.
type ParticlesDrawOrder int

const (
	ParticlesDrawOrderIndex     ParticlesDrawOrder = 0
	ParticlesDrawOrderLifetime  ParticlesDrawOrder = 1
	ParticlesDrawOrderViewDepth ParticlesDrawOrder = 2
)

//func NewParticlesFromPointer(ptr gdnative.Pointer) Particles {
func newParticlesFromPointer(ptr gdnative.Pointer) Particles {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Particles{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type Particles struct {
	GeometryInstance
	owner gdnative.Object
}

func (o *Particles) BaseClass() string {
	return "Particles"
}

/*
        Undocumented
	Args: [], Returns: AABB
*/
func (o *Particles) CaptureAabb() gdnative.Aabb {
	//log.Println("Calling Particles.CaptureAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "capture_aabb")

	// Call the parent method.
	// AABB
	retPtr := gdnative.NewEmptyAabb()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewAabbFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Particles) GetAmount() gdnative.Int {
	//log.Println("Calling Particles.GetAmount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_amount")

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
	Args: [], Returns: enum.Particles::DrawOrder
*/
func (o *Particles) GetDrawOrder() ParticlesDrawOrder {
	//log.Println("Calling Particles.GetDrawOrder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_draw_order")

	// Call the parent method.
	// enum.Particles::DrawOrder
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ParticlesDrawOrder(ret)
}

/*
        Undocumented
	Args: [{ false pass int}], Returns: Mesh
*/
func (o *Particles) GetDrawPassMesh(pass gdnative.Int) MeshImplementer {
	//log.Println("Calling Particles.GetDrawPassMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(pass)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_draw_pass_mesh")

	// Call the parent method.
	// Mesh
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMeshFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MeshImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Mesh" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(MeshImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Particles) GetDrawPasses() gdnative.Int {
	//log.Println("Calling Particles.GetDrawPasses()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_draw_passes")

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
func (o *Particles) GetExplosivenessRatio() gdnative.Real {
	//log.Println("Calling Particles.GetExplosivenessRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_explosiveness_ratio")

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
	Args: [], Returns: int
*/
func (o *Particles) GetFixedFps() gdnative.Int {
	//log.Println("Calling Particles.GetFixedFps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_fixed_fps")

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
	Args: [], Returns: bool
*/
func (o *Particles) GetFractionalDelta() gdnative.Bool {
	//log.Println("Calling Particles.GetFractionalDelta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_fractional_delta")

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
func (o *Particles) GetLifetime() gdnative.Real {
	//log.Println("Calling Particles.GetLifetime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_lifetime")

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
func (o *Particles) GetOneShot() gdnative.Bool {
	//log.Println("Calling Particles.GetOneShot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_one_shot")

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
func (o *Particles) GetPreProcessTime() gdnative.Real {
	//log.Println("Calling Particles.GetPreProcessTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_pre_process_time")

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
func (o *Particles) GetProcessMaterial() MaterialImplementer {
	//log.Println("Calling Particles.GetProcessMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_process_material")

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
	Args: [], Returns: float
*/
func (o *Particles) GetRandomnessRatio() gdnative.Real {
	//log.Println("Calling Particles.GetRandomnessRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_randomness_ratio")

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
func (o *Particles) GetSpeedScale() gdnative.Real {
	//log.Println("Calling Particles.GetSpeedScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_speed_scale")

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
func (o *Particles) GetUseLocalCoordinates() gdnative.Bool {
	//log.Println("Calling Particles.GetUseLocalCoordinates()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_use_local_coordinates")

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
	Args: [], Returns: AABB
*/
func (o *Particles) GetVisibilityAabb() gdnative.Aabb {
	//log.Println("Calling Particles.GetVisibilityAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "get_visibility_aabb")

	// Call the parent method.
	// AABB
	retPtr := gdnative.NewEmptyAabb()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewAabbFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Particles) IsEmitting() gdnative.Bool {
	//log.Println("Calling Particles.IsEmitting()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "is_emitting")

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
	Args: [], Returns: void
*/
func (o *Particles) Restart() {
	//log.Println("Calling Particles.Restart()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "restart")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount int}], Returns: void
*/
func (o *Particles) SetAmount(amount gdnative.Int) {
	//log.Println("Calling Particles.SetAmount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_amount")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false order int}], Returns: void
*/
func (o *Particles) SetDrawOrder(order gdnative.Int) {
	//log.Println("Calling Particles.SetDrawOrder()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(order)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_draw_order")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pass int} { false mesh Mesh}], Returns: void
*/
func (o *Particles) SetDrawPassMesh(pass gdnative.Int, mesh MeshImplementer) {
	//log.Println("Calling Particles.SetDrawPassMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(pass)
	ptrArguments[1] = gdnative.NewPointerFromObject(mesh.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_draw_pass_mesh")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false passes int}], Returns: void
*/
func (o *Particles) SetDrawPasses(passes gdnative.Int) {
	//log.Println("Calling Particles.SetDrawPasses()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(passes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_draw_passes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false emitting bool}], Returns: void
*/
func (o *Particles) SetEmitting(emitting gdnative.Bool) {
	//log.Println("Calling Particles.SetEmitting()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(emitting)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_emitting")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ratio float}], Returns: void
*/
func (o *Particles) SetExplosivenessRatio(ratio gdnative.Real) {
	//log.Println("Calling Particles.SetExplosivenessRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(ratio)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_explosiveness_ratio")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false fps int}], Returns: void
*/
func (o *Particles) SetFixedFps(fps gdnative.Int) {
	//log.Println("Calling Particles.SetFixedFps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(fps)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_fixed_fps")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Particles) SetFractionalDelta(enable gdnative.Bool) {
	//log.Println("Calling Particles.SetFractionalDelta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_fractional_delta")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false secs float}], Returns: void
*/
func (o *Particles) SetLifetime(secs gdnative.Real) {
	//log.Println("Calling Particles.SetLifetime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(secs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_lifetime")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Particles) SetOneShot(enable gdnative.Bool) {
	//log.Println("Calling Particles.SetOneShot()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_one_shot")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false secs float}], Returns: void
*/
func (o *Particles) SetPreProcessTime(secs gdnative.Real) {
	//log.Println("Calling Particles.SetPreProcessTime()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(secs)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_pre_process_time")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false material Material}], Returns: void
*/
func (o *Particles) SetProcessMaterial(material MaterialImplementer) {
	//log.Println("Calling Particles.SetProcessMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_process_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ratio float}], Returns: void
*/
func (o *Particles) SetRandomnessRatio(ratio gdnative.Real) {
	//log.Println("Calling Particles.SetRandomnessRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(ratio)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_randomness_ratio")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false scale float}], Returns: void
*/
func (o *Particles) SetSpeedScale(scale gdnative.Real) {
	//log.Println("Calling Particles.SetSpeedScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(scale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_speed_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Particles) SetUseLocalCoordinates(enable gdnative.Bool) {
	//log.Println("Calling Particles.SetUseLocalCoordinates()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_use_local_coordinates")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false aabb AABB}], Returns: void
*/
func (o *Particles) SetVisibilityAabb(aabb gdnative.Aabb) {
	//log.Println("Calling Particles.SetVisibilityAabb()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromAabb(aabb)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Particles", "set_visibility_aabb")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ParticlesImplementer is an interface that implements the methods
// of the Particles class.
type ParticlesImplementer interface {
	GeometryInstanceImplementer
	CaptureAabb() gdnative.Aabb
	GetAmount() gdnative.Int
	GetDrawPassMesh(pass gdnative.Int) MeshImplementer
	GetDrawPasses() gdnative.Int
	GetExplosivenessRatio() gdnative.Real
	GetFixedFps() gdnative.Int
	GetFractionalDelta() gdnative.Bool
	GetLifetime() gdnative.Real
	GetOneShot() gdnative.Bool
	GetPreProcessTime() gdnative.Real
	GetProcessMaterial() MaterialImplementer
	GetRandomnessRatio() gdnative.Real
	GetSpeedScale() gdnative.Real
	GetUseLocalCoordinates() gdnative.Bool
	GetVisibilityAabb() gdnative.Aabb
	IsEmitting() gdnative.Bool
	Restart()
	SetAmount(amount gdnative.Int)
	SetDrawOrder(order gdnative.Int)
	SetDrawPassMesh(pass gdnative.Int, mesh MeshImplementer)
	SetDrawPasses(passes gdnative.Int)
	SetEmitting(emitting gdnative.Bool)
	SetExplosivenessRatio(ratio gdnative.Real)
	SetFixedFps(fps gdnative.Int)
	SetFractionalDelta(enable gdnative.Bool)
	SetLifetime(secs gdnative.Real)
	SetOneShot(enable gdnative.Bool)
	SetPreProcessTime(secs gdnative.Real)
	SetProcessMaterial(material MaterialImplementer)
	SetRandomnessRatio(ratio gdnative.Real)
	SetSpeedScale(scale gdnative.Real)
	SetUseLocalCoordinates(enable gdnative.Bool)
	SetVisibilityAabb(aabb gdnative.Aabb)
}
