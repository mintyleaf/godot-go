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

//func NewEditorPropertyFromPointer(ptr gdnative.Pointer) EditorProperty {
func newEditorPropertyFromPointer(ptr gdnative.Pointer) EditorProperty {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorProperty{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This control allows property editing for one or multiple properties into [EditorInspector]. It is added via [EditorInspectorPlugin].
*/
type EditorProperty struct {
	Container
	owner gdnative.Object
}

func (o *EditorProperty) BaseClass() string {
	return "EditorProperty"
}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *EditorProperty) X_FocusableFocused(arg0 gdnative.Int) {
	//log.Println("Calling EditorProperty.X_FocusableFocused()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "_focusable_focused")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *EditorProperty) X_GuiInput(arg0 InputEventImplementer) {
	//log.Println("Calling EditorProperty.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If any of the controls added can gain keyboard focus, add it here. This ensures that focus will be restored if the inspector is refreshed.
	Args: [{ false control Control}], Returns: void
*/
func (o *EditorProperty) AddFocusable(control ControlImplementer) {
	//log.Println("Calling EditorProperty.AddFocusable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(control.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "add_focusable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        If one or several properties have changed, this must be called. [code]field[/code] is used in case your editor can modify fields separately (as an example, Vector3.x). The [code]changing[/code] argument avoids the editor requesting this property to be refreshed (leave as [code]false[/code] if unsure).
	Args: [{ false property String} { false value Variant} { true field String} {False true changing bool}], Returns: void
*/
func (o *EditorProperty) EmitChanged(property gdnative.String, value gdnative.Variant, field gdnative.String, changing gdnative.Bool) {
	//log.Println("Calling EditorProperty.EmitChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(property)
	ptrArguments[1] = gdnative.NewPointerFromVariant(value)
	ptrArguments[2] = gdnative.NewPointerFromString(field)
	ptrArguments[3] = gdnative.NewPointerFromBool(changing)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "emit_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Gets the edited object.
	Args: [], Returns: Object
*/
func (o *EditorProperty) GetEditedObject() ObjectImplementer {
	//log.Println("Calling EditorProperty.GetEditedObject()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "get_edited_object")

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
        Gets the edited property. If your editor is for a single property (added via [method EditorInspectorPlugin.parse_property]), then this will return the property.
	Args: [], Returns: String
*/
func (o *EditorProperty) GetEditedProperty() gdnative.String {
	//log.Println("Calling EditorProperty.GetEditedProperty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "get_edited_property")

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
func (o *EditorProperty) GetLabel() gdnative.String {
	//log.Println("Calling EditorProperty.GetLabel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "get_label")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Override if you want to allow a custom tooltip over your property.
	Args: [], Returns: String
*/
func (o *EditorProperty) GetTooltipText() gdnative.String {
	//log.Println("Calling EditorProperty.GetTooltipText()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "get_tooltip_text")

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
func (o *EditorProperty) IsCheckable() gdnative.Bool {
	//log.Println("Calling EditorProperty.IsCheckable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "is_checkable")

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
func (o *EditorProperty) IsChecked() gdnative.Bool {
	//log.Println("Calling EditorProperty.IsChecked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "is_checked")

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
func (o *EditorProperty) IsDrawRed() gdnative.Bool {
	//log.Println("Calling EditorProperty.IsDrawRed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "is_draw_red")

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
func (o *EditorProperty) IsKeying() gdnative.Bool {
	//log.Println("Calling EditorProperty.IsKeying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "is_keying")

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
func (o *EditorProperty) IsReadOnly() gdnative.Bool {
	//log.Println("Calling EditorProperty.IsReadOnly()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "is_read_only")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Adds controls with this function if you want them on the bottom (below the label).
	Args: [{ false editor Control}], Returns: void
*/
func (o *EditorProperty) SetBottomEditor(editor ControlImplementer) {
	//log.Println("Calling EditorProperty.SetBottomEditor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(editor.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "set_bottom_editor")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false checkable bool}], Returns: void
*/
func (o *EditorProperty) SetCheckable(checkable gdnative.Bool) {
	//log.Println("Calling EditorProperty.SetCheckable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(checkable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "set_checkable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false checked bool}], Returns: void
*/
func (o *EditorProperty) SetChecked(checked gdnative.Bool) {
	//log.Println("Calling EditorProperty.SetChecked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(checked)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "set_checked")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false draw_red bool}], Returns: void
*/
func (o *EditorProperty) SetDrawRed(drawRed gdnative.Bool) {
	//log.Println("Calling EditorProperty.SetDrawRed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(drawRed)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "set_draw_red")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false keying bool}], Returns: void
*/
func (o *EditorProperty) SetKeying(keying gdnative.Bool) {
	//log.Println("Calling EditorProperty.SetKeying()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(keying)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "set_keying")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false text String}], Returns: void
*/
func (o *EditorProperty) SetLabel(text gdnative.String) {
	//log.Println("Calling EditorProperty.SetLabel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(text)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "set_label")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false read_only bool}], Returns: void
*/
func (o *EditorProperty) SetReadOnly(readOnly gdnative.Bool) {
	//log.Println("Calling EditorProperty.SetReadOnly()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(readOnly)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "set_read_only")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        When this virtual function is called, you must update your editor.
	Args: [], Returns: void
*/
func (o *EditorProperty) UpdateProperty() {
	//log.Println("Calling EditorProperty.UpdateProperty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorProperty", "update_property")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorPropertyImplementer is an interface that implements the methods
// of the EditorProperty class.
type EditorPropertyImplementer interface {
	ContainerImplementer
	X_FocusableFocused(arg0 gdnative.Int)
	AddFocusable(control ControlImplementer)
	EmitChanged(property gdnative.String, value gdnative.Variant, field gdnative.String, changing gdnative.Bool)
	GetEditedObject() ObjectImplementer
	GetEditedProperty() gdnative.String
	GetLabel() gdnative.String
	GetTooltipText() gdnative.String
	IsCheckable() gdnative.Bool
	IsChecked() gdnative.Bool
	IsDrawRed() gdnative.Bool
	IsKeying() gdnative.Bool
	IsReadOnly() gdnative.Bool
	SetBottomEditor(editor ControlImplementer)
	SetCheckable(checkable gdnative.Bool)
	SetChecked(checked gdnative.Bool)
	SetDrawRed(drawRed gdnative.Bool)
	SetKeying(keying gdnative.Bool)
	SetLabel(text gdnative.String)
	SetReadOnly(readOnly gdnative.Bool)
	UpdateProperty()
}
