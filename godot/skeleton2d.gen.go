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

//func NewSkeleton2DFromPointer(ptr gdnative.Pointer) Skeleton2D {
func newSkeleton2DFromPointer(ptr gdnative.Pointer) Skeleton2D {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Skeleton2D{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Skeleton2D parents a hierarchy of [Bone2D] objects. It is a requirement of [Bone2D]. Skeleton2D holds a reference to the rest pose of its children and acts as a single point of access to its bones.
*/
type Skeleton2D struct {
	Node2D
	owner gdnative.Object
}

func (o *Skeleton2D) BaseClass() string {
	return "Skeleton2D"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Skeleton2D) X_UpdateBoneSetup() {
	//log.Println("Calling Skeleton2D.X_UpdateBoneSetup()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Skeleton2D", "_update_bone_setup")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Skeleton2D) X_UpdateTransform() {
	//log.Println("Calling Skeleton2D.X_UpdateTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Skeleton2D", "_update_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns a [Bone2D] from the node hierarchy parented by Skeleton2D. The object to return is identified by the parameter [code]idx[/code]. Bones are indexed by descending the node hierarchy from top to bottom, adding the children of each branch before moving to the next sibling.
	Args: [{ false idx int}], Returns: Bone2D
*/
func (o *Skeleton2D) GetBone(idx gdnative.Int) Bone2DImplementer {
	//log.Println("Calling Skeleton2D.GetBone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Skeleton2D", "get_bone")

	// Call the parent method.
	// Bone2D
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newBone2DFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(Bone2DImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Bone2D" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(Bone2DImplementer)
	}

	return &ret
}

/*
        Returns the number of [Bone2D] nodes in the node hierarchy parented by Skeleton2D.
	Args: [], Returns: int
*/
func (o *Skeleton2D) GetBoneCount() gdnative.Int {
	//log.Println("Calling Skeleton2D.GetBoneCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Skeleton2D", "get_bone_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns the [RID] of a Skeleton2D instance.
	Args: [], Returns: RID
*/
func (o *Skeleton2D) GetSkeleton() gdnative.Rid {
	//log.Println("Calling Skeleton2D.GetSkeleton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Skeleton2D", "get_skeleton")

	// Call the parent method.
	// RID
	retPtr := gdnative.NewEmptyRid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRidFromPointer(retPtr)
	return ret
}

// Skeleton2DImplementer is an interface that implements the methods
// of the Skeleton2D class.
type Skeleton2DImplementer interface {
	Node2DImplementer
	X_UpdateBoneSetup()
	X_UpdateTransform()
	GetBone(idx gdnative.Int) Bone2DImplementer
	GetBoneCount() gdnative.Int
	GetSkeleton() gdnative.Rid
}
