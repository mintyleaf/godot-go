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

//func NewSceneTreeTimerFromPointer(ptr gdnative.Pointer) SceneTreeTimer {
func newSceneTreeTimerFromPointer(ptr gdnative.Pointer) SceneTreeTimer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SceneTreeTimer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A one-shot timer managed by the scene tree, which emits [signal timeout] on completion. See also [method SceneTree.create_timer]. As opposed to [Timer], it does not require the instantiation of a node. Commonly used to create a one-shot delay timer as in the following example: [codeblock] func some_function(): print("Timer started.") yield(get_tree().create_timer(1.0), "timeout") print("Timer ended.") [/codeblock]
*/
type SceneTreeTimer struct {
	Reference
	owner gdnative.Object
}

func (o *SceneTreeTimer) BaseClass() string {
	return "SceneTreeTimer"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *SceneTreeTimer) GetTimeLeft() gdnative.Real {
	//log.Println("Calling SceneTreeTimer.GetTimeLeft()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTreeTimer", "get_time_left")

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
	Args: [{ false time float}], Returns: void
*/
func (o *SceneTreeTimer) SetTimeLeft(time gdnative.Real) {
	//log.Println("Calling SceneTreeTimer.SetTimeLeft()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(time)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SceneTreeTimer", "set_time_left")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// SceneTreeTimerImplementer is an interface that implements the methods
// of the SceneTreeTimer class.
type SceneTreeTimerImplementer interface {
	ReferenceImplementer
	GetTimeLeft() gdnative.Real
	SetTimeLeft(time gdnative.Real)
}
