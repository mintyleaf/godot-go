package godot

import (
	"github.com/gabstv/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewResourceFormatLoaderCryptoFromPointer(ptr gdnative.Pointer) ResourceFormatLoaderCrypto {
func newResourceFormatLoaderCryptoFromPointer(ptr gdnative.Pointer) ResourceFormatLoaderCrypto {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ResourceFormatLoaderCrypto{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type ResourceFormatLoaderCrypto struct {
	ResourceFormatLoader
	owner gdnative.Object
}

func (o *ResourceFormatLoaderCrypto) BaseClass() string {
	return "ResourceFormatLoaderCrypto"
}

// ResourceFormatLoaderCryptoImplementer is an interface that implements the methods
// of the ResourceFormatLoaderCrypto class.
type ResourceFormatLoaderCryptoImplementer interface {
	ResourceFormatLoaderImplementer
}
