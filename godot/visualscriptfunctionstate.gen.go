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

//func NewVisualScriptFunctionStateFromPointer(ptr gdnative.Pointer) VisualScriptFunctionState {
func newVisualScriptFunctionStateFromPointer(ptr gdnative.Pointer) VisualScriptFunctionState {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptFunctionState{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptFunctionState struct {
	Reference
	owner gdnative.Object
}

func (o *VisualScriptFunctionState) BaseClass() string {
	return "VisualScriptFunctionState"
}

/*
        Undocumented
	Args: [], Returns: Variant
*/
func (o *VisualScriptFunctionState) X_SignalCallback(args ...gdnative.Variant) gdnative.Variant {
	//log.Println("Calling VisualScriptFunctionState.X_SignalCallback()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0+len(args), 0+len(args))

	for i, arg := range args {
		ptrArguments[i+0] = gdnative.NewPointerFromVariant(arg)
	}

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionState", "_signal_callback")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false obj Object} { false signals String} { false args Array}], Returns: void
*/
func (o *VisualScriptFunctionState) ConnectToSignal(obj ObjectImplementer, signals gdnative.String, args gdnative.Array) {
	//log.Println("Calling VisualScriptFunctionState.ConnectToSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(obj.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromString(signals)
	ptrArguments[2] = gdnative.NewPointerFromArray(args)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionState", "connect_to_signal")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *VisualScriptFunctionState) IsValid() gdnative.Bool {
	//log.Println("Calling VisualScriptFunctionState.IsValid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionState", "is_valid")

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
	Args: [{Null true args Array}], Returns: Variant
*/
func (o *VisualScriptFunctionState) Resume(args gdnative.Array) gdnative.Variant {
	//log.Println("Calling VisualScriptFunctionState.Resume()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(args)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptFunctionState", "resume")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

// VisualScriptFunctionStateImplementer is an interface that implements the methods
// of the VisualScriptFunctionState class.
type VisualScriptFunctionStateImplementer interface {
	ReferenceImplementer
	X_SignalCallback(args ...gdnative.Variant) gdnative.Variant
	ConnectToSignal(obj ObjectImplementer, signals gdnative.String, args gdnative.Array)
	IsValid() gdnative.Bool
	Resume(args gdnative.Array) gdnative.Variant
}
