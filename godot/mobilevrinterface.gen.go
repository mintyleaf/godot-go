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

//func NewMobileVRInterfaceFromPointer(ptr gdnative.Pointer) MobileVRInterface {
func newMobileVRInterfaceFromPointer(ptr gdnative.Pointer) MobileVRInterface {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := MobileVRInterface{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type MobileVRInterface struct {
	ARVRInterface
	owner gdnative.Object
}

func (o *MobileVRInterface) BaseClass() string {
	return "MobileVRInterface"
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *MobileVRInterface) GetDisplayToLens() gdnative.Real {
	//log.Println("Calling MobileVRInterface.GetDisplayToLens()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "get_display_to_lens")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *MobileVRInterface) GetDisplayWidth() gdnative.Real {
	//log.Println("Calling MobileVRInterface.GetDisplayWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "get_display_width")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *MobileVRInterface) GetEyeHeight() gdnative.Real {
	//log.Println("Calling MobileVRInterface.GetEyeHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "get_eye_height")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *MobileVRInterface) GetIod() gdnative.Real {
	//log.Println("Calling MobileVRInterface.GetIod()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "get_iod")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *MobileVRInterface) GetK1() gdnative.Real {
	//log.Println("Calling MobileVRInterface.GetK1()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "get_k1")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *MobileVRInterface) GetK2() gdnative.Real {
	//log.Println("Calling MobileVRInterface.GetK2()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "get_k2")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: float
*/
func (o *MobileVRInterface) GetOversample() gdnative.Real {
	//log.Println("Calling MobileVRInterface.GetOversample()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "get_oversample")

	// Call the parent method.
	// float
	retPtr := gdnative.NewEmptyReal()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewRealFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [{ false display_to_lens float}], Returns: void
*/
func (o *MobileVRInterface) SetDisplayToLens(displayToLens gdnative.Real) {
	//log.Println("Calling MobileVRInterface.SetDisplayToLens()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(displayToLens)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "set_display_to_lens")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false display_width float}], Returns: void
*/
func (o *MobileVRInterface) SetDisplayWidth(displayWidth gdnative.Real) {
	//log.Println("Calling MobileVRInterface.SetDisplayWidth()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(displayWidth)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "set_display_width")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false eye_height float}], Returns: void
*/
func (o *MobileVRInterface) SetEyeHeight(eyeHeight gdnative.Real) {
	//log.Println("Calling MobileVRInterface.SetEyeHeight()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(eyeHeight)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "set_eye_height")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false iod float}], Returns: void
*/
func (o *MobileVRInterface) SetIod(iod gdnative.Real) {
	//log.Println("Calling MobileVRInterface.SetIod()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(iod)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "set_iod")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false k float}], Returns: void
*/
func (o *MobileVRInterface) SetK1(k gdnative.Real) {
	//log.Println("Calling MobileVRInterface.SetK1()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(k)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "set_k1")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false k float}], Returns: void
*/
func (o *MobileVRInterface) SetK2(k gdnative.Real) {
	//log.Println("Calling MobileVRInterface.SetK2()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(k)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "set_k2")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false oversample float}], Returns: void
*/
func (o *MobileVRInterface) SetOversample(oversample gdnative.Real) {
	//log.Println("Calling MobileVRInterface.SetOversample()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromReal(oversample)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("MobileVRInterface", "set_oversample")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// MobileVRInterfaceImplementer is an interface that implements the methods
// of the MobileVRInterface class.
type MobileVRInterfaceImplementer interface {
	ARVRInterfaceImplementer
	GetDisplayToLens() gdnative.Real
	GetDisplayWidth() gdnative.Real
	GetEyeHeight() gdnative.Real
	GetIod() gdnative.Real
	GetK1() gdnative.Real
	GetK2() gdnative.Real
	GetOversample() gdnative.Real
	SetDisplayToLens(displayToLens gdnative.Real)
	SetDisplayWidth(displayWidth gdnative.Real)
	SetEyeHeight(eyeHeight gdnative.Real)
	SetIod(iod gdnative.Real)
	SetK1(k gdnative.Real)
	SetK2(k gdnative.Real)
	SetOversample(oversample gdnative.Real)
}
