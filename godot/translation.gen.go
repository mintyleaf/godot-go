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

//func NewTranslationFromPointer(ptr gdnative.Pointer) Translation {
func newTranslationFromPointer(ptr gdnative.Pointer) Translation {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Translation{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Translations are resources that can be loaded and unloaded on demand. They map a string to another string.
*/
type Translation struct {
	Resource
	owner gdnative.Object
}

func (o *Translation) BaseClass() string {
	return "Translation"
}

/*
        Undocumented
	Args: [], Returns: PoolStringArray
*/
func (o *Translation) X_GetMessages() gdnative.PoolStringArray {
	//log.Println("Calling Translation.X_GetMessages()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Translation", "_get_messages")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false arg0 PoolStringArray}], Returns: void
*/
func (o *Translation) X_SetMessages(arg0 gdnative.PoolStringArray) {
	//log.Println("Calling Translation.X_SetMessages()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolStringArray(arg0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Translation", "_set_messages")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Adds a message if nonexistent, followed by its translation.
	Args: [{ false src_message String} { false xlated_message String}], Returns: void
*/
func (o *Translation) AddMessage(srcMessage gdnative.String, xlatedMessage gdnative.String) {
	//log.Println("Calling Translation.AddMessage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(srcMessage)
	ptrArguments[1] = gdnative.NewPointerFromString(xlatedMessage)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Translation", "add_message")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Erases a message.
	Args: [{ false src_message String}], Returns: void
*/
func (o *Translation) EraseMessage(srcMessage gdnative.String) {
	//log.Println("Calling Translation.EraseMessage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(srcMessage)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Translation", "erase_message")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [], Returns: String
*/
func (o *Translation) GetLocale() gdnative.String {
	//log.Println("Calling Translation.GetLocale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Translation", "get_locale")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns a message's translation.
	Args: [{ false src_message String}], Returns: String
*/
func (o *Translation) GetMessage(srcMessage gdnative.String) gdnative.String {
	//log.Println("Calling Translation.GetMessage()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(srcMessage)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Translation", "get_message")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Returns the number of existing messages.
	Args: [], Returns: int
*/
func (o *Translation) GetMessageCount() gdnative.Int {
	//log.Println("Calling Translation.GetMessageCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Translation", "get_message_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Returns all the messages (keys).
	Args: [], Returns: PoolStringArray
*/
func (o *Translation) GetMessageList() gdnative.PoolStringArray {
	//log.Println("Calling Translation.GetMessageList()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Translation", "get_message_list")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false locale String}], Returns: void
*/
func (o *Translation) SetLocale(locale gdnative.String) {
	//log.Println("Calling Translation.SetLocale()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(locale)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Translation", "set_locale")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// TranslationImplementer is an interface that implements the methods
// of the Translation class.
type TranslationImplementer interface {
	ResourceImplementer
	X_GetMessages() gdnative.PoolStringArray
	X_SetMessages(arg0 gdnative.PoolStringArray)
	AddMessage(srcMessage gdnative.String, xlatedMessage gdnative.String)
	EraseMessage(srcMessage gdnative.String)
	GetLocale() gdnative.String
	GetMessage(srcMessage gdnative.String) gdnative.String
	GetMessageCount() gdnative.Int
	GetMessageList() gdnative.PoolStringArray
	SetLocale(locale gdnative.String)
}
