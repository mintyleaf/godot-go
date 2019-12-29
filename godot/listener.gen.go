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

//func NewListenerFromPointer(ptr gdnative.Pointer) Listener {
func newListenerFromPointer(ptr gdnative.Pointer) Listener {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Listener{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Once added to the scene tree and enabled using [method make_current], this node will override the location sounds are heard from. This can be used to listen from a location different from the [Camera]. [b]Note:[/b] There is no 2D equivalent for this node yet.
*/
type Listener struct {
	Spatial
	owner gdnative.Object
}

func (o *Listener) BaseClass() string {
	return "Listener"
}

/*
        Disables the listener to use the current camera's listener instead.
	Args: [], Returns: void
*/
func (o *Listener) ClearCurrent() {
	//log.Println("Calling Listener.ClearCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Listener", "clear_current")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the listener's global orthonormalized [Transform].
	Args: [], Returns: Transform
*/
func (o *Listener) GetListenerTransform() gdnative.Transform {
	//log.Println("Calling Listener.GetListenerTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Listener", "get_listener_transform")

	// Call the parent method.
	// Transform
	retPtr := gdnative.NewEmptyTransform()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransformFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the listener was made current using [method make_current], [code]false[/code] otherwise. [b]Note:[/b] There may be more than one Listener marked as "current" in the scene tree, but only the one that was made current last will be used.
	Args: [], Returns: bool
*/
func (o *Listener) IsCurrent() gdnative.Bool {
	//log.Println("Calling Listener.IsCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Listener", "is_current")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Enables the listener. This will override the current camera's listener.
	Args: [], Returns: void
*/
func (o *Listener) MakeCurrent() {
	//log.Println("Calling Listener.MakeCurrent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Listener", "make_current")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ListenerImplementer is an interface that implements the methods
// of the Listener class.
type ListenerImplementer interface {
	SpatialImplementer
	ClearCurrent()
	GetListenerTransform() gdnative.Transform
	IsCurrent() gdnative.Bool
	MakeCurrent()
}
