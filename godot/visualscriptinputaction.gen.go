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

// VisualScriptInputActionMode is an enum for Mode values.
type VisualScriptInputActionMode int

const (
	VisualScriptInputActionModeJustPressed  VisualScriptInputActionMode = 2
	VisualScriptInputActionModeJustReleased VisualScriptInputActionMode = 3
	VisualScriptInputActionModePressed      VisualScriptInputActionMode = 0
	VisualScriptInputActionModeReleased     VisualScriptInputActionMode = 1
)

//func NewVisualScriptInputActionFromPointer(ptr gdnative.Pointer) VisualScriptInputAction {
func newVisualScriptInputActionFromPointer(ptr gdnative.Pointer) VisualScriptInputAction {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptInputAction{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptInputAction struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptInputAction) BaseClass() string {
	return "VisualScriptInputAction"
}

/*
        Undocumented
	Args: [], Returns: enum.VisualScriptInputAction::Mode
*/
func (o *VisualScriptInputAction) GetActionMode() VisualScriptInputActionMode {
	//log.Println("Calling VisualScriptInputAction.GetActionMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptInputAction", "get_action_mode")

	// Call the parent method.
	// enum.VisualScriptInputAction::Mode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualScriptInputActionMode(ret)
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptInputAction) GetActionName() gdnative.String {
	//log.Println("Calling VisualScriptInputAction.GetActionName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptInputAction", "get_action_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *VisualScriptInputAction) SetActionMode(mode gdnative.Int) {
	//log.Println("Calling VisualScriptInputAction.SetActionMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptInputAction", "set_action_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false name String}], Returns: void
*/
func (o *VisualScriptInputAction) SetActionName(name gdnative.String) {
	//log.Println("Calling VisualScriptInputAction.SetActionName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptInputAction", "set_action_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptInputActionImplementer is an interface that implements the methods
// of the VisualScriptInputAction class.
type VisualScriptInputActionImplementer interface {
	VisualScriptNodeImplementer
	GetActionName() gdnative.String
	SetActionMode(mode gdnative.Int)
	SetActionName(name gdnative.String)
}
