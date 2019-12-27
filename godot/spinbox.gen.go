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

//func NewSpinBoxFromPointer(ptr gdnative.Pointer) SpinBox {
func newSpinBoxFromPointer(ptr gdnative.Pointer) SpinBox {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := SpinBox{}
	obj.SetBaseObject(owner)

	return obj
}

/*
SpinBox is a numerical input text field. It allows entering integers and floats. [b]Example:[/b] [codeblock] var spin_box = SpinBox.new() add_child(spin_box) var line_edit = spin_box.get_line_edit() line_edit.context_menu_enabled = false spin_box.align = LineEdit.ALIGN_RIGHT [/codeblock] The above code will create a [SpinBox], disable context menu on it and set the text alignment to right. See [Range] class for more options over the [SpinBox].
*/
type SpinBox struct {
	Range
	owner gdnative.Object
}

func (o *SpinBox) BaseClass() string {
	return "SpinBox"
}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *SpinBox) X_GuiInput(arg0 InputEventImplementer) {
	//log.Println("Calling SpinBox.X_GuiInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "_gui_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *SpinBox) X_LineEditFocusExit() {
	//log.Println("Calling SpinBox.X_LineEditFocusExit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "_line_edit_focus_exit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *SpinBox) X_LineEditInput(arg0 InputEventImplementer) {
	//log.Println("Calling SpinBox.X_LineEditInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "_line_edit_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *SpinBox) X_RangeClickTimeout() {
	//log.Println("Calling SpinBox.X_RangeClickTimeout()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "_range_click_timeout")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String}], Returns: void
*/
func (o *SpinBox) X_TextEntered(arg0 gdnative.String) {
	//log.Println("Calling SpinBox.X_TextEntered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "_text_entered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.LineEdit::Align
*/
func (o *SpinBox) GetAlign() LineEditAlign {
	//log.Println("Calling SpinBox.GetAlign()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "get_align")

	// Call the parent method.
	// enum.LineEdit::Align
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return LineEditAlign(ret)
}

/*
        Returns the [LineEdit] instance from this [SpinBox]. You can use it to access properties and methods of [LineEdit].
	Args: [], Returns: LineEdit
*/
func (o *SpinBox) GetLineEdit() LineEditImplementer {
	//log.Println("Calling SpinBox.GetLineEdit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "get_line_edit")

	// Call the parent method.
	// LineEdit
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newLineEditFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(LineEditImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "LineEdit" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(LineEditImplementer)
	}

	return &ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *SpinBox) GetPrefix() gdnative.String {
	//log.Println("Calling SpinBox.GetPrefix()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "get_prefix")

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
func (o *SpinBox) GetSuffix() gdnative.String {
	//log.Println("Calling SpinBox.GetSuffix()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "get_suffix")

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
func (o *SpinBox) IsEditable() gdnative.Bool {
	//log.Println("Calling SpinBox.IsEditable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "is_editable")

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
	Args: [{ false align int}], Returns: void
*/
func (o *SpinBox) SetAlign(align gdnative.Int) {
	//log.Println("Calling SpinBox.SetAlign()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(align)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "set_align")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false editable bool}], Returns: void
*/
func (o *SpinBox) SetEditable(editable gdnative.Bool) {
	//log.Println("Calling SpinBox.SetEditable()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(editable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "set_editable")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false prefix String}], Returns: void
*/
func (o *SpinBox) SetPrefix(prefix gdnative.String) {
	//log.Println("Calling SpinBox.SetPrefix()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(prefix)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "set_prefix")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false suffix String}], Returns: void
*/
func (o *SpinBox) SetSuffix(suffix gdnative.String) {
	//log.Println("Calling SpinBox.SetSuffix()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(suffix)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("SpinBox", "set_suffix")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// SpinBoxImplementer is an interface that implements the methods
// of the SpinBox class.
type SpinBoxImplementer interface {
	RangeImplementer
	X_LineEditFocusExit()
	X_LineEditInput(arg0 InputEventImplementer)
	X_RangeClickTimeout()
	X_TextEntered(arg0 gdnative.String)
	GetLineEdit() LineEditImplementer
	GetPrefix() gdnative.String
	GetSuffix() gdnative.String
	IsEditable() gdnative.Bool
	SetAlign(align gdnative.Int)
	SetEditable(editable gdnative.Bool)
	SetPrefix(prefix gdnative.String)
	SetSuffix(suffix gdnative.String)
}
