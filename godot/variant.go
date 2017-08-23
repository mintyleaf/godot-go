package godot

/*
#include <stdio.h>
#include <stdlib.h>
#include <godot/gdnative.h>
#include <godot_nativescript.h>
*/
import "C"

const (
	VariantTypeNil = iota

	// atomic types
	VariantTypeBool
	VariantTypeInt
	VariantTypeReal
	VariantTypeString

	// math types
	VariantTypeVector2 // 5
	VariantTypeRect2
	VariantTypeVector3
	VariantTypeTransform2d
	VariantTypePlane
	VariantTypeQuat // 10
	VariantTypeRect3
	VariantTypeBasis
	VariantTypeTransform

	// misc types
	VariantTypeColor
	VariantTypenodePath // 15
	VariantTyperId
	VariantTypeObject
	VariantTypeDictionary
	VariantTypeArray // 20

	// ARRAYS
	VariantTypePoolByteArray
	VariantTypePoolIntArray
	VariantTypePoolRealArray
	VariantTypePoolStringArray
	VariantTypePoolVector2Array // 25
	VariantTypePoolVector3Array
	VariantTypePoolColorArray
)

func variantAsBool(variant *C.godot_variant) bool {
	godotBool := C.godot_variant_as_bool(variant)
	return bool(godotBool)
}

func boolAsVariant(value bool) *C.godot_variant {
	var variant C.godot_variant
	godotBool := C.godot_bool(value)
	C.godot_variant_new_bool(&variant, godotBool)

	return &variant
}

func variantAsUInt(variant *C.godot_variant) uint64 {
	godotUInt := C.godot_variant_as_uint(variant)
	return uint64(godotUInt)
}

func uIntAsVariant(value uint64) *C.godot_variant {
	var variant C.godot_variant
	cUInt64 := C.uint64_t(value)
	C.godot_variant_new_uint(&variant, cUInt64)

	return &variant
}

func variantAsInt(variant *C.godot_variant) int64 {
	godotInt := C.godot_variant_as_int(variant)
	return int64(godotInt)
}

func intAsVariant(value int64) *C.godot_variant {
	var variant C.godot_variant
	cInt64 := C.int64_t(value)
	C.godot_variant_new_int(&variant, cInt64)

	return &variant
}

func variantAsReal(variant *C.godot_variant) float64 {
	godotFloat := C.godot_variant_as_real(variant)
	return float64(godotFloat)
}

func realAsVariant(value float64) *C.godot_variant {
	var variant C.godot_variant
	cDouble := C.double(value)
	C.godot_variant_new_real(&variant, cDouble)

	return &variant
}

func variantAsString(variant *C.godot_variant) string {
	godotString := C.godot_variant_as_string(variant)
	godotCString := C.godot_string_c_str(&godotString)
	return C.GoString(godotCString)
}

func stringAsVariant(value string) *C.godot_variant {
	var variant C.godot_variant
	var godotString C.godot_string
	cString := C.CString(value)
	C.godot_string_new_data(&godotString, cString, C.int(len(value)))
	C.godot_variant_new_string(&variant, &godotString)

	return &variant
}

/*
TODO:
godot_array GDAPI godot_variant_as_array(const godot_variant *p_self);
godot_basis GDAPI godot_variant_as_basis(const godot_variant *p_self);
godot_color GDAPI godot_variant_as_color(const godot_variant *p_self);
godot_dictionary GDAPI godot_variant_as_dictionary(const godot_variant *p_self);
godot_node_path GDAPI godot_variant_as_node_path(const godot_variant *p_self);
godot_object GDAPI *godot_variant_as_object(const godot_variant *p_self);
godot_plane GDAPI godot_variant_as_plane(const godot_variant *p_self);
godot_pool_byte_array GDAPI godot_variant_as_pool_byte_array(const godot_variant *p_self);
godot_pool_color_array GDAPI godot_variant_as_pool_color_array(const godot_variant *p_self);
godot_pool_int_array GDAPI godot_variant_as_pool_int_array(const godot_variant *p_self);
godot_pool_real_array GDAPI godot_variant_as_pool_real_array(const godot_variant *p_self);
godot_pool_string_array GDAPI godot_variant_as_pool_string_array(const godot_variant *p_self);
godot_pool_vector2_array GDAPI godot_variant_as_pool_vector2_array(const godot_variant *p_self);
godot_pool_vector3_array GDAPI godot_variant_as_pool_vector3_array(const godot_variant *p_self);
godot_quat GDAPI godot_variant_as_quat(const godot_variant *p_self);
godot_rect2 GDAPI godot_variant_as_rect2(const godot_variant *p_self);
godot_rect3 GDAPI godot_variant_as_rect3(const godot_variant *p_self);
godot_rid GDAPI godot_variant_as_rid(const godot_variant *p_self);
godot_transform GDAPI godot_variant_as_transform(const godot_variant *p_self);
godot_transform2d GDAPI godot_variant_as_transform2d(const godot_variant *p_self);
godot_vector2 GDAPI godot_variant_as_vector2(const godot_variant *p_self);
godot_vector3 GDAPI godot_variant_as_vector3(const godot_variant *p_self);
*/