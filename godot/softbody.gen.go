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

//func NewSoftBodyFromPointer(ptr gdnative.Pointer) SoftBody {
func newSoftBodyFromPointer(ptr gdnative.Pointer) SoftBody {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SoftBody{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type SoftBody struct {
	MeshInstance
	owner gdnative.Object
}

func (o *SoftBody) BaseClass() string {
	return "SoftBody"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *SoftBody) X_DrawSoftMesh() {
	//log.Println("Calling SoftBody.X_DrawSoftMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "_draw_soft_mesh")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false body Node}], Returns: void
*/
func (o *SoftBody) AddCollisionExceptionWith(body NodeImplementer) {
	//log.Println("Calling SoftBody.AddCollisionExceptionWith()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(body.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "add_collision_exception_with")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *SoftBody) GetAreaAngularStiffness() gdnative.Real {
	//log.Println("Calling SoftBody.GetAreaAngularStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_areaAngular_stiffness")

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
	Args: [], Returns: Array
*/
func (o *SoftBody) GetCollisionExceptions() gdnative.Array {
	//log.Println("Calling SoftBody.GetCollisionExceptions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_collision_exceptions")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *SoftBody) GetCollisionLayer() gdnative.Int {
	//log.Println("Calling SoftBody.GetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_collision_layer")

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
	Args: [{ false bit int}], Returns: bool
*/
func (o *SoftBody) GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling SoftBody.GetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_collision_layer_bit")

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
	Args: [], Returns: int
*/
func (o *SoftBody) GetCollisionMask() gdnative.Int {
	//log.Println("Calling SoftBody.GetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_collision_mask")

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
	Args: [{ false bit int}], Returns: bool
*/
func (o *SoftBody) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling SoftBody.GetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_collision_mask_bit")

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
func (o *SoftBody) GetDampingCoefficient() gdnative.Real {
	//log.Println("Calling SoftBody.GetDampingCoefficient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_damping_coefficient")

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
func (o *SoftBody) GetDragCoefficient() gdnative.Real {
	//log.Println("Calling SoftBody.GetDragCoefficient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_drag_coefficient")

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
func (o *SoftBody) GetLinearStiffness() gdnative.Real {
	//log.Println("Calling SoftBody.GetLinearStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_linear_stiffness")

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
	Args: [], Returns: NodePath
*/
func (o *SoftBody) GetParentCollisionIgnore() gdnative.NodePath {
	//log.Println("Calling SoftBody.GetParentCollisionIgnore()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_parent_collision_ignore")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *SoftBody) GetPoseMatchingCoefficient() gdnative.Real {
	//log.Println("Calling SoftBody.GetPoseMatchingCoefficient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_pose_matching_coefficient")

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
func (o *SoftBody) GetPressureCoefficient() gdnative.Real {
	//log.Println("Calling SoftBody.GetPressureCoefficient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_pressure_coefficient")

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
func (o *SoftBody) GetSimulationPrecision() gdnative.Int {
	//log.Println("Calling SoftBody.GetSimulationPrecision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_simulation_precision")

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
func (o *SoftBody) GetTotalMass() gdnative.Real {
	//log.Println("Calling SoftBody.GetTotalMass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_total_mass")

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
func (o *SoftBody) GetVolumeStiffness() gdnative.Real {
	//log.Println("Calling SoftBody.GetVolumeStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "get_volume_stiffness")

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
func (o *SoftBody) IsRayPickable() gdnative.Bool {
	//log.Println("Calling SoftBody.IsRayPickable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "is_ray_pickable")

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
	Args: [{ false body Node}], Returns: void
*/
func (o *SoftBody) RemoveCollisionExceptionWith(body NodeImplementer) {
	//log.Println("Calling SoftBody.RemoveCollisionExceptionWith()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(body.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "remove_collision_exception_with")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false areaAngular_stiffness float}], Returns: void
*/
func (o *SoftBody) SetAreaAngularStiffness(areaAngularStiffness gdnative.Real) {
	//log.Println("Calling SoftBody.SetAreaAngularStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(areaAngularStiffness)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_areaAngular_stiffness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false collision_layer int}], Returns: void
*/
func (o *SoftBody) SetCollisionLayer(collisionLayer gdnative.Int) {
	//log.Println("Calling SoftBody.SetCollisionLayer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(collisionLayer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_collision_layer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *SoftBody) SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling SoftBody.SetCollisionLayerBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_collision_layer_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false collision_mask int}], Returns: void
*/
func (o *SoftBody) SetCollisionMask(collisionMask gdnative.Int) {
	//log.Println("Calling SoftBody.SetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(collisionMask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_collision_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *SoftBody) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling SoftBody.SetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_collision_mask_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false damping_coefficient float}], Returns: void
*/
func (o *SoftBody) SetDampingCoefficient(dampingCoefficient gdnative.Real) {
	//log.Println("Calling SoftBody.SetDampingCoefficient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(dampingCoefficient)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_damping_coefficient")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false drag_coefficient float}], Returns: void
*/
func (o *SoftBody) SetDragCoefficient(dragCoefficient gdnative.Real) {
	//log.Println("Calling SoftBody.SetDragCoefficient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(dragCoefficient)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_drag_coefficient")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false linear_stiffness float}], Returns: void
*/
func (o *SoftBody) SetLinearStiffness(linearStiffness gdnative.Real) {
	//log.Println("Calling SoftBody.SetLinearStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(linearStiffness)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_linear_stiffness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false parent_collision_ignore NodePath}], Returns: void
*/
func (o *SoftBody) SetParentCollisionIgnore(parentCollisionIgnore gdnative.NodePath) {
	//log.Println("Calling SoftBody.SetParentCollisionIgnore()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(parentCollisionIgnore)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_parent_collision_ignore")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pose_matching_coefficient float}], Returns: void
*/
func (o *SoftBody) SetPoseMatchingCoefficient(poseMatchingCoefficient gdnative.Real) {
	//log.Println("Calling SoftBody.SetPoseMatchingCoefficient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(poseMatchingCoefficient)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_pose_matching_coefficient")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false pressure_coefficient float}], Returns: void
*/
func (o *SoftBody) SetPressureCoefficient(pressureCoefficient gdnative.Real) {
	//log.Println("Calling SoftBody.SetPressureCoefficient()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(pressureCoefficient)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_pressure_coefficient")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false ray_pickable bool}], Returns: void
*/
func (o *SoftBody) SetRayPickable(rayPickable gdnative.Bool) {
	//log.Println("Calling SoftBody.SetRayPickable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(rayPickable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_ray_pickable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false simulation_precision int}], Returns: void
*/
func (o *SoftBody) SetSimulationPrecision(simulationPrecision gdnative.Int) {
	//log.Println("Calling SoftBody.SetSimulationPrecision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(simulationPrecision)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_simulation_precision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mass float}], Returns: void
*/
func (o *SoftBody) SetTotalMass(mass gdnative.Real) {
	//log.Println("Calling SoftBody.SetTotalMass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(mass)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_total_mass")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false volume_stiffness float}], Returns: void
*/
func (o *SoftBody) SetVolumeStiffness(volumeStiffness gdnative.Real) {
	//log.Println("Calling SoftBody.SetVolumeStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(volumeStiffness)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SoftBody", "set_volume_stiffness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// SoftBodyImplementer is an interface that implements the methods
// of the SoftBody class.
type SoftBodyImplementer interface {
	MeshInstanceImplementer
	X_DrawSoftMesh()
	AddCollisionExceptionWith(body NodeImplementer)
	GetAreaAngularStiffness() gdnative.Real
	GetCollisionExceptions() gdnative.Array
	GetCollisionLayer() gdnative.Int
	GetCollisionLayerBit(bit gdnative.Int) gdnative.Bool
	GetCollisionMask() gdnative.Int
	GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool
	GetDampingCoefficient() gdnative.Real
	GetDragCoefficient() gdnative.Real
	GetLinearStiffness() gdnative.Real
	GetParentCollisionIgnore() gdnative.NodePath
	GetPoseMatchingCoefficient() gdnative.Real
	GetPressureCoefficient() gdnative.Real
	GetSimulationPrecision() gdnative.Int
	GetTotalMass() gdnative.Real
	GetVolumeStiffness() gdnative.Real
	IsRayPickable() gdnative.Bool
	RemoveCollisionExceptionWith(body NodeImplementer)
	SetAreaAngularStiffness(areaAngularStiffness gdnative.Real)
	SetCollisionLayer(collisionLayer gdnative.Int)
	SetCollisionLayerBit(bit gdnative.Int, value gdnative.Bool)
	SetCollisionMask(collisionMask gdnative.Int)
	SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool)
	SetDampingCoefficient(dampingCoefficient gdnative.Real)
	SetDragCoefficient(dragCoefficient gdnative.Real)
	SetLinearStiffness(linearStiffness gdnative.Real)
	SetParentCollisionIgnore(parentCollisionIgnore gdnative.NodePath)
	SetPoseMatchingCoefficient(poseMatchingCoefficient gdnative.Real)
	SetPressureCoefficient(pressureCoefficient gdnative.Real)
	SetRayPickable(rayPickable gdnative.Bool)
	SetSimulationPrecision(simulationPrecision gdnative.Int)
	SetTotalMass(mass gdnative.Real)
	SetVolumeStiffness(volumeStiffness gdnative.Real)
}
