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

//func NewConfirmationDialogFromPointer(ptr gdnative.Pointer) ConfirmationDialog {
func newConfirmationDialogFromPointer(ptr gdnative.Pointer) ConfirmationDialog {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ConfirmationDialog{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Dialog for confirmation of actions. This dialog inherits from [AcceptDialog], but has by default an OK and Cancel button (in host OS order). To get cancel action, you can use: [codeblock] get_cancel().connect("pressed", self, "cancelled") [/codeblock].
*/
type ConfirmationDialog struct {
	AcceptDialog
	owner gdnative.Object
}

func (o *ConfirmationDialog) BaseClass() string {
	return "ConfirmationDialog"
}

/*
        Returns the cancel button.
	Args: [], Returns: Button
*/
func (o *ConfirmationDialog) GetCancel() ButtonImplementer {
	//log.Println("Calling ConfirmationDialog.GetCancel()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfirmationDialog", "get_cancel")

	// Call the parent method.
	// Button
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newButtonFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(ButtonImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "Button" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(ButtonImplementer)
	}

	return &ret
}

// ConfirmationDialogImplementer is an interface that implements the methods
// of the ConfirmationDialog class.
type ConfirmationDialogImplementer interface {
	AcceptDialogImplementer
	GetCancel() ButtonImplementer
}
