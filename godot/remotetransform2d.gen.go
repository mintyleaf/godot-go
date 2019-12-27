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

//func NewRemoteTransform2DFromPointer(ptr gdnative.Pointer) RemoteTransform2D {
func newRemoteTransform2DFromPointer(ptr gdnative.Pointer) RemoteTransform2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RemoteTransform2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
RemoteTransform2D pushes its own [Transform2D] to another [CanvasItem] derived Node (called the remote node) in the scene. It can be set to update another Node's position, rotation and/or scale. It can use either global or local coordinates.
*/
type RemoteTransform2D struct {
	Node2D
	owner gdnative.Object
}

func (o *RemoteTransform2D) BaseClass() string {
	return "RemoteTransform2D"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *RemoteTransform2D) ForceUpdateCache() {
	//log.Println("Calling RemoteTransform2D.ForceUpdateCache()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "force_update_cache")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *RemoteTransform2D) GetRemoteNode() gdnative.NodePath {
	//log.Println("Calling RemoteTransform2D.GetRemoteNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "get_remote_node")

	// Call the parent method.
	// NodePath
	retPtr := gdnative.NewEmptyNodePath()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewNodePathFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RemoteTransform2D) GetUpdatePosition() gdnative.Bool {
	//log.Println("Calling RemoteTransform2D.GetUpdatePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "get_update_position")

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
	Args: [], Returns: bool
*/
func (o *RemoteTransform2D) GetUpdateRotation() gdnative.Bool {
	//log.Println("Calling RemoteTransform2D.GetUpdateRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "get_update_rotation")

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
	Args: [], Returns: bool
*/
func (o *RemoteTransform2D) GetUpdateScale() gdnative.Bool {
	//log.Println("Calling RemoteTransform2D.GetUpdateScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "get_update_scale")

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
	Args: [], Returns: bool
*/
func (o *RemoteTransform2D) GetUseGlobalCoordinates() gdnative.Bool {
	//log.Println("Calling RemoteTransform2D.GetUseGlobalCoordinates()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "get_use_global_coordinates")

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
	Args: [{ false path NodePath}], Returns: void
*/
func (o *RemoteTransform2D) SetRemoteNode(path gdnative.NodePath) {
	//log.Println("Calling RemoteTransform2D.SetRemoteNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "set_remote_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false update_remote_position bool}], Returns: void
*/
func (o *RemoteTransform2D) SetUpdatePosition(updateRemotePosition gdnative.Bool) {
	//log.Println("Calling RemoteTransform2D.SetUpdatePosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(updateRemotePosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "set_update_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false update_remote_rotation bool}], Returns: void
*/
func (o *RemoteTransform2D) SetUpdateRotation(updateRemoteRotation gdnative.Bool) {
	//log.Println("Calling RemoteTransform2D.SetUpdateRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(updateRemoteRotation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "set_update_rotation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false update_remote_scale bool}], Returns: void
*/
func (o *RemoteTransform2D) SetUpdateScale(updateRemoteScale gdnative.Bool) {
	//log.Println("Calling RemoteTransform2D.SetUpdateScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(updateRemoteScale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "set_update_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false use_global_coordinates bool}], Returns: void
*/
func (o *RemoteTransform2D) SetUseGlobalCoordinates(useGlobalCoordinates gdnative.Bool) {
	//log.Println("Calling RemoteTransform2D.SetUseGlobalCoordinates()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(useGlobalCoordinates)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RemoteTransform2D", "set_use_global_coordinates")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// RemoteTransform2DImplementer is an interface that implements the methods
// of the RemoteTransform2D class.
type RemoteTransform2DImplementer interface {
	Node2DImplementer
	ForceUpdateCache()
	GetRemoteNode() gdnative.NodePath
	GetUpdatePosition() gdnative.Bool
	GetUpdateRotation() gdnative.Bool
	GetUpdateScale() gdnative.Bool
	GetUseGlobalCoordinates() gdnative.Bool
	SetRemoteNode(path gdnative.NodePath)
	SetUpdatePosition(updateRemotePosition gdnative.Bool)
	SetUpdateRotation(updateRemoteRotation gdnative.Bool)
	SetUpdateScale(updateRemoteScale gdnative.Bool)
	SetUseGlobalCoordinates(useGlobalCoordinates gdnative.Bool)
}
