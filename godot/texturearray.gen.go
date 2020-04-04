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

//func NewTextureArrayFromPointer(ptr gdnative.Pointer) TextureArray {
func newTextureArrayFromPointer(ptr gdnative.Pointer) TextureArray {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := TextureArray{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type TextureArray struct {
	TextureLayered
	owner gdnative.Object
}

func (o *TextureArray) BaseClass() string {
	return "TextureArray"
}

// TextureArrayImplementer is an interface that implements the methods
// of the TextureArray class.
type TextureArrayImplementer interface {
	TextureLayeredImplementer
}
