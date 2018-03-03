package gdnativelibrary

import (
	"log"
	"reflect"

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

/*
Undocumented
*/
type GDNativeLibrary struct {
	Resource
}

func (o *GDNativeLibrary) BaseClass() string {
	return "GDNativeLibrary"
}

/*
   Undocumented
*/
func (o *GDNativeLibrary) GetConfigFile() *ConfigFile {
	log.Println("Calling GDNativeLibrary.GetConfigFile()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_config_file", goArguments, "*ConfigFile")

	returnValue := goRet.Interface().(*ConfigFile)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) GetCurrentDependencies() *PoolStringArray {
	log.Println("Calling GDNativeLibrary.GetCurrentDependencies()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_current_dependencies", goArguments, "*PoolStringArray")

	returnValue := goRet.Interface().(*PoolStringArray)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) GetCurrentLibraryPath() gdnative.String {
	log.Println("Calling GDNativeLibrary.GetCurrentLibraryPath()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_current_library_path", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) GetSymbolPrefix() gdnative.String {
	log.Println("Calling GDNativeLibrary.GetSymbolPrefix()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_symbol_prefix", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) IsReloadable() gdnative.Bool {
	log.Println("Calling GDNativeLibrary.IsReloadable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_reloadable", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) IsSingleton() gdnative.Bool {
	log.Println("Calling GDNativeLibrary.IsSingleton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_singleton", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) SetLoadOnce(loadOnce gdnative.Bool) {
	log.Println("Calling GDNativeLibrary.SetLoadOnce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(loadOnce)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_load_once", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) SetReloadable(reloadable gdnative.Bool) {
	log.Println("Calling GDNativeLibrary.SetReloadable()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(reloadable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_reloadable", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) SetSingleton(singleton gdnative.Bool) {
	log.Println("Calling GDNativeLibrary.SetSingleton()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(singleton)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_singleton", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) SetSymbolPrefix(symbolPrefix gdnative.String) {
	log.Println("Calling GDNativeLibrary.SetSymbolPrefix()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(symbolPrefix)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_symbol_prefix", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *GDNativeLibrary) ShouldLoadOnce() gdnative.Bool {
	log.Println("Calling GDNativeLibrary.ShouldLoadOnce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "should_load_once", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   GDNativeLibraryImplementer is an interface for GDNativeLibrary objects.
*/
type GDNativeLibraryImplementer interface {
	Class
}