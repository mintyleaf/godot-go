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
#include "gdnative.gen.h"
#include <gdnative/node_path.h>
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

// NewEmptyNodePath will return a pointer to an empty
// initialized NodePath. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyNodePath() Pointer {
	var obj C.godot_node_path
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromNodePath will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromNodePath(obj NodePath) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewNodePathFromPointer will return a NodePath from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewNodePathFromPointer(ptr Pointer) NodePath {

	return NodePath{base: (*C.godot_node_path)(ptr.getBase())}
}

type NodePath struct {
	base *C.godot_node_path
}

func (gdt NodePath) getBase() *C.godot_node_path {
	return gdt.base
}

// NewNodePath godot_node_path_new [[godot_node_path * r_dest] [const godot_string * p_from]] void
func NewNodePath(from String) NodePath {
	var dest C.godot_node_path
	arg1 := from.getBase()
	C.go_godot_node_path_new(GDNative.api, &dest, arg1)
	return NodePath{base: &dest}
}

// NewNodePathCopy godot_node_path_new_copy [[godot_node_path * r_dest] [const godot_node_path * p_src]] void
func NewNodePathCopy(src NodePath) NodePath {
	var dest C.godot_node_path
	arg1 := src.getBase()
	C.go_godot_node_path_new_copy(GDNative.api, &dest, arg1)
	return NodePath{base: &dest}
}

// Destroy godot_node_path_destroy [[godot_node_path * p_self]] void
func (gdt *NodePath) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_node_path_destroy(GDNative.api, arg0)
}

// AsString godot_node_path_as_string [[const godot_node_path * p_self]] godot_string
func (gdt *NodePath) AsString() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_node_path_as_string(GDNative.api, arg0)

	length := C.go_godot_string_length(GDNative.api, &ret)
	wchar := C.go_godot_string_wide_str(GDNative.api, &ret)
	goWchar := newWcharTWithLength(wchar, int(length))
	return String(goWchar.AsString())

}

// IsAbsolute godot_node_path_is_absolute [[const godot_node_path * p_self]] godot_bool
func (gdt *NodePath) IsAbsolute() Bool {
	arg0 := gdt.getBase()

	ret := C.go_godot_node_path_is_absolute(GDNative.api, arg0)

	return Bool(ret)
}

// GetNameCount godot_node_path_get_name_count [[const godot_node_path * p_self]] godot_int
func (gdt *NodePath) GetNameCount() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_node_path_get_name_count(GDNative.api, arg0)

	return Int(ret)
}

// GetName godot_node_path_get_name [[const godot_node_path * p_self] [const godot_int p_idx]] godot_string
func (gdt *NodePath) GetName(idx Int) String {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_node_path_get_name(GDNative.api, arg0, arg1)

	length := C.go_godot_string_length(GDNative.api, &ret)
	wchar := C.go_godot_string_wide_str(GDNative.api, &ret)
	goWchar := newWcharTWithLength(wchar, int(length))
	return String(goWchar.AsString())

}

// GetSubnameCount godot_node_path_get_subname_count [[const godot_node_path * p_self]] godot_int
func (gdt *NodePath) GetSubnameCount() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_node_path_get_subname_count(GDNative.api, arg0)

	return Int(ret)
}

// GetSubname godot_node_path_get_subname [[const godot_node_path * p_self] [const godot_int p_idx]] godot_string
func (gdt *NodePath) GetSubname(idx Int) String {
	arg0 := gdt.getBase()
	arg1 := idx.getBase()

	ret := C.go_godot_node_path_get_subname(GDNative.api, arg0, arg1)

	length := C.go_godot_string_length(GDNative.api, &ret)
	wchar := C.go_godot_string_wide_str(GDNative.api, &ret)
	goWchar := newWcharTWithLength(wchar, int(length))
	return String(goWchar.AsString())

}

// GetConcatenatedSubnames godot_node_path_get_concatenated_subnames [[const godot_node_path * p_self]] godot_string
func (gdt *NodePath) GetConcatenatedSubnames() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_node_path_get_concatenated_subnames(GDNative.api, arg0)

	length := C.go_godot_string_length(GDNative.api, &ret)
	wchar := C.go_godot_string_wide_str(GDNative.api, &ret)
	goWchar := newWcharTWithLength(wchar, int(length))
	return String(goWchar.AsString())

}

// IsEmpty godot_node_path_is_empty [[const godot_node_path * p_self]] godot_bool
func (gdt *NodePath) IsEmpty() Bool {
	arg0 := gdt.getBase()

	ret := C.go_godot_node_path_is_empty(GDNative.api, arg0)

	return Bool(ret)
}

// OperatorEqual godot_node_path_operator_equal [[const godot_node_path * p_self] [const godot_node_path * p_b]] godot_bool
func (gdt *NodePath) OperatorEqual(b NodePath) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_node_path_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}
