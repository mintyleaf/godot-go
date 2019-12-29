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

// PackedSceneGenEditState is an enum for GenEditState values.
type PackedSceneGenEditState int

const (
	PackedSceneGenEditStateDisabled PackedSceneGenEditState = 0
	PackedSceneGenEditStateInstance PackedSceneGenEditState = 1
	PackedSceneGenEditStateMain     PackedSceneGenEditState = 2
)

//func NewPackedSceneFromPointer(ptr gdnative.Pointer) PackedScene {
func newPackedSceneFromPointer(ptr gdnative.Pointer) PackedScene {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := PackedScene{}
	obj.SetBaseObject(owner)

	return obj
}

/*
A simplified interface to a scene file. Provides access to operations and checks that can be performed on the scene resource itself. Can be used to save a node to a file. When saving, the node as well as all the node it owns get saved (see [code]owner[/code] property on [Node]). [b]Note:[/b] The node doesn't need to own itself. [b]Example of saving a node with different owners:[/b] The following example creates 3 objects: [code]Node2D[/code] ([code]node[/code]), [code]RigidBody2D[/code] ([code]rigid[/code]) and [code]CollisionObject2D[/code] ([code]collision[/code]). [code]collision[/code] is a child of [code]rigid[/code] which is a child of [code]node[/code]. Only [code]rigid[/code] is owned by [code]node[/code] and [code]pack[/code] will therefore only save those two nodes, but not [code]collision[/code]. [codeblock] # Create the objects var node = Node2D.new() var rigid = RigidBody2D.new() var collision = CollisionShape2D.new() # Create the object hierarchy rigid.add_child(collision) node.add_child(rigid) # Change owner of rigid, but not of collision rigid.owner = node var scene = PackedScene.new() # Only node and rigid are now packed var result = scene.pack(node) if result == OK: ResourceSaver.save("res://path/name.scn", scene) # Or "user://..." [/codeblock]
*/
type PackedScene struct {
	Resource
	owner gdnative.Object
}

func (o *PackedScene) BaseClass() string {
	return "PackedScene"
}

/*
        Undocumented
	Args: [], Returns: Dictionary
*/
func (o *PackedScene) X_GetBundledScene() gdnative.Dictionary {
	//log.Println("Calling PackedScene.X_GetBundledScene()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "_get_bundled_scene")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 Dictionary}], Returns: void
*/
func (o *PackedScene) X_SetBundledScene(arg0 gdnative.Dictionary) {
	//log.Println("Calling PackedScene.X_SetBundledScene()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromDictionary(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "_set_bundled_scene")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns [code]true[/code] if the scene file has nodes.
	Args: [], Returns: bool
*/
func (o *PackedScene) CanInstance() gdnative.Bool {
	//log.Println("Calling PackedScene.CanInstance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "can_instance")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns the [code]SceneState[/code] representing the scene file contents.
	Args: [], Returns: SceneState
*/
func (o *PackedScene) GetState() SceneStateImplementer {
	//log.Println("Calling PackedScene.GetState()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "get_state")

	// Call the parent method.
	// SceneState
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newSceneStateFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(SceneStateImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "SceneState" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(SceneStateImplementer)
	}

	return &ret
}

/*
        Instantiates the scene's node hierarchy. Triggers child scene instantiation(s). Triggers a [constant Node.NOTIFICATION_INSTANCED] notification on the root node.
	Args: [{0 true edit_state int}], Returns: Node
*/
func (o *PackedScene) Instance(editState gdnative.Int) NodeImplementer {
	//log.Println("Calling PackedScene.Instance()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(editState)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "instance")

	// Call the parent method.
	// Node
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newNodeFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(NodeImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Node" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(NodeImplementer)
	}

	return &ret
}

/*
        Pack will ignore any sub-nodes not owned by given node. See [member Node.owner].
	Args: [{ false path Node}], Returns: enum.Error
*/
func (o *PackedScene) Pack(path NodeImplementer) gdnative.Error {
	//log.Println("Calling PackedScene.Pack()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(path.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("PackedScene", "pack")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

// PackedSceneImplementer is an interface that implements the methods
// of the PackedScene class.
type PackedSceneImplementer interface {
	ResourceImplementer
	X_GetBundledScene() gdnative.Dictionary
	X_SetBundledScene(arg0 gdnative.Dictionary)
	CanInstance() gdnative.Bool
	GetState() SceneStateImplementer
	Instance(editState gdnative.Int) NodeImplementer
}
