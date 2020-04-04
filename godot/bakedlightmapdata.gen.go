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

//func NewBakedLightmapDataFromPointer(ptr gdnative.Pointer) BakedLightmapData {
func newBakedLightmapDataFromPointer(ptr gdnative.Pointer) BakedLightmapData {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := BakedLightmapData{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type BakedLightmapData struct {
	Resource
	owner gdnative.Object
}

func (o *BakedLightmapData) BaseClass() string {
	return "BakedLightmapData"
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *BakedLightmapData) X_GetUserData() gdnative.Array {
	//log.Println("Calling BakedLightmapData.X_GetUserData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "_get_user_data")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false data Array}], Returns: void
*/
func (o *BakedLightmapData) X_SetUserData(data gdnative.Array) {
	//log.Println("Calling BakedLightmapData.X_SetUserData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(data)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "_set_user_data")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false path NodePath} { false lightmap Texture} { false instance int}], Returns: void
*/
func (o *BakedLightmapData) AddUser(path gdnative.NodePath, lightmap TextureImplementer, instance gdnative.Int) {
	//log.Println("Calling BakedLightmapData.AddUser()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(path)
	ptrArguments[1] = gdnative.NewPointerFromObject(lightmap.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromInt(instance)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "add_user")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *BakedLightmapData) ClearUsers() {
	//log.Println("Calling BakedLightmapData.ClearUsers()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "clear_users")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: AABB
*/
func (o *BakedLightmapData) GetBounds() gdnative.Aabb {
	//log.Println("Calling BakedLightmapData.GetBounds()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "get_bounds")

	// Call the parent method.
	// AABB
	retPtr := gdnative.NewEmptyAabb()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewAabbFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: Transform
*/
func (o *BakedLightmapData) GetCellSpaceTransform() gdnative.Transform {
	//log.Println("Calling BakedLightmapData.GetCellSpaceTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "get_cell_space_transform")

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
	Args: [], Returns: int
*/
func (o *BakedLightmapData) GetCellSubdiv() gdnative.Int {
	//log.Println("Calling BakedLightmapData.GetCellSubdiv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "get_cell_subdiv")

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
func (o *BakedLightmapData) GetEnergy() gdnative.Real {
	//log.Println("Calling BakedLightmapData.GetEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "get_energy")

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
	Args: [], Returns: PoolByteArray
*/
func (o *BakedLightmapData) GetOctree() gdnative.PoolByteArray {
	//log.Println("Calling BakedLightmapData.GetOctree()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "get_octree")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: int
*/
func (o *BakedLightmapData) GetUserCount() gdnative.Int {
	//log.Println("Calling BakedLightmapData.GetUserCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "get_user_count")

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
	Args: [{ false user_idx int}], Returns: Texture
*/
func (o *BakedLightmapData) GetUserLightmap(userIdx gdnative.Int) TextureImplementer {
	//log.Println("Calling BakedLightmapData.GetUserLightmap()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(userIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "get_user_lightmap")

	// Call the parent method.
	// Texture
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newTextureFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(TextureImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Texture" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(TextureImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false user_idx int}], Returns: NodePath
*/
func (o *BakedLightmapData) GetUserPath(userIdx gdnative.Int) gdnative.NodePath {
	//log.Println("Calling BakedLightmapData.GetUserPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(userIdx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "get_user_path")

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
	Args: [{ false bounds AABB}], Returns: void
*/
func (o *BakedLightmapData) SetBounds(bounds gdnative.Aabb) {
	//log.Println("Calling BakedLightmapData.SetBounds()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromAabb(bounds)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "set_bounds")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false xform Transform}], Returns: void
*/
func (o *BakedLightmapData) SetCellSpaceTransform(xform gdnative.Transform) {
	//log.Println("Calling BakedLightmapData.SetCellSpaceTransform()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromTransform(xform)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "set_cell_space_transform")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false cell_subdiv int}], Returns: void
*/
func (o *BakedLightmapData) SetCellSubdiv(cellSubdiv gdnative.Int) {
	//log.Println("Calling BakedLightmapData.SetCellSubdiv()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(cellSubdiv)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "set_cell_subdiv")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false energy float}], Returns: void
*/
func (o *BakedLightmapData) SetEnergy(energy gdnative.Real) {
	//log.Println("Calling BakedLightmapData.SetEnergy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(energy)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "set_energy")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false octree PoolByteArray}], Returns: void
*/
func (o *BakedLightmapData) SetOctree(octree gdnative.PoolByteArray) {
	//log.Println("Calling BakedLightmapData.SetOctree()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolByteArray(octree)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("BakedLightmapData", "set_octree")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// BakedLightmapDataImplementer is an interface that implements the methods
// of the BakedLightmapData class.
type BakedLightmapDataImplementer interface {
	ResourceImplementer
	X_GetUserData() gdnative.Array
	X_SetUserData(data gdnative.Array)
	AddUser(path gdnative.NodePath, lightmap TextureImplementer, instance gdnative.Int)
	ClearUsers()
	GetBounds() gdnative.Aabb
	GetCellSpaceTransform() gdnative.Transform
	GetCellSubdiv() gdnative.Int
	GetEnergy() gdnative.Real
	GetOctree() gdnative.PoolByteArray
	GetUserCount() gdnative.Int
	GetUserLightmap(userIdx gdnative.Int) TextureImplementer
	GetUserPath(userIdx gdnative.Int) gdnative.NodePath
	SetBounds(bounds gdnative.Aabb)
	SetCellSpaceTransform(xform gdnative.Transform)
	SetCellSubdiv(cellSubdiv gdnative.Int)
	SetEnergy(energy gdnative.Real)
	SetOctree(octree gdnative.PoolByteArray)
}
