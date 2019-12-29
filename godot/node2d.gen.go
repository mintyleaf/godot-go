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

//func NewNode2DFromPointer(ptr gdnative.Pointer) Node2D {
func newNode2DFromPointer(ptr gdnative.Pointer) Node2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Node2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A 2D game object, with a transform (position, rotation, and scale). All 2D nodes, including physics objects and sprites, inherit from Node2D. Use Node2D as a parent node to move, scale and rotate children in a 2D project. Also gives control of the node's render order.
*/
type Node2D struct {
	CanvasItem
	owner gdnative.Object
}

func (o *Node2D) BaseClass() string {
	return "Node2D"
}

/*
        Multiplies the current scale by the [code]ratio[/code] vector.
	Args: [{ false ratio Vector2}], Returns: void
*/
func (o *Node2D) ApplyScale(ratio gdnative.Vector2) {
	//log.Println("Calling Node2D.ApplyScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(ratio)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "apply_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns the angle between the node and the [code]point[/code] in radians.
	Args: [{ false point Vector2}], Returns: float
*/
func (o *Node2D) GetAngleTo(point gdnative.Vector2) gdnative.Real {
	//log.Println("Calling Node2D.GetAngleTo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_angle_to")

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
	Args: [], Returns: Vector2
*/
func (o *Node2D) GetGlobalPosition() gdnative.Vector2 {
	//log.Println("Calling Node2D.GetGlobalPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_global_position")

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
	Args: [], Returns: float
*/
func (o *Node2D) GetGlobalRotation() gdnative.Real {
	//log.Println("Calling Node2D.GetGlobalRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_global_rotation")

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
	Args: [], Returns: float
*/
func (o *Node2D) GetGlobalRotationDegrees() gdnative.Real {
	//log.Println("Calling Node2D.GetGlobalRotationDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_global_rotation_degrees")

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
	Args: [], Returns: Vector2
*/
func (o *Node2D) GetGlobalScale() gdnative.Vector2 {
	//log.Println("Calling Node2D.GetGlobalScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_global_scale")

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
	Args: [], Returns: Vector2
*/
func (o *Node2D) GetPosition() gdnative.Vector2 {
	//log.Println("Calling Node2D.GetPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_position")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Returns the [Transform2D] relative to this node's parent.
	Args: [{ false parent Node}], Returns: Transform2D
*/
func (o *Node2D) GetRelativeTransformToParent(parent NodeImplementer) gdnative.Transform2D {
	//log.Println("Calling Node2D.GetRelativeTransformToParent()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(parent.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_relative_transform_to_parent")

	// Call the parent method.
	// Transform2D
	retPtr := gdnative.NewEmptyTransform2D()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransform2DFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *Node2D) GetRotation() gdnative.Real {
	//log.Println("Calling Node2D.GetRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_rotation")

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
	Args: [], Returns: float
*/
func (o *Node2D) GetRotationDegrees() gdnative.Real {
	//log.Println("Calling Node2D.GetRotationDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_rotation_degrees")

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
	Args: [], Returns: Vector2
*/
func (o *Node2D) GetScale() gdnative.Vector2 {
	//log.Println("Calling Node2D.GetScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_scale")

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
func (o *Node2D) GetZIndex() gdnative.Int {
	//log.Println("Calling Node2D.GetZIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "get_z_index")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Adds the [code]offset[/code] vector to the node's global position.
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *Node2D) GlobalTranslate(offset gdnative.Vector2) {
	//log.Println("Calling Node2D.GlobalTranslate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "global_translate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Node2D) IsZRelative() gdnative.Bool {
	//log.Println("Calling Node2D.IsZRelative()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "is_z_relative")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Rotates the node so it points towards the [code]point[/code], which is expected to use global coordinates.
	Args: [{ false point Vector2}], Returns: void
*/
func (o *Node2D) LookAt(point gdnative.Vector2) {
	//log.Println("Calling Node2D.LookAt()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "look_at")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Applies a local translation on the node's X axis based on the [method Node._process]'s [code]delta[/code]. If [code]scaled[/code] is [code]false[/code], normalizes the movement.
	Args: [{ false delta float} {False true scaled bool}], Returns: void
*/
func (o *Node2D) MoveLocalX(delta gdnative.Real, scaled gdnative.Bool) {
	//log.Println("Calling Node2D.MoveLocalX()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromReal(delta)
	ptrArguments[1] = gdnative.NewPointerFromBool(scaled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "move_local_x")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Applies a local translation on the node's Y axis based on the [method Node._process]'s [code]delta[/code]. If [code]scaled[/code] is [code]false[/code], normalizes the movement.
	Args: [{ false delta float} {False true scaled bool}], Returns: void
*/
func (o *Node2D) MoveLocalY(delta gdnative.Real, scaled gdnative.Bool) {
	//log.Println("Calling Node2D.MoveLocalY()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromReal(delta)
	ptrArguments[1] = gdnative.NewPointerFromBool(scaled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "move_local_y")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Applies a rotation to the node, in radians, starting from its current rotation.
	Args: [{ false radians float}], Returns: void
*/
func (o *Node2D) Rotate(radians gdnative.Real) {
	//log.Println("Calling Node2D.Rotate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(radians)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "rotate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false position Vector2}], Returns: void
*/
func (o *Node2D) SetGlobalPosition(position gdnative.Vector2) {
	//log.Println("Calling Node2D.SetGlobalPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_global_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radians float}], Returns: void
*/
func (o *Node2D) SetGlobalRotation(radians gdnative.Real) {
	//log.Println("Calling Node2D.SetGlobalRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(radians)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_global_rotation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *Node2D) SetGlobalRotationDegrees(degrees gdnative.Real) {
	//log.Println("Calling Node2D.SetGlobalRotationDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_global_rotation_degrees")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false scale Vector2}], Returns: void
*/
func (o *Node2D) SetGlobalScale(scale gdnative.Vector2) {
	//log.Println("Calling Node2D.SetGlobalScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(scale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_global_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false xform Transform2D}], Returns: void
*/
func (o *Node2D) SetGlobalTransform(xform gdnative.Transform2D) {
	//log.Println("Calling Node2D.SetGlobalTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(xform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_global_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false position Vector2}], Returns: void
*/
func (o *Node2D) SetPosition(position gdnative.Vector2) {
	//log.Println("Calling Node2D.SetPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false radians float}], Returns: void
*/
func (o *Node2D) SetRotation(radians gdnative.Real) {
	//log.Println("Calling Node2D.SetRotation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(radians)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_rotation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false degrees float}], Returns: void
*/
func (o *Node2D) SetRotationDegrees(degrees gdnative.Real) {
	//log.Println("Calling Node2D.SetRotationDegrees()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(degrees)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_rotation_degrees")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false scale Vector2}], Returns: void
*/
func (o *Node2D) SetScale(scale gdnative.Vector2) {
	//log.Println("Calling Node2D.SetScale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(scale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_scale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false xform Transform2D}], Returns: void
*/
func (o *Node2D) SetTransform(xform gdnative.Transform2D) {
	//log.Println("Calling Node2D.SetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromTransform2D(xform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *Node2D) SetZAsRelative(enable gdnative.Bool) {
	//log.Println("Calling Node2D.SetZAsRelative()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_z_as_relative")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false z_index int}], Returns: void
*/
func (o *Node2D) SetZIndex(zIndex gdnative.Int) {
	//log.Println("Calling Node2D.SetZIndex()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(zIndex)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "set_z_index")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Converts a local point's coordinates to global coordinates.
	Args: [{ false local_point Vector2}], Returns: Vector2
*/
func (o *Node2D) ToGlobal(localPoint gdnative.Vector2) gdnative.Vector2 {
	//log.Println("Calling Node2D.ToGlobal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(localPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "to_global")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Converts a global point's coordinates to local coordinates.
	Args: [{ false global_point Vector2}], Returns: Vector2
*/
func (o *Node2D) ToLocal(globalPoint gdnative.Vector2) gdnative.Vector2 {
	//log.Println("Calling Node2D.ToLocal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(globalPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "to_local")

	// Call the parent method.
	// Vector2
	retPtr := gdnative.NewEmptyVector2()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector2FromPointer(retPtr)
	return ret
}

/*
        Translates the node by the given [code]offset[/code] in local coordinates.
	Args: [{ false offset Vector2}], Returns: void
*/
func (o *Node2D) Translate(offset gdnative.Vector2) {
	//log.Println("Calling Node2D.Translate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(offset)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Node2D", "translate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// Node2DImplementer is an interface that implements the methods
// of the Node2D class.
type Node2DImplementer interface {
	CanvasItemImplementer
	ApplyScale(ratio gdnative.Vector2)
	GetAngleTo(point gdnative.Vector2) gdnative.Real
	GetGlobalPosition() gdnative.Vector2
	GetGlobalRotation() gdnative.Real
	GetGlobalRotationDegrees() gdnative.Real
	GetGlobalScale() gdnative.Vector2
	GetPosition() gdnative.Vector2
	GetRelativeTransformToParent(parent NodeImplementer) gdnative.Transform2D
	GetRotation() gdnative.Real
	GetRotationDegrees() gdnative.Real
	GetScale() gdnative.Vector2
	GetZIndex() gdnative.Int
	GlobalTranslate(offset gdnative.Vector2)
	IsZRelative() gdnative.Bool
	LookAt(point gdnative.Vector2)
	MoveLocalX(delta gdnative.Real, scaled gdnative.Bool)
	MoveLocalY(delta gdnative.Real, scaled gdnative.Bool)
	Rotate(radians gdnative.Real)
	SetGlobalPosition(position gdnative.Vector2)
	SetGlobalRotation(radians gdnative.Real)
	SetGlobalRotationDegrees(degrees gdnative.Real)
	SetGlobalScale(scale gdnative.Vector2)
	SetGlobalTransform(xform gdnative.Transform2D)
	SetPosition(position gdnative.Vector2)
	SetRotation(radians gdnative.Real)
	SetRotationDegrees(degrees gdnative.Real)
	SetScale(scale gdnative.Vector2)
	SetTransform(xform gdnative.Transform2D)
	SetZAsRelative(enable gdnative.Bool)
	SetZIndex(zIndex gdnative.Int)
	ToGlobal(localPoint gdnative.Vector2) gdnative.Vector2
	ToLocal(globalPoint gdnative.Vector2) gdnative.Vector2
	Translate(offset gdnative.Vector2)
}
