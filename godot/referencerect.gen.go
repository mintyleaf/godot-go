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

//func NewReferenceRectFromPointer(ptr gdnative.Pointer) ReferenceRect {
func newReferenceRectFromPointer(ptr gdnative.Pointer) ReferenceRect {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ReferenceRect{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A rectangle box that displays only a [member border_color] border color around its rectangle. [ReferenceRect] has no fill [Color].
*/
type ReferenceRect struct {
	Control
	owner gdnative.Object
}

func (o *ReferenceRect) BaseClass() string {
	return "ReferenceRect"
}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *ReferenceRect) GetBorderColor() gdnative.Color {
	//log.Println("Calling ReferenceRect.GetBorderColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ReferenceRect", "get_border_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *ReferenceRect) GetEditorOnly() gdnative.Bool {
	//log.Println("Calling ReferenceRect.GetEditorOnly()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ReferenceRect", "get_editor_only")

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
	Args: [{ false color Color}], Returns: void
*/
func (o *ReferenceRect) SetBorderColor(color gdnative.Color) {
	//log.Println("Calling ReferenceRect.SetBorderColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ReferenceRect", "set_border_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *ReferenceRect) SetEditorOnly(enabled gdnative.Bool) {
	//log.Println("Calling ReferenceRect.SetEditorOnly()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ReferenceRect", "set_editor_only")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ReferenceRectImplementer is an interface that implements the methods
// of the ReferenceRect class.
type ReferenceRectImplementer interface {
	ControlImplementer
	GetBorderColor() gdnative.Color
	GetEditorOnly() gdnative.Bool
	SetBorderColor(color gdnative.Color)
	SetEditorOnly(enabled gdnative.Bool)
}
