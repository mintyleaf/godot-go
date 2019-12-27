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

//func NewPopupFromPointer(ptr gdnative.Pointer) Popup {
func newPopupFromPointer(ptr gdnative.Pointer) Popup {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Popup{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Popup is a base [Control] used to show dialogs and popups. It's a subwindow and modal by default (see [Control]) and has helpers for custom popup behavior.
*/
type Popup struct {
	Control
	owner gdnative.Object
}

func (o *Popup) BaseClass() string {
	return "Popup"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Popup) IsExclusive() gdnative.Bool {
	//log.Println("Calling Popup.IsExclusive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Popup", "is_exclusive")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Popup (show the control in modal form).
	Args: [{(0, 0, 0, 0) true bounds Rect2}], Returns: void
*/
func (o *Popup) PopupMethod(bounds gdnative.Rect2) {
	//log.Println("Calling Popup.PopupMethod()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRect2(bounds)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Popup", "popup")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Popup (show the control in modal form) in the center of the screen relative to its current canvas transform, at the current size, or at a size determined by "size".
	Args: [{(0, 0) true size Vector2}], Returns: void
*/
func (o *Popup) PopupCentered(size gdnative.Vector2) {
	//log.Println("Calling Popup.PopupCentered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Popup", "popup_centered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{(0, 0) true size Vector2} {0.75 true fallback_ratio float}], Returns: void
*/
func (o *Popup) PopupCenteredClamped(size gdnative.Vector2, fallbackRatio gdnative.Real) {
	//log.Println("Calling Popup.PopupCenteredClamped()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromVector2(size)
	ptrArguments[1] = gdnative.NewPointerFromReal(fallbackRatio)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Popup", "popup_centered_clamped")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Popup (show the control in modal form) in the center of the screen relative to the current canvas transform, ensuring the size is never smaller than [code]minsize[/code].
	Args: [{(0, 0) true minsize Vector2}], Returns: void
*/
func (o *Popup) PopupCenteredMinsize(minsize gdnative.Vector2) {
	//log.Println("Calling Popup.PopupCenteredMinsize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(minsize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Popup", "popup_centered_minsize")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Popup (show the control in modal form) in the center of the screen relative to the current canvas transform, scaled at a ratio of size of the screen.
	Args: [{0.75 true ratio float}], Returns: void
*/
func (o *Popup) PopupCenteredRatio(ratio gdnative.Real) {
	//log.Println("Calling Popup.PopupCenteredRatio()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(ratio)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Popup", "popup_centered_ratio")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Popup) SetAsMinsize() {
	//log.Println("Calling Popup.SetAsMinsize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Popup", "set_as_minsize")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Popup) SetExclusive(enable gdnative.Bool) {
	//log.Println("Calling Popup.SetExclusive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Popup", "set_exclusive")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// PopupImplementer is an interface that implements the methods
// of the Popup class.
type PopupImplementer interface {
	ControlImplementer
	IsExclusive() gdnative.Bool
	PopupMethod(bounds gdnative.Rect2)
	PopupCentered(size gdnative.Vector2)
	PopupCenteredClamped(size gdnative.Vector2, fallbackRatio gdnative.Real)
	PopupCenteredMinsize(minsize gdnative.Vector2)
	PopupCenteredRatio(ratio gdnative.Real)
	SetAsMinsize()
	SetExclusive(enable gdnative.Bool)
}
