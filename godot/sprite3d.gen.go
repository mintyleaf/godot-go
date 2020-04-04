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

//func NewSprite3DFromPointer(ptr gdnative.Pointer) Sprite3D {
func newSprite3DFromPointer(ptr gdnative.Pointer) Sprite3D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Sprite3D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A node that displays a 2D texture in a 3D environment. The texture displayed can be a region from a larger atlas texture, or a frame from a sprite sheet animation. [b]Note:[/b] There are [url=https://github.com/godotengine/godot/issues/20855]known performance issues[/url] when using [Sprite3D]. Consider using a [MeshInstance3D] with a [QuadMesh] as the mesh instead. You can still have billboarding by enabling billboard properties in the QuadMesh's [StandardMaterial3D].
*/
type Sprite3D struct {
	SpriteBase3D
	owner gdnative.Object
}

func (o *Sprite3D) BaseClass() string {
	return "Sprite3D"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Sprite3D) GetFrame() gdnative.Int {
	//log.Println("Calling Sprite3D.GetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "get_frame")

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
	Args: [], Returns: Vector2
*/
func (o *Sprite3D) GetFrameCoords() gdnative.Vector2 {
	//log.Println("Calling Sprite3D.GetFrameCoords()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "get_frame_coords")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Sprite3D) GetHframes() gdnative.Int {
	//log.Println("Calling Sprite3D.GetHframes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "get_hframes")

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
	Args: [], Returns: Rect2
*/
func (o *Sprite3D) GetRegionRect() gdnative.Rect2 {
	//log.Println("Calling Sprite3D.GetRegionRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "get_region_rect")

	// Call the parent method.
	// Rect2
	retPtr := gdnative.NewEmptyRect2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRect2FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *Sprite3D) GetTexture() TextureImplementer {
	//log.Println("Calling Sprite3D.GetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "get_texture")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *Sprite3D) GetVframes() gdnative.Int {
	//log.Println("Calling Sprite3D.GetVframes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "get_vframes")

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
	Args: [], Returns: bool
*/
func (o *Sprite3D) IsRegion() gdnative.Bool {
	//log.Println("Calling Sprite3D.IsRegion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "is_region")

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
	Args: [{ false frame int}], Returns: void
*/
func (o *Sprite3D) SetFrame(frame gdnative.Int) {
	//log.Println("Calling Sprite3D.SetFrame()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(frame)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "set_frame")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false coords Vector2}], Returns: void
*/
func (o *Sprite3D) SetFrameCoords(coords gdnative.Vector2) {
	//log.Println("Calling Sprite3D.SetFrameCoords()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(coords)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "set_frame_coords")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false hframes int}], Returns: void
*/
func (o *Sprite3D) SetHframes(hframes gdnative.Int) {
	//log.Println("Calling Sprite3D.SetHframes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(hframes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "set_hframes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *Sprite3D) SetRegion(enabled gdnative.Bool) {
	//log.Println("Calling Sprite3D.SetRegion()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "set_region")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false rect Rect2}], Returns: void
*/
func (o *Sprite3D) SetRegionRect(rect gdnative.Rect2) {
	//log.Println("Calling Sprite3D.SetRegionRect()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRect2(rect)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "set_region_rect")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *Sprite3D) SetTexture(texture TextureImplementer) {
	//log.Println("Calling Sprite3D.SetTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "set_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false vframes int}], Returns: void
*/
func (o *Sprite3D) SetVframes(vframes gdnative.Int) {
	//log.Println("Calling Sprite3D.SetVframes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(vframes)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Sprite3D", "set_vframes")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Sprite3DImplementer is an interface that implements the methods
// of the Sprite3D class.
type Sprite3DImplementer interface {
	SpriteBase3DImplementer
	GetFrame() gdnative.Int
	GetFrameCoords() gdnative.Vector2
	GetHframes() gdnative.Int
	GetRegionRect() gdnative.Rect2
	GetTexture() TextureImplementer
	GetVframes() gdnative.Int
	IsRegion() gdnative.Bool
	SetFrame(frame gdnative.Int)
	SetFrameCoords(coords gdnative.Vector2)
	SetHframes(hframes gdnative.Int)
	SetRegion(enabled gdnative.Bool)
	SetRegionRect(rect gdnative.Rect2)
	SetTexture(texture TextureImplementer)
	SetVframes(vframes gdnative.Int)
}
