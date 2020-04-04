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

//func NewInputEventPanGestureFromPointer(ptr gdnative.Pointer) InputEventPanGesture {
func newInputEventPanGestureFromPointer(ptr gdnative.Pointer) InputEventPanGesture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := InputEventPanGesture{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type InputEventPanGesture struct {
	InputEventGesture
	owner gdnative.Object
}

func (o *InputEventPanGesture) BaseClass() string {
	return "InputEventPanGesture"
}

/*
        Undocumented
	Args: [], Returns: Vector2
*/
func (o *InputEventPanGesture) GetDelta() gdnative.Vector2 {
	//log.Println("Calling InputEventPanGesture.GetDelta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventPanGesture", "get_delta")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false delta Vector2}], Returns: void
*/
func (o *InputEventPanGesture) SetDelta(delta gdnative.Vector2) {
	//log.Println("Calling InputEventPanGesture.SetDelta()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(delta)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("InputEventPanGesture", "set_delta")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// InputEventPanGestureImplementer is an interface that implements the methods
// of the InputEventPanGesture class.
type InputEventPanGestureImplementer interface {
	InputEventGestureImplementer
	GetDelta() gdnative.Vector2
	SetDelta(delta gdnative.Vector2)
}
