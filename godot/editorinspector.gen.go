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

//func NewEditorInspectorFromPointer(ptr gdnative.Pointer) EditorInspector {
func newEditorInspectorFromPointer(ptr gdnative.Pointer) EditorInspector {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorInspector{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type EditorInspector struct {
	ScrollContainer
	owner gdnative.Object
}

func (o *EditorInspector) BaseClass() string {
	return "EditorInspector"
}

/*
        Undocumented
	Args: [{ false arg0 Object} { false arg1 String}], Returns: void
*/
func (o *EditorInspector) X_EditRequestChange(arg0 ObjectImplementer, arg1 gdnative.String) {
	//log.Println("Calling EditorInspector.X_EditRequestChange()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromString(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_edit_request_change")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorInspector) X_FeatureProfileChanged() {
	//log.Println("Calling EditorInspector.X_FeatureProfileChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_feature_profile_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String}], Returns: void
*/
func (o *EditorInspector) X_FilterChanged(arg0 gdnative.String) {
	//log.Println("Calling EditorInspector.X_FilterChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_filter_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 PoolStringArray} { false arg1 Array}], Returns: void
*/
func (o *EditorInspector) X_MultiplePropertiesChanged(arg0 gdnative.PoolStringArray, arg1 gdnative.Array) {
	//log.Println("Calling EditorInspector.X_MultiplePropertiesChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromPoolStringArray(arg0)
	ptrArguments[1] = gdnative.NewPointerFromArray(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_multiple_properties_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Node}], Returns: void
*/
func (o *EditorInspector) X_NodeRemoved(arg0 NodeImplementer) {
	//log.Println("Calling EditorInspector.X_NodeRemoved()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_node_removed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 int}], Returns: void
*/
func (o *EditorInspector) X_ObjectIdSelected(arg0 gdnative.String, arg1 gdnative.Int) {
	//log.Println("Calling EditorInspector.X_ObjectIdSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromInt(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_object_id_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 Variant} { true arg2 String} {False true arg3 bool}], Returns: void
*/
func (o *EditorInspector) X_PropertyChanged(arg0 gdnative.String, arg1 gdnative.Variant, arg2 gdnative.String, arg3 gdnative.Bool) {
	//log.Println("Calling EditorInspector.X_PropertyChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromVariant(arg1)
	ptrArguments[2] = gdnative.NewPointerFromString(arg2)
	ptrArguments[3] = gdnative.NewPointerFromBool(arg3)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_property_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 Variant} { false arg2 String} { false arg3 bool}], Returns: void
*/
func (o *EditorInspector) X_PropertyChangedUpdateAll(arg0 gdnative.String, arg1 gdnative.Variant, arg2 gdnative.String, arg3 gdnative.Bool) {
	//log.Println("Calling EditorInspector.X_PropertyChangedUpdateAll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromVariant(arg1)
	ptrArguments[2] = gdnative.NewPointerFromString(arg2)
	ptrArguments[3] = gdnative.NewPointerFromBool(arg3)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_property_changed_update_all")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 bool}], Returns: void
*/
func (o *EditorInspector) X_PropertyChecked(arg0 gdnative.String, arg1 gdnative.Bool) {
	//log.Println("Calling EditorInspector.X_PropertyChecked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromBool(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_property_checked")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 bool}], Returns: void
*/
func (o *EditorInspector) X_PropertyKeyed(arg0 gdnative.String, arg1 gdnative.Bool) {
	//log.Println("Calling EditorInspector.X_PropertyKeyed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromBool(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_property_keyed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 Variant} { false arg2 bool}], Returns: void
*/
func (o *EditorInspector) X_PropertyKeyedWithValue(arg0 gdnative.String, arg1 gdnative.Variant, arg2 gdnative.Bool) {
	//log.Println("Calling EditorInspector.X_PropertyKeyedWithValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromVariant(arg1)
	ptrArguments[2] = gdnative.NewPointerFromBool(arg2)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_property_keyed_with_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 int}], Returns: void
*/
func (o *EditorInspector) X_PropertySelected(arg0 gdnative.String, arg1 gdnative.Int) {
	//log.Println("Calling EditorInspector.X_PropertySelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromInt(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_property_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 Resource}], Returns: void
*/
func (o *EditorInspector) X_ResourceSelected(arg0 gdnative.String, arg1 ResourceImplementer) {
	//log.Println("Calling EditorInspector.X_ResourceSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromObject(arg1.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_resource_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *EditorInspector) X_VscrollChanged(arg0 gdnative.Real) {
	//log.Println("Calling EditorInspector.X_VscrollChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "_vscroll_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*

	Args: [], Returns: void
*/
func (o *EditorInspector) Refresh() {
	//log.Println("Calling EditorInspector.Refresh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorInspector", "refresh")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorInspectorImplementer is an interface that implements the methods
// of the EditorInspector class.
type EditorInspectorImplementer interface {
	ScrollContainerImplementer
	X_EditRequestChange(arg0 ObjectImplementer, arg1 gdnative.String)
	X_FeatureProfileChanged()
	X_FilterChanged(arg0 gdnative.String)
	X_MultiplePropertiesChanged(arg0 gdnative.PoolStringArray, arg1 gdnative.Array)
	X_NodeRemoved(arg0 NodeImplementer)
	X_ObjectIdSelected(arg0 gdnative.String, arg1 gdnative.Int)
	X_PropertyChanged(arg0 gdnative.String, arg1 gdnative.Variant, arg2 gdnative.String, arg3 gdnative.Bool)
	X_PropertyChangedUpdateAll(arg0 gdnative.String, arg1 gdnative.Variant, arg2 gdnative.String, arg3 gdnative.Bool)
	X_PropertyChecked(arg0 gdnative.String, arg1 gdnative.Bool)
	X_PropertyKeyed(arg0 gdnative.String, arg1 gdnative.Bool)
	X_PropertyKeyedWithValue(arg0 gdnative.String, arg1 gdnative.Variant, arg2 gdnative.Bool)
	X_PropertySelected(arg0 gdnative.String, arg1 gdnative.Int)
	X_ResourceSelected(arg0 gdnative.String, arg1 ResourceImplementer)
	X_VscrollChanged(arg0 gdnative.Real)
	Refresh()
}
