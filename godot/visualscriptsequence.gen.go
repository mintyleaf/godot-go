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

//func NewVisualScriptSequenceFromPointer(ptr gdnative.Pointer) VisualScriptSequence {
func newVisualScriptSequenceFromPointer(ptr gdnative.Pointer) VisualScriptSequence {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptSequence{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptSequence struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptSequence) BaseClass() string {
	return "VisualScriptSequence"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *VisualScriptSequence) GetSteps() gdnative.Int {
	//log.Println("Calling VisualScriptSequence.GetSteps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptSequence", "get_steps")

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
	Args: [{ false steps int}], Returns: void
*/
func (o *VisualScriptSequence) SetSteps(steps gdnative.Int) {
	//log.Println("Calling VisualScriptSequence.SetSteps()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(steps)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptSequence", "set_steps")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptSequenceImplementer is an interface that implements the methods
// of the VisualScriptSequence class.
type VisualScriptSequenceImplementer interface {
	VisualScriptNodeImplementer
	GetSteps() gdnative.Int
	SetSteps(steps gdnative.Int)
}
