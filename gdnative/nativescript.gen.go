package gdnative

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "types.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
#cgo CFLAGS: -I../godot_headers
#include "gdnative.gen.h"
// #include <godot_headers/nativescript/godot_nativescript.h>
// Include all headers for now. TODO: Look up all the required
// headers we need to import based on the method arguments and return types.
#include <gdnative/aabb.h>
#include <gdnative/array.h>
#include <gdnative/basis.h>
#include <gdnative/color.h>
#include <gdnative/dictionary.h>
#include <gdnative/gdnative.h>
#include <gdnative/node_path.h>
#include <gdnative/plane.h>
#include <gdnative/pool_arrays.h>
#include <gdnative/quat.h>
#include <gdnative/rect2.h>
#include <gdnative/rid.h>
#include <gdnative/string.h>
#include <gdnative/string_name.h>
#include <gdnative/transform.h>
#include <gdnative/transform2d.h>
#include <gdnative/variant.h>
#include <gdnative/vector2.h>
#include <gdnative/vector3.h>
#include <gdnative_api_struct.gen.h>
*/
import "C"
import "unsafe"

// MethodRpcMode is a Go wrapper for the C.godot_method_rpc_mode enum type.
type MethodRpcMode int

func (e MethodRpcMode) getBase() C.godot_method_rpc_mode {
	return C.godot_method_rpc_mode(e)
}

const (
	MethodRpcModeDisabled   MethodRpcMode = 0
	MethodRpcModeRemote     MethodRpcMode = 1
	MethodRpcModeMaster     MethodRpcMode = 2
	MethodRpcModePuppet     MethodRpcMode = 3
	MethodRpcModeSlave      MethodRpcMode = 4
	MethodRpcModeRemotesync MethodRpcMode = 5
	MethodRpcModeSync       MethodRpcMode = 6
	MethodRpcModeMastersync MethodRpcMode = 7
	MethodRpcModePuppetsync MethodRpcMode = 8
)

// MethodRpcModeLookupMap is a string-based lookup table of constants for MethodRpcMode.
var MethodRpcModeLookupMap = map[string]MethodRpcMode{
	"MethodRpcModeDisabled":   MethodRpcModeDisabled,
	"MethodRpcModeRemote":     MethodRpcModeRemote,
	"MethodRpcModeMaster":     MethodRpcModeMaster,
	"MethodRpcModePuppet":     MethodRpcModePuppet,
	"MethodRpcModeSlave":      MethodRpcModeSlave,
	"MethodRpcModeRemotesync": MethodRpcModeRemotesync,
	"MethodRpcModeSync":       MethodRpcModeSync,
	"MethodRpcModeMastersync": MethodRpcModeMastersync,
	"MethodRpcModePuppetsync": MethodRpcModePuppetsync,
}

// PropertyHint is a Go wrapper for the C.godot_property_hint enum type.
type PropertyHint int

func (e PropertyHint) getBase() C.godot_property_hint {
	return C.godot_property_hint(e)
}

