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

//func NewVisualScriptEmitSignalFromPointer(ptr gdnative.Pointer) VisualScriptEmitSignal {
func newVisualScriptEmitSignalFromPointer(ptr gdnative.Pointer) VisualScriptEmitSignal {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptEmitSignal{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptEmitSignal struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptEmitSignal) BaseClass() string {
	return "VisualScriptEmitSignal"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptEmitSignal) GetSignal() gdnative.String {
	//log.Println("Calling VisualScriptEmitSignal.GetSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptEmitSignal", "get_signal")

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
	Args: [{ false name String}], Returns: void
*/
func (o *VisualScriptEmitSignal) SetSignal(name gdnative.String) {
	//log.Println("Calling VisualScriptEmitSignal.SetSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptEmitSignal", "set_signal")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptEmitSignalImplementer is an interface that implements the methods
// of the VisualScriptEmitSignal class.
type VisualScriptEmitSignalImplementer interface {
	VisualScriptNodeImplementer
	GetSignal() gdnative.String
	SetSignal(name gdnative.String)
}
