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

//func NewCryptoKeyFromPointer(ptr gdnative.Pointer) CryptoKey {
func newCryptoKeyFromPointer(ptr gdnative.Pointer) CryptoKey {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := CryptoKey{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The CryptoKey class represents a cryptographic key. Keys can be loaded and saved like any other [Resource]. They can be used to generate a self-signed [X509Certificate] via [method Crypto.generate_self_signed_certificate] and as private key in [method StreamPeerSSL.accept_stream] along with the appropriate certificate. [b]Note:[/b] Not available in HTML5 exports.
*/
type CryptoKey struct {
	Resource
	owner gdnative.Object
}

func (o *CryptoKey) BaseClass() string {
	return "CryptoKey"
}

/*
        Loads a key from [code]path[/code] ("*.key" file).
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *CryptoKey) Load(path gdnative.String) gdnative.Error {
	//log.Println("Calling CryptoKey.Load()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CryptoKey", "load")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Saves a key to the given [code]path[/code] (should be a "*.key" file).
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *CryptoKey) Save(path gdnative.String) gdnative.Error {
	//log.Println("Calling CryptoKey.Save()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("CryptoKey", "save")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

// CryptoKeyImplementer is an interface that implements the methods
// of the CryptoKey class.
type CryptoKeyImplementer interface {
	ResourceImplementer
}
