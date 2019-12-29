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

//func NewColorPickerFromPointer(ptr gdnative.Pointer) ColorPicker {
func newColorPickerFromPointer(ptr gdnative.Pointer) ColorPicker {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ColorPicker{}
	obj.SetBaseObject(owner)

	return obj
}

/*
[Control] node displaying a color picker widget. It's useful for selecting a color from an RGB/RGBA colorspace.
*/
type ColorPicker struct {
	BoxContainer
	owner gdnative.Object
}

func (o *ColorPicker) BaseClass() string {
	return "ColorPicker"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ColorPicker) X_AddPresetPressed() {
	//log.Println("Calling ColorPicker.X_AddPresetPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_add_preset_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ColorPicker) X_FocusEnter() {
	//log.Println("Calling ColorPicker.X_FocusEnter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_focus_enter")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ColorPicker) X_FocusExit() {
	//log.Println("Calling ColorPicker.X_FocusExit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_focus_exit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int} { false arg1 Control}], Returns: void
*/
func (o *ColorPicker) X_HsvDraw(arg0 gdnative.Int, arg1 ControlImplementer) {
	//log.Println("Calling ColorPicker.X_HsvDraw()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)
	ptrArguments[1] = gdnative.NewPointerFromObject(arg1.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_hsv_draw")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String}], Returns: void
*/
func (o *ColorPicker) X_HtmlEntered(arg0 gdnative.String) {
	//log.Println("Calling ColorPicker.X_HtmlEntered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_html_entered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ColorPicker) X_HtmlFocusExit() {
	//log.Println("Calling ColorPicker.X_HtmlFocusExit()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_html_focus_exit")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *ColorPicker) X_PresetInput(arg0 InputEventImplementer) {
	//log.Println("Calling ColorPicker.X_PresetInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_preset_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ColorPicker) X_SampleDraw() {
	//log.Println("Calling ColorPicker.X_SampleDraw()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_sample_draw")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *ColorPicker) X_ScreenInput(arg0 InputEventImplementer) {
	//log.Println("Calling ColorPicker.X_ScreenInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_screen_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ColorPicker) X_ScreenPickPressed() {
	//log.Println("Calling ColorPicker.X_ScreenPickPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_screen_pick_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ColorPicker) X_TextTypeToggled() {
	//log.Println("Calling ColorPicker.X_TextTypeToggled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_text_type_toggled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *ColorPicker) X_UpdatePresets() {
	//log.Println("Calling ColorPicker.X_UpdatePresets()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_update_presets")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *ColorPicker) X_UvInput(arg0 InputEventImplementer) {
	//log.Println("Calling ColorPicker.X_UvInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_uv_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 float}], Returns: void
*/
func (o *ColorPicker) X_ValueChanged(arg0 gdnative.Real) {
	//log.Println("Calling ColorPicker.X_ValueChanged()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_value_changed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *ColorPicker) X_WInput(arg0 InputEventImplementer) {
	//log.Println("Calling ColorPicker.X_WInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "_w_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds the given color to a list of color presets. The presets are displayed in the color picker and the user will be able to select them. [b]Note:[/b] the presets list is only for [i]this[/i] color picker.
	Args: [{ false color Color}], Returns: void
*/
func (o *ColorPicker) AddPreset(color gdnative.Color) {
	//log.Println("Calling ColorPicker.AddPreset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "add_preset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *ColorPicker) ArePresetsEnabled() gdnative.Bool {
	//log.Println("Calling ColorPicker.ArePresetsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "are_presets_enabled")

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
func (o *ColorPicker) ArePresetsVisible() gdnative.Bool {
	//log.Println("Calling ColorPicker.ArePresetsVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "are_presets_visible")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Removes the given color from the list of color presets of this color picker.
	Args: [{ false color Color}], Returns: void
*/
func (o *ColorPicker) ErasePreset(color gdnative.Color) {
	//log.Println("Calling ColorPicker.ErasePreset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "erase_preset")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: Color
*/
func (o *ColorPicker) GetPickColor() gdnative.Color {
	//log.Println("Calling ColorPicker.GetPickColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "get_pick_color")

	// Call the parent method.
	// Color
	retPtr := gdnative.NewEmptyColor()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewColorFromPointer(retPtr)
	return ret
}

/*
        Returns the list of colors in the presets of the color picker.
	Args: [], Returns: PoolColorArray
*/
func (o *ColorPicker) GetPresets() gdnative.PoolColorArray {
	//log.Println("Calling ColorPicker.GetPresets()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "get_presets")

	// Call the parent method.
	// PoolColorArray
	retPtr := gdnative.NewEmptyPoolColorArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolColorArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *ColorPicker) IsDeferredMode() gdnative.Bool {
	//log.Println("Calling ColorPicker.IsDeferredMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "is_deferred_mode")

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
func (o *ColorPicker) IsEditingAlpha() gdnative.Bool {
	//log.Println("Calling ColorPicker.IsEditingAlpha()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "is_editing_alpha")

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
func (o *ColorPicker) IsHsvMode() gdnative.Bool {
	//log.Println("Calling ColorPicker.IsHsvMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "is_hsv_mode")

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
func (o *ColorPicker) IsRawMode() gdnative.Bool {
	//log.Println("Calling ColorPicker.IsRawMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "is_raw_mode")

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
	Args: [{ false mode bool}], Returns: void
*/
func (o *ColorPicker) SetDeferredMode(mode gdnative.Bool) {
	//log.Println("Calling ColorPicker.SetDeferredMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "set_deferred_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false show bool}], Returns: void
*/
func (o *ColorPicker) SetEditAlpha(show gdnative.Bool) {
	//log.Println("Calling ColorPicker.SetEditAlpha()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(show)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "set_edit_alpha")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode bool}], Returns: void
*/
func (o *ColorPicker) SetHsvMode(mode gdnative.Bool) {
	//log.Println("Calling ColorPicker.SetHsvMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "set_hsv_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false color Color}], Returns: void
*/
func (o *ColorPicker) SetPickColor(color gdnative.Color) {
	//log.Println("Calling ColorPicker.SetPickColor()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromColor(color)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "set_pick_color")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *ColorPicker) SetPresetsEnabled(enabled gdnative.Bool) {
	//log.Println("Calling ColorPicker.SetPresetsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "set_presets_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false visible bool}], Returns: void
*/
func (o *ColorPicker) SetPresetsVisible(visible gdnative.Bool) {
	//log.Println("Calling ColorPicker.SetPresetsVisible()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(visible)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "set_presets_visible")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode bool}], Returns: void
*/
func (o *ColorPicker) SetRawMode(mode gdnative.Bool) {
	//log.Println("Calling ColorPicker.SetRawMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ColorPicker", "set_raw_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ColorPickerImplementer is an interface that implements the methods
// of the ColorPicker class.
type ColorPickerImplementer interface {
	BoxContainerImplementer
	X_AddPresetPressed()
	X_FocusEnter()
	X_FocusExit()
	X_HsvDraw(arg0 gdnative.Int, arg1 ControlImplementer)
	X_HtmlEntered(arg0 gdnative.String)
	X_HtmlFocusExit()
	X_PresetInput(arg0 InputEventImplementer)
	X_SampleDraw()
	X_ScreenInput(arg0 InputEventImplementer)
	X_ScreenPickPressed()
	X_TextTypeToggled()
	X_UpdatePresets()
	X_UvInput(arg0 InputEventImplementer)
	X_ValueChanged(arg0 gdnative.Real)
	X_WInput(arg0 InputEventImplementer)
	AddPreset(color gdnative.Color)
	ArePresetsEnabled() gdnative.Bool
	ArePresetsVisible() gdnative.Bool
	ErasePreset(color gdnative.Color)
	GetPickColor() gdnative.Color
	GetPresets() gdnative.PoolColorArray
	IsDeferredMode() gdnative.Bool
	IsEditingAlpha() gdnative.Bool
	IsHsvMode() gdnative.Bool
	IsRawMode() gdnative.Bool
	SetDeferredMode(mode gdnative.Bool)
	SetEditAlpha(show gdnative.Bool)
	SetHsvMode(mode gdnative.Bool)
	SetPickColor(color gdnative.Color)
	SetPresetsEnabled(enabled gdnative.Bool)
	SetPresetsVisible(visible gdnative.Bool)
	SetRawMode(mode gdnative.Bool)
}