const (
	PropertyHintNone                  PropertyHint = 0 // < no hint provided.
	PropertyHintRange                 PropertyHint = 1 // < hint_text = "min,max,step,slider;
	PropertyHintExpRange              PropertyHint = 2 // < hint_text = "min,max,step", exponential edit
	PropertyHintEnum                  PropertyHint = 3 // < hint_text= "val1,val2,val3,etc"
	PropertyHintExpEasing             PropertyHint = 4 //  exponential easing function (Math::ease)
	PropertyHintLength                PropertyHint = 5 // < hint_text= "length" (as integer)
	PropertyHintSpriteFrame           PropertyHint = 6 // FIXME: Obsolete: drop whenever we can break compat
	PropertyHintKeyAccel              PropertyHint = 7 // < hint_text= "length" (as integer)
	PropertyHintFlags                 PropertyHint = 8 // < hint_text= "flag1,flag2,etc" (as bit flags)
	PropertyHintLayers2DRender        PropertyHint = 9
	PropertyHintLayers2DPhysics       PropertyHint = 10
	PropertyHintLayers3DRender        PropertyHint = 11
	PropertyHintLayers3DPhysics       PropertyHint = 12
	PropertyHintFile                  PropertyHint = 13 // < a file path must be passed, hint_text (optionally) is a filter "*.png,*.wav,*.doc,"
	PropertyHintDir                   PropertyHint = 14 // < a directory path must be passed
	PropertyHintGlobalFile            PropertyHint = 15 // < a file path must be passed, hint_text (optionally) is a filter "*.png,*.wav,*.doc,"
	PropertyHintGlobalDir             PropertyHint = 16 // < a directory path must be passed
	PropertyHintResourceType          PropertyHint = 17 // < a resource object type
	PropertyHintMultilineText         PropertyHint = 18 // < used for string properties that can contain multiple lines
	PropertyHintPlaceholderText       PropertyHint = 19 // < used to set a placeholder text for string properties
	PropertyHintColorNoAlpha          PropertyHint = 20 // < used for ignoring alpha component when editing a color
	PropertyHintImageCompressLossy    PropertyHint = 21
	PropertyHintImageCompressLossless PropertyHint = 22
	PropertyHintObjectId              PropertyHint = 23
	PropertyHintTypeString            PropertyHint = 24 // < a type string, the hint is the base type to choose
	PropertyHintNodePathToEditedNode  PropertyHint = 25 // < so something else can provide this (used in scripts)
	PropertyHintMethodOfVariantType   PropertyHint = 26 // < a method of a type
	PropertyHintMethodOfBaseType      PropertyHint = 27 // < a method of a base type
	PropertyHintMethodOfInstance      PropertyHint = 28 // < a method of an instance
	PropertyHintMethodOfScript        PropertyHint = 29 // < a method of a script & base
	PropertyHintPropertyOfVariantType PropertyHint = 30 // < a property of a type
	PropertyHintPropertyOfBaseType    PropertyHint = 31 // < a property of a base type
	PropertyHintPropertyOfInstance    PropertyHint = 32 // < a property of an instance
	PropertyHintPropertyOfScript      PropertyHint = 33 // < a property of a script & base
	PropertyHintMax                   PropertyHint = 34
)

// PropertyHintLookupMap is a string-based lookup table of constants for PropertyHint.
var PropertyHintLookupMap = map[string]PropertyHint{
	"PropertyHintNone":                  PropertyHintNone,
	"PropertyHintRange":                 PropertyHintRange,
	"PropertyHintExpRange":              PropertyHintExpRange,
	"PropertyHintEnum":                  PropertyHintEnum,
	"PropertyHintExpEasing":             PropertyHintExpEasing,
	"PropertyHintLength":                PropertyHintLength,
	"PropertyHintSpriteFrame":           PropertyHintSpriteFrame,
	"PropertyHintKeyAccel":              PropertyHintKeyAccel,
	"PropertyHintFlags":                 PropertyHintFlags,
	"PropertyHintLayers2DRender":        PropertyHintLayers2DRender,
	"PropertyHintLayers2DPhysics":       PropertyHintLayers2DPhysics,
	"PropertyHintLayers3DRender":        PropertyHintLayers3DRender,
	"PropertyHintLayers3DPhysics":       PropertyHintLayers3DPhysics,
	"PropertyHintFile":                  PropertyHintFile,
	"PropertyHintDir":                   PropertyHintDir,
	"PropertyHintGlobalFile":            PropertyHintGlobalFile,
	"PropertyHintGlobalDir":             PropertyHintGlobalDir,
	"PropertyHintResourceType":          PropertyHintResourceType,
	"PropertyHintMultilineText":         PropertyHintMultilineText,
	"PropertyHintPlaceholderText":       PropertyHintPlaceholderText,
	"PropertyHintColorNoAlpha":          PropertyHintColorNoAlpha,
	"PropertyHintImageCompressLossy":    PropertyHintImageCompressLossy,
	"PropertyHintImageCompressLossless": PropertyHintImageCompressLossless,
	"PropertyHintObjectId":              PropertyHintObjectId,
	"PropertyHintTypeString":            PropertyHintTypeString,
	"PropertyHintNodePathToEditedNode":  PropertyHintNodePathToEditedNode,
	"PropertyHintMethodOfVariantType":   PropertyHintMethodOfVariantType,
	"PropertyHintMethodOfBaseType":      PropertyHintMethodOfBaseType,
	"PropertyHintMethodOfInstance":      PropertyHintMethodOfInstance,
	"PropertyHintMethodOfScript":        PropertyHintMethodOfScript,
	"PropertyHintPropertyOfVariantType": PropertyHintPropertyOfVariantType,
	"PropertyHintPropertyOfBaseType":    PropertyHintPropertyOfBaseType,
	"PropertyHintPropertyOfInstance":    PropertyHintPropertyOfInstance,
	"PropertyHintPropertyOfScript":      PropertyHintPropertyOfScript,
	"PropertyHintMax":                   PropertyHintMax,
}

