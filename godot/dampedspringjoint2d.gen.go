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

//func NewDampedSpringJoint2DFromPointer(ptr gdnative.Pointer) DampedSpringJoint2D {
func newDampedSpringJoint2DFromPointer(ptr gdnative.Pointer) DampedSpringJoint2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := DampedSpringJoint2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Damped spring constraint for 2D physics. This resembles a spring joint that always wants to go back to a given length.
*/
type DampedSpringJoint2D struct {
	Joint2D
	owner gdnative.Object
}

func (o *DampedSpringJoint2D) BaseClass() string {
	return "DampedSpringJoint2D"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *DampedSpringJoint2D) GetDamping() gdnative.Real {
	//log.Println("Calling DampedSpringJoint2D.GetDamping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DampedSpringJoint2D", "get_damping")

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
func (o *DampedSpringJoint2D) GetLength() gdnative.Real {
	//log.Println("Calling DampedSpringJoint2D.GetLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DampedSpringJoint2D", "get_length")

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
func (o *DampedSpringJoint2D) GetRestLength() gdnative.Real {
	//log.Println("Calling DampedSpringJoint2D.GetRestLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DampedSpringJoint2D", "get_rest_length")

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
func (o *DampedSpringJoint2D) GetStiffness() gdnative.Real {
	//log.Println("Calling DampedSpringJoint2D.GetStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DampedSpringJoint2D", "get_stiffness")

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
	Args: [{ false damping float}], Returns: void
*/
func (o *DampedSpringJoint2D) SetDamping(damping gdnative.Real) {
	//log.Println("Calling DampedSpringJoint2D.SetDamping()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(damping)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DampedSpringJoint2D", "set_damping")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false length float}], Returns: void
*/
func (o *DampedSpringJoint2D) SetLength(length gdnative.Real) {
	//log.Println("Calling DampedSpringJoint2D.SetLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(length)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DampedSpringJoint2D", "set_length")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rest_length float}], Returns: void
*/
func (o *DampedSpringJoint2D) SetRestLength(restLength gdnative.Real) {
	//log.Println("Calling DampedSpringJoint2D.SetRestLength()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(restLength)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DampedSpringJoint2D", "set_rest_length")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false stiffness float}], Returns: void
*/
func (o *DampedSpringJoint2D) SetStiffness(stiffness gdnative.Real) {
	//log.Println("Calling DampedSpringJoint2D.SetStiffness()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(stiffness)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("DampedSpringJoint2D", "set_stiffness")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// DampedSpringJoint2DImplementer is an interface that implements the methods
// of the DampedSpringJoint2D class.
type DampedSpringJoint2DImplementer interface {
	Joint2DImplementer
	GetDamping() gdnative.Real
	GetLength() gdnative.Real
	GetRestLength() gdnative.Real
	GetStiffness() gdnative.Real
	SetDamping(damping gdnative.Real)
	SetLength(length gdnative.Real)
	SetRestLength(restLength gdnative.Real)
	SetStiffness(stiffness gdnative.Real)
}
