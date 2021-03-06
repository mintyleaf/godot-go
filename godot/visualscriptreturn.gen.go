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

//func NewVisualScriptReturnFromPointer(ptr gdnative.Pointer) VisualScriptReturn {
func newVisualScriptReturnFromPointer(ptr gdnative.Pointer) VisualScriptReturn {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptReturn{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptReturn struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptReturn) BaseClass() string {
	return "VisualScriptReturn"
}

/*
        Undocumented
	Args: [], Returns: enum.Variant::Type
*/
func (o *VisualScriptReturn) GetReturnType() gdnative.VariantType {
	//log.Println("Calling VisualScriptReturn.GetReturnType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptReturn", "get_return_type")

	// Call the parent method.
	// enum.Variant::Type
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.VariantType(ret)
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *VisualScriptReturn) IsReturnValueEnabled() gdnative.Bool {
	//log.Println("Calling VisualScriptReturn.IsReturnValueEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptReturn", "is_return_value_enabled")

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
	Args: [{ false enable bool}], Returns: void
*/
func (o *VisualScriptReturn) SetEnableReturnValue(enable gdnative.Bool) {
	//log.Println("Calling VisualScriptReturn.SetEnableReturnValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptReturn", "set_enable_return_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false type int}], Returns: void
*/
func (o *VisualScriptReturn) SetReturnType(aType gdnative.Int) {
	//log.Println("Calling VisualScriptReturn.SetReturnType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptReturn", "set_return_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptReturnImplementer is an interface that implements the methods
// of the VisualScriptReturn class.
type VisualScriptReturnImplementer interface {
	VisualScriptNodeImplementer
	IsReturnValueEnabled() gdnative.Bool
	SetEnableReturnValue(enable gdnative.Bool)
	SetReturnType(aType gdnative.Int)
}
