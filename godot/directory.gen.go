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

//func NewDirectoryFromPointer(ptr gdnative.Pointer) Directory {
func new_DirectoryFromPointer(ptr gdnative.Pointer) Directory {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Directory{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Directory type. It is used to manage directories and their content (not restricted to the project folder). When creating a new [Directory], its default opened directory will be [code]res://[/code]. This may change in the future, so it is advised to always use [method open] to initialize your [Directory] where you want to operate, with explicit error checking. Here is an example on how to iterate through the files of a directory: [codeblock] func dir_contents(path): var dir = Directory.new() if dir.open(path) == OK: dir.list_dir_begin() var file_name = dir.get_next() while (file_name != ""): if dir.current_is_dir(): print("Found directory: " + file_name) else: print("Found file: " + file_name) file_name = dir.get_next() else: print("An error occurred when trying to access the path.") [/codeblock]
*/
type Directory struct {
	Reference
	owner gdnative.Object
}

func (o *Directory) BaseClass() string {
	return "_Directory"
}

/*
        Undocumented
	Args: [{ false todir String}], Returns: enum.Error
*/
func (o *Directory) ChangeDir(todir gdnative.String) gdnative.Error {
	//log.Println("Calling _Directory.ChangeDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(todir)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "change_dir")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false from String} { false to String}], Returns: enum.Error
*/
func (o *Directory) Copy(from gdnative.String, to gdnative.String) gdnative.Error {
	//log.Println("Calling _Directory.Copy()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(from)
	ptrArguments[1] = gdnative.NewPointerFromString(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "copy")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *Directory) CurrentIsDir() gdnative.Bool {
	//log.Println("Calling _Directory.CurrentIsDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "current_is_dir")

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
	Args: [{ false path String}], Returns: bool
*/
func (o *Directory) DirExists(path gdnative.String) gdnative.Bool {
	//log.Println("Calling _Directory.DirExists()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "dir_exists")

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
	Args: [{ false path String}], Returns: bool
*/
func (o *Directory) FileExists(path gdnative.String) gdnative.Bool {
	//log.Println("Calling _Directory.FileExists()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "file_exists")

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
	Args: [], Returns: String
*/
func (o *Directory) GetCurrentDir() gdnative.String {
	//log.Println("Calling _Directory.GetCurrentDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "get_current_dir")

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
	Args: [], Returns: int
*/
func (o *Directory) GetCurrentDrive() gdnative.Int {
	//log.Println("Calling _Directory.GetCurrentDrive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "get_current_drive")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false idx int}], Returns: String
*/
func (o *Directory) GetDrive(idx gdnative.Int) gdnative.String {
	//log.Println("Calling _Directory.GetDrive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "get_drive")

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
	Args: [], Returns: int
*/
func (o *Directory) GetDriveCount() gdnative.Int {
	//log.Println("Calling _Directory.GetDriveCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "get_drive_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *Directory) GetNext() gdnative.String {
	//log.Println("Calling _Directory.GetNext()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "get_next")

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
	Args: [], Returns: int
*/
func (o *Directory) GetSpaceLeft() gdnative.Int {
	//log.Println("Calling _Directory.GetSpaceLeft()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "get_space_left")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{False true skip_navigational bool} {False true skip_hidden bool}], Returns: enum.Error
*/
func (o *Directory) ListDirBegin(skipNavigational gdnative.Bool, skipHidden gdnative.Bool) gdnative.Error {
	//log.Println("Calling _Directory.ListDirBegin()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromBool(skipNavigational)
	ptrArguments[1] = gdnative.NewPointerFromBool(skipHidden)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "list_dir_begin")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [], Returns: void
*/
func (o *Directory) ListDirEnd() {
	//log.Println("Calling _Directory.ListDirEnd()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "list_dir_end")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *Directory) MakeDir(path gdnative.String) gdnative.Error {
	//log.Println("Calling _Directory.MakeDir()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "make_dir")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *Directory) MakeDirRecursive(path gdnative.String) gdnative.Error {
	//log.Println("Calling _Directory.MakeDirRecursive()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "make_dir_recursive")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *Directory) Open(path gdnative.String) gdnative.Error {
	//log.Println("Calling _Directory.Open()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "open")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *Directory) Remove(path gdnative.String) gdnative.Error {
	//log.Println("Calling _Directory.Remove()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "remove")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Undocumented
	Args: [{ false from String} { false to String}], Returns: enum.Error
*/
func (o *Directory) Rename(from gdnative.String, to gdnative.String) gdnative.Error {
	//log.Println("Calling _Directory.Rename()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(from)
	ptrArguments[1] = gdnative.NewPointerFromString(to)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("_Directory", "rename")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

// DirectoryImplementer is an interface that implements the methods
// of the Directory class.
type DirectoryImplementer interface {
	ReferenceImplementer
	CurrentIsDir() gdnative.Bool
	DirExists(path gdnative.String) gdnative.Bool
	FileExists(path gdnative.String) gdnative.Bool
	GetCurrentDir() gdnative.String
	GetCurrentDrive() gdnative.Int
	GetDrive(idx gdnative.Int) gdnative.String
	GetDriveCount() gdnative.Int
	GetNext() gdnative.String
	GetSpaceLeft() gdnative.Int
	ListDirEnd()
}
