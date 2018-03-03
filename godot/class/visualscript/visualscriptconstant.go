package visualscript

import (
	"log"
	"reflect"

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

/*
Undocumented
*/
type VisualScriptConstant struct {
	VisualScriptNode
}

func (o *VisualScriptConstant) BaseClass() string {
	return "VisualScriptConstant"
}

/*
   Undocumented
*/
func (o *VisualScriptConstant) GetConstantType() gdnative.Int {
	log.Println("Calling VisualScriptConstant.GetConstantType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_constant_type", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptConstant) GetConstantValue() *Variant {
	log.Println("Calling VisualScriptConstant.GetConstantValue()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_constant_value", goArguments, "*Variant")

	returnValue := goRet.Interface().(*Variant)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptConstant) SetConstantType(aType gdnative.Int) {
	log.Println("Calling VisualScriptConstant.SetConstantType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(aType)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_constant_type", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VisualScriptConstant) SetConstantValue(value *Variant) {
	log.Println("Calling VisualScriptConstant.SetConstantValue()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_constant_value", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptConstantImplementer is an interface for VisualScriptConstant objects.
*/
type VisualScriptConstantImplementer interface {
	Class
}