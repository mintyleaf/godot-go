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

// VisualScriptPropertyGetCallMode is an enum for CallMode values.
type VisualScriptPropertyGetCallMode int

const (
	VisualScriptPropertyGetCallModeInstance VisualScriptPropertyGetCallMode = 2
	VisualScriptPropertyGetCallModeNodePath VisualScriptPropertyGetCallMode = 1
	VisualScriptPropertyGetCallModeSelf     VisualScriptPropertyGetCallMode = 0
)

//func NewVisualScriptPropertyGetFromPointer(ptr gdnative.Pointer) VisualScriptPropertyGet {
func newVisualScriptPropertyGetFromPointer(ptr gdnative.Pointer) VisualScriptPropertyGet {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VisualScriptPropertyGet{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VisualScriptPropertyGet struct {
	VisualScriptNode
	owner gdnative.Object
}

func (o *VisualScriptPropertyGet) BaseClass() string {
	return "VisualScriptPropertyGet"
}

/*
        Undocumented
	Args: [], Returns: enum.Variant::Type
*/
func (o *VisualScriptPropertyGet) X_GetTypeCache() gdnative.VariantType {
	//log.Println("Calling VisualScriptPropertyGet.X_GetTypeCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "_get_type_cache")

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
	Args: [{ false type_cache int}], Returns: void
*/
func (o *VisualScriptPropertyGet) X_SetTypeCache(typeCache gdnative.Int) {
	//log.Println("Calling VisualScriptPropertyGet.X_SetTypeCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(typeCache)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "_set_type_cache")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *VisualScriptPropertyGet) GetBasePath() gdnative.NodePath {
	//log.Println("Calling VisualScriptPropertyGet.GetBasePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "get_base_path")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptPropertyGet) GetBaseScript() gdnative.String {
	//log.Println("Calling VisualScriptPropertyGet.GetBaseScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "get_base_script")

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
func (o *VisualScriptPropertyGet) GetBaseType() gdnative.String {
	//log.Println("Calling VisualScriptPropertyGet.GetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "get_base_type")

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
	Args: [], Returns: enum.Variant::Type
*/
func (o *VisualScriptPropertyGet) GetBasicType() gdnative.VariantType {
	//log.Println("Calling VisualScriptPropertyGet.GetBasicType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "get_basic_type")

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
	Args: [], Returns: enum.VisualScriptPropertyGet::CallMode
*/
func (o *VisualScriptPropertyGet) GetCallMode() VisualScriptPropertyGetCallMode {
	//log.Println("Calling VisualScriptPropertyGet.GetCallMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "get_call_mode")

	// Call the parent method.
	// enum.VisualScriptPropertyGet::CallMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return VisualScriptPropertyGetCallMode(ret)
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VisualScriptPropertyGet) GetIndex() gdnative.String {
	//log.Println("Calling VisualScriptPropertyGet.GetIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "get_index")

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
func (o *VisualScriptPropertyGet) GetProperty() gdnative.String {
	//log.Println("Calling VisualScriptPropertyGet.GetProperty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "get_property")

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
	Args: [{ false base_path NodePath}], Returns: void
*/
func (o *VisualScriptPropertyGet) SetBasePath(basePath gdnative.NodePath) {
	//log.Println("Calling VisualScriptPropertyGet.SetBasePath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(basePath)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "set_base_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false base_script String}], Returns: void
*/
func (o *VisualScriptPropertyGet) SetBaseScript(baseScript gdnative.String) {
	//log.Println("Calling VisualScriptPropertyGet.SetBaseScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(baseScript)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "set_base_script")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false base_type String}], Returns: void
*/
func (o *VisualScriptPropertyGet) SetBaseType(baseType gdnative.String) {
	//log.Println("Calling VisualScriptPropertyGet.SetBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(baseType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "set_base_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false basic_type int}], Returns: void
*/
func (o *VisualScriptPropertyGet) SetBasicType(basicType gdnative.Int) {
	//log.Println("Calling VisualScriptPropertyGet.SetBasicType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(basicType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "set_basic_type")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *VisualScriptPropertyGet) SetCallMode(mode gdnative.Int) {
	//log.Println("Calling VisualScriptPropertyGet.SetCallMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "set_call_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false index String}], Returns: void
*/
func (o *VisualScriptPropertyGet) SetIndex(index gdnative.String) {
	//log.Println("Calling VisualScriptPropertyGet.SetIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "set_index")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false property String}], Returns: void
*/
func (o *VisualScriptPropertyGet) SetProperty(property gdnative.String) {
	//log.Println("Calling VisualScriptPropertyGet.SetProperty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(property)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VisualScriptPropertyGet", "set_property")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VisualScriptPropertyGetImplementer is an interface that implements the methods
// of the VisualScriptPropertyGet class.
type VisualScriptPropertyGetImplementer interface {
	VisualScriptNodeImplementer
	X_SetTypeCache(typeCache gdnative.Int)
	GetBasePath() gdnative.NodePath
	GetBaseScript() gdnative.String
	GetBaseType() gdnative.String
	GetIndex() gdnative.String
	GetProperty() gdnative.String
	SetBasePath(basePath gdnative.NodePath)
	SetBaseScript(baseScript gdnative.String)
	SetBaseType(baseType gdnative.String)
	SetBasicType(basicType gdnative.Int)
	SetCallMode(mode gdnative.Int)
	SetIndex(index gdnative.String)
	SetProperty(property gdnative.String)
}
