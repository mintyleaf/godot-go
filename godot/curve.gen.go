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

// CurveTangentMode is an enum for TangentMode values.
type CurveTangentMode int

const (
	CurveTangentFree      CurveTangentMode = 0
	CurveTangentLinear    CurveTangentMode = 1
	CurveTangentModeCount CurveTangentMode = 2
)

//func NewCurveFromPointer(ptr gdnative.Pointer) Curve {
func newCurveFromPointer(ptr gdnative.Pointer) Curve {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Curve{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A curve that can be saved and re-used for other objects. By default it ranges between [code]0[/code] and [code]1[/code] on the y-axis and positions points relative to the [code]0.5[/code] y-position.
*/
type Curve struct {
	Resource
	owner gdnative.Object
}

func (o *Curve) BaseClass() string {
	return "Curve"
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *Curve) X_GetData() gdnative.Array {
	//log.Println("Calling Curve.X_GetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "_get_data")

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
	Args: [{ false data Array}], Returns: void
*/
func (o *Curve) X_SetData(data gdnative.Array) {
	//log.Println("Calling Curve.X_SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "_set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a point to the curve. For each side, if the [code]*_mode[/code] is [constant TANGENT_LINEAR], the [code]*_tangent[/code] angle (in degrees) uses the slope of the curve halfway to the adjacent point. Allows custom assignments to the [code]*_tangent[/code] angle if [code]*_mode[/code] is set to [constant TANGENT_FREE].
	Args: [{ false position Vector2} {0 true left_tangent float} {0 true right_tangent float} {0 true left_mode int} {0 true right_mode int}], Returns: int
*/
func (o *Curve) AddPoint(position gdnative.Vector2, leftTangent gdnative.Real, rightTangent gdnative.Real, leftMode gdnative.Int, rightMode gdnative.Int) gdnative.Int {
	//log.Println("Calling Curve.AddPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)
	ptrArguments[1] = gdnative.NewPointerFromReal(leftTangent)
	ptrArguments[2] = gdnative.NewPointerFromReal(rightTangent)
	ptrArguments[3] = gdnative.NewPointerFromInt(leftMode)
	ptrArguments[4] = gdnative.NewPointerFromInt(rightMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "add_point")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Recomputes the baked cache of points for the curve.
	Args: [], Returns: void
*/
func (o *Curve) Bake() {
	//log.Println("Calling Curve.Bake()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "bake")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes points that are closer than [code]CMP_EPSILON[/code] (0.00001) units to their neighbor on the curve.
	Args: [], Returns: void
*/
func (o *Curve) CleanDupes() {
	//log.Println("Calling Curve.CleanDupes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "clean_dupes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all points from the curve.
	Args: [], Returns: void
*/
func (o *Curve) ClearPoints() {
	//log.Println("Calling Curve.ClearPoints()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "clear_points")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Curve) GetBakeResolution() gdnative.Int {
	//log.Println("Calling Curve.GetBakeResolution()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "get_bake_resolution")

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
func (o *Curve) GetMaxValue() gdnative.Real {
	//log.Println("Calling Curve.GetMaxValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "get_max_value")

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
func (o *Curve) GetMinValue() gdnative.Real {
	//log.Println("Calling Curve.GetMinValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "get_min_value")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the number of points describing the curve.
	Args: [], Returns: int
*/
func (o *Curve) GetPointCount() gdnative.Int {
	//log.Println("Calling Curve.GetPointCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "get_point_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the left [code]TangentMode[/code] for the point at [code]index[/code].
	Args: [{ false index int}], Returns: enum.Curve::TangentMode
*/
func (o *Curve) GetPointLeftMode(index gdnative.Int) CurveTangentMode {
	//log.Println("Calling Curve.GetPointLeftMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "get_point_left_mode")

	// Call the parent method.
	// enum.Curve::TangentMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return CurveTangentMode(ret)
}

/*
        Returns the left tangent angle (in degrees) for the point at [code]index[/code].
	Args: [{ false index int}], Returns: float
*/
func (o *Curve) GetPointLeftTangent(index gdnative.Int) gdnative.Real {
	//log.Println("Calling Curve.GetPointLeftTangent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "get_point_left_tangent")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the curve coordinates for the point at [code]index[/code].
	Args: [{ false index int}], Returns: Vector2
*/
func (o *Curve) GetPointPosition(index gdnative.Int) gdnative.Vector2 {
	//log.Println("Calling Curve.GetPointPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "get_point_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the right [code]TangentMode[/code] for the point at [code]index[/code].
	Args: [{ false index int}], Returns: enum.Curve::TangentMode
*/
func (o *Curve) GetPointRightMode(index gdnative.Int) CurveTangentMode {
	//log.Println("Calling Curve.GetPointRightMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "get_point_right_mode")

	// Call the parent method.
	// enum.Curve::TangentMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return CurveTangentMode(ret)
}

/*
        Returns the right tangent angle (in degrees) for the point at [code]index[/code].
	Args: [{ false index int}], Returns: float
*/
func (o *Curve) GetPointRightTangent(index gdnative.Int) gdnative.Real {
	//log.Println("Calling Curve.GetPointRightTangent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "get_point_right_tangent")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the y value for the point that would exist at x-position [code]offset[/code] along the curve.
	Args: [{ false offset float}], Returns: float
*/
func (o *Curve) Interpolate(offset gdnative.Real) gdnative.Real {
	//log.Println("Calling Curve.Interpolate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "interpolate")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Returns the y value for the point that would exist at x-position [code]offset[/code] along the curve using the baked cache. Bakes the curve's points if not already baked.
	Args: [{ false offset float}], Returns: float
*/
func (o *Curve) InterpolateBaked(offset gdnative.Real) gdnative.Real {
	//log.Println("Calling Curve.InterpolateBaked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "interpolate_baked")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Removes the point at [code]index[/code] from the curve.
	Args: [{ false index int}], Returns: void
*/
func (o *Curve) RemovePoint(index gdnative.Int) {
	//log.Println("Calling Curve.RemovePoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "remove_point")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false resolution int}], Returns: void
*/
func (o *Curve) SetBakeResolution(resolution gdnative.Int) {
	//log.Println("Calling Curve.SetBakeResolution()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(resolution)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "set_bake_resolution")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false max float}], Returns: void
*/
func (o *Curve) SetMaxValue(max gdnative.Real) {
	//log.Println("Calling Curve.SetMaxValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(max)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "set_max_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false min float}], Returns: void
*/
func (o *Curve) SetMinValue(min gdnative.Real) {
	//log.Println("Calling Curve.SetMinValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(min)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "set_min_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the left [code]TangentMode[/code] for the point at [code]index[/code] to [code]mode[/code].
	Args: [{ false index int} { false mode int}], Returns: void
*/
func (o *Curve) SetPointLeftMode(index gdnative.Int, mode gdnative.Int) {
	//log.Println("Calling Curve.SetPointLeftMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "set_point_left_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the left tangent angle for the point at [code]index[/code] to [code]tangent[/code].
	Args: [{ false index int} { false tangent float}], Returns: void
*/
func (o *Curve) SetPointLeftTangent(index gdnative.Int, tangent gdnative.Real) {
	//log.Println("Calling Curve.SetPointLeftTangent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromReal(tangent)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "set_point_left_tangent")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the offset from [code]0.5[/code]
	Args: [{ false index int} { false offset float}], Returns: int
*/
func (o *Curve) SetPointOffset(index gdnative.Int, offset gdnative.Real) gdnative.Int {
	//log.Println("Calling Curve.SetPointOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromReal(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "set_point_offset")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Sets the right [code]TangentMode[/code] for the point at [code]index[/code] to [code]mode[/code].
	Args: [{ false index int} { false mode int}], Returns: void
*/
func (o *Curve) SetPointRightMode(index gdnative.Int, mode gdnative.Int) {
	//log.Println("Calling Curve.SetPointRightMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "set_point_right_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the right tangent angle for the point at [code]index[/code] to [code]tangent[/code].
	Args: [{ false index int} { false tangent float}], Returns: void
*/
func (o *Curve) SetPointRightTangent(index gdnative.Int, tangent gdnative.Real) {
	//log.Println("Calling Curve.SetPointRightTangent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromReal(tangent)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "set_point_right_tangent")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Assigns the vertical position [code]y[/code] to the point at [code]index[/code].
	Args: [{ false index int} { false y float}], Returns: void
*/
func (o *Curve) SetPointValue(index gdnative.Int, y gdnative.Real) {
	//log.Println("Calling Curve.SetPointValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromReal(y)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Curve", "set_point_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CurveImplementer is an interface that implements the methods
// of the Curve class.
type CurveImplementer interface {
	ResourceImplementer
	X_GetData() gdnative.Array
	X_SetData(data gdnative.Array)
	AddPoint(position gdnative.Vector2, leftTangent gdnative.Real, rightTangent gdnative.Real, leftMode gdnative.Int, rightMode gdnative.Int) gdnative.Int
	Bake()
	CleanDupes()
	ClearPoints()
	GetBakeResolution() gdnative.Int
	GetMaxValue() gdnative.Real
	GetMinValue() gdnative.Real
	GetPointCount() gdnative.Int
	GetPointLeftTangent(index gdnative.Int) gdnative.Real
	GetPointPosition(index gdnative.Int) gdnative.Vector2
	GetPointRightTangent(index gdnative.Int) gdnative.Real
	Interpolate(offset gdnative.Real) gdnative.Real
	InterpolateBaked(offset gdnative.Real) gdnative.Real
	RemovePoint(index gdnative.Int)
	SetBakeResolution(resolution gdnative.Int)
	SetMaxValue(max gdnative.Real)
	SetMinValue(min gdnative.Real)
	SetPointLeftMode(index gdnative.Int, mode gdnative.Int)
	SetPointLeftTangent(index gdnative.Int, tangent gdnative.Real)
	SetPointOffset(index gdnative.Int, offset gdnative.Real) gdnative.Int
	SetPointRightMode(index gdnative.Int, mode gdnative.Int)
	SetPointRightTangent(index gdnative.Int, tangent gdnative.Real)
	SetPointValue(index gdnative.Int, y gdnative.Real)
}
