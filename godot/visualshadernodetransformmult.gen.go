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

// VisualShaderNodeTransformMultOperator is an enum for Operator values.
type VisualShaderNodeTransformMultOperator int

const (
	VisualShaderNodeTransformMultOpAxB     VisualShaderNodeTransformMultOperator = 0
	VisualShaderNodeTransformMultOpAxBComp VisualShaderNodeTransformMultOperator = 2
	VisualShaderNodeTransformMultOpBxA     VisualShaderNodeTransformMultOperator = 1
	VisualShaderNodeTransformMultOpBxAComp VisualShaderNodeTransformMultOperator = 3
)

//func NewVisualShaderNodeTransformMultFromPointer(ptr gdnative.Pointer) VisualShaderNodeTransformMult {
func newVisualShaderNodeTransformMultFromPointer(ptr gdnative.Pointer) VisualShaderNodeTransformMult {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeTransformMult{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeTransformMult struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeTransformMult) BaseClass() string {
	return "VisualShaderNodeTransformMult"
}

/*
        Undocumented
	Args: [], Returns: enum.VisualShaderNodeTransformMult::Operator
*/
func (o *VisualShaderNodeTransformMult) GetOperator() VisualShaderNodeTransformMultOperator {
	//log.Println("Calling VisualShaderNodeTransformMult.GetOperator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeTransformMult", "get_operator")

	// Call the parent method.
	// enum.VisualShaderNodeTransformMult::Operator
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualShaderNodeTransformMultOperator(ret)
}

/*
        Undocumented
	Args: [{ false op int}], Returns: void
*/
func (o *VisualShaderNodeTransformMult) SetOperator(op gdnative.Int) {
	//log.Println("Calling VisualShaderNodeTransformMult.SetOperator()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(op)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeTransformMult", "set_operator")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeTransformMultImplementer is an interface that implements the methods
// of the VisualShaderNodeTransformMult class.
type VisualShaderNodeTransformMultImplementer interface {
	VisualShaderNodeImplementer
	SetOperator(op gdnative.Int)
}
