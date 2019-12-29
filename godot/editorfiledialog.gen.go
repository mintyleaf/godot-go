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

// EditorFileDialogAccess is an enum for Access values.
type EditorFileDialogAccess int

const (
	EditorFileDialogAccessFilesystem EditorFileDialogAccess = 2
	EditorFileDialogAccessResources  EditorFileDialogAccess = 0
	EditorFileDialogAccessUserdata   EditorFileDialogAccess = 1
)

// EditorFileDialogDisplayMode is an enum for DisplayMode values.
type EditorFileDialogDisplayMode int

const (
	EditorFileDialogDisplayList       EditorFileDialogDisplayMode = 1
	EditorFileDialogDisplayThumbnails EditorFileDialogDisplayMode = 0
)

// EditorFileDialogMode is an enum for Mode values.
type EditorFileDialogMode int

const (
	EditorFileDialogModeOpenAny   EditorFileDialogMode = 3
	EditorFileDialogModeOpenDir   EditorFileDialogMode = 2
	EditorFileDialogModeOpenFile  EditorFileDialogMode = 0
	EditorFileDialogModeOpenFiles EditorFileDialogMode = 1
	EditorFileDialogModeSaveFile  EditorFileDialogMode = 4
)

//func NewEditorFileDialogFromPointer(ptr gdnative.Pointer) EditorFileDialog {
func newEditorFileDialogFromPointer(ptr gdnative.Pointer) EditorFileDialog {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := EditorFileDialog{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type EditorFileDialog struct {
	ConfirmationDialog
	owner gdnative.Object
}

func (o *EditorFileDialog) BaseClass() string {
	return "EditorFileDialog"
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_ActionPressed() {
	//log.Println("Calling EditorFileDialog.X_ActionPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_action_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_CancelPressed() {
	//log.Println("Calling EditorFileDialog.X_CancelPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_cancel_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String}], Returns: void
*/
func (o *EditorFileDialog) X_DirEntered(arg0 gdnative.String) {
	//log.Println("Calling EditorFileDialog.X_DirEntered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_dir_entered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_FavoriteMoveDown() {
	//log.Println("Calling EditorFileDialog.X_FavoriteMoveDown()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_favorite_move_down")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_FavoriteMoveUp() {
	//log.Println("Calling EditorFileDialog.X_FavoriteMoveUp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_favorite_move_up")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_FavoritePressed() {
	//log.Println("Calling EditorFileDialog.X_FavoritePressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_favorite_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *EditorFileDialog) X_FavoriteSelected(arg0 gdnative.Int) {
	//log.Println("Calling EditorFileDialog.X_FavoriteSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_favorite_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String}], Returns: void
*/
func (o *EditorFileDialog) X_FileEntered(arg0 gdnative.String) {
	//log.Println("Calling EditorFileDialog.X_FileEntered()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_file_entered")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *EditorFileDialog) X_FilterSelected(arg0 gdnative.Int) {
	//log.Println("Calling EditorFileDialog.X_FilterSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_filter_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_GoBack() {
	//log.Println("Calling EditorFileDialog.X_GoBack()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_go_back")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_GoForward() {
	//log.Println("Calling EditorFileDialog.X_GoForward()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_go_forward")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_GoUp() {
	//log.Println("Calling EditorFileDialog.X_GoUp()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_go_up")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *EditorFileDialog) X_ItemDbSelected(arg0 gdnative.Int) {
	//log.Println("Calling EditorFileDialog.X_ItemDbSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_item_db_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int} { false arg1 Vector2}], Returns: void
*/
func (o *EditorFileDialog) X_ItemListItemRmbSelected(arg0 gdnative.Int, arg1 gdnative.Vector2) {
	//log.Println("Calling EditorFileDialog.X_ItemListItemRmbSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)
	ptrArguments[1] = gdnative.NewPointerFromVector2(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_item_list_item_rmb_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 Vector2}], Returns: void
*/
func (o *EditorFileDialog) X_ItemListRmbClicked(arg0 gdnative.Vector2) {
	//log.Println("Calling EditorFileDialog.X_ItemListRmbClicked()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromVector2(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_item_list_rmb_clicked")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *EditorFileDialog) X_ItemMenuIdPressed(arg0 gdnative.Int) {
	//log.Println("Calling EditorFileDialog.X_ItemMenuIdPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_item_menu_id_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *EditorFileDialog) X_ItemSelected(arg0 gdnative.Int) {
	//log.Println("Calling EditorFileDialog.X_ItemSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_item_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_ItemsClearSelection() {
	//log.Println("Calling EditorFileDialog.X_ItemsClearSelection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_items_clear_selection")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_MakeDir() {
	//log.Println("Calling EditorFileDialog.X_MakeDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_make_dir")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_MakeDirConfirm() {
	//log.Println("Calling EditorFileDialog.X_MakeDirConfirm()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_make_dir_confirm")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int} { false arg1 bool}], Returns: void
*/
func (o *EditorFileDialog) X_MultiSelected(arg0 gdnative.Int, arg1 gdnative.Bool) {
	//log.Println("Calling EditorFileDialog.X_MultiSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)
	ptrArguments[1] = gdnative.NewPointerFromBool(arg1)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_multi_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *EditorFileDialog) X_RecentSelected(arg0 gdnative.Int) {
	//log.Println("Calling EditorFileDialog.X_RecentSelected()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_recent_selected")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_SaveConfirmPressed() {
	//log.Println("Calling EditorFileDialog.X_SaveConfirmPressed()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_save_confirm_pressed")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 int}], Returns: void
*/
func (o *EditorFileDialog) X_SelectDrive(arg0 gdnative.Int) {
	//log.Println("Calling EditorFileDialog.X_SelectDrive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_select_drive")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 Texture} { false arg2 Texture} { false arg3 Variant}], Returns: void
*/
func (o *EditorFileDialog) X_ThumbnailDone(arg0 gdnative.String, arg1 TextureImplementer, arg2 TextureImplementer, arg3 gdnative.Variant) {
	//log.Println("Calling EditorFileDialog.X_ThumbnailDone()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromObject(arg1.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromObject(arg2.GetBaseObject())
	ptrArguments[3] = gdnative.NewPointerFromVariant(arg3)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_thumbnail_done")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 String} { false arg1 Texture} { false arg2 Texture} { false arg3 Variant}], Returns: void
*/
func (o *EditorFileDialog) X_ThumbnailResult(arg0 gdnative.String, arg1 TextureImplementer, arg2 TextureImplementer, arg3 gdnative.Variant) {
	//log.Println("Calling EditorFileDialog.X_ThumbnailResult()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromString(arg0)
	ptrArguments[1] = gdnative.NewPointerFromObject(arg1.GetBaseObject())
	ptrArguments[2] = gdnative.NewPointerFromObject(arg2.GetBaseObject())
	ptrArguments[3] = gdnative.NewPointerFromVariant(arg3)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_thumbnail_result")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false arg0 InputEvent}], Returns: void
*/
func (o *EditorFileDialog) X_UnhandledInput(arg0 InputEventImplementer) {
	//log.Println("Calling EditorFileDialog.X_UnhandledInput()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(arg0.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_unhandled_input")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_UpdateDir() {
	//log.Println("Calling EditorFileDialog.X_UpdateDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_update_dir")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_UpdateFileList() {
	//log.Println("Calling EditorFileDialog.X_UpdateFileList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_update_file_list")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *EditorFileDialog) X_UpdateFileName() {
	//log.Println("Calling EditorFileDialog.X_UpdateFileName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "_update_file_name")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a comma-delimited file extension filter option to the [EditorFileDialog] with an optional semi-colon-delimited label. For example, [code]"*.tscn, *.scn; Scenes"[/code] results in filter text "Scenes (*.tscn, *.scn)".
	Args: [{ false filter String}], Returns: void
*/
func (o *EditorFileDialog) AddFilter(filter gdnative.String) {
	//log.Println("Calling EditorFileDialog.AddFilter()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(filter)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "add_filter")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Removes all filters except for "All Files (*)".
	Args: [], Returns: void
*/
func (o *EditorFileDialog) ClearFilters() {
	//log.Println("Calling EditorFileDialog.ClearFilters()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "clear_filters")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: enum.EditorFileDialog::Access
*/
func (o *EditorFileDialog) GetAccess() EditorFileDialogAccess {
	//log.Println("Calling EditorFileDialog.GetAccess()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "get_access")

	// Call the parent method.
	// enum.EditorFileDialog::Access
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return EditorFileDialogAccess(ret)
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *EditorFileDialog) GetCurrentDir() gdnative.String {
	//log.Println("Calling EditorFileDialog.GetCurrentDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "get_current_dir")

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
func (o *EditorFileDialog) GetCurrentFile() gdnative.String {
	//log.Println("Calling EditorFileDialog.GetCurrentFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "get_current_file")

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
func (o *EditorFileDialog) GetCurrentPath() gdnative.String {
	//log.Println("Calling EditorFileDialog.GetCurrentPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "get_current_path")

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
	Args: [], Returns: enum.EditorFileDialog::DisplayMode
*/
func (o *EditorFileDialog) GetDisplayMode() EditorFileDialogDisplayMode {
	//log.Println("Calling EditorFileDialog.GetDisplayMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "get_display_mode")

	// Call the parent method.
	// enum.EditorFileDialog::DisplayMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return EditorFileDialogDisplayMode(ret)
}

/*
        Undocumented
	Args: [], Returns: enum.EditorFileDialog::Mode
*/
func (o *EditorFileDialog) GetMode() EditorFileDialogMode {
	//log.Println("Calling EditorFileDialog.GetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "get_mode")

	// Call the parent method.
	// enum.EditorFileDialog::Mode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return EditorFileDialogMode(ret)
}

/*
        Returns the [code]VBoxContainer[/code] used to display the file system.
	Args: [], Returns: VBoxContainer
*/
func (o *EditorFileDialog) GetVbox() VBoxContainerImplementer {
	//log.Println("Calling EditorFileDialog.GetVbox()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "get_vbox")

	// Call the parent method.
	// VBoxContainer
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newVBoxContainerFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(VBoxContainerImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "VBoxContainer" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(VBoxContainerImplementer)
	}

	return &ret
}

/*
        Notify the [EditorFileDialog] that its view of the data is no longer accurate. Updates the view contents on next view update.
	Args: [], Returns: void
*/
func (o *EditorFileDialog) Invalidate() {
	//log.Println("Calling EditorFileDialog.Invalidate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "invalidate")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *EditorFileDialog) IsOverwriteWarningDisabled() gdnative.Bool {
	//log.Println("Calling EditorFileDialog.IsOverwriteWarningDisabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "is_overwrite_warning_disabled")

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
func (o *EditorFileDialog) IsShowingHiddenFiles() gdnative.Bool {
	//log.Println("Calling EditorFileDialog.IsShowingHiddenFiles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "is_showing_hidden_files")

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
	Args: [{ false access int}], Returns: void
*/
func (o *EditorFileDialog) SetAccess(access gdnative.Int) {
	//log.Println("Calling EditorFileDialog.SetAccess()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(access)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "set_access")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false dir String}], Returns: void
*/
func (o *EditorFileDialog) SetCurrentDir(dir gdnative.String) {
	//log.Println("Calling EditorFileDialog.SetCurrentDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(dir)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "set_current_dir")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false file String}], Returns: void
*/
func (o *EditorFileDialog) SetCurrentFile(file gdnative.String) {
	//log.Println("Calling EditorFileDialog.SetCurrentFile()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(file)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "set_current_file")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false path String}], Returns: void
*/
func (o *EditorFileDialog) SetCurrentPath(path gdnative.String) {
	//log.Println("Calling EditorFileDialog.SetCurrentPath()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "set_current_path")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false disable bool}], Returns: void
*/
func (o *EditorFileDialog) SetDisableOverwriteWarning(disable gdnative.Bool) {
	//log.Println("Calling EditorFileDialog.SetDisableOverwriteWarning()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(disable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "set_disable_overwrite_warning")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *EditorFileDialog) SetDisplayMode(mode gdnative.Int) {
	//log.Println("Calling EditorFileDialog.SetDisplayMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "set_display_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *EditorFileDialog) SetMode(mode gdnative.Int) {
	//log.Println("Calling EditorFileDialog.SetMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "set_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false show bool}], Returns: void
*/
func (o *EditorFileDialog) SetShowHiddenFiles(show gdnative.Bool) {
	//log.Println("Calling EditorFileDialog.SetShowHiddenFiles()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(show)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("EditorFileDialog", "set_show_hidden_files")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// EditorFileDialogImplementer is an interface that implements the methods
// of the EditorFileDialog class.
type EditorFileDialogImplementer interface {
	ConfirmationDialogImplementer
	X_ActionPressed()
	X_CancelPressed()
	X_DirEntered(arg0 gdnative.String)
	X_FavoriteMoveDown()
	X_FavoriteMoveUp()
	X_FavoritePressed()
	X_FavoriteSelected(arg0 gdnative.Int)
	X_FileEntered(arg0 gdnative.String)
	X_FilterSelected(arg0 gdnative.Int)
	X_GoBack()
	X_GoForward()
	X_GoUp()
	X_ItemDbSelected(arg0 gdnative.Int)
	X_ItemListItemRmbSelected(arg0 gdnative.Int, arg1 gdnative.Vector2)
	X_ItemListRmbClicked(arg0 gdnative.Vector2)
	X_ItemMenuIdPressed(arg0 gdnative.Int)
	X_ItemSelected(arg0 gdnative.Int)
	X_ItemsClearSelection()
	X_MakeDir()
	X_MakeDirConfirm()
	X_MultiSelected(arg0 gdnative.Int, arg1 gdnative.Bool)
	X_RecentSelected(arg0 gdnative.Int)
	X_SaveConfirmPressed()
	X_SelectDrive(arg0 gdnative.Int)
	X_ThumbnailDone(arg0 gdnative.String, arg1 TextureImplementer, arg2 TextureImplementer, arg3 gdnative.Variant)
	X_ThumbnailResult(arg0 gdnative.String, arg1 TextureImplementer, arg2 TextureImplementer, arg3 gdnative.Variant)
	X_UpdateDir()
	X_UpdateFileList()
	X_UpdateFileName()
	AddFilter(filter gdnative.String)
	ClearFilters()
	GetCurrentDir() gdnative.String
	GetCurrentFile() gdnative.String
	GetCurrentPath() gdnative.String
	GetVbox() VBoxContainerImplementer
	Invalidate()
	IsOverwriteWarningDisabled() gdnative.Bool
	IsShowingHiddenFiles() gdnative.Bool
	SetAccess(access gdnative.Int)
	SetCurrentDir(dir gdnative.String)
	SetCurrentFile(file gdnative.String)
	SetCurrentPath(path gdnative.String)
	SetDisableOverwriteWarning(disable gdnative.Bool)
	SetDisplayMode(mode gdnative.Int)
	SetMode(mode gdnative.Int)
	SetShowHiddenFiles(show gdnative.Bool)
}
