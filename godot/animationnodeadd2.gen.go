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

//func NewAnimationNodeAdd2FromPointer(ptr gdnative.Pointer) AnimationNodeAdd2 {
func newAnimationNodeAdd2FromPointer(ptr gdnative.Pointer) AnimationNodeAdd2 {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNodeAdd2{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A resource to add to an [AnimationNodeBlendTree]. Blends two animations additively based on an amount value in the [code][0.0, 1.0][/code] range.
*/
type AnimationNodeAdd2 struct {
	AnimationNode
	owner gdnative.Object
}

func (o *AnimationNodeAdd2) BaseClass() string {
	return "AnimationNodeAdd2"
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *AnimationNodeAdd2) IsUsingSync() gdnative.Bool {
	//log.Println("Calling AnimationNodeAdd2.IsUsingSync()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeAdd2", "is_using_sync")

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
	Args: [{ false enable bool}], Returns: void
*/
func (o *AnimationNodeAdd2) SetUseSync(enable gdnative.Bool) {
	//log.Println("Calling AnimationNodeAdd2.SetUseSync()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNodeAdd2", "set_use_sync")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationNodeAdd2Implementer is an interface that implements the methods
// of the AnimationNodeAdd2 class.
type AnimationNodeAdd2Implementer interface {
	AnimationNodeImplementer
	IsUsingSync() gdnative.Bool
	SetUseSync(enable gdnative.Bool)
}
