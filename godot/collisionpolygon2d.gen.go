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

// CollisionPolygon2DBuildMode is an enum for BuildMode values.
type CollisionPolygon2DBuildMode int

const (
	CollisionPolygon2DBuildSegments CollisionPolygon2DBuildMode = 1
	CollisionPolygon2DBuildSolids   CollisionPolygon2DBuildMode = 0
)

//func NewCollisionPolygon2DFromPointer(ptr gdnative.Pointer) CollisionPolygon2D {
func newCollisionPolygon2DFromPointer(ptr gdnative.Pointer) CollisionPolygon2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CollisionPolygon2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Provides a 2D collision polygon to a [CollisionObject2D] parent. Polygons can be drawn in the editor or specified by a list of vertices.
*/
type CollisionPolygon2D struct {
	Node2D
	owner gdnative.Object
}

func (o *CollisionPolygon2D) BaseClass() string {
	return "CollisionPolygon2D"
}

/*
        Undocumented
	Args: [], Returns: enum.CollisionPolygon2D::BuildMode
*/
func (o *CollisionPolygon2D) GetBuildMode() CollisionPolygon2DBuildMode {
	//log.Println("Calling CollisionPolygon2D.GetBuildMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "get_build_mode")

	// Call the parent method.
	// enum.CollisionPolygon2D::BuildMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return CollisionPolygon2DBuildMode(ret)
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *CollisionPolygon2D) GetOneWayCollisionMargin() gdnative.Real {
	//log.Println("Calling CollisionPolygon2D.GetOneWayCollisionMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "get_one_way_collision_margin")

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
	Args: [], Returns: PoolVector2Array
*/
func (o *CollisionPolygon2D) GetPolygon() gdnative.PoolVector2Array {
	//log.Println("Calling CollisionPolygon2D.GetPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "get_polygon")

	// Call the parent method.
	// PoolVector2Array
	retPtr := gdnative.NewEmptyPoolVector2Array()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolVector2ArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *CollisionPolygon2D) IsDisabled() gdnative.Bool {
	//log.Println("Calling CollisionPolygon2D.IsDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "is_disabled")

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
func (o *CollisionPolygon2D) IsOneWayCollisionEnabled() gdnative.Bool {
	//log.Println("Calling CollisionPolygon2D.IsOneWayCollisionEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "is_one_way_collision_enabled")

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
	Args: [{ false build_mode int}], Returns: void
*/
func (o *CollisionPolygon2D) SetBuildMode(buildMode gdnative.Int) {
	//log.Println("Calling CollisionPolygon2D.SetBuildMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(buildMode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "set_build_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false disabled bool}], Returns: void
*/
func (o *CollisionPolygon2D) SetDisabled(disabled gdnative.Bool) {
	//log.Println("Calling CollisionPolygon2D.SetDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(disabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "set_disabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *CollisionPolygon2D) SetOneWayCollision(enabled gdnative.Bool) {
	//log.Println("Calling CollisionPolygon2D.SetOneWayCollision()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "set_one_way_collision")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false margin float}], Returns: void
*/
func (o *CollisionPolygon2D) SetOneWayCollisionMargin(margin gdnative.Real) {
	//log.Println("Calling CollisionPolygon2D.SetOneWayCollisionMargin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(margin)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "set_one_way_collision_margin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false polygon PoolVector2Array}], Returns: void
*/
func (o *CollisionPolygon2D) SetPolygon(polygon gdnative.PoolVector2Array) {
	//log.Println("Calling CollisionPolygon2D.SetPolygon()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector2Array(polygon)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CollisionPolygon2D", "set_polygon")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// CollisionPolygon2DImplementer is an interface that implements the methods
// of the CollisionPolygon2D class.
type CollisionPolygon2DImplementer interface {
	Node2DImplementer
	GetOneWayCollisionMargin() gdnative.Real
	GetPolygon() gdnative.PoolVector2Array
	IsDisabled() gdnative.Bool
	IsOneWayCollisionEnabled() gdnative.Bool
	SetBuildMode(buildMode gdnative.Int)
	SetDisabled(disabled gdnative.Bool)
	SetOneWayCollision(enabled gdnative.Bool)
	SetOneWayCollisionMargin(margin gdnative.Real)
	SetPolygon(polygon gdnative.PoolVector2Array)
}
