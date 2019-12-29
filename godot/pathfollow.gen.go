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

// PathFollowRotationMode is an enum for RotationMode values.
type PathFollowRotationMode int

const (
	PathFollowRotationNone     PathFollowRotationMode = 0
	PathFollowRotationOriented PathFollowRotationMode = 4
	PathFollowRotationXy       PathFollowRotationMode = 2
	PathFollowRotationXyz      PathFollowRotationMode = 3
	PathFollowRotationY        PathFollowRotationMode = 1
)

//func NewPathFollowFromPointer(ptr gdnative.Pointer) PathFollow {
func newPathFollowFromPointer(ptr gdnative.Pointer) PathFollow {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PathFollow{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This node takes its parent [Path], and returns the coordinates of a point within it, given a distance from the first vertex. It is useful for making other nodes follow a path, without coding the movement pattern. For that, the nodes must be children of this node. The descendant nodes will then move accordingly when setting an offset in this node.
*/
type PathFollow struct {
	Spatial
	owner gdnative.Object
}

func (o *PathFollow) BaseClass() string {
	return "PathFollow"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *PathFollow) GetCubicInterpolation() gdnative.Bool {
	//log.Println("Calling PathFollow.GetCubicInterpolation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "get_cubic_interpolation")

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
func (o *PathFollow) GetHOffset() gdnative.Real {
	//log.Println("Calling PathFollow.GetHOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "get_h_offset")

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
func (o *PathFollow) GetOffset() gdnative.Real {
	//log.Println("Calling PathFollow.GetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "get_offset")

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
	Args: [], Returns: enum.PathFollow::RotationMode
*/
func (o *PathFollow) GetRotationMode() PathFollowRotationMode {
	//log.Println("Calling PathFollow.GetRotationMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "get_rotation_mode")

	// Call the parent method.
	// enum.PathFollow::RotationMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return PathFollowRotationMode(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *PathFollow) GetUnitOffset() gdnative.Real {
	//log.Println("Calling PathFollow.GetUnitOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "get_unit_offset")

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
func (o *PathFollow) GetVOffset() gdnative.Real {
	//log.Println("Calling PathFollow.GetVOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "get_v_offset")

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
func (o *PathFollow) HasLoop() gdnative.Bool {
	//log.Println("Calling PathFollow.HasLoop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "has_loop")

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
func (o *PathFollow) SetCubicInterpolation(enable gdnative.Bool) {
	//log.Println("Calling PathFollow.SetCubicInterpolation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "set_cubic_interpolation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false h_offset float}], Returns: void
*/
func (o *PathFollow) SetHOffset(hOffset gdnative.Real) {
	//log.Println("Calling PathFollow.SetHOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(hOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "set_h_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false loop bool}], Returns: void
*/
func (o *PathFollow) SetLoop(loop gdnative.Bool) {
	//log.Println("Calling PathFollow.SetLoop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(loop)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "set_loop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false offset float}], Returns: void
*/
func (o *PathFollow) SetOffset(offset gdnative.Real) {
	//log.Println("Calling PathFollow.SetOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "set_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rotation_mode int}], Returns: void
*/
func (o *PathFollow) SetRotationMode(rotationMode gdnative.Int) {
	//log.Println("Calling PathFollow.SetRotationMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(rotationMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "set_rotation_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false unit_offset float}], Returns: void
*/
func (o *PathFollow) SetUnitOffset(unitOffset gdnative.Real) {
	//log.Println("Calling PathFollow.SetUnitOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(unitOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "set_unit_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false v_offset float}], Returns: void
*/
func (o *PathFollow) SetVOffset(vOffset gdnative.Real) {
	//log.Println("Calling PathFollow.SetVOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(vOffset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PathFollow", "set_v_offset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PathFollowImplementer is an interface that implements the methods
// of the PathFollow class.
type PathFollowImplementer interface {
	SpatialImplementer
	GetCubicInterpolation() gdnative.Bool
	GetHOffset() gdnative.Real
	GetOffset() gdnative.Real
	GetUnitOffset() gdnative.Real
	GetVOffset() gdnative.Real
	HasLoop() gdnative.Bool
	SetCubicInterpolation(enable gdnative.Bool)
	SetHOffset(hOffset gdnative.Real)
	SetLoop(loop gdnative.Bool)
	SetOffset(offset gdnative.Real)
	SetRotationMode(rotationMode gdnative.Int)
	SetUnitOffset(unitOffset gdnative.Real)
	SetVOffset(vOffset gdnative.Real)
}
