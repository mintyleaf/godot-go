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

//func NewViewportContainerFromPointer(ptr gdnative.Pointer) ViewportContainer {
func newViewportContainerFromPointer(ptr gdnative.Pointer) ViewportContainer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ViewportContainer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A [Container] node that holds a [Viewport], automatically setting its size.
*/
type ViewportContainer struct {
	Container
	owner gdnative.Object
}

func (o *ViewportContainer) BaseClass() string {
	return "ViewportContainer"
}

/*
        Undocumented
	Args: [{ false event InputEvent}], Returns: void
*/
func (o *ViewportContainer) X_Input(event InputEventImplementer) {
	//log.Println("Calling ViewportContainer.X_Input()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ViewportContainer", "_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false event InputEvent}], Returns: void
*/
func (o *ViewportContainer) X_UnhandledInput(event InputEventImplementer) {
	//log.Println("Calling ViewportContainer.X_UnhandledInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(event.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ViewportContainer", "_unhandled_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *ViewportContainer) GetStretchShrink() gdnative.Int {
	//log.Println("Calling ViewportContainer.GetStretchShrink()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ViewportContainer", "get_stretch_shrink")

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
func (o *ViewportContainer) IsStretchEnabled() gdnative.Bool {
	//log.Println("Calling ViewportContainer.IsStretchEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ViewportContainer", "is_stretch_enabled")

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
func (o *ViewportContainer) SetStretch(enable gdnative.Bool) {
	//log.Println("Calling ViewportContainer.SetStretch()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ViewportContainer", "set_stretch")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false amount int}], Returns: void
*/
func (o *ViewportContainer) SetStretchShrink(amount gdnative.Int) {
	//log.Println("Calling ViewportContainer.SetStretchShrink()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(amount)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ViewportContainer", "set_stretch_shrink")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ViewportContainerImplementer is an interface that implements the methods
// of the ViewportContainer class.
type ViewportContainerImplementer interface {
	ContainerImplementer
	GetStretchShrink() gdnative.Int
	IsStretchEnabled() gdnative.Bool
	SetStretch(enable gdnative.Bool)
	SetStretchShrink(amount gdnative.Int)
}
