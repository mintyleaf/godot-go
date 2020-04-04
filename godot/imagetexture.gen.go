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

// ImageTextureStorage is an enum for Storage values.
type ImageTextureStorage int

const (
	ImageTextureStorageCompressLossless ImageTextureStorage = 2
	ImageTextureStorageCompressLossy    ImageTextureStorage = 1
	ImageTextureStorageRaw              ImageTextureStorage = 0
)

//func NewImageTextureFromPointer(ptr gdnative.Pointer) ImageTexture {
func newImageTextureFromPointer(ptr gdnative.Pointer) ImageTexture {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ImageTexture{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A [Texture2D] based on an [Image]. Can be created from an [Image] with [method create_from_image].
*/
type ImageTexture struct {
	Texture
	owner gdnative.Object
}

func (o *ImageTexture) BaseClass() string {
	return "ImageTexture"
}

/*
        Undocumented
	Args: [{ false rid RID}], Returns: void
*/
func (o *ImageTexture) X_ReloadHook(rid gdnative.Rid) {
	//log.Println("Calling ImageTexture.X_ReloadHook()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRid(rid)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "_reload_hook")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false width int} { false height int} { false format int} {7 true flags int}], Returns: void
*/
func (o *ImageTexture) Create(width gdnative.Int, height gdnative.Int, format gdnative.Int, flags gdnative.Int) {
	//log.Println("Calling ImageTexture.Create()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromInt(width)
	ptrArguments[1] = gdnative.NewPointerFromInt(height)
	ptrArguments[2] = gdnative.NewPointerFromInt(format)
	ptrArguments[3] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "create")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Create a new [ImageTexture] from an [Image].
	Args: [{ false image Image} {7 true flags int}], Returns: void
*/
func (o *ImageTexture) CreateFromImage(image ImageImplementer, flags gdnative.Int) {
	//log.Println("Calling ImageTexture.CreateFromImage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(image.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(flags)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "create_from_image")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the format of the [ImageTexture], one of [enum Image.Format].
	Args: [], Returns: enum.Image::Format
*/
func (o *ImageTexture) GetFormat() ImageFormat {
	//log.Println("Calling ImageTexture.GetFormat()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "get_format")

	// Call the parent method.
	// enum.Image::Format
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ImageFormat(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *ImageTexture) GetLossyStorageQuality() gdnative.Real {
	//log.Println("Calling ImageTexture.GetLossyStorageQuality()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "get_lossy_storage_quality")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.ImageTexture::Storage
*/
func (o *ImageTexture) GetStorage() ImageTextureStorage {
	//log.Println("Calling ImageTexture.GetStorage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "get_storage")

	// Call the parent method.
	// enum.ImageTexture::Storage
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ImageTextureStorage(ret)
}

/*
        Undocumented
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *ImageTexture) Load(path gdnative.String) gdnative.Error {
	//log.Println("Calling ImageTexture.Load()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "load")

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
	Args: [{ false image Image}], Returns: void
*/
func (o *ImageTexture) SetData(image ImageImplementer) {
	//log.Println("Calling ImageTexture.SetData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(image.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "set_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false quality float}], Returns: void
*/
func (o *ImageTexture) SetLossyStorageQuality(quality gdnative.Real) {
	//log.Println("Calling ImageTexture.SetLossyStorageQuality()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(quality)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "set_lossy_storage_quality")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Resizes the [ImageTexture] to the specified dimensions.
	Args: [{ false size Vector2}], Returns: void
*/
func (o *ImageTexture) SetSizeOverride(size gdnative.Vector2) {
	//log.Println("Calling ImageTexture.SetSizeOverride()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "set_size_override")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *ImageTexture) SetStorage(mode gdnative.Int) {
	//log.Println("Calling ImageTexture.SetStorage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ImageTexture", "set_storage")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ImageTextureImplementer is an interface that implements the methods
// of the ImageTexture class.
type ImageTextureImplementer interface {
	TextureImplementer
	X_ReloadHook(rid gdnative.Rid)
	Create(width gdnative.Int, height gdnative.Int, format gdnative.Int, flags gdnative.Int)
	CreateFromImage(image ImageImplementer, flags gdnative.Int)
	GetLossyStorageQuality() gdnative.Real
	SetData(image ImageImplementer)
	SetLossyStorageQuality(quality gdnative.Real)
	SetSizeOverride(size gdnative.Vector2)
	SetStorage(mode gdnative.Int)
}
