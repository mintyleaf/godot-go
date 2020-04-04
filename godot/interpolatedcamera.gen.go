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

//func NewInterpolatedCameraFromPointer(ptr gdnative.Pointer) InterpolatedCamera {
func newInterpolatedCameraFromPointer(ptr gdnative.Pointer) InterpolatedCamera {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InterpolatedCamera{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type InterpolatedCamera struct {
	Camera
	owner gdnative.Object
}

func (o *InterpolatedCamera) BaseClass() string {
	return "InterpolatedCamera"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *InterpolatedCamera) GetSpeed() gdnative.Real {
	//log.Println("Calling InterpolatedCamera.GetSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "get_speed")

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
func (o *InterpolatedCamera) GetTargetPath() gdnative.NodePath {
	//log.Println("Calling InterpolatedCamera.GetTargetPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "get_target_path")

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
	Args: [], Returns: bool
*/
func (o *InterpolatedCamera) IsInterpolationEnabled() gdnative.Bool {
	//log.Println("Calling InterpolatedCamera.IsInterpolationEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "is_interpolation_enabled")

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
	Args: [{ false target_path bool}], Returns: void
*/
func (o *InterpolatedCamera) SetInterpolationEnabled(targetPath gdnative.Bool) {
	//log.Println("Calling InterpolatedCamera.SetInterpolationEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(targetPath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "set_interpolation_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false speed float}], Returns: void
*/
func (o *InterpolatedCamera) SetSpeed(speed gdnative.Real) {
	//log.Println("Calling InterpolatedCamera.SetSpeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(speed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "set_speed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false target Object}], Returns: void
*/
func (o *InterpolatedCamera) SetTarget(target ObjectImplementer) {
	//log.Println("Calling InterpolatedCamera.SetTarget()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(target.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "set_target")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false target_path NodePath}], Returns: void
*/
func (o *InterpolatedCamera) SetTargetPath(targetPath gdnative.NodePath) {
	//log.Println("Calling InterpolatedCamera.SetTargetPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(targetPath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InterpolatedCamera", "set_target_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// InterpolatedCameraImplementer is an interface that implements the methods
// of the InterpolatedCamera class.
type InterpolatedCameraImplementer interface {
	CameraImplementer
	GetSpeed() gdnative.Real
	GetTargetPath() gdnative.NodePath
	IsInterpolationEnabled() gdnative.Bool
	SetInterpolationEnabled(targetPath gdnative.Bool)
	SetSpeed(speed gdnative.Real)
	SetTarget(target ObjectImplementer)
	SetTargetPath(targetPath gdnative.NodePath)
}
