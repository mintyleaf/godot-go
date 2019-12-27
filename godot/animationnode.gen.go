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

// AnimationNodeFilterAction is an enum for FilterAction values.
type AnimationNodeFilterAction int

const (
	AnimationNodeFilterBlend  AnimationNodeFilterAction = 3
	AnimationNodeFilterIgnore AnimationNodeFilterAction = 0
	AnimationNodeFilterPass   AnimationNodeFilterAction = 1
	AnimationNodeFilterStop   AnimationNodeFilterAction = 2
)

//func NewAnimationNodeFromPointer(ptr gdnative.Pointer) AnimationNode {
func newAnimationNodeFromPointer(ptr gdnative.Pointer) AnimationNode {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AnimationNode{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base resource for [AnimationTree] nodes. In general it's not used directly but you can create custom ones with custom blending formulas. Inherit this when creating nodes mainly for use in [AnimationNodeBlendTree], otherwise [AnimationRootNode] should be used instead.
*/
type AnimationNode struct {
	Resource
	owner gdnative.Object
}

func (o *AnimationNode) BaseClass() string {
	return "AnimationNode"
}

/*
        Undocumented
	Args: [], Returns: Array
*/
func (o *AnimationNode) X_GetFilters() gdnative.Array {
	//log.Println("Calling AnimationNode.X_GetFilters()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "_get_filters")

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
	Args: [{ false filters Array}], Returns: void
*/
func (o *AnimationNode) X_SetFilters(filters gdnative.Array) {
	//log.Println("Calling AnimationNode.X_SetFilters()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromArray(filters)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "_set_filters")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Add an input to the node. This is only useful for nodes created for use in an [AnimationNodeBlendTree]
	Args: [{ false name String}], Returns: void
*/
func (o *AnimationNode) AddInput(name gdnative.String) {
	//log.Println("Calling AnimationNode.AddInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "add_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Blend an animation by "blend" amount (name must be valid in the linked [AnimationPlayer]). A time and delta mas be passed, as well as whether seek happened.
	Args: [{ false animation String} { false time float} { false delta float} { false seeked bool} { false blend float}], Returns: void
*/
func (o *AnimationNode) BlendAnimation(animation gdnative.String, time gdnative.Real, delta gdnative.Real, seeked gdnative.Bool, blend gdnative.Real) {
	//log.Println("Calling AnimationNode.BlendAnimation()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromString(animation)
	ptrArguments[1] = gdnative.NewPointerFromReal(time)
	ptrArguments[2] = gdnative.NewPointerFromReal(delta)
	ptrArguments[3] = gdnative.NewPointerFromBool(seeked)
	ptrArguments[4] = gdnative.NewPointerFromReal(blend)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "blend_animation")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Blend an input. This is only useful for nodes created for an [AnimationNodeBlendTree]. Time is a delta, unless "seek" is [code]true[/code], in which case it is absolute. A filter mode may be optionally passed.
	Args: [{ false input_index int} { false time float} { false seek bool} { false blend float} {0 true filter int} {True true optimize bool}], Returns: float
*/
func (o *AnimationNode) BlendInput(inputIndex gdnative.Int, time gdnative.Real, seek gdnative.Bool, blend gdnative.Real, filter gdnative.Int, optimize gdnative.Bool) gdnative.Real {
	//log.Println("Calling AnimationNode.BlendInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 6, 6)
	ptrArguments[0] = gdnative.NewPointerFromInt(inputIndex)
	ptrArguments[1] = gdnative.NewPointerFromReal(time)
	ptrArguments[2] = gdnative.NewPointerFromBool(seek)
	ptrArguments[3] = gdnative.NewPointerFromReal(blend)
	ptrArguments[4] = gdnative.NewPointerFromInt(filter)
	ptrArguments[5] = gdnative.NewPointerFromBool(optimize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "blend_input")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Blend another animaiton node (in case this node contains children animation nodes). This function is only useful if you inherit from [AnimationRootNode] instead, else editors will not display your node for addition.
	Args: [{ false name String} { false node AnimationNode} { false time float} { false seek bool} { false blend float} {0 true filter int} {True true optimize bool}], Returns: float
*/
func (o *AnimationNode) BlendNode(name gdnative.String, node AnimationNodeImplementer, time gdnative.Real, seek gdnative.Bool, blend gdnative.Real, filter gdnative.Int, optimize gdnative.Bool) gdnative.Real {
	//log.Println("Calling AnimationNode.BlendNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 7, 7)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromObject(node.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromReal(time)
	ptrArguments[3] = gdnative.NewPointerFromBool(seek)
	ptrArguments[4] = gdnative.NewPointerFromReal(blend)
	ptrArguments[5] = gdnative.NewPointerFromInt(filter)
	ptrArguments[6] = gdnative.NewPointerFromBool(optimize)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "blend_node")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Get the text caption for this node (used by some editors)
	Args: [], Returns: String
*/
func (o *AnimationNode) GetCaption() gdnative.String {
	//log.Println("Calling AnimationNode.GetCaption()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "get_caption")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Get the a child node by index (used by editors inheriting from [AnimationRootNode]).
	Args: [{ false name String}], Returns: Object
*/
func (o *AnimationNode) GetChildByName(name gdnative.String) ObjectImplementer {
	//log.Println("Calling AnimationNode.GetChildByName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "get_child_by_name")

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
        Get all children nodes, in order as a name:node dictionary. Only useful when inheriting [AnimationRootNode].
	Args: [], Returns: Dictionary
*/
func (o *AnimationNode) GetChildNodes() gdnative.Dictionary {
	//log.Println("Calling AnimationNode.GetChildNodes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "get_child_nodes")

	// Call the parent method.
	// Dictionary
	retPtr := gdnative.NewEmptyDictionary()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewDictionaryFromPointer(retPtr)
	return ret
}

/*
        Amount of inputs in this node, only useful for nodes that go into [AnimationNodeBlendTree].
	Args: [], Returns: int
*/
func (o *AnimationNode) GetInputCount() gdnative.Int {
	//log.Println("Calling AnimationNode.GetInputCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "get_input_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Get the name of an input by index.
	Args: [{ false input int}], Returns: String
*/
func (o *AnimationNode) GetInputName(input gdnative.Int) gdnative.String {
	//log.Println("Calling AnimationNode.GetInputName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(input)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "get_input_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Get the value of a parameter. Parameters are custom local memory used for your nodes, given a resource can be reused in multiple trees.
	Args: [{ false name String}], Returns: Variant
*/
func (o *AnimationNode) GetParameter(name gdnative.String) gdnative.Variant {
	//log.Println("Calling AnimationNode.GetParameter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "get_parameter")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Get the default value of a parameter. Parameters are custom local memory used for your nodes, given a resource can be reused in multiple trees.
	Args: [{ false name String}], Returns: Variant
*/
func (o *AnimationNode) GetParameterDefaultValue(name gdnative.String) gdnative.Variant {
	//log.Println("Calling AnimationNode.GetParameterDefaultValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "get_parameter_default_value")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Get the property information for parameter. Parameters are custom local memory used for your nodes, given a resource can be reused in multiple trees. Format is similar to [method Object.get_property_list].
	Args: [], Returns: Array
*/
func (o *AnimationNode) GetParameterList() gdnative.Array {
	//log.Println("Calling AnimationNode.GetParameterList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "get_parameter_list")

	// Call the parent method.
	// Array
	retPtr := gdnative.NewEmptyArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewArrayFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] whether you want the blend tree editor to display filter editing on this node.
	Args: [], Returns: String
*/
func (o *AnimationNode) HasFilter() gdnative.String {
	//log.Println("Calling AnimationNode.HasFilter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "has_filter")

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
func (o *AnimationNode) IsFilterEnabled() gdnative.Bool {
	//log.Println("Calling AnimationNode.IsFilterEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "is_filter_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] whether a given path is filtered.
	Args: [{ false path NodePath}], Returns: bool
*/
func (o *AnimationNode) IsPathFiltered(path gdnative.NodePath) gdnative.Bool {
	//log.Println("Calling AnimationNode.IsPathFiltered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "is_path_filtered")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Called when a custom node is processed. The argument "time" is relative, unless "seek" is [code]true[/code] (in which case it is absolute). Here, call the [method blend_input], [method blend_node] or [method blend_animation] functions. You can also use [method get_parameter] and [method set_parameter] to modify local memory. This function returns the time left for the current animation to finish (if unsure, just pass the value from the main blend being called).
	Args: [{ false time float} { false seek bool}], Returns: void
*/
func (o *AnimationNode) Process(time gdnative.Real, seek gdnative.Bool) {
	//log.Println("Calling AnimationNode.Process()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromReal(time)
	ptrArguments[1] = gdnative.NewPointerFromBool(seek)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "process")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Remove an input, call this only when inactive.
	Args: [{ false index int}], Returns: void
*/
func (o *AnimationNode) RemoveInput(index gdnative.Int) {
	//log.Println("Calling AnimationNode.RemoveInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "remove_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *AnimationNode) SetFilterEnabled(enable gdnative.Bool) {
	//log.Println("Calling AnimationNode.SetFilterEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "set_filter_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Add/Remove a path for the filter.
	Args: [{ false path NodePath} { false enable bool}], Returns: void
*/
func (o *AnimationNode) SetFilterPath(path gdnative.NodePath, enable gdnative.Bool) {
	//log.Println("Calling AnimationNode.SetFilterPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromNodePath(path)
	ptrArguments[1] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "set_filter_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Set a custom parameter. These are used as local storage, because resources can be reused across the tree or scenes.
	Args: [{ false name String} { false value Variant}], Returns: void
*/
func (o *AnimationNode) SetParameter(name gdnative.String, value gdnative.Variant) {
	//log.Println("Calling AnimationNode.SetParameter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("AnimationNode", "set_parameter")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// AnimationNodeImplementer is an interface that implements the methods
// of the AnimationNode class.
type AnimationNodeImplementer interface {
	ResourceImplementer
	X_GetFilters() gdnative.Array
	X_SetFilters(filters gdnative.Array)
	AddInput(name gdnative.String)
	BlendAnimation(animation gdnative.String, time gdnative.Real, delta gdnative.Real, seeked gdnative.Bool, blend gdnative.Real)
	BlendInput(inputIndex gdnative.Int, time gdnative.Real, seek gdnative.Bool, blend gdnative.Real, filter gdnative.Int, optimize gdnative.Bool) gdnative.Real
	BlendNode(name gdnative.String, node AnimationNodeImplementer, time gdnative.Real, seek gdnative.Bool, blend gdnative.Real, filter gdnative.Int, optimize gdnative.Bool) gdnative.Real
	GetCaption() gdnative.String
	GetChildByName(name gdnative.String) ObjectImplementer
	GetChildNodes() gdnative.Dictionary
	GetInputCount() gdnative.Int
	GetInputName(input gdnative.Int) gdnative.String
	GetParameter(name gdnative.String) gdnative.Variant
	GetParameterDefaultValue(name gdnative.String) gdnative.Variant
	GetParameterList() gdnative.Array
	HasFilter() gdnative.String
	IsFilterEnabled() gdnative.Bool
	IsPathFiltered(path gdnative.NodePath) gdnative.Bool
	Process(time gdnative.Real, seek gdnative.Bool)
	RemoveInput(index gdnative.Int)
	SetFilterEnabled(enable gdnative.Bool)
	SetFilterPath(path gdnative.NodePath, enable gdnative.Bool)
	SetParameter(name gdnative.String, value gdnative.Variant)
}
