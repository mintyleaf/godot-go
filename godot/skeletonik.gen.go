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

//func NewSkeletonIKFromPointer(ptr gdnative.Pointer) SkeletonIK {
func newSkeletonIKFromPointer(ptr gdnative.Pointer) SkeletonIK {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SkeletonIK{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type SkeletonIK struct {
	Node
	owner gdnative.Object
}

func (o *SkeletonIK) BaseClass() string {
	return "SkeletonIK"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *SkeletonIK) GetInterpolation() gdnative.Real {
	//log.Println("Calling SkeletonIK.GetInterpolation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "get_interpolation")

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
	Args: [], Returns: Vector3
*/
func (o *SkeletonIK) GetMagnetPosition() gdnative.Vector3 {
	//log.Println("Calling SkeletonIK.GetMagnetPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "get_magnet_position")

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
	Args: [], Returns: int
*/
func (o *SkeletonIK) GetMaxIterations() gdnative.Int {
	//log.Println("Calling SkeletonIK.GetMaxIterations()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "get_max_iterations")

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
	Args: [], Returns: float
*/
func (o *SkeletonIK) GetMinDistance() gdnative.Real {
	//log.Println("Calling SkeletonIK.GetMinDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "get_min_distance")

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
	Args: [], Returns: Skeleton
*/
func (o *SkeletonIK) GetParentSkeleton() SkeletonImplementer {
	//log.Println("Calling SkeletonIK.GetParentSkeleton()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "get_parent_skeleton")

	// Call the parent method.
	// Skeleton
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newSkeletonFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(SkeletonImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Skeleton" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(SkeletonImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *SkeletonIK) GetRootBone() gdnative.String {
	//log.Println("Calling SkeletonIK.GetRootBone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "get_root_bone")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: NodePath
*/
func (o *SkeletonIK) GetTargetNode() gdnative.NodePath {
	//log.Println("Calling SkeletonIK.GetTargetNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "get_target_node")

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
	Args: [], Returns: Transform
*/
func (o *SkeletonIK) GetTargetTransform() gdnative.Transform {
	//log.Println("Calling SkeletonIK.GetTargetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "get_target_transform")

	// Call the parent method.
	// Transform
	retPtr := gdnative.NewEmptyTransform()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewTransformFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *SkeletonIK) GetTipBone() gdnative.String {
	//log.Println("Calling SkeletonIK.GetTipBone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "get_tip_bone")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *SkeletonIK) IsOverrideTipBasis() gdnative.Bool {
	//log.Println("Calling SkeletonIK.IsOverrideTipBasis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "is_override_tip_basis")

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
func (o *SkeletonIK) IsRunning() gdnative.Bool {
	//log.Println("Calling SkeletonIK.IsRunning()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "is_running")

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
func (o *SkeletonIK) IsUsingMagnet() gdnative.Bool {
	//log.Println("Calling SkeletonIK.IsUsingMagnet()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "is_using_magnet")

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
	Args: [{ false interpolation float}], Returns: void
*/
func (o *SkeletonIK) SetInterpolation(interpolation gdnative.Real) {
	//log.Println("Calling SkeletonIK.SetInterpolation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(interpolation)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_interpolation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false local_position Vector3}], Returns: void
*/
func (o *SkeletonIK) SetMagnetPosition(localPosition gdnative.Vector3) {
	//log.Println("Calling SkeletonIK.SetMagnetPosition()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector3(localPosition)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_magnet_position")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false iterations int}], Returns: void
*/
func (o *SkeletonIK) SetMaxIterations(iterations gdnative.Int) {
	//log.Println("Calling SkeletonIK.SetMaxIterations()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(iterations)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_max_iterations")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false min_distance float}], Returns: void
*/
func (o *SkeletonIK) SetMinDistance(minDistance gdnative.Real) {
	//log.Println("Calling SkeletonIK.SetMinDistance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(minDistance)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_min_distance")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false override bool}], Returns: void
*/
func (o *SkeletonIK) SetOverrideTipBasis(override gdnative.Bool) {
	//log.Println("Calling SkeletonIK.SetOverrideTipBasis()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(override)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_override_tip_basis")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false root_bone String}], Returns: void
*/
func (o *SkeletonIK) SetRootBone(rootBone gdnative.String) {
	//log.Println("Calling SkeletonIK.SetRootBone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(rootBone)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_root_bone")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false node NodePath}], Returns: void
*/
func (o *SkeletonIK) SetTargetNode(node gdnative.NodePath) {
	//log.Println("Calling SkeletonIK.SetTargetNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(node)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_target_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false target Transform}], Returns: void
*/
func (o *SkeletonIK) SetTargetTransform(target gdnative.Transform) {
	//log.Println("Calling SkeletonIK.SetTargetTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromTransform(target)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_target_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false tip_bone String}], Returns: void
*/
func (o *SkeletonIK) SetTipBone(tipBone gdnative.String) {
	//log.Println("Calling SkeletonIK.SetTipBone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(tipBone)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_tip_bone")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false use bool}], Returns: void
*/
func (o *SkeletonIK) SetUseMagnet(use gdnative.Bool) {
	//log.Println("Calling SkeletonIK.SetUseMagnet()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(use)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "set_use_magnet")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{False true one_time bool}], Returns: void
*/
func (o *SkeletonIK) Start(oneTime gdnative.Bool) {
	//log.Println("Calling SkeletonIK.Start()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(oneTime)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "start")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *SkeletonIK) Stop() {
	//log.Println("Calling SkeletonIK.Stop()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SkeletonIK", "stop")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// SkeletonIKImplementer is an interface that implements the methods
// of the SkeletonIK class.
type SkeletonIKImplementer interface {
	NodeImplementer
	GetInterpolation() gdnative.Real
	GetMagnetPosition() gdnative.Vector3
	GetMaxIterations() gdnative.Int
	GetMinDistance() gdnative.Real
	GetParentSkeleton() SkeletonImplementer
	GetRootBone() gdnative.String
	GetTargetNode() gdnative.NodePath
	GetTargetTransform() gdnative.Transform
	GetTipBone() gdnative.String
	IsOverrideTipBasis() gdnative.Bool
	IsRunning() gdnative.Bool
	IsUsingMagnet() gdnative.Bool
	SetInterpolation(interpolation gdnative.Real)
	SetMagnetPosition(localPosition gdnative.Vector3)
	SetMaxIterations(iterations gdnative.Int)
	SetMinDistance(minDistance gdnative.Real)
	SetOverrideTipBasis(override gdnative.Bool)
	SetRootBone(rootBone gdnative.String)
	SetTargetNode(node gdnative.NodePath)
	SetTargetTransform(target gdnative.Transform)
	SetTipBone(tipBone gdnative.String)
	SetUseMagnet(use gdnative.Bool)
	Start(oneTime gdnative.Bool)
	Stop()
}
