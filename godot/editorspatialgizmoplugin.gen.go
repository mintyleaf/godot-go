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

//func NewEditorSpatialGizmoPluginFromPointer(ptr gdnative.Pointer) EditorSpatialGizmoPlugin {
func newEditorSpatialGizmoPluginFromPointer(ptr gdnative.Pointer) EditorSpatialGizmoPlugin {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorSpatialGizmoPlugin{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type EditorSpatialGizmoPlugin struct {
	Resource
	owner gdnative.Object
}

func (o *EditorSpatialGizmoPlugin) BaseClass() string {
	return "EditorSpatialGizmoPlugin"
}

/*
        Undocumented
	Args: [{ false name String} { false material SpatialMaterial}], Returns: void
*/
func (o *EditorSpatialGizmoPlugin) AddMaterial(name gdnative.String, material SpatialMaterialImplementer) {
	//log.Println("Calling EditorSpatialGizmoPlugin.AddMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromObject(material.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "add_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *EditorSpatialGizmoPlugin) CanBeHidden() gdnative.Bool {
	//log.Println("Calling EditorSpatialGizmoPlugin.CanBeHidden()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "can_be_hidden")

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
	Args: [{ false gizmo EditorSpatialGizmo} { false index int} { false restore Variant} { false cancel bool}], Returns: void
*/
func (o *EditorSpatialGizmoPlugin) CommitHandle(gizmo EditorSpatialGizmoImplementer, index gdnative.Int, restore gdnative.Variant, cancel gdnative.Bool) {
	//log.Println("Calling EditorSpatialGizmoPlugin.CommitHandle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(gizmo.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(index)
	ptrArguments[2] = gdnative.NewPointerFromVariant(restore)
	ptrArguments[3] = gdnative.NewPointerFromBool(cancel)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "commit_handle")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false spatial Spatial}], Returns: EditorSpatialGizmo
*/
func (o *EditorSpatialGizmoPlugin) CreateGizmo(spatial SpatialImplementer) EditorSpatialGizmoImplementer {
	//log.Println("Calling EditorSpatialGizmoPlugin.CreateGizmo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(spatial.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "create_gizmo")

	// Call the parent method.
	// EditorSpatialGizmo
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newEditorSpatialGizmoFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(EditorSpatialGizmoImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "EditorSpatialGizmo" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(EditorSpatialGizmoImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [{ false name String} {False true billboard bool}], Returns: void
*/
func (o *EditorSpatialGizmoPlugin) CreateHandleMaterial(name gdnative.String, billboard gdnative.Bool) {
	//log.Println("Calling EditorSpatialGizmoPlugin.CreateHandleMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromBool(billboard)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "create_handle_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false name String} { false texture Texture} {False true on_top bool} {1,1,1,1 true color Color}], Returns: void
*/
func (o *EditorSpatialGizmoPlugin) CreateIconMaterial(name gdnative.String, texture TextureImplementer, onTop gdnative.Bool, color gdnative.Color) {
	//log.Println("Calling EditorSpatialGizmoPlugin.CreateIconMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromObject(texture.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromBool(onTop)
	ptrArguments[3] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "create_icon_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false name String} { false color Color} {False true billboard bool} {False true on_top bool} {False true use_vertex_color bool}], Returns: void
*/
func (o *EditorSpatialGizmoPlugin) CreateMaterial(name gdnative.String, color gdnative.Color, billboard gdnative.Bool, onTop gdnative.Bool, useVertexColor gdnative.Bool) {
	//log.Println("Calling EditorSpatialGizmoPlugin.CreateMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 5, 5)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromColor(color)
	ptrArguments[2] = gdnative.NewPointerFromBool(billboard)
	ptrArguments[3] = gdnative.NewPointerFromBool(onTop)
	ptrArguments[4] = gdnative.NewPointerFromBool(useVertexColor)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "create_material")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false gizmo EditorSpatialGizmo} { false index int}], Returns: String
*/
func (o *EditorSpatialGizmoPlugin) GetHandleName(gizmo EditorSpatialGizmoImplementer, index gdnative.Int) gdnative.String {
	//log.Println("Calling EditorSpatialGizmoPlugin.GetHandleName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(gizmo.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "get_handle_name")

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
	Args: [{ false gizmo EditorSpatialGizmo} { false index int}], Returns: Variant
*/
func (o *EditorSpatialGizmoPlugin) GetHandleValue(gizmo EditorSpatialGizmoImplementer, index gdnative.Int) gdnative.Variant {
	//log.Println("Calling EditorSpatialGizmoPlugin.GetHandleValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(gizmo.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "get_handle_value")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false name String} { false gizmo EditorSpatialGizmo}], Returns: SpatialMaterial
*/
func (o *EditorSpatialGizmoPlugin) GetMaterial(name gdnative.String, gizmo EditorSpatialGizmoImplementer) SpatialMaterialImplementer {
	//log.Println("Calling EditorSpatialGizmoPlugin.GetMaterial()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(name)
	ptrArguments[1] = gdnative.NewPointerFromObject(gizmo.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "get_material")

	// Call the parent method.
	// SpatialMaterial
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newSpatialMaterialFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(SpatialMaterialImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "SpatialMaterial" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(SpatialMaterialImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *EditorSpatialGizmoPlugin) GetName() gdnative.String {
	//log.Println("Calling EditorSpatialGizmoPlugin.GetName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "get_name")

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
	Args: [], Returns: String
*/
func (o *EditorSpatialGizmoPlugin) GetPriority() gdnative.String {
	//log.Println("Calling EditorSpatialGizmoPlugin.GetPriority()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "get_priority")

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
	Args: [{ false spatial Spatial}], Returns: bool
*/
func (o *EditorSpatialGizmoPlugin) HasGizmo(spatial SpatialImplementer) gdnative.Bool {
	//log.Println("Calling EditorSpatialGizmoPlugin.HasGizmo()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(spatial.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "has_gizmo")

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
	Args: [{ false gizmo EditorSpatialGizmo} { false index int}], Returns: bool
*/
func (o *EditorSpatialGizmoPlugin) IsHandleHighlighted(gizmo EditorSpatialGizmoImplementer, index gdnative.Int) gdnative.Bool {
	//log.Println("Calling EditorSpatialGizmoPlugin.IsHandleHighlighted()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(gizmo.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(index)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "is_handle_highlighted")

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
func (o *EditorSpatialGizmoPlugin) IsSelectableWhenHidden() gdnative.Bool {
	//log.Println("Calling EditorSpatialGizmoPlugin.IsSelectableWhenHidden()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "is_selectable_when_hidden")

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
	Args: [{ false gizmo EditorSpatialGizmo}], Returns: void
*/
func (o *EditorSpatialGizmoPlugin) Redraw(gizmo EditorSpatialGizmoImplementer) {
	//log.Println("Calling EditorSpatialGizmoPlugin.Redraw()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(gizmo.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "redraw")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false gizmo EditorSpatialGizmo} { false index int} { false camera Camera} { false point Vector2}], Returns: void
*/
func (o *EditorSpatialGizmoPlugin) SetHandle(gizmo EditorSpatialGizmoImplementer, index gdnative.Int, camera CameraImplementer, point gdnative.Vector2) {
	//log.Println("Calling EditorSpatialGizmoPlugin.SetHandle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(gizmo.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(index)
	ptrArguments[2] = gdnative.NewPointerFromObject(camera.GetBaseObject())
	ptrArguments[3] = gdnative.NewPointerFromVector2(point)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorSpatialGizmoPlugin", "set_handle")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorSpatialGizmoPluginImplementer is an interface that implements the methods
// of the EditorSpatialGizmoPlugin class.
type EditorSpatialGizmoPluginImplementer interface {
	ResourceImplementer
	AddMaterial(name gdnative.String, material SpatialMaterialImplementer)
	CanBeHidden() gdnative.Bool
	CommitHandle(gizmo EditorSpatialGizmoImplementer, index gdnative.Int, restore gdnative.Variant, cancel gdnative.Bool)
	CreateGizmo(spatial SpatialImplementer) EditorSpatialGizmoImplementer
	CreateHandleMaterial(name gdnative.String, billboard gdnative.Bool)
	CreateIconMaterial(name gdnative.String, texture TextureImplementer, onTop gdnative.Bool, color gdnative.Color)
	CreateMaterial(name gdnative.String, color gdnative.Color, billboard gdnative.Bool, onTop gdnative.Bool, useVertexColor gdnative.Bool)
	GetHandleName(gizmo EditorSpatialGizmoImplementer, index gdnative.Int) gdnative.String
	GetHandleValue(gizmo EditorSpatialGizmoImplementer, index gdnative.Int) gdnative.Variant
	GetMaterial(name gdnative.String, gizmo EditorSpatialGizmoImplementer) SpatialMaterialImplementer
	GetPriority() gdnative.String
	HasGizmo(spatial SpatialImplementer) gdnative.Bool
	IsHandleHighlighted(gizmo EditorSpatialGizmoImplementer, index gdnative.Int) gdnative.Bool
	IsSelectableWhenHidden() gdnative.Bool
	Redraw(gizmo EditorSpatialGizmoImplementer)
	SetHandle(gizmo EditorSpatialGizmoImplementer, index gdnative.Int, camera CameraImplementer, point gdnative.Vector2)
}
