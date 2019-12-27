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

//func NewRichTextEffectFromPointer(ptr gdnative.Pointer) RichTextEffect {
func newRichTextEffectFromPointer(ptr gdnative.Pointer) RichTextEffect {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RichTextEffect{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type RichTextEffect struct {
	Resource
	owner gdnative.Object
}

func (o *RichTextEffect) BaseClass() string {
	return "RichTextEffect"
}

/*
        Undocumented
	Args: [{ false char_fx CharFXTransform}], Returns: bool
*/
func (o *RichTextEffect) X_ProcessCustomFx(charFx CharFXTransformImplementer) gdnative.Bool {
	//log.Println("Calling RichTextEffect.X_ProcessCustomFx()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(charFx.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RichTextEffect", "_process_custom_fx")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

// RichTextEffectImplementer is an interface that implements the methods
// of the RichTextEffect class.
type RichTextEffectImplementer interface {
	ResourceImplementer
	X_ProcessCustomFx(charFx CharFXTransformImplementer) gdnative.Bool
}
