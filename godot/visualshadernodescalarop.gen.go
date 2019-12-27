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

// VisualShaderNodeScalarOpOperator is an enum for Operator values.
type VisualShaderNodeScalarOpOperator int

const (
	VisualShaderNodeScalarOpOpAdd   VisualShaderNodeScalarOpOperator = 0
	VisualShaderNodeScalarOpOpAtan2 VisualShaderNodeScalarOpOperator = 8
	VisualShaderNodeScalarOpOpDiv   VisualShaderNodeScalarOpOperator = 3
	VisualShaderNodeScalarOpOpMax   VisualShaderNodeScalarOpOperator = 6
	VisualShaderNodeScalarOpOpMin   VisualShaderNodeScalarOpOperator = 7
	VisualShaderNodeScalarOpOpMod   VisualShaderNodeScalarOpOperator = 4
	VisualShaderNodeScalarOpOpMul   VisualShaderNodeScalarOpOperator = 2
	VisualShaderNodeScalarOpOpPow   VisualShaderNodeScalarOpOperator = 5
	VisualShaderNodeScalarOpOpStep  VisualShaderNodeScalarOpOperator = 9
	VisualShaderNodeScalarOpOpSub   VisualShaderNodeScalarOpOperator = 1
)

//func NewVisualShaderNodeScalarOpFromPointer(ptr gdnative.Pointer) VisualShaderNodeScalarOp {
func newVisualShaderNodeScalarOpFromPointer(ptr gdnative.Pointer) VisualShaderNodeScalarOp {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeScalarOp{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeScalarOp struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeScalarOp) BaseClass() string {
	return "VisualShaderNodeScalarOp"
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeScalarOp::Operator
*/
func (o *VisualShaderNodeScalarOp) GetOperator() VisualShaderNodeScalarOpOperator {
	//log.Println("Calling VisualShaderNodeScalarOp.GetOperator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeScalarOp", "get_operator")

	// Call the parent method.
	// enum.VisualShaderNodeScalarOp::Operator
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeScalarOpOperator(ret)
}

/*
        Undocumented
	Args: [{ false op int}], Returns: void
*/
func (o *VisualShaderNodeScalarOp) SetOperator(op gdnative.Int) {
	//log.Println("Calling VisualShaderNodeScalarOp.SetOperator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(op)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeScalarOp", "set_operator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeScalarOpImplementer is an interface that implements the methods
// of the VisualShaderNodeScalarOp class.
type VisualShaderNodeScalarOpImplementer interface {
	VisualShaderNodeImplementer
	SetOperator(op gdnative.Int)
}
