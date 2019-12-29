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

// Generic6DOFJointFlag is an enum for Flag values.
type Generic6DOFJointFlag int

const (
	Generic6DOFJointFlagEnableAngularLimit  Generic6DOFJointFlag = 1
	Generic6DOFJointFlagEnableAngularSpring Generic6DOFJointFlag = 2
	Generic6DOFJointFlagEnableLinearLimit   Generic6DOFJointFlag = 0
	Generic6DOFJointFlagEnableLinearMotor   Generic6DOFJointFlag = 5
	Generic6DOFJointFlagEnableLinearSpring  Generic6DOFJointFlag = 3
	Generic6DOFJointFlagEnableMotor         Generic6DOFJointFlag = 4
	Generic6DOFJointFlagMax                 Generic6DOFJointFlag = 6
)

// Generic6DOFJointParam is an enum for Param values.
type Generic6DOFJointParam int

const (
	Generic6DOFJointParamAngularDamping             Generic6DOFJointParam = 13
	Generic6DOFJointParamAngularErp                 Generic6DOFJointParam = 16
	Generic6DOFJointParamAngularForceLimit          Generic6DOFJointParam = 15
	Generic6DOFJointParamAngularLimitSoftness       Generic6DOFJointParam = 12
	Generic6DOFJointParamAngularLowerLimit          Generic6DOFJointParam = 10
	Generic6DOFJointParamAngularMotorForceLimit     Generic6DOFJointParam = 18
	Generic6DOFJointParamAngularMotorTargetVelocity Generic6DOFJointParam = 17
	Generic6DOFJointParamAngularRestitution         Generic6DOFJointParam = 14
	Generic6DOFJointParamAngularUpperLimit          Generic6DOFJointParam = 11
	Generic6DOFJointParamLinearDamping              Generic6DOFJointParam = 4
	Generic6DOFJointParamLinearLimitSoftness        Generic6DOFJointParam = 2
	Generic6DOFJointParamLinearLowerLimit           Generic6DOFJointParam = 0
	Generic6DOFJointParamLinearMotorForceLimit      Generic6DOFJointParam = 6
	Generic6DOFJointParamLinearMotorTargetVelocity  Generic6DOFJointParam = 5
	Generic6DOFJointParamLinearRestitution          Generic6DOFJointParam = 3
	Generic6DOFJointParamLinearUpperLimit           Generic6DOFJointParam = 1
	Generic6DOFJointParamMax                        Generic6DOFJointParam = 22
)

