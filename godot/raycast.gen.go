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

//func NewRayCastFromPointer(ptr gdnative.Pointer) RayCast {
func newRayCastFromPointer(ptr gdnative.Pointer) RayCast {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := RayCast{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A RayCast represents a line from its origin to its destination position, [code]cast_to[/code]. It is used to query the 3D space in order to find the closest object along the path of the ray. RayCast can ignore some objects by adding them to the exception list via [code]add_exception[/code] or by setting proper filtering with collision layers and masks. RayCast can be configured to report collisions with [Area]s ([member collide_with_areas]) and/or [PhysicsBody]s ([member collide_with_bodies]). Only enabled raycasts will be able to query the space and report collisions. RayCast calculates intersection every physics frame (see [Node]), and the result is cached so it can be used later until the next frame. If multiple queries are required between physics frames (or during the same frame), use [method force_raycast_update] after adjusting the raycast.
*/
type RayCast struct {
	Spatial
	owner gdnative.Object
}

func (o *RayCast) BaseClass() string {
	return "RayCast"
}

/*
        Adds a collision exception so the ray does not report collisions with the specified node.
	Args: [{ false node Object}], Returns: void
*/
func (o *RayCast) AddException(node ObjectImplementer) {
	//log.Println("Calling RayCast.AddException()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "add_exception")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a collision exception so the ray does not report collisions with the specified [RID].
	Args: [{ false rid RID}], Returns: void
*/
func (o *RayCast) AddExceptionRid(rid gdnative.Rid) {
	//log.Println("Calling RayCast.AddExceptionRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRid(rid)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "add_exception_rid")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all collision exceptions for this ray.
	Args: [], Returns: void
*/
func (o *RayCast) ClearExceptions() {
	//log.Println("Calling RayCast.ClearExceptions()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "clear_exceptions")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Updates the collision information for the ray. Use this method to update the collision information immediately instead of waiting for the next [code]_physics_process[/code] call, for example if the ray or its parent has changed state. [b]Note:[/b] [code]enabled == true[/code] is not required for this to work.
	Args: [], Returns: void
*/
func (o *RayCast) ForceRaycastUpdate() {
	//log.Println("Calling RayCast.ForceRaycastUpdate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "force_raycast_update")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Vector3
*/
func (o *RayCast) GetCastTo() gdnative.Vector3 {
	//log.Println("Calling RayCast.GetCastTo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_cast_to")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the first object that the ray intersects, or [code]null[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
	Args: [], Returns: Object
*/
func (o *RayCast) GetCollider() ObjectImplementer {
	//log.Println("Calling RayCast.GetCollider()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collider")

	// Call the parent method.
	// Object
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newObjectFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ObjectImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Object" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ObjectImplementer)
	}

	return &ret
}

/*
        Returns the shape ID of the first object that the ray intersects, or [code]0[/code] if no object is intersecting the ray (i.e. [method is_colliding] returns [code]false[/code]).
	Args: [], Returns: int
*/
func (o *RayCast) GetColliderShape() gdnative.Int {
	//log.Println("Calling RayCast.GetColliderShape()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collider_shape")

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
	Args: [], Returns: int
*/
func (o *RayCast) GetCollisionMask() gdnative.Int {
	//log.Println("Calling RayCast.GetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collision_mask")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the bit index passed is turned on. [b]Note:[/b] Bit indices range from 0-19.
	Args: [{ false bit int}], Returns: bool
*/
func (o *RayCast) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	//log.Println("Calling RayCast.GetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collision_mask_bit")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns the normal of the intersecting object's shape at the collision point.
	Args: [], Returns: Vector3
*/
func (o *RayCast) GetCollisionNormal() gdnative.Vector3 {
	//log.Println("Calling RayCast.GetCollisionNormal()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collision_normal")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Returns the collision point at which the ray intersects the closest object. [b]Note:[/b] This point is in the [b]global[/b] coordinate system.
	Args: [], Returns: Vector3
*/
func (o *RayCast) GetCollisionPoint() gdnative.Vector3 {
	//log.Println("Calling RayCast.GetCollisionPoint()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_collision_point")

	// Call the parent method.
	// Vector3
	retPtr := gdnative.NewEmptyVector3()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVector3FromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *RayCast) GetExcludeParentBody() gdnative.Bool {
	//log.Println("Calling RayCast.GetExcludeParentBody()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "get_exclude_parent_body")

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
func (o *RayCast) IsCollideWithAreasEnabled() gdnative.Bool {
	//log.Println("Calling RayCast.IsCollideWithAreasEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "is_collide_with_areas_enabled")

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
func (o *RayCast) IsCollideWithBodiesEnabled() gdnative.Bool {
	//log.Println("Calling RayCast.IsCollideWithBodiesEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "is_collide_with_bodies_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns whether any object is intersecting with the ray's vector (considering the vector length).
	Args: [], Returns: bool
*/
func (o *RayCast) IsColliding() gdnative.Bool {
	//log.Println("Calling RayCast.IsColliding()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "is_colliding")

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
func (o *RayCast) IsEnabled() gdnative.Bool {
	//log.Println("Calling RayCast.IsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "is_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes a collision exception so the ray does report collisions with the specified node.
	Args: [{ false node Object}], Returns: void
*/
func (o *RayCast) RemoveException(node ObjectImplementer) {
	//log.Println("Calling RayCast.RemoveException()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "remove_exception")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes a collision exception so the ray does report collisions with the specified [RID].
	Args: [{ false rid RID}], Returns: void
*/
func (o *RayCast) RemoveExceptionRid(rid gdnative.Rid) {
	//log.Println("Calling RayCast.RemoveExceptionRid()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromRid(rid)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "remove_exception_rid")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false local_point Vector3}], Returns: void
*/
func (o *RayCast) SetCastTo(localPoint gdnative.Vector3) {
	//log.Println("Calling RayCast.SetCastTo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(localPoint)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_cast_to")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *RayCast) SetCollideWithAreas(enable gdnative.Bool) {
	//log.Println("Calling RayCast.SetCollideWithAreas()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_collide_with_areas")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *RayCast) SetCollideWithBodies(enable gdnative.Bool) {
	//log.Println("Calling RayCast.SetCollideWithBodies()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_collide_with_bodies")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask int}], Returns: void
*/
func (o *RayCast) SetCollisionMask(mask gdnative.Int) {
	//log.Println("Calling RayCast.SetCollisionMask()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_collision_mask")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the bit index passed to the [code]value[/code] passed. [b]Note:[/b] Bit indexes range from 0-19.
	Args: [{ false bit int} { false value bool}], Returns: void
*/
func (o *RayCast) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	//log.Println("Calling RayCast.SetCollisionMaskBit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(bit)
	ptrArguments[1] = gdnative.NewPointerFromBool(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_collision_mask_bit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *RayCast) SetEnabled(enabled gdnative.Bool) {
	//log.Println("Calling RayCast.SetEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mask bool}], Returns: void
*/
func (o *RayCast) SetExcludeParentBody(mask gdnative.Bool) {
	//log.Println("Calling RayCast.SetExcludeParentBody()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(mask)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("RayCast", "set_exclude_parent_body")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// RayCastImplementer is an interface that implements the methods
// of the RayCast class.
type RayCastImplementer interface {
	SpatialImplementer
	AddException(node ObjectImplementer)
	AddExceptionRid(rid gdnative.Rid)
	ClearExceptions()
	ForceRaycastUpdate()
	GetCastTo() gdnative.Vector3
	GetCollider() ObjectImplementer
	GetColliderShape() gdnative.Int
	GetCollisionMask() gdnative.Int
	GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool
	GetCollisionNormal() gdnative.Vector3
	GetCollisionPoint() gdnative.Vector3
	GetExcludeParentBody() gdnative.Bool
	IsCollideWithAreasEnabled() gdnative.Bool
	IsCollideWithBodiesEnabled() gdnative.Bool
	IsColliding() gdnative.Bool
	IsEnabled() gdnative.Bool
	RemoveException(node ObjectImplementer)
	RemoveExceptionRid(rid gdnative.Rid)
	SetCastTo(localPoint gdnative.Vector3)
	SetCollideWithAreas(enable gdnative.Bool)
	SetCollideWithBodies(enable gdnative.Bool)
	SetCollisionMask(mask gdnative.Int)
	SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool)
	SetEnabled(enabled gdnative.Bool)
	SetExcludeParentBody(mask gdnative.Bool)
}
