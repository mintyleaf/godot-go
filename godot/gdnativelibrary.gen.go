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

//func NewGDNativeLibraryFromPointer(ptr gdnative.Pointer) GDNativeLibrary {
func newGDNativeLibraryFromPointer(ptr gdnative.Pointer) GDNativeLibrary {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := GDNativeLibrary{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type GDNativeLibrary struct {
	Resource
	owner gdnative.Object
}

func (o *GDNativeLibrary) BaseClass() string {
	return "GDNativeLibrary"
}

/*
        Undocumented
	Args: [], Returns: ConfigFile
*/
func (o *GDNativeLibrary) GetConfigFile() ConfigFileImplementer {
	//log.Println("Calling GDNativeLibrary.GetConfigFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "get_config_file")

	// Call the parent method.
	// ConfigFile
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newConfigFileFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ConfigFileImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "ConfigFile" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ConfigFileImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: PoolStringArray
*/
func (o *GDNativeLibrary) GetCurrentDependencies() gdnative.PoolStringArray {
	//log.Println("Calling GDNativeLibrary.GetCurrentDependencies()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "get_current_dependencies")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *GDNativeLibrary) GetCurrentLibraryPath() gdnative.String {
	//log.Println("Calling GDNativeLibrary.GetCurrentLibraryPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "get_current_library_path")

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
func (o *GDNativeLibrary) GetSymbolPrefix() gdnative.String {
	//log.Println("Calling GDNativeLibrary.GetSymbolPrefix()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "get_symbol_prefix")

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
	Args: [], Returns: bool
*/
func (o *GDNativeLibrary) IsReloadable() gdnative.Bool {
	//log.Println("Calling GDNativeLibrary.IsReloadable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "is_reloadable")

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
	Args: [], Returns: bool
*/
func (o *GDNativeLibrary) IsSingleton() gdnative.Bool {
	//log.Println("Calling GDNativeLibrary.IsSingleton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "is_singleton")

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
	Args: [{ false config_file ConfigFile}], Returns: void
*/
func (o *GDNativeLibrary) SetConfigFile(configFile ConfigFileImplementer) {
	//log.Println("Calling GDNativeLibrary.SetConfigFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(configFile.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "set_config_file")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false load_once bool}], Returns: void
*/
func (o *GDNativeLibrary) SetLoadOnce(loadOnce gdnative.Bool) {
	//log.Println("Calling GDNativeLibrary.SetLoadOnce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(loadOnce)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "set_load_once")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false reloadable bool}], Returns: void
*/
func (o *GDNativeLibrary) SetReloadable(reloadable gdnative.Bool) {
	//log.Println("Calling GDNativeLibrary.SetReloadable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(reloadable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "set_reloadable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false singleton bool}], Returns: void
*/
func (o *GDNativeLibrary) SetSingleton(singleton gdnative.Bool) {
	//log.Println("Calling GDNativeLibrary.SetSingleton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(singleton)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "set_singleton")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false symbol_prefix String}], Returns: void
*/
func (o *GDNativeLibrary) SetSymbolPrefix(symbolPrefix gdnative.String) {
	//log.Println("Calling GDNativeLibrary.SetSymbolPrefix()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(symbolPrefix)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "set_symbol_prefix")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *GDNativeLibrary) ShouldLoadOnce() gdnative.Bool {
	//log.Println("Calling GDNativeLibrary.ShouldLoadOnce()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("GDNativeLibrary", "should_load_once")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

// GDNativeLibraryImplementer is an interface that implements the methods
// of the GDNativeLibrary class.
type GDNativeLibraryImplementer interface {
	ResourceImplementer
	GetConfigFile() ConfigFileImplementer
	GetCurrentDependencies() gdnative.PoolStringArray
	GetCurrentLibraryPath() gdnative.String
	GetSymbolPrefix() gdnative.String
	IsReloadable() gdnative.Bool
	IsSingleton() gdnative.Bool
	SetConfigFile(configFile ConfigFileImplementer)
	SetLoadOnce(loadOnce gdnative.Bool)
	SetReloadable(reloadable gdnative.Bool)
	SetSingleton(singleton gdnative.Bool)
	SetSymbolPrefix(symbolPrefix gdnative.String)
	ShouldLoadOnce() gdnative.Bool
}
