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

//func NewEditorSpatialGizmoFromPointer(ptr gdnative.Pointer) EditorSpatialGizmo {
func newEditorSpatialGizmoFromPointer(ptr gdnative.Pointer) EditorSpatialGizmo {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorSpatialGizmo{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Custom gizmo that is used for providing custom visualization and editing (handles) for 3D Spatial objects. See [EditorSpatialGizmoPlugin] for more information.
*/
type EditorSpatialGizmo struct {
	SpatialGizmo
	owner gdnative.Object
}

func (o *EditorSpatialGizmo) BaseClass() string {
	return "EditorSpatialGizmo"
}

/*

	Args: [{ false segments PoolVector3Array}], Returns: void
*/
func (o *EditorSpatialGizmo) AddCollisionSegments(segments gdnative.PoolVector3Array) {
	//log.Println("Calling EditorSpatialGizmo.AddCollisionSegments()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector3Array(segments)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "add_collision_segments")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds collision triangles to the gizmo for picking. A [TriangleMesh] can be generated from a regular [Mesh] too. Call this function during [method redraw].
	Args: [{ false triangles TriangleMesh}], Returns: void
*/
func (o *EditorSpatialGizmo) AddCollisionTriangles(triangles TriangleMeshImplementer) {
	//log.Println("Calling EditorSpatialGizmo.AddCollisionTriangles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(triangles.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "add_collision_triangles")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a list of handles (points) which can be used to deform the object being edited. There are virtual functions which will be called upon editing of these handles. Call this function during [method redraw].
	Args: [{ false handles PoolVector3Array} { false material Material} {False true billboard bool} {False true secondary bool}], Returns: void
*/
func (o *EditorSpatialGizmo) AddHandles(handles gdnative.PoolVector3Array, material MaterialImplementer, billboard gdnative.Bool, secondary gdnative.Bool) {
	//log.Println("Calling EditorSpatialGizmo.AddHandles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector3Array(handles)
	ptrArguments[1] = gdnative.NewPointerFromObject(material.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromBool(billboard)
	ptrArguments[3] = gdnative.NewPointerFromBool(secondary)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "add_handles")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds lines to the gizmo (as sets of 2 points), with a given material. The lines are used for visualizing the gizmo. Call this function during [method redraw].
	Args: [{ false lines PoolVector3Array} { false material Material} {False true billboard bool} {1,1,1,1 true modulate Color}], Returns: void
*/
func (o *EditorSpatialGizmo) AddLines(lines gdnative.PoolVector3Array, material MaterialImplementer, billboard gdnative.Bool, modulate gdnative.Color) {
	//log.Println("Calling EditorSpatialGizmo.AddLines()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromPoolVector3Array(lines)
	ptrArguments[1] = gdnative.NewPointerFromObject(material.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromBool(billboard)
	ptrArguments[3] = gdnative.NewPointerFromColor(modulate)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "add_lines")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false mesh ArrayMesh} {False true billboard bool} {[Object:null] true skeleton SkinReference} {Null true material Material}], Returns: void
*/
func (o *EditorSpatialGizmo) AddMesh(mesh ArrayMeshImplementer, billboard gdnative.Bool, skeleton SkinReferenceImplementer, material MaterialImplementer) {
	//log.Println("Calling EditorSpatialGizmo.AddMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(mesh.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromBool(billboard)
	ptrArguments[2] = gdnative.NewPointerFromObject(skeleton.GetBaseObject())
	ptrArguments[3] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "add_mesh")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds an unscaled billboard for visualization. Call this function during [method redraw].
	Args: [{ false material Material} {1 true default_scale float} {1,1,1,1 true modulate Color}], Returns: void
*/
func (o *EditorSpatialGizmo) AddUnscaledBillboard(material MaterialImplementer, defaultScale gdnative.Real, modulate gdnative.Color) {
	//log.Println("Calling EditorSpatialGizmo.AddUnscaledBillboard()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromObject(material.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromReal(defaultScale)
	ptrArguments[2] = gdnative.NewPointerFromColor(modulate)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "add_unscaled_billboard")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: void
*/
func (o *EditorSpatialGizmo) Clear() {
	//log.Println("Calling EditorSpatialGizmo.Clear()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "clear")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Commit a handle being edited (handles must have been previously added by [method add_handles]). If the [code]cancel[/code] parameter is [code]true[/code], an option to restore the edited value to the original is provided.
	Args: [{ false index int} { false restore Variant} { false cancel bool}], Returns: void
*/
func (o *EditorSpatialGizmo) CommitHandle(index gdnative.Int, restore gdnative.Variant, cancel gdnative.Bool) {
	//log.Println("Calling EditorSpatialGizmo.CommitHandle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromVariant(restore)
	ptrArguments[2] = gdnative.NewPointerFromBool(cancel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "commit_handle")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Gets the name of an edited handle (handles must have been previously added by [method add_handles]). Handles can be named for reference to the user when editing.
	Args: [{ false index int}], Returns: String
*/
func (o *EditorSpatialGizmo) GetHandleName(index gdnative.Int) gdnative.String {
	//log.Println("Calling EditorSpatialGizmo.GetHandleName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "get_handle_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Gets actual value of a handle. This value can be anything and used for eventually undoing the motion when calling [method commit_handle].
	Args: [{ false index int}], Returns: Variant
*/
func (o *EditorSpatialGizmo) GetHandleValue(index gdnative.Int) gdnative.Variant {
	//log.Println("Calling EditorSpatialGizmo.GetHandleValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "get_handle_value")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns the [EditorSpatialGizmoPlugin] that owns this gizmo. It's useful to retrieve materials using [method EditorSpatialGizmoPlugin.get_material].
	Args: [], Returns: EditorSpatialGizmoPlugin
*/
func (o *EditorSpatialGizmo) GetPlugin() EditorSpatialGizmoPluginImplementer {
	//log.Println("Calling EditorSpatialGizmo.GetPlugin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "get_plugin")

	// Call the parent method.
	// EditorSpatialGizmoPlugin
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newEditorSpatialGizmoPluginFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(EditorSpatialGizmoPluginImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "EditorSpatialGizmoPlugin" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(EditorSpatialGizmoPluginImplementer)
	}

	return &ret
}

/*
        Returns the Spatial node associated with this gizmo.
	Args: [], Returns: Spatial
*/
func (o *EditorSpatialGizmo) GetSpatialNode() SpatialImplementer {
	//log.Println("Calling EditorSpatialGizmo.GetSpatialNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "get_spatial_node")

	// Call the parent method.
	// Spatial
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newSpatialFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(SpatialImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Spatial" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(SpatialImplementer)
	}

	return &ret
}

/*
        Gets whether a handle is highlighted or not.
	Args: [{ false index int}], Returns: bool
*/
func (o *EditorSpatialGizmo) IsHandleHighlighted(index gdnative.Int) gdnative.Bool {
	//log.Println("Calling EditorSpatialGizmo.IsHandleHighlighted()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "is_handle_highlighted")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        This function is called when the Spatial this gizmo refers to changes (the [method Spatial.update_gizmo] is called).
	Args: [], Returns: void
*/
func (o *EditorSpatialGizmo) Redraw() {
	//log.Println("Calling EditorSpatialGizmo.Redraw()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "redraw")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        This function is used when the user drags a gizmo handle (previously added with [method add_handles]) in screen coordinates. The [Camera] is also provided so screen coordinates can be converted to raycasts.
	Args: [{ false index int} { false camera Camera} { false point Vector2}], Returns: void
*/
func (o *EditorSpatialGizmo) SetHandle(index gdnative.Int, camera CameraImplementer, point gdnative.Vector2) {
	//log.Println("Calling EditorSpatialGizmo.SetHandle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromInt(index)
	ptrArguments[1] = gdnative.NewPointerFromObject(camera.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromVector2(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "set_handle")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false hidden bool}], Returns: void
*/
func (o *EditorSpatialGizmo) SetHidden(hidden gdnative.Bool) {
	//log.Println("Calling EditorSpatialGizmo.SetHidden()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(hidden)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "set_hidden")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false node Node}], Returns: void
*/
func (o *EditorSpatialGizmo) SetSpatialNode(node NodeImplementer) {
	//log.Println("Calling EditorSpatialGizmo.SetSpatialNode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(node.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmo", "set_spatial_node")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorSpatialGizmoImplementer is an interface that implements the methods
// of the EditorSpatialGizmo class.
type EditorSpatialGizmoImplementer interface {
	SpatialGizmoImplementer
	AddCollisionSegments(segments gdnative.PoolVector3Array)
	AddCollisionTriangles(triangles TriangleMeshImplementer)
	AddHandles(handles gdnative.PoolVector3Array, material MaterialImplementer, billboard gdnative.Bool, secondary gdnative.Bool)
	AddLines(lines gdnative.PoolVector3Array, material MaterialImplementer, billboard gdnative.Bool, modulate gdnative.Color)
	AddMesh(mesh ArrayMeshImplementer, billboard gdnative.Bool, skeleton SkinReferenceImplementer, material MaterialImplementer)
	AddUnscaledBillboard(material MaterialImplementer, defaultScale gdnative.Real, modulate gdnative.Color)
	Clear()
	CommitHandle(index gdnative.Int, restore gdnative.Variant, cancel gdnative.Bool)
	GetHandleName(index gdnative.Int) gdnative.String
	GetHandleValue(index gdnative.Int) gdnative.Variant
	GetPlugin() EditorSpatialGizmoPluginImplementer
	GetSpatialNode() SpatialImplementer
	IsHandleHighlighted(index gdnative.Int) gdnative.Bool
	Redraw()
	SetHandle(index gdnative.Int, camera CameraImplementer, point gdnative.Vector2)
	SetHidden(hidden gdnative.Bool)
	SetSpatialNode(node NodeImplementer)
}
