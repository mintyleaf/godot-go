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

//func NewVisualScriptClassConstantFromPointer(ptr gdnative.Pointer) VisualScriptClassConstant {
func newVisualScriptClassConstantFromPointer(ptr gdnative.Pointer) VisualScriptClassConstant {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptClassConstant{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptClassConstant struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptClassConstant) BaseClass() string {
	return "VisualScriptClassConstant"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptClassConstant) GetBaseType() gdnative.String {
	//log.Println("Calling VisualScriptClassConstant.GetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptClassConstant", "get_base_type")

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
func (o *VisualScriptClassConstant) GetClassConstant() gdnative.String {
	//log.Println("Calling VisualScriptClassConstant.GetClassConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptClassConstant", "get_class_constant")

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
func (o *VisualScriptClassConstant) SetBaseType(name gdnative.String) {
	//log.Println("Calling VisualScriptClassConstant.SetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptClassConstant", "set_base_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false name String}], Returns: void
*/
func (o *VisualScriptClassConstant) SetClassConstant(name gdnative.String) {
	//log.Println("Calling VisualScriptClassConstant.SetClassConstant()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptClassConstant", "set_class_constant")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptClassConstantImplementer is an interface that implements the methods
// of the VisualScriptClassConstant class.
type VisualScriptClassConstantImplementer interface {
	VisualScriptNodeImplementer
	GetBaseType() gdnative.String
	GetClassConstant() gdnative.String
	SetBaseType(name gdnative.String)
	SetClassConstant(name gdnative.String)
}