// NewEmptyPropertyAttributes will return a pointer to an empty
// initialized PropertyAttributes. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyPropertyAttributes() Pointer {
	var obj C.godot_property_attributes
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromPropertyAttributes will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromPropertyAttributes(obj PropertyAttributes) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewPropertyAttributesFromPointer will return a PropertyAttributes from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewPropertyAttributesFromPointer(ptr Pointer) PropertyAttributes {

	return PropertyAttributes{base: (*C.godot_property_attributes)(ptr.getBase())}
}

type PropertyAttributes struct {
	base *C.godot_property_attributes

	RsetType     MethodRpcMode
	Type         Int
	Hint         PropertyHint
	HintString   String
	Usage        PropertyUsageFlags
	DefaultValue Variant
}

func (gdt PropertyAttributes) getBase() *C.godot_property_attributes {
	return gdt.base
}

// NewEmptySignalArgument will return a pointer to an empty
// initialized SignalArgument. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptySignalArgument() Pointer {
	var obj C.godot_signal_argument
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromSignalArgument will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromSignalArgument(obj SignalArgument) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewSignalArgumentFromPointer will return a SignalArgument from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewSignalArgumentFromPointer(ptr Pointer) SignalArgument {

	return SignalArgument{base: (*C.godot_signal_argument)(ptr.getBase())}
}

type SignalArgument struct {
	base *C.godot_signal_argument

	Name         String
	Type         Int
	Hint         PropertyHint
	HintString   String
	Usage        PropertyUsageFlags
	DefaultValue Variant
}

func (gdt SignalArgument) getBase() *C.godot_signal_argument {
	return gdt.base
}

// NewEmptySignal will return a pointer to an empty
// initialized Signal. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptySignal() Pointer {
	var obj C.godot_signal
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromSignal will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromSignal(obj Signal) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewSignalFromPointer will return a Signal from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewSignalFromPointer(ptr Pointer) Signal {

	return Signal{base: (*C.godot_signal)(ptr.getBase())}
}

type Signal struct {
	base *C.godot_signal

	Name           String
	NumArgs        Int
	Args           []SignalArgument
	NumDefaultArgs Int
	DefaultArgs    []Variant
}

func (gdt Signal) getBase() *C.godot_signal {
	return gdt.base
}

// NewEmptyMethodArg will return a pointer to an empty
// initialized MethodArg. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyMethodArg() Pointer {
	var obj C.godot_method_arg
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromMethodArg will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromMethodArg(obj MethodArg) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewMethodArgFromPointer will return a MethodArg from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewMethodArgFromPointer(ptr Pointer) MethodArg {

	return MethodArg{base: (*C.godot_method_arg)(ptr.getBase())}
}

type MethodArg struct {
	base *C.godot_method_arg

	Name       String
	Type       VariantType
	Hint       PropertyHint
	HintString String
}

func (gdt MethodArg) getBase() *C.godot_method_arg {
	return gdt.base
}
