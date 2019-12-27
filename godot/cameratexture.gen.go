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

//func NewCameraTextureFromPointer(ptr gdnative.Pointer) CameraTexture {
func newCameraTextureFromPointer(ptr gdnative.Pointer) CameraTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CameraTexture{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type CameraTexture struct {
	Texture
	owner gdnative.Object
}

func (o *CameraTexture) BaseClass() string {
	return "CameraTexture"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *CameraTexture) GetCameraActive() gdnative.Bool {
	//log.Println("Calling CameraTexture.GetCameraActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CameraTexture", "get_camera_active")

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
	Args: [], Returns: int
*/
func (o *CameraTexture) GetCameraFeedId() gdnative.Int {
	//log.Println("Calling CameraTexture.GetCameraFeedId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CameraTexture", "get_camera_feed_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.CameraServer::FeedImage
*/
func (o *CameraTexture) GetWhichFeed() CameraServerFeedImage {
	//log.Println("Calling CameraTexture.GetWhichFeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CameraTexture", "get_which_feed")

	// Call the parent method.
	// enum.CameraServer::FeedImage
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return CameraServerFeedImage(ret)
}

/*
        Undocumented
	Args: [{ false active bool}], Returns: void
*/
func (o *CameraTexture) SetCameraActive(active gdnative.Bool) {
	//log.Println("Calling CameraTexture.SetCameraActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(active)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CameraTexture", "set_camera_active")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false feed_id int}], Returns: void
*/
func (o *CameraTexture) SetCameraFeedId(feedId gdnative.Int) {
	//log.Println("Calling CameraTexture.SetCameraFeedId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(feedId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CameraTexture", "set_camera_feed_id")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false which_feed int}], Returns: void
*/
func (o *CameraTexture) SetWhichFeed(whichFeed gdnative.Int) {
	//log.Println("Calling CameraTexture.SetWhichFeed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(whichFeed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CameraTexture", "set_which_feed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CameraTextureImplementer is an interface that implements the methods
// of the CameraTexture class.
type CameraTextureImplementer interface {
	TextureImplementer
	GetCameraActive() gdnative.Bool
	GetCameraFeedId() gdnative.Int
	SetCameraActive(active gdnative.Bool)
	SetCameraFeedId(feedId gdnative.Int)
	SetWhichFeed(whichFeed gdnative.Int)
}
