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

// _ResourceSaverSaverFlags is an enum for SaverFlags values.
type _ResourceSaverSaverFlags int

const (
	_ResourceSaverFlagBundleResources         _ResourceSaverSaverFlags = 2
	_ResourceSaverFlagChangePath              _ResourceSaverSaverFlags = 4
	_ResourceSaverFlagCompress                _ResourceSaverSaverFlags = 32
	_ResourceSaverFlagOmitEditorProperties    _ResourceSaverSaverFlags = 8
	_ResourceSaverFlagRelativePaths           _ResourceSaverSaverFlags = 1
	_ResourceSaverFlagReplaceSubresourcePaths _ResourceSaverSaverFlags = 64
	_ResourceSaverFlagSaveBigEndian           _ResourceSaverSaverFlags = 16
)

//func NewresourceSaverFromPointer(ptr gdnative.Pointer) resourceSaver {
func new_ResourceSaverFromPointer(ptr gdnative.Pointer) resourceSaver {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := resourceSaver{}
	obj.SetBaseObject(owner)

	return obj
}

func newSingletonResourceSaver() *resourceSaver {
	return &resourceSaver{}
}

/*
   Singleton for saving Godot-specific resource types to the filesystem. It uses the many [ResourceFormatSaver] classes registered in the engine (either built-in or from a plugin) to save engine-specific resource data to text-based (e.g. [code].tres[/code] or [code].tscn[/code]) or binary files (e.g. [code].res[/code] or [code].scn[/code]).
*/
var ResourceSaver = newSingletonResourceSaver()

/*
Singleton for saving Godot-specific resource types to the filesystem. It uses the many [ResourceFormatSaver] classes registered in the engine (either built-in or from a plugin) to save engine-specific resource data to text-based (e.g. [code].tres[/code] or [code].tscn[/code]) or binary files (e.g. [code].res[/code] or [code].scn[/code]).
*/
type resourceSaver struct {
	Object
	owner       gdnative.Object
	initialized bool
}

// EnsureSingleton will check to see if we have an object for it. If not, it will fetch its
// GDNative object and set it.
func (o *resourceSaver) ensureSingleton() {
	if o.initialized == true {
		return
	}
	//log.Println("Singleton not found. Fetching from GDNative...")
	base := gdnative.GetSingleton("_ResourceSaver")
	o.SetBaseObject(base)
	o.initialized = true
}

func (o *resourceSaver) BaseClass() string {
	return "_ResourceSaver"
}

/*
        Undocumented
	Args: [{ false type Resource}], Returns: PoolStringArray
*/
func (o *resourceSaver) GetRecognizedExtensions(aType ResourceImplementer) gdnative.PoolStringArray {
	o.ensureSingleton()
	//log.Println("Calling _ResourceSaver.GetRecognizedExtensions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(aType.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_ResourceSaver", "get_recognized_extensions")

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
	Args: [{ false path String} { false resource Resource} {0 true flags int}], Returns: enum.Error
*/
func (o *resourceSaver) Save(path gdnative.String, resource ResourceImplementer, flags gdnative.Int) gdnative.Error {
	o.ensureSingleton()
	//log.Println("Calling _ResourceSaver.Save()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromObject(resource.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_ResourceSaver", "save")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

// ResourceSaverImplementer is an interface that implements the methods
// of the ResourceSaver class.
type ResourceSaverImplementer interface {
	ObjectImplementer
	GetRecognizedExtensions(aType ResourceImplementer) gdnative.PoolStringArray
}
