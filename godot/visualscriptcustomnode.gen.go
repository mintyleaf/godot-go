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

// VisualScriptCustomNodeStartMode is an enum for StartMode values.
type VisualScriptCustomNodeStartMode int

const (
	VisualScriptCustomNodeStartModeBeginSequence    VisualScriptCustomNodeStartMode = 0
	VisualScriptCustomNodeStartModeContinueSequence VisualScriptCustomNodeStartMode = 1
	VisualScriptCustomNodeStartModeResumeYield      VisualScriptCustomNodeStartMode = 2
)

//func NewVisualScriptCustomNodeFromPointer(ptr gdnative.Pointer) VisualScriptCustomNode {
func newVisualScriptCustomNodeFromPointer(ptr gdnative.Pointer) VisualScriptCustomNode {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptCustomNode{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptCustomNode struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptCustomNode) BaseClass() string {
	return "VisualScriptCustomNode"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptCustomNode) X_GetCaption() gdnative.String {
	//log.Println("Calling VisualScriptCustomNode.X_GetCaption()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_caption")

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
	Args: [], Returns: String
*/
func (o *VisualScriptCustomNode) X_GetCategory() gdnative.String {
	//log.Println("Calling VisualScriptCustomNode.X_GetCategory()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_category")

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
	Args: [], Returns: int
*/
func (o *VisualScriptCustomNode) X_GetInputValuePortCount() gdnative.Int {
	//log.Println("Calling VisualScriptCustomNode.X_GetInputValuePortCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_input_value_port_count")

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
	Args: [{ false idx int}], Returns: String
*/
func (o *VisualScriptCustomNode) X_GetInputValuePortName(idx gdnative.Int) gdnative.String {
	//log.Println("Calling VisualScriptCustomNode.X_GetInputValuePortName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_input_value_port_name")

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
	Args: [{ false idx int}], Returns: int
*/
func (o *VisualScriptCustomNode) X_GetInputValuePortType(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling VisualScriptCustomNode.X_GetInputValuePortType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_input_value_port_type")

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
	Args: [], Returns: int
*/
func (o *VisualScriptCustomNode) X_GetOutputSequencePortCount() gdnative.Int {
	//log.Println("Calling VisualScriptCustomNode.X_GetOutputSequencePortCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_output_sequence_port_count")

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
	Args: [{ false idx int}], Returns: String
*/
func (o *VisualScriptCustomNode) X_GetOutputSequencePortText(idx gdnative.Int) gdnative.String {
	//log.Println("Calling VisualScriptCustomNode.X_GetOutputSequencePortText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_output_sequence_port_text")

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
	Args: [], Returns: int
*/
func (o *VisualScriptCustomNode) X_GetOutputValuePortCount() gdnative.Int {
	//log.Println("Calling VisualScriptCustomNode.X_GetOutputValuePortCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_output_value_port_count")

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
	Args: [{ false idx int}], Returns: String
*/
func (o *VisualScriptCustomNode) X_GetOutputValuePortName(idx gdnative.Int) gdnative.String {
	//log.Println("Calling VisualScriptCustomNode.X_GetOutputValuePortName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_output_value_port_name")

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
	Args: [{ false idx int}], Returns: int
*/
func (o *VisualScriptCustomNode) X_GetOutputValuePortType(idx gdnative.Int) gdnative.Int {
	//log.Println("Calling VisualScriptCustomNode.X_GetOutputValuePortType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_output_value_port_type")

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
	Args: [], Returns: String
*/
func (o *VisualScriptCustomNode) X_GetText() gdnative.String {
	//log.Println("Calling VisualScriptCustomNode.X_GetText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_text")

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
	Args: [], Returns: int
*/
func (o *VisualScriptCustomNode) X_GetWorkingMemorySize() gdnative.Int {
	//log.Println("Calling VisualScriptCustomNode.X_GetWorkingMemorySize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_get_working_memory_size")

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
func (o *VisualScriptCustomNode) X_HasInputSequencePort() gdnative.Bool {
	//log.Println("Calling VisualScriptCustomNode.X_HasInputSequencePort()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_has_input_sequence_port")

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
	Args: [], Returns: void
*/
func (o *VisualScriptCustomNode) X_ScriptChanged() {
	//log.Println("Calling VisualScriptCustomNode.X_ScriptChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_script_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false inputs Array} { false outputs Array} { false start_mode int} { false working_mem Array}], Returns: Variant
*/
func (o *VisualScriptCustomNode) X_Step(inputs gdnative.Array, outputs gdnative.Array, startMode gdnative.Int, workingMem gdnative.Array) gdnative.Variant {
	//log.Println("Calling VisualScriptCustomNode.X_Step()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromArray(inputs)
	ptrArguments[1] = gdnative.NewPointerFromArray(outputs)
	ptrArguments[2] = gdnative.NewPointerFromInt(startMode)
	ptrArguments[3] = gdnative.NewPointerFromArray(workingMem)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptCustomNode", "_step")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

// VisualScriptCustomNodeImplementer is an interface that implements the methods
// of the VisualScriptCustomNode class.
type VisualScriptCustomNodeImplementer interface {
	VisualScriptNodeImplementer
	X_GetCaption() gdnative.String
	X_GetCategory() gdnative.String
	X_GetInputValuePortCount() gdnative.Int
	X_GetInputValuePortName(idx gdnative.Int) gdnative.String
	X_GetInputValuePortType(idx gdnative.Int) gdnative.Int
	X_GetOutputSequencePortCount() gdnative.Int
	X_GetOutputSequencePortText(idx gdnative.Int) gdnative.String
	X_GetOutputValuePortCount() gdnative.Int
	X_GetOutputValuePortName(idx gdnative.Int) gdnative.String
	X_GetOutputValuePortType(idx gdnative.Int) gdnative.Int
	X_GetText() gdnative.String
	X_GetWorkingMemorySize() gdnative.Int
	X_HasInputSequencePort() gdnative.Bool
	X_ScriptChanged()
	X_Step(inputs gdnative.Array, outputs gdnative.Array, startMode gdnative.Int, workingMem gdnative.Array) gdnative.Variant
}
