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
// #include <godot_headers/gdnative_api_struct.gen.h>
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

// NewEmptyGdnativeExtAndroidApiStruct will return a pointer to an empty
// initialized GdnativeExtAndroidApiStruct. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyGdnativeExtAndroidApiStruct() Pointer {
	var obj C.godot_gdnative_ext_android_api_struct
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromGdnativeExtAndroidApiStruct will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromGdnativeExtAndroidApiStruct(obj GdnativeExtAndroidApiStruct) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewGdnativeExtAndroidApiStructFromPointer will return a GdnativeExtAndroidApiStruct from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewGdnativeExtAndroidApiStructFromPointer(ptr Pointer) GdnativeExtAndroidApiStruct {

	return GdnativeExtAndroidApiStruct{base: (*C.godot_gdnative_ext_android_api_struct)(ptr.getBase())}
}

type GdnativeExtAndroidApiStruct struct {
	base *C.godot_gdnative_ext_android_api_struct

	Type    Uint
	Version GdnativeApiVersion
}

func (gdt GdnativeExtAndroidApiStruct) getBase() *C.godot_gdnative_ext_android_api_struct {
	return gdt.base
}

// NewEmptyGdnativeExtVideodecoderApiStruct will return a pointer to an empty
// initialized GdnativeExtVideodecoderApiStruct. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyGdnativeExtVideodecoderApiStruct() Pointer {
	var obj C.godot_gdnative_ext_videodecoder_api_struct
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromGdnativeExtVideodecoderApiStruct will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromGdnativeExtVideodecoderApiStruct(obj GdnativeExtVideodecoderApiStruct) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewGdnativeExtVideodecoderApiStructFromPointer will return a GdnativeExtVideodecoderApiStruct from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewGdnativeExtVideodecoderApiStructFromPointer(ptr Pointer) GdnativeExtVideodecoderApiStruct {

	return GdnativeExtVideodecoderApiStruct{base: (*C.godot_gdnative_ext_videodecoder_api_struct)(ptr.getBase())}
}

type GdnativeExtVideodecoderApiStruct struct {
	base *C.godot_gdnative_ext_videodecoder_api_struct

	Type    Uint
	Version GdnativeApiVersion
}

func (gdt GdnativeExtVideodecoderApiStruct) getBase() *C.godot_gdnative_ext_videodecoder_api_struct {
	return gdt.base
}

// NewEmptyGdnativeExtNet32ApiStruct will return a pointer to an empty
// initialized GdnativeExtNet32ApiStruct. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyGdnativeExtNet32ApiStruct() Pointer {
	var obj C.godot_gdnative_ext_net_3_2_api_struct
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromGdnativeExtNet32ApiStruct will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromGdnativeExtNet32ApiStruct(obj GdnativeExtNet32ApiStruct) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewGdnativeExtNet32ApiStructFromPointer will return a GdnativeExtNet32ApiStruct from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewGdnativeExtNet32ApiStructFromPointer(ptr Pointer) GdnativeExtNet32ApiStruct {

	return GdnativeExtNet32ApiStruct{base: (*C.godot_gdnative_ext_net_3_2_api_struct)(ptr.getBase())}
}

type GdnativeExtNet32ApiStruct struct {
	base *C.godot_gdnative_ext_net_3_2_api_struct

	Type    Uint
	Version GdnativeApiVersion
}

func (gdt GdnativeExtNet32ApiStruct) getBase() *C.godot_gdnative_ext_net_3_2_api_struct {
	return gdt.base
}

// NewEmptyGdnativeExtNetApiStruct will return a pointer to an empty
// initialized GdnativeExtNetApiStruct. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyGdnativeExtNetApiStruct() Pointer {
	var obj C.godot_gdnative_ext_net_api_struct
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromGdnativeExtNetApiStruct will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromGdnativeExtNetApiStruct(obj GdnativeExtNetApiStruct) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewGdnativeExtNetApiStructFromPointer will return a GdnativeExtNetApiStruct from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewGdnativeExtNetApiStructFromPointer(ptr Pointer) GdnativeExtNetApiStruct {

	return GdnativeExtNetApiStruct{base: (*C.godot_gdnative_ext_net_api_struct)(ptr.getBase())}
}

type GdnativeExtNetApiStruct struct {
	base *C.godot_gdnative_ext_net_api_struct

	Type    Uint
	Version GdnativeApiVersion
}

func (gdt GdnativeExtNetApiStruct) getBase() *C.godot_gdnative_ext_net_api_struct {
	return gdt.base
}

// NewEmptyGdnativeCore12ApiStruct will return a pointer to an empty
// initialized GdnativeCore12ApiStruct. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyGdnativeCore12ApiStruct() Pointer {
	var obj C.godot_gdnative_core_1_2_api_struct
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromGdnativeCore12ApiStruct will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromGdnativeCore12ApiStruct(obj GdnativeCore12ApiStruct) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewGdnativeCore12ApiStructFromPointer will return a GdnativeCore12ApiStruct from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewGdnativeCore12ApiStructFromPointer(ptr Pointer) GdnativeCore12ApiStruct {

	return GdnativeCore12ApiStruct{base: (*C.godot_gdnative_core_1_2_api_struct)(ptr.getBase())}
}

type GdnativeCore12ApiStruct struct {
	base *C.godot_gdnative_core_1_2_api_struct

	Type    Uint
	Version GdnativeApiVersion
}

func (gdt GdnativeCore12ApiStruct) getBase() *C.godot_gdnative_core_1_2_api_struct {
	return gdt.base
}

// NewEmptyGdnativeCore11ApiStruct will return a pointer to an empty
// initialized GdnativeCore11ApiStruct. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyGdnativeCore11ApiStruct() Pointer {
	var obj C.godot_gdnative_core_1_1_api_struct
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromGdnativeCore11ApiStruct will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromGdnativeCore11ApiStruct(obj GdnativeCore11ApiStruct) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewGdnativeCore11ApiStructFromPointer will return a GdnativeCore11ApiStruct from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewGdnativeCore11ApiStructFromPointer(ptr Pointer) GdnativeCore11ApiStruct {

	return GdnativeCore11ApiStruct{base: (*C.godot_gdnative_core_1_1_api_struct)(ptr.getBase())}
}

type GdnativeCore11ApiStruct struct {
	base *C.godot_gdnative_core_1_1_api_struct

	Type    Uint
	Version GdnativeApiVersion
}

func (gdt GdnativeCore11ApiStruct) getBase() *C.godot_gdnative_core_1_1_api_struct {
	return gdt.base
}
