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
#include <gdnative/dictionary.h>
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

// NewEmptyDictionary will return a pointer to an empty
// initialized Dictionary. This is primarily used in
// conjunction with MethodBindPtrCall.
func NewEmptyDictionary() Pointer {
	var obj C.godot_dictionary
	return Pointer{base: unsafe.Pointer(&obj)}
}

// NewPointerFromDictionary will return an unsafe pointer to the given
// object. This is primarily used in conjunction with MethodBindPtrCall.
func NewPointerFromDictionary(obj Dictionary) Pointer {
	return Pointer{base: unsafe.Pointer(obj.getBase())}
}

// NewDictionaryFromPointer will return a Dictionary from the
// given unsafe pointer. This is primarily used in conjunction with MethodBindPtrCall.
func NewDictionaryFromPointer(ptr Pointer) Dictionary {

	return Dictionary{base: (*C.godot_dictionary)(ptr.getBase())}
}

type Dictionary struct {
	base *C.godot_dictionary
}

func (gdt Dictionary) getBase() *C.godot_dictionary {
	return gdt.base
}

// NewDictionary godot_dictionary_new [[godot_dictionary * r_dest]] void
func NewDictionary() Dictionary {
	var dest C.godot_dictionary
	C.go_godot_dictionary_new(GDNative.api, &dest)
	return Dictionary{base: &dest}
}

// NewDictionaryCopy godot_dictionary_new_copy [[godot_dictionary * r_dest] [const godot_dictionary * p_src]] void
func NewDictionaryCopy(src Dictionary) Dictionary {
	var dest C.godot_dictionary
	arg1 := src.getBase()
	C.go_godot_dictionary_new_copy(GDNative.api, &dest, arg1)
	return Dictionary{base: &dest}
}

// Destroy godot_dictionary_destroy [[godot_dictionary * p_self]] void
func (gdt *Dictionary) Destroy() {
	arg0 := gdt.getBase()

	C.go_godot_dictionary_destroy(GDNative.api, arg0)
}

// Size godot_dictionary_size [[const godot_dictionary * p_self]] godot_int
func (gdt *Dictionary) Size() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_dictionary_size(GDNative.api, arg0)

	return Int(ret)
}

// Empty godot_dictionary_empty [[const godot_dictionary * p_self]] godot_bool
func (gdt *Dictionary) Empty() Bool {
	arg0 := gdt.getBase()

	ret := C.go_godot_dictionary_empty(GDNative.api, arg0)

	return Bool(ret)
}

// Clear godot_dictionary_clear [[godot_dictionary * p_self]] void
func (gdt *Dictionary) Clear() {
	arg0 := gdt.getBase()

	C.go_godot_dictionary_clear(GDNative.api, arg0)
}

// Has godot_dictionary_has [[const godot_dictionary * p_self] [const godot_variant * p_key]] godot_bool
func (gdt *Dictionary) Has(key Variant) Bool {
	arg0 := gdt.getBase()
	arg1 := key.getBase()

	ret := C.go_godot_dictionary_has(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// HasAll godot_dictionary_has_all [[const godot_dictionary * p_self] [const godot_array * p_keys]] godot_bool
func (gdt *Dictionary) HasAll(keys Array) Bool {
	arg0 := gdt.getBase()
	arg1 := keys.getBase()

	ret := C.go_godot_dictionary_has_all(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// Erase godot_dictionary_erase [[godot_dictionary * p_self] [const godot_variant * p_key]] void
func (gdt *Dictionary) Erase(key Variant) {
	arg0 := gdt.getBase()
	arg1 := key.getBase()

	C.go_godot_dictionary_erase(GDNative.api, arg0, arg1)
}

// Hash godot_dictionary_hash [[const godot_dictionary * p_self]] godot_int
func (gdt *Dictionary) Hash() Int {
	arg0 := gdt.getBase()

	ret := C.go_godot_dictionary_hash(GDNative.api, arg0)

	return Int(ret)
}

// Keys godot_dictionary_keys [[const godot_dictionary * p_self]] godot_array
func (gdt *Dictionary) Keys() Array {
	arg0 := gdt.getBase()

	ret := C.go_godot_dictionary_keys(GDNative.api, arg0)

	return Array{base: &ret}

}

// Values godot_dictionary_values [[const godot_dictionary * p_self]] godot_array
func (gdt *Dictionary) Values() Array {
	arg0 := gdt.getBase()

	ret := C.go_godot_dictionary_values(GDNative.api, arg0)

	return Array{base: &ret}

}

// Get godot_dictionary_get [[const godot_dictionary * p_self] [const godot_variant * p_key]] godot_variant
func (gdt *Dictionary) Get(key Variant) Variant {
	arg0 := gdt.getBase()
	arg1 := key.getBase()

	ret := C.go_godot_dictionary_get(GDNative.api, arg0, arg1)

	return Variant{base: &ret}

}

// Set godot_dictionary_set [[godot_dictionary * p_self] [const godot_variant * p_key] [const godot_variant * p_value]] void
func (gdt *Dictionary) Set(key Variant, value Variant) {
	arg0 := gdt.getBase()
	arg1 := key.getBase()
	arg2 := value.getBase()

	C.go_godot_dictionary_set(GDNative.api, arg0, arg1, arg2)
}

// OperatorIndex godot_dictionary_operator_index [[godot_dictionary * p_self] [const godot_variant * p_key]] godot_variant *
func (gdt *Dictionary) OperatorIndex(key Variant) Variant {
	arg0 := gdt.getBase()
	arg1 := key.getBase()

	ret := C.go_godot_dictionary_operator_index(GDNative.api, arg0, arg1)

	return Variant{base: ret}

}

// OperatorIndexConst godot_dictionary_operator_index_const [[const godot_dictionary * p_self] [const godot_variant * p_key]] const godot_variant *
func (gdt *Dictionary) OperatorIndexConst(key Variant) Variant {
	arg0 := gdt.getBase()
	arg1 := key.getBase()

	ret := C.go_godot_dictionary_operator_index_const(GDNative.api, arg0, arg1)

	return Variant{base: ret}

}

// Next godot_dictionary_next [[const godot_dictionary * p_self] [const godot_variant * p_key]] godot_variant *
func (gdt *Dictionary) Next(key Variant) Variant {
	arg0 := gdt.getBase()
	arg1 := key.getBase()

	ret := C.go_godot_dictionary_next(GDNative.api, arg0, arg1)

	return Variant{base: ret}

}

// OperatorEqual godot_dictionary_operator_equal [[const godot_dictionary * p_self] [const godot_dictionary * p_b]] godot_bool
func (gdt *Dictionary) OperatorEqual(b Dictionary) Bool {
	arg0 := gdt.getBase()
	arg1 := b.getBase()

	ret := C.go_godot_dictionary_operator_equal(GDNative.api, arg0, arg1)

	return Bool(ret)
}

// ToJson godot_dictionary_to_json [[const godot_dictionary * p_self]] godot_string
func (gdt *Dictionary) ToJson() String {
	arg0 := gdt.getBase()

	ret := C.go_godot_dictionary_to_json(GDNative.api, arg0)

	length := C.go_godot_string_length(GDNative.api, &ret)
	wchar := C.go_godot_string_wide_str(GDNative.api, &ret)
	goWchar := newWcharTWithLength(wchar, int(length))
	return String(goWchar.AsString())

}
