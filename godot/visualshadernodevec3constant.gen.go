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

//func NewVisualShaderNodeVec3ConstantFromPointer(ptr gdnative.Pointer) VisualShaderNodeVec3Constant {
func newVisualShaderNodeVec3ConstantFromPointer(ptr gdnative.Pointer) VisualShaderNodeVec3Constant {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualShaderNodeVec3Constant{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VisualShaderNodeVec3Constant struct {
	VisualShaderNode
	owner gdnative.Object
}

func (o *VisualShaderNodeVec3Constant) BaseClass() string {
	return "VisualShaderNodeVec3Constant"
}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *VisualShaderNodeVec3Constant) GetConstant() gdnative.Vector3 {
	//log.Println("Calling VisualShaderNodeVec3Constant.GetConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeVec3Constant", "get_constant")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false value Vector3}], Returns: void
*/
func (o *VisualShaderNodeVec3Constant) SetConstant(value gdnative.Vector3) {
	//log.Println("Calling VisualShaderNodeVec3Constant.SetConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualShaderNodeVec3Constant", "set_constant")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualShaderNodeVec3ConstantImplementer is an interface that implements the methods
// of the VisualShaderNodeVec3Constant class.
type VisualShaderNodeVec3ConstantImplementer interface {
	VisualShaderNodeImplementer
	GetConstant() gdnative.Vector3
	SetConstant(value gdnative.Vector3)
}
