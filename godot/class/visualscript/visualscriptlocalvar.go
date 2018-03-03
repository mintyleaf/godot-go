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
type VisualScriptLocalVar struct {
	VisualScriptNode
}

func (o *VisualScriptLocalVar) BaseClass() string {
	return "VisualScriptLocalVar"
}

/*
   Undocumented
*/
func (o *VisualScriptLocalVar) GetVarName() gdnative.String {
	log.Println("Calling VisualScriptLocalVar.GetVarName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_var_name", goArguments, "gdnative.String")

	returnValue := goRet.Interface().(gdnative.String)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptLocalVar) GetVarType() gdnative.Int {
	log.Println("Calling VisualScriptLocalVar.GetVarType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_var_type", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VisualScriptLocalVar) SetVarName(name gdnative.String) {
	log.Println("Calling VisualScriptLocalVar.SetVarName()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(name)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_var_name", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VisualScriptLocalVar) SetVarType(aType gdnative.Int) {
	log.Println("Calling VisualScriptLocalVar.SetVarType()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(aType)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_var_type", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VisualScriptLocalVarImplementer is an interface for VisualScriptLocalVar objects.
*/
type VisualScriptLocalVarImplementer interface {
	Class
}