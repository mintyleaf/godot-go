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

// EditorFeatureProfileFeature is an enum for Feature values.
type EditorFeatureProfileFeature int

const (
	EditorFeatureProfileFeature3D             EditorFeatureProfileFeature = 0
	EditorFeatureProfileFeatureAssetLib       EditorFeatureProfileFeature = 2
	EditorFeatureProfileFeatureFilesystemDock EditorFeatureProfileFeature = 6
	EditorFeatureProfileFeatureImportDock     EditorFeatureProfileFeature = 4
	EditorFeatureProfileFeatureMax            EditorFeatureProfileFeature = 7
	EditorFeatureProfileFeatureNodeDock       EditorFeatureProfileFeature = 5
	EditorFeatureProfileFeatureSceneTree      EditorFeatureProfileFeature = 3
	EditorFeatureProfileFeatureScript         EditorFeatureProfileFeature = 1
)

//func NewEditorFeatureProfileFromPointer(ptr gdnative.Pointer) EditorFeatureProfile {
func newEditorFeatureProfileFromPointer(ptr gdnative.Pointer) EditorFeatureProfile {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorFeatureProfile{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type EditorFeatureProfile struct {
	Reference
	owner gdnative.Object
}

func (o *EditorFeatureProfile) BaseClass() string {
	return "EditorFeatureProfile"
}

/*

	Args: [{ false feature int}], Returns: String
*/
func (o *EditorFeatureProfile) GetFeatureName(feature gdnative.Int) gdnative.String {
	//log.Println("Calling EditorFeatureProfile.GetFeatureName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(feature)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "get_feature_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false class_name String}], Returns: bool
*/
func (o *EditorFeatureProfile) IsClassDisabled(className gdnative.String) gdnative.Bool {
	//log.Println("Calling EditorFeatureProfile.IsClassDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(className)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "is_class_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false class_name String}], Returns: bool
*/
func (o *EditorFeatureProfile) IsClassEditorDisabled(className gdnative.String) gdnative.Bool {
	//log.Println("Calling EditorFeatureProfile.IsClassEditorDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(className)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "is_class_editor_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false class_name String} { false property String}], Returns: bool
*/
func (o *EditorFeatureProfile) IsClassPropertyDisabled(className gdnative.String, property gdnative.String) gdnative.Bool {
	//log.Println("Calling EditorFeatureProfile.IsClassPropertyDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(className)
	ptrArguments[1] = gdnative.NewPointerFromString(property)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "is_class_property_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false feature int}], Returns: bool
*/
func (o *EditorFeatureProfile) IsFeatureDisabled(feature gdnative.Int) gdnative.Bool {
	//log.Println("Calling EditorFeatureProfile.IsFeatureDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(feature)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "is_feature_disabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false path String}], Returns: enum.Error
*/
func (o *EditorFeatureProfile) LoadFromFile(path gdnative.String) gdnative.Error {
	//log.Println("Calling EditorFeatureProfile.LoadFromFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "load_from_file")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*

	Args: [{ false path String}], Returns: enum.Error
*/
func (o *EditorFeatureProfile) SaveToFile(path gdnative.String) gdnative.Error {
	//log.Println("Calling EditorFeatureProfile.SaveToFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "save_to_file")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*

	Args: [{ false class_name String} { false disable bool}], Returns: void
*/
func (o *EditorFeatureProfile) SetDisableClass(className gdnative.String, disable gdnative.Bool) {
	//log.Println("Calling EditorFeatureProfile.SetDisableClass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(className)
	ptrArguments[1] = gdnative.NewPointerFromBool(disable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "set_disable_class")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false class_name String} { false disable bool}], Returns: void
*/
func (o *EditorFeatureProfile) SetDisableClassEditor(className gdnative.String, disable gdnative.Bool) {
	//log.Println("Calling EditorFeatureProfile.SetDisableClassEditor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(className)
	ptrArguments[1] = gdnative.NewPointerFromBool(disable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "set_disable_class_editor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false class_name String} { false property String} { false disable bool}], Returns: void
*/
func (o *EditorFeatureProfile) SetDisableClassProperty(className gdnative.String, property gdnative.String, disable gdnative.Bool) {
	//log.Println("Calling EditorFeatureProfile.SetDisableClassProperty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(className)
	ptrArguments[1] = gdnative.NewPointerFromString(property)
	ptrArguments[2] = gdnative.NewPointerFromBool(disable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "set_disable_class_property")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false feature int} { false disable bool}], Returns: void
*/
func (o *EditorFeatureProfile) SetDisableFeature(feature gdnative.Int, disable gdnative.Bool) {
	//log.Println("Calling EditorFeatureProfile.SetDisableFeature()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(feature)
	ptrArguments[1] = gdnative.NewPointerFromBool(disable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFeatureProfile", "set_disable_feature")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorFeatureProfileImplementer is an interface that implements the methods
// of the EditorFeatureProfile class.
type EditorFeatureProfileImplementer interface {
	ReferenceImplementer
	GetFeatureName(feature gdnative.Int) gdnative.String
	IsClassDisabled(className gdnative.String) gdnative.Bool
	IsClassEditorDisabled(className gdnative.String) gdnative.Bool
	IsClassPropertyDisabled(className gdnative.String, property gdnative.String) gdnative.Bool
	IsFeatureDisabled(feature gdnative.Int) gdnative.Bool
	SetDisableClass(className gdnative.String, disable gdnative.Bool)
	SetDisableClassEditor(className gdnative.String, disable gdnative.Bool)
	SetDisableClassProperty(className gdnative.String, property gdnative.String, disable gdnative.Bool)
	SetDisableFeature(feature gdnative.Int, disable gdnative.Bool)
}
