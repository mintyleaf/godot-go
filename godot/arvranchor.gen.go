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

//func NewARVRAnchorFromPointer(ptr gdnative.Pointer) ARVRAnchor {
func newARVRAnchorFromPointer(ptr gdnative.Pointer) ARVRAnchor {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ARVRAnchor{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The ARVR Anchor point is a spatial node that maps a real world location identified by the AR platform to a position within the game world. For example, as long as plane detection in ARKit is on, ARKit will identify and update the position of planes (tables, floors, etc) and create anchors for them. This node is mapped to one of the anchors through its unique ID. When you receive a signal that a new anchor is available, you should add this node to your scene for that anchor. You can predefine nodes and set the ID; the nodes will simply remain on 0,0,0 until a plane is recognized. Keep in mind that, as long as plane detection is enabled, the size, placing and orientation of an anchor will be updated as the detection logic learns more about the real world out there especially if only part of the surface is in view.
*/
type ARVRAnchor struct {
	Spatial
	owner gdnative.Object
}

func (o *ARVRAnchor) BaseClass() string {
	return "ARVRAnchor"
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *ARVRAnchor) GetAnchorId() gdnative.Int {
	//log.Println("Calling ARVRAnchor.GetAnchorId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRAnchor", "get_anchor_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the name given to this anchor.
	Args: [], Returns: String
*/
func (o *ARVRAnchor) GetAnchorName() gdnative.String {
	//log.Println("Calling ARVRAnchor.GetAnchorName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRAnchor", "get_anchor_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the anchor is being tracked and [code]false[/code] if no anchor with this ID is currently known.
	Args: [], Returns: bool
*/
func (o *ARVRAnchor) GetIsActive() gdnative.Bool {
	//log.Println("Calling ARVRAnchor.GetIsActive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRAnchor", "get_is_active")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        If provided by the ARVR Interface, this returns a mesh object for the anchor. For an anchor, this can be a shape related to the object being tracked or it can be a mesh that provides topology related to the anchor and can be used to create shadows/reflections on surfaces or for generating collision shapes.
	Args: [], Returns: Mesh
*/
func (o *ARVRAnchor) GetMesh() MeshImplementer {
	//log.Println("Calling ARVRAnchor.GetMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRAnchor", "get_mesh")

	// Call the parent method.
	// Mesh
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newMeshFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(MeshImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Mesh" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(MeshImplementer)
	}

	return &ret
}

/*
        Returns a plane aligned with our anchor; handy for intersection testing.
	Args: [], Returns: Plane
*/
func (o *ARVRAnchor) GetPlane() gdnative.Plane {
	//log.Println("Calling ARVRAnchor.GetPlane()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRAnchor", "get_plane")

	// Call the parent method.
	// Plane
	retPtr := gdnative.NewEmptyPlane()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPlaneFromPointer(retPtr)
	return ret
}

/*
        Returns the estimated size of the plane that was detected. Say when the anchor relates to a table in the real world, this is the estimated size of the surface of that table.
	Args: [], Returns: Vector3
*/
func (o *ARVRAnchor) GetSize() gdnative.Vector3 {
	//log.Println("Calling ARVRAnchor.GetSize()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRAnchor", "get_size")

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
	Args: [{ false anchor_id int}], Returns: void
*/
func (o *ARVRAnchor) SetAnchorId(anchorId gdnative.Int) {
	//log.Println("Calling ARVRAnchor.SetAnchorId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(anchorId)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ARVRAnchor", "set_anchor_id")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ARVRAnchorImplementer is an interface that implements the methods
// of the ARVRAnchor class.
type ARVRAnchorImplementer interface {
	SpatialImplementer
	GetAnchorId() gdnative.Int
	GetAnchorName() gdnative.String
	GetIsActive() gdnative.Bool
	GetMesh() MeshImplementer
	GetPlane() gdnative.Plane
	GetSize() gdnative.Vector3
	SetAnchorId(anchorId gdnative.Int)
}
