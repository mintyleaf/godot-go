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

//func NewVideoStreamGDNativeFromPointer(ptr gdnative.Pointer) VideoStreamGDNative {
func newVideoStreamGDNativeFromPointer(ptr gdnative.Pointer) VideoStreamGDNative {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VideoStreamGDNative{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type VideoStreamGDNative struct {
	VideoStream
	owner gdnative.Object
}

func (o *VideoStreamGDNative) BaseClass() string {
	return "VideoStreamGDNative"
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *VideoStreamGDNative) GetFile() gdnative.String {
	//log.Println("Calling VideoStreamGDNative.GetFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoStreamGDNative", "get_file")

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
	Args: [{ false file String}], Returns: void
*/
func (o *VideoStreamGDNative) SetFile(file gdnative.String) {
	//log.Println("Calling VideoStreamGDNative.SetFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(file)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("VideoStreamGDNative", "set_file")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// VideoStreamGDNativeImplementer is an interface that implements the methods
// of the VideoStreamGDNative class.
type VideoStreamGDNativeImplementer interface {
	VideoStreamImplementer
	GetFile() gdnative.String
	SetFile(file gdnative.String)
}