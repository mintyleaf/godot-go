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

//func NewScriptFromPointer(ptr gdnative.Pointer) Script {
func newScriptFromPointer(ptr gdnative.Pointer) Script {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Script{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A class stored as a resource. A script exends the functionality of all objects that instance it. The [code]new[/code] method of a script subclass creates a new instance. [method Object.set_script] extends an existing object, if that object's class matches one of the script's base classes.
*/
type Script struct {
	Resource
	owner gdnative.Object
}

func (o *Script) BaseClass() string {
	return "Script"
}

/*
        Returns [code]true[/code] if the script can be instanced.
	Args: [], Returns: bool
*/
func (o *Script) CanInstance() gdnative.Bool {
	//log.Println("Calling Script.CanInstance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "can_instance")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns the script directly inherited by this script.
	Args: [], Returns: Script
*/
func (o *Script) GetBaseScript() ScriptImplementer {
	//log.Println("Calling Script.GetBaseScript()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "get_base_script")

	// Call the parent method.
	// Script
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newScriptFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ScriptImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Script" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ScriptImplementer)
	}

	return &ret
}

/*
        Returns the script's base type.
	Args: [], Returns: String
*/
func (o *Script) GetInstanceBaseType() gdnative.String {
	//log.Println("Calling Script.GetInstanceBaseType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "get_instance_base_type")

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
	Args: [{ false property String}], Returns: Variant
*/
func (o *Script) GetPropertyDefaultValue(property gdnative.String) gdnative.Variant {
	//log.Println("Calling Script.GetPropertyDefaultValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(property)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "get_property_default_value")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *Script) GetScriptConstantMap() gdnative.Dictionary {
	//log.Println("Calling Script.GetScriptConstantMap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "get_script_constant_map")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *Script) GetScriptMethodList() gdnative.Array {
	//log.Println("Calling Script.GetScriptMethodList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "get_script_method_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *Script) GetScriptPropertyList() gdnative.Array {
	//log.Println("Calling Script.GetScriptPropertyList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "get_script_property_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *Script) GetScriptSignalList() gdnative.Array {
	//log.Println("Calling Script.GetScriptSignalList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "get_script_signal_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *Script) GetSourceCode() gdnative.String {
	//log.Println("Calling Script.GetSourceCode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "get_source_code")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the script, or a base class, defines a signal with the given name.
	Args: [{ false signal_name String}], Returns: bool
*/
func (o *Script) HasScriptSignal(signalName gdnative.String) gdnative.Bool {
	//log.Println("Calling Script.HasScriptSignal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(signalName)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "has_script_signal")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the script contains non-empty source code.
	Args: [], Returns: bool
*/
func (o *Script) HasSourceCode() gdnative.Bool {
	//log.Println("Calling Script.HasSourceCode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "has_source_code")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if [code]base_object[/code] is an instance of this script.
	Args: [{ false base_object Object}], Returns: bool
*/
func (o *Script) InstanceHas(baseObject ObjectImplementer) gdnative.Bool {
	//log.Println("Calling Script.InstanceHas()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(baseObject.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "instance_has")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the script is a tool script. A tool script can run in the editor.
	Args: [], Returns: bool
*/
func (o *Script) IsTool() gdnative.Bool {
	//log.Println("Calling Script.IsTool()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "is_tool")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Reloads the script's class implementation. Returns an error code.
	Args: [{False true keep_state bool}], Returns: enum.Error
*/
func (o *Script) Reload(keepState gdnative.Bool) gdnative.Error {
	//log.Println("Calling Script.Reload()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(keepState)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "reload")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false source String}], Returns: void
*/
func (o *Script) SetSourceCode(source gdnative.String) {
	//log.Println("Calling Script.SetSourceCode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(source)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Script", "set_source_code")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ScriptImplementer is an interface that implements the methods
// of the Script class.
type ScriptImplementer interface {
	ResourceImplementer
	CanInstance() gdnative.Bool
	GetBaseScript() ScriptImplementer
	GetInstanceBaseType() gdnative.String
	GetPropertyDefaultValue(property gdnative.String) gdnative.Variant
	GetScriptConstantMap() gdnative.Dictionary
	GetScriptMethodList() gdnative.Array
	GetScriptPropertyList() gdnative.Array
	GetScriptSignalList() gdnative.Array
	GetSourceCode() gdnative.String
	HasScriptSignal(signalName gdnative.String) gdnative.Bool
	HasSourceCode() gdnative.Bool
	InstanceHas(baseObject ObjectImplementer) gdnative.Bool
	IsTool() gdnative.Bool
	SetSourceCode(source gdnative.String)
}