//func NewGeneric6DOFJointFromPointer(ptr gdnative.Pointer) Generic6DOFJoint {
func newGeneric6DOFJointFromPointer(ptr gdnative.Pointer) Generic6DOFJoint {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Generic6DOFJoint{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The first 3 DOF axes are linear axes, which represent translation of Bodies, and the latter 3 DOF axes represent the angular motion. Each axis can be either locked, or limited.
*/
type Generic6DOFJoint struct {
	Joint
	owner gdnative.Object
}

func (o *Generic6DOFJoint) BaseClass() string {
	return "Generic6DOFJoint"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Generic6DOFJoint) X_GetAngularHiLimitX() gdnative.Real {
	//log.Println("Calling Generic6DOFJoint.X_GetAngularHiLimitX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_get_angular_hi_limit_x")

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
func (o *Generic6DOFJoint) X_GetAngularHiLimitY() gdnative.Real {
	//log.Println("Calling Generic6DOFJoint.X_GetAngularHiLimitY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_get_angular_hi_limit_y")

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
func (o *Generic6DOFJoint) X_GetAngularHiLimitZ() gdnative.Real {
	//log.Println("Calling Generic6DOFJoint.X_GetAngularHiLimitZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_get_angular_hi_limit_z")

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
func (o *Generic6DOFJoint) X_GetAngularLoLimitX() gdnative.Real {
	//log.Println("Calling Generic6DOFJoint.X_GetAngularLoLimitX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_get_angular_lo_limit_x")

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
func (o *Generic6DOFJoint) X_GetAngularLoLimitY() gdnative.Real {
	//log.Println("Calling Generic6DOFJoint.X_GetAngularLoLimitY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_get_angular_lo_limit_y")

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
func (o *Generic6DOFJoint) X_GetAngularLoLimitZ() gdnative.Real {
	//log.Println("Calling Generic6DOFJoint.X_GetAngularLoLimitZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_get_angular_lo_limit_z")

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
	Args: [{ false angle float}], Returns: void
*/
func (o *Generic6DOFJoint) X_SetAngularHiLimitX(angle gdnative.Real) {
	//log.Println("Calling Generic6DOFJoint.X_SetAngularHiLimitX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_set_angular_hi_limit_x")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false angle float}], Returns: void
*/
func (o *Generic6DOFJoint) X_SetAngularHiLimitY(angle gdnative.Real) {
	//log.Println("Calling Generic6DOFJoint.X_SetAngularHiLimitY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_set_angular_hi_limit_y")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false angle float}], Returns: void
*/
func (o *Generic6DOFJoint) X_SetAngularHiLimitZ(angle gdnative.Real) {
	//log.Println("Calling Generic6DOFJoint.X_SetAngularHiLimitZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_set_angular_hi_limit_z")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false angle float}], Returns: void
*/
func (o *Generic6DOFJoint) X_SetAngularLoLimitX(angle gdnative.Real) {
	//log.Println("Calling Generic6DOFJoint.X_SetAngularLoLimitX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_set_angular_lo_limit_x")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false angle float}], Returns: void
*/
func (o *Generic6DOFJoint) X_SetAngularLoLimitY(angle gdnative.Real) {
	//log.Println("Calling Generic6DOFJoint.X_SetAngularLoLimitY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_set_angular_lo_limit_y")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false angle float}], Returns: void
*/
func (o *Generic6DOFJoint) X_SetAngularLoLimitZ(angle gdnative.Real) {
	//log.Println("Calling Generic6DOFJoint.X_SetAngularLoLimitZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(angle)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "_set_angular_lo_limit_z")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false flag int}], Returns: bool
*/
func (o *Generic6DOFJoint) GetFlagX(flag gdnative.Int) gdnative.Bool {
	//log.Println("Calling Generic6DOFJoint.GetFlagX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "get_flag_x")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false flag int}], Returns: bool
*/
func (o *Generic6DOFJoint) GetFlagY(flag gdnative.Int) gdnative.Bool {
	//log.Println("Calling Generic6DOFJoint.GetFlagY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "get_flag_y")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false flag int}], Returns: bool
*/
func (o *Generic6DOFJoint) GetFlagZ(flag gdnative.Int) gdnative.Bool {
	//log.Println("Calling Generic6DOFJoint.GetFlagZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "get_flag_z")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false param int}], Returns: float
*/
func (o *Generic6DOFJoint) GetParamX(param gdnative.Int) gdnative.Real {
	//log.Println("Calling Generic6DOFJoint.GetParamX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "get_param_x")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false param int}], Returns: float
*/
func (o *Generic6DOFJoint) GetParamY(param gdnative.Int) gdnative.Real {
	//log.Println("Calling Generic6DOFJoint.GetParamY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "get_param_y")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false param int}], Returns: float
*/
func (o *Generic6DOFJoint) GetParamZ(param gdnative.Int) gdnative.Real {
	//log.Println("Calling Generic6DOFJoint.GetParamZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "get_param_z")

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
func (o *Generic6DOFJoint) GetPrecision() gdnative.Int {
	//log.Println("Calling Generic6DOFJoint.GetPrecision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "get_precision")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false flag int} { false value bool}], Returns: void
*/
func (o *Generic6DOFJoint) SetFlagX(flag gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling Generic6DOFJoint.SetFlagX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "set_flag_x")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false flag int} { false value bool}], Returns: void
*/
func (o *Generic6DOFJoint) SetFlagY(flag gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling Generic6DOFJoint.SetFlagY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "set_flag_y")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false flag int} { false value bool}], Returns: void
*/
func (o *Generic6DOFJoint) SetFlagZ(flag gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling Generic6DOFJoint.SetFlagZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(flag)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "set_flag_z")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false param int} { false value float}], Returns: void
*/
func (o *Generic6DOFJoint) SetParamX(param gdnative.Int, value gdnative.Real) {
	//log.Println("Calling Generic6DOFJoint.SetParamX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "set_param_x")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false param int} { false value float}], Returns: void
*/
func (o *Generic6DOFJoint) SetParamY(param gdnative.Int, value gdnative.Real) {
	//log.Println("Calling Generic6DOFJoint.SetParamY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "set_param_y")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false param int} { false value float}], Returns: void
*/
func (o *Generic6DOFJoint) SetParamZ(param gdnative.Int, value gdnative.Real) {
	//log.Println("Calling Generic6DOFJoint.SetParamZ()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(param)
	ptrArguments[1] = gdnative.NewPointerFromReal(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "set_param_z")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false precision int}], Returns: void
*/
func (o *Generic6DOFJoint) SetPrecision(precision gdnative.Int) {
	//log.Println("Calling Generic6DOFJoint.SetPrecision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(precision)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Generic6DOFJoint", "set_precision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Generic6DOFJointImplementer is an interface that implements the methods
// of the Generic6DOFJoint class.
type Generic6DOFJointImplementer interface {
	JointImplementer
	X_GetAngularHiLimitX() gdnative.Real
	X_GetAngularHiLimitY() gdnative.Real
	X_GetAngularHiLimitZ() gdnative.Real
	X_GetAngularLoLimitX() gdnative.Real
	X_GetAngularLoLimitY() gdnative.Real
	X_GetAngularLoLimitZ() gdnative.Real
	X_SetAngularHiLimitX(angle gdnative.Real)
	X_SetAngularHiLimitY(angle gdnative.Real)
	X_SetAngularHiLimitZ(angle gdnative.Real)
	X_SetAngularLoLimitX(angle gdnative.Real)
	X_SetAngularLoLimitY(angle gdnative.Real)
	X_SetAngularLoLimitZ(angle gdnative.Real)
	GetFlagX(flag gdnative.Int) gdnative.Bool
	GetFlagY(flag gdnative.Int) gdnative.Bool
	GetFlagZ(flag gdnative.Int) gdnative.Bool
	GetParamX(param gdnative.Int) gdnative.Real
	GetParamY(param gdnative.Int) gdnative.Real
	GetParamZ(param gdnative.Int) gdnative.Real
	GetPrecision() gdnative.Int
	SetFlagX(flag gdnative.Int, value gdnative.Bool)
	SetFlagY(flag gdnative.Int, value gdnative.Bool)
	SetFlagZ(flag gdnative.Int, value gdnative.Bool)
	SetParamX(param gdnative.Int, value gdnative.Real)
	SetParamY(param gdnative.Int, value gdnative.Real)
	SetParamZ(param gdnative.Int, value gdnative.Real)
	SetPrecision(precision gdnative.Int)
}
