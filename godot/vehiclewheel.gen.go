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

//func NewVehicleWheelFromPointer(ptr gdnative.Pointer) VehicleWheel {
func newVehicleWheelFromPointer(ptr gdnative.Pointer) VehicleWheel {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VehicleWheel{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This node needs to be used as a child node of [VehicleBody] and simulates the behavior of one of its wheels. This node also acts as a collider to detect if the wheel is touching a surface.
*/
type VehicleWheel struct {
	Spatial
	owner gdnative.Object
}

func (o *VehicleWheel) BaseClass() string {
	return "VehicleWheel"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *VehicleWheel) GetBrake() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetBrake()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_brake")

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
func (o *VehicleWheel) GetDampingCompression() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetDampingCompression()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_damping_compression")

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
func (o *VehicleWheel) GetDampingRelaxation() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetDampingRelaxation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_damping_relaxation")

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
func (o *VehicleWheel) GetEngineForce() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetEngineForce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_engine_force")

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
func (o *VehicleWheel) GetFrictionSlip() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetFrictionSlip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_friction_slip")

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
func (o *VehicleWheel) GetRadius() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_radius")

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
func (o *VehicleWheel) GetRollInfluence() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetRollInfluence()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_roll_influence")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the rotational speed of the wheel in revolutions per minute.
	Args: [], Returns: float
*/
func (o *VehicleWheel) GetRpm() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetRpm()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_rpm")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns a value between 0.0 and 1.0 that indicates whether this wheel is skidding. 0.0 is skidding (the wheel has lost grip, e.g. icy terrain), 1.0 means not skidding (the wheel has full grip, e.g. dry asphalt road).
	Args: [], Returns: float
*/
func (o *VehicleWheel) GetSkidinfo() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetSkidinfo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_skidinfo")

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
func (o *VehicleWheel) GetSteering() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetSteering()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_steering")

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
func (o *VehicleWheel) GetSuspensionMaxForce() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetSuspensionMaxForce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_suspension_max_force")

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
func (o *VehicleWheel) GetSuspensionRestLength() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetSuspensionRestLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_suspension_rest_length")

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
func (o *VehicleWheel) GetSuspensionStiffness() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetSuspensionStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_suspension_stiffness")

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
func (o *VehicleWheel) GetSuspensionTravel() gdnative.Real {
	//log.Println("Calling VehicleWheel.GetSuspensionTravel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "get_suspension_travel")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if this wheel is in contact with a surface.
	Args: [], Returns: bool
*/
func (o *VehicleWheel) IsInContact() gdnative.Bool {
	//log.Println("Calling VehicleWheel.IsInContact()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "is_in_contact")

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
func (o *VehicleWheel) IsUsedAsSteering() gdnative.Bool {
	//log.Println("Calling VehicleWheel.IsUsedAsSteering()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "is_used_as_steering")

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
func (o *VehicleWheel) IsUsedAsTraction() gdnative.Bool {
	//log.Println("Calling VehicleWheel.IsUsedAsTraction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "is_used_as_traction")

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
	Args: [{ false brake float}], Returns: void
*/
func (o *VehicleWheel) SetBrake(brake gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetBrake()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(brake)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_brake")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *VehicleWheel) SetDampingCompression(length gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetDampingCompression()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_damping_compression")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *VehicleWheel) SetDampingRelaxation(length gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetDampingRelaxation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_damping_relaxation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false engine_force float}], Returns: void
*/
func (o *VehicleWheel) SetEngineForce(engineForce gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetEngineForce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(engineForce)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_engine_force")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *VehicleWheel) SetFrictionSlip(length gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetFrictionSlip()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_friction_slip")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *VehicleWheel) SetRadius(length gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetRadius()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_radius")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false roll_influence float}], Returns: void
*/
func (o *VehicleWheel) SetRollInfluence(rollInfluence gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetRollInfluence()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(rollInfluence)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_roll_influence")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false steering float}], Returns: void
*/
func (o *VehicleWheel) SetSteering(steering gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetSteering()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(steering)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_steering")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *VehicleWheel) SetSuspensionMaxForce(length gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetSuspensionMaxForce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_suspension_max_force")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *VehicleWheel) SetSuspensionRestLength(length gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetSuspensionRestLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_suspension_rest_length")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *VehicleWheel) SetSuspensionStiffness(length gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetSuspensionStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_suspension_stiffness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *VehicleWheel) SetSuspensionTravel(length gdnative.Real) {
	//log.Println("Calling VehicleWheel.SetSuspensionTravel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_suspension_travel")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *VehicleWheel) SetUseAsSteering(enable gdnative.Bool) {
	//log.Println("Calling VehicleWheel.SetUseAsSteering()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_use_as_steering")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *VehicleWheel) SetUseAsTraction(enable gdnative.Bool) {
	//log.Println("Calling VehicleWheel.SetUseAsTraction()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VehicleWheel", "set_use_as_traction")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VehicleWheelImplementer is an interface that implements the methods
// of the VehicleWheel class.
type VehicleWheelImplementer interface {
	SpatialImplementer
	GetBrake() gdnative.Real
	GetDampingCompression() gdnative.Real
	GetDampingRelaxation() gdnative.Real
	GetEngineForce() gdnative.Real
	GetFrictionSlip() gdnative.Real
	GetRadius() gdnative.Real
	GetRollInfluence() gdnative.Real
	GetRpm() gdnative.Real
	GetSkidinfo() gdnative.Real
	GetSteering() gdnative.Real
	GetSuspensionMaxForce() gdnative.Real
	GetSuspensionRestLength() gdnative.Real
	GetSuspensionStiffness() gdnative.Real
	GetSuspensionTravel() gdnative.Real
	IsInContact() gdnative.Bool
	IsUsedAsSteering() gdnative.Bool
	IsUsedAsTraction() gdnative.Bool
	SetBrake(brake gdnative.Real)
	SetDampingCompression(length gdnative.Real)
	SetDampingRelaxation(length gdnative.Real)
	SetEngineForce(engineForce gdnative.Real)
	SetFrictionSlip(length gdnative.Real)
	SetRadius(length gdnative.Real)
	SetRollInfluence(rollInfluence gdnative.Real)
	SetSteering(steering gdnative.Real)
	SetSuspensionMaxForce(length gdnative.Real)
	SetSuspensionRestLength(length gdnative.Real)
	SetSuspensionStiffness(length gdnative.Real)
	SetSuspensionTravel(length gdnative.Real)
	SetUseAsSteering(enable gdnative.Bool)
	SetUseAsTraction(enable gdnative.Bool)
}
