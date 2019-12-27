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

//func NewEditorInspectorPluginFromPointer(ptr gdnative.Pointer) EditorInspectorPlugin {
func newEditorInspectorPluginFromPointer(ptr gdnative.Pointer) EditorInspectorPlugin {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorInspectorPlugin{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type EditorInspectorPlugin struct {
	Reference
	owner gdnative.Object
}

func (o *EditorInspectorPlugin) BaseClass() string {
	return "EditorInspectorPlugin"
}

/*

	Args: [{ false control Control}], Returns: void
*/
func (o *EditorInspectorPlugin) AddCustomControl(control ControlImplementer) {
	//log.Println("Calling EditorInspectorPlugin.AddCustomControl()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(control.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspectorPlugin", "add_custom_control")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false property String} { false editor Control}], Returns: void
*/
func (o *EditorInspectorPlugin) AddPropertyEditor(property gdnative.String, editor ControlImplementer) {
	//log.Println("Calling EditorInspectorPlugin.AddPropertyEditor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(property)
	ptrArguments[1] = gdnative.NewPointerFromObject(editor.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspectorPlugin", "add_property_editor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false label String} { false properties PoolStringArray} { false editor Control}], Returns: void
*/
func (o *EditorInspectorPlugin) AddPropertyEditorForMultipleProperties(label gdnative.String, properties gdnative.PoolStringArray, editor ControlImplementer) {
	//log.Println("Calling EditorInspectorPlugin.AddPropertyEditorForMultipleProperties()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(label)
	ptrArguments[1] = gdnative.NewPointerFromPoolStringArray(properties)
	ptrArguments[2] = gdnative.NewPointerFromObject(editor.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspectorPlugin", "add_property_editor_for_multiple_properties")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false object Object}], Returns: bool
*/
func (o *EditorInspectorPlugin) CanHandle(object ObjectImplementer) gdnative.Bool {
	//log.Println("Calling EditorInspectorPlugin.CanHandle()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(object.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspectorPlugin", "can_handle")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*

	Args: [{ false object Object}], Returns: void
*/
func (o *EditorInspectorPlugin) ParseBegin(object ObjectImplementer) {
	//log.Println("Calling EditorInspectorPlugin.ParseBegin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(object.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspectorPlugin", "parse_begin")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false object Object} { false category String}], Returns: void
*/
func (o *EditorInspectorPlugin) ParseCategory(object ObjectImplementer, category gdnative.String) {
	//log.Println("Calling EditorInspectorPlugin.ParseCategory()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(object.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromString(category)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspectorPlugin", "parse_category")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: void
*/
func (o *EditorInspectorPlugin) ParseEnd() {
	//log.Println("Calling EditorInspectorPlugin.ParseEnd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspectorPlugin", "parse_end")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [{ false object Object} { false type int} { false path String} { false hint int} { false hint_text String} { false usage int}], Returns: bool
*/
func (o *EditorInspectorPlugin) ParseProperty(object ObjectImplementer, aType gdnative.Int, path gdnative.String, hint gdnative.Int, hintText gdnative.String, usage gdnative.Int) gdnative.Bool {
	//log.Println("Calling EditorInspectorPlugin.ParseProperty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 6, 6)
	ptrArguments[0] = gdnative.NewPointerFromObject(object.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromInt(aType)
	ptrArguments[2] = gdnative.NewPointerFromString(path)
	ptrArguments[3] = gdnative.NewPointerFromInt(hint)
	ptrArguments[4] = gdnative.NewPointerFromString(hintText)
	ptrArguments[5] = gdnative.NewPointerFromInt(usage)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspectorPlugin", "parse_property")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

// EditorInspectorPluginImplementer is an interface that implements the methods
// of the EditorInspectorPlugin class.
type EditorInspectorPluginImplementer interface {
	ReferenceImplementer
	AddCustomControl(control ControlImplementer)
	AddPropertyEditor(property gdnative.String, editor ControlImplementer)
	AddPropertyEditorForMultipleProperties(label gdnative.String, properties gdnative.PoolStringArray, editor ControlImplementer)
	CanHandle(object ObjectImplementer) gdnative.Bool
	ParseBegin(object ObjectImplementer)
	ParseCategory(object ObjectImplementer, category gdnative.String)
	ParseEnd()
	ParseProperty(object ObjectImplementer, aType gdnative.Int, path gdnative.String, hint gdnative.Int, hintText gdnative.String, usage gdnative.Int) gdnative.Bool
}
