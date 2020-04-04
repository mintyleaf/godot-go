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

//func NewResourceFormatSaverFromPointer(ptr gdnative.Pointer) ResourceFormatSaver {
func newResourceFormatSaverFromPointer(ptr gdnative.Pointer) ResourceFormatSaver {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatSaver{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The engine can save resources when you do it from the editor, or when you use the [ResourceSaver] singleton. This is accomplished thanks to multiple [ResourceFormatSaver]s, each handling its own format and called automatically by the engine. By default, Godot saves resources as [code].tres[/code] (text-based), [code].res[/code] (binary) or another built-in format, but you can choose to create your own format by extending this class. Be sure to respect the documented return types and values. You should give it a global class name with [code]class_name[/code] for it to be registered. Like built-in ResourceFormatSavers, it will be called automatically when saving resources of its recognized type(s). You may also implement a [ResourceFormatLoader].
*/
type ResourceFormatSaver struct {
	Reference
	owner gdnative.Object
}

func (o *ResourceFormatSaver) BaseClass() string {
	return "ResourceFormatSaver"
}

/*
        Returns the list of extensions available for saving the resource object, provided it is recognized (see [method recognize]).
	Args: [{ false resource Resource}], Returns: PoolStringArray
*/
func (o *ResourceFormatSaver) GetRecognizedExtensions(resource ResourceImplementer) gdnative.PoolStringArray {
	//log.Println("Calling ResourceFormatSaver.GetRecognizedExtensions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(resource.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ResourceFormatSaver", "get_recognized_extensions")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Returns whether the given resource object can be saved by this saver.
	Args: [{ false resource Resource}], Returns: bool
*/
func (o *ResourceFormatSaver) Recognize(resource ResourceImplementer) gdnative.Bool {
	//log.Println("Calling ResourceFormatSaver.Recognize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(resource.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ResourceFormatSaver", "recognize")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Saves the given resource object to a file at the target [code]path[/code]. [code]flags[/code] is a bitmask composed with [enum ResourceSaver.SaverFlags] constants. Returns [constant OK] on success, or an [enum Error] constant in case of failure.
	Args: [{ false path String} { false resource Resource} { false flags int}], Returns: int
*/
func (o *ResourceFormatSaver) Save(path gdnative.String, resource ResourceImplementer, flags gdnative.Int) gdnative.Int {
	//log.Println("Calling ResourceFormatSaver.Save()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromObject(resource.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ResourceFormatSaver", "save")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

// ResourceFormatSaverImplementer is an interface that implements the methods
// of the ResourceFormatSaver class.
type ResourceFormatSaverImplementer interface {
	ReferenceImplementer
	GetRecognizedExtensions(resource ResourceImplementer) gdnative.PoolStringArray
	Recognize(resource ResourceImplementer) gdnative.Bool
	Save(path gdnative.String, resource ResourceImplementer, flags gdnative.Int) gdnative.Int
}
