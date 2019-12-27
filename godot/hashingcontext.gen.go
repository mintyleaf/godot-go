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

// HashingContextHashType is an enum for HashType values.
type HashingContextHashType int

const (
	HashingContextHashMd5    HashingContextHashType = 0
	HashingContextHashSha1   HashingContextHashType = 1
	HashingContextHashSha256 HashingContextHashType = 2
)

//func NewHashingContextFromPointer(ptr gdnative.Pointer) HashingContext {
func newHashingContextFromPointer(ptr gdnative.Pointer) HashingContext {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := HashingContext{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type HashingContext struct {
	Reference
	owner gdnative.Object
}

func (o *HashingContext) BaseClass() string {
	return "HashingContext"
}

/*
        Undocumented
	Args: [], Returns: PoolByteArray
*/
func (o *HashingContext) Finish() gdnative.PoolByteArray {
	//log.Println("Calling HashingContext.Finish()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HashingContext", "finish")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false type int}], Returns: enum.Error
*/
func (o *HashingContext) Start(aType gdnative.Int) gdnative.Error {
	//log.Println("Calling HashingContext.Start()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(aType)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HashingContext", "start")

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
	Args: [{ false chunk PoolByteArray}], Returns: enum.Error
*/
func (o *HashingContext) Update(chunk gdnative.PoolByteArray) gdnative.Error {
	//log.Println("Calling HashingContext.Update()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolByteArray(chunk)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("HashingContext", "update")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

// HashingContextImplementer is an interface that implements the methods
// of the HashingContext class.
type HashingContextImplementer interface {
	ReferenceImplementer
	Finish() gdnative.PoolByteArray
}
