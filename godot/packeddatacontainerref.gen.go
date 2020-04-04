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

//func NewPackedDataContainerRefFromPointer(ptr gdnative.Pointer) PackedDataContainerRef {
func newPackedDataContainerRefFromPointer(ptr gdnative.Pointer) PackedDataContainerRef {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PackedDataContainerRef{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type PackedDataContainerRef struct {
	Reference
	owner gdnative.Object
}

func (o *PackedDataContainerRef) BaseClass() string {
	return "PackedDataContainerRef"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *PackedDataContainerRef) X_IsDictionary() gdnative.Bool {
	//log.Println("Calling PackedDataContainerRef.X_IsDictionary()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedDataContainerRef", "_is_dictionary")

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
	Args: [{ false arg0 Variant}], Returns: Variant
*/
func (o *PackedDataContainerRef) X_IterGet(arg0 gdnative.Variant) gdnative.Variant {
	//log.Println("Calling PackedDataContainerRef.X_IterGet()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVariant(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedDataContainerRef", "_iter_get")

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
	Args: [{ false arg0 Array}], Returns: Variant
*/
func (o *PackedDataContainerRef) X_IterInit(arg0 gdnative.Array) gdnative.Variant {
	//log.Println("Calling PackedDataContainerRef.X_IterInit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedDataContainerRef", "_iter_init")

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
	Args: [{ false arg0 Array}], Returns: Variant
*/
func (o *PackedDataContainerRef) X_IterNext(arg0 gdnative.Array) gdnative.Variant {
	//log.Println("Calling PackedDataContainerRef.X_IterNext()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedDataContainerRef", "_iter_next")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*

	Args: [], Returns: int
*/
func (o *PackedDataContainerRef) Size() gdnative.Int {
	//log.Println("Calling PackedDataContainerRef.Size()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedDataContainerRef", "size")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

// PackedDataContainerRefImplementer is an interface that implements the methods
// of the PackedDataContainerRef class.
type PackedDataContainerRefImplementer interface {
	ReferenceImplementer
	X_IsDictionary() gdnative.Bool
	X_IterGet(arg0 gdnative.Variant) gdnative.Variant
	X_IterInit(arg0 gdnative.Array) gdnative.Variant
	X_IterNext(arg0 gdnative.Array) gdnative.Variant
	Size() gdnative.Int
}
