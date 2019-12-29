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

// TextureButtonStretchMode is an enum for StretchMode values.
type TextureButtonStretchMode int

const (
	TextureButtonStretchKeep               TextureButtonStretchMode = 2
	TextureButtonStretchKeepAspect         TextureButtonStretchMode = 4
	TextureButtonStretchKeepAspectCentered TextureButtonStretchMode = 5
	TextureButtonStretchKeepAspectCovered  TextureButtonStretchMode = 6
	TextureButtonStretchKeepCentered       TextureButtonStretchMode = 3
	TextureButtonStretchScale              TextureButtonStretchMode = 0
	TextureButtonStretchTile               TextureButtonStretchMode = 1
)

//func NewTextureButtonFromPointer(ptr gdnative.Pointer) TextureButton {
func newTextureButtonFromPointer(ptr gdnative.Pointer) TextureButton {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := TextureButton{}
	obj.SetBaseObject(owner)

	return obj
}

/*
[TextureButton] has the same functionality as [Button], except it uses sprites instead of Godot's [Theme] resource. It is faster to create, but it doesn't support localization like more complex [Control]s. The "normal" state must contain a texture ([member texture_normal]); other textures are optional.
*/
type TextureButton struct {
	BaseButton
	owner gdnative.Object
}

func (o *TextureButton) BaseClass() string {
	return "TextureButton"
}

/*
        Undocumented
	Args: [], Returns: BitMap
*/
func (o *TextureButton) GetClickMask() BitMapImplementer {
	//log.Println("Calling TextureButton.GetClickMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "get_click_mask")

	// Call the parent method.
	// BitMap
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newBitMapFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(BitMapImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "BitMap" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(BitMapImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: Texture
*/
func (o *TextureButton) GetDisabledTexture() TextureImplementer {
	//log.Println("Calling TextureButton.GetDisabledTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "get_disabled_texture")

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
	Args: [], Returns: bool
*/
func (o *TextureButton) GetExpand() gdnative.Bool {
	//log.Println("Calling TextureButton.GetExpand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "get_expand")

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
	Args: [], Returns: Texture
*/
func (o *TextureButton) GetFocusedTexture() TextureImplementer {
	//log.Println("Calling TextureButton.GetFocusedTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "get_focused_texture")

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
	Args: [], Returns: Texture
*/
func (o *TextureButton) GetHoverTexture() TextureImplementer {
	//log.Println("Calling TextureButton.GetHoverTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "get_hover_texture")

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
	Args: [], Returns: Texture
*/
func (o *TextureButton) GetNormalTexture() TextureImplementer {
	//log.Println("Calling TextureButton.GetNormalTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "get_normal_texture")

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
	Args: [], Returns: Texture
*/
func (o *TextureButton) GetPressedTexture() TextureImplementer {
	//log.Println("Calling TextureButton.GetPressedTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "get_pressed_texture")

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
	Args: [], Returns: enum.TextureButton::StretchMode
*/
func (o *TextureButton) GetStretchMode() TextureButtonStretchMode {
	//log.Println("Calling TextureButton.GetStretchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "get_stretch_mode")

	// Call the parent method.
	// enum.TextureButton::StretchMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return TextureButtonStretchMode(ret)
}

/*
        Undocumented
	Args: [{ false mask BitMap}], Returns: void
*/
func (o *TextureButton) SetClickMask(mask BitMapImplementer) {
	//log.Println("Calling TextureButton.SetClickMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(mask.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "set_click_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *TextureButton) SetDisabledTexture(texture TextureImplementer) {
	//log.Println("Calling TextureButton.SetDisabledTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "set_disabled_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false p_expand bool}], Returns: void
*/
func (o *TextureButton) SetExpand(pExpand gdnative.Bool) {
	//log.Println("Calling TextureButton.SetExpand()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(pExpand)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "set_expand")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *TextureButton) SetFocusedTexture(texture TextureImplementer) {
	//log.Println("Calling TextureButton.SetFocusedTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "set_focused_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *TextureButton) SetHoverTexture(texture TextureImplementer) {
	//log.Println("Calling TextureButton.SetHoverTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "set_hover_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *TextureButton) SetNormalTexture(texture TextureImplementer) {
	//log.Println("Calling TextureButton.SetNormalTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "set_normal_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false texture Texture}], Returns: void
*/
func (o *TextureButton) SetPressedTexture(texture TextureImplementer) {
	//log.Println("Calling TextureButton.SetPressedTexture()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(texture.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "set_pressed_texture")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false p_mode int}], Returns: void
*/
func (o *TextureButton) SetStretchMode(pMode gdnative.Int) {
	//log.Println("Calling TextureButton.SetStretchMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(pMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("TextureButton", "set_stretch_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// TextureButtonImplementer is an interface that implements the methods
// of the TextureButton class.
type TextureButtonImplementer interface {
	BaseButtonImplementer
	GetClickMask() BitMapImplementer
	GetDisabledTexture() TextureImplementer
	GetExpand() gdnative.Bool
	GetFocusedTexture() TextureImplementer
	GetHoverTexture() TextureImplementer
	GetNormalTexture() TextureImplementer
	GetPressedTexture() TextureImplementer
	SetClickMask(mask BitMapImplementer)
	SetDisabledTexture(texture TextureImplementer)
	SetExpand(pExpand gdnative.Bool)
	SetFocusedTexture(texture TextureImplementer)
	SetHoverTexture(texture TextureImplementer)
	SetNormalTexture(texture TextureImplementer)
	SetPressedTexture(texture TextureImplementer)
	SetStretchMode(pMode gdnative.Int)
}
