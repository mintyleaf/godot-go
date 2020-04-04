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

//func NewCryptoFromPointer(ptr gdnative.Pointer) Crypto {
func newCryptoFromPointer(ptr gdnative.Pointer) Crypto {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := Crypto{}
	obj.SetBaseObject(owner)

	return obj
}

/*
The Crypto class allows you to access some more advanced cryptographic functionalities in Godot. For now, this includes generating cryptographically secure random bytes, and RSA keys and self-signed X509 certificates generation. More functionalities are planned for future releases. [codeblock] extends Node var crypto = Crypto.new() var key = CryptoKey.new() var cert = X509Certificate.new() func _ready(): # Generate new RSA key. key = crypto.generate_rsa(4096) # Generate new self-signed certificate with the given key. cert = crypto.generate_self_signed_certificate(key, "CN=mydomain.com,O=My Game Company,C=IT") # Save key and certificate in the user folder. key.save("user://generated.key") cert.save("user://generated.crt") [/codeblock] [b]Note:[/b] Not available in HTML5 exports.
*/
type Crypto struct {
	Reference
	owner gdnative.Object
}

func (o *Crypto) BaseClass() string {
	return "Crypto"
}

/*
        Generates a [PackedByteArray] of cryptographically secure random bytes with given [code]size[/code].
	Args: [{ false size int}], Returns: PoolByteArray
*/
func (o *Crypto) GenerateRandomBytes(size gdnative.Int) gdnative.PoolByteArray {
	//log.Println("Calling Crypto.GenerateRandomBytes()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Crypto", "generate_random_bytes")

	// Call the parent method.
	// PoolByteArray
	retPtr := gdnative.NewEmptyPoolByteArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolByteArrayFromPointer(retPtr)
	return ret
}

/*
        Generates an RSA [CryptoKey] that can be used for creating self-signed certificates and passed to [method StreamPeerSSL.accept_stream].
	Args: [{ false size int}], Returns: CryptoKey
*/
func (o *Crypto) GenerateRsa(size gdnative.Int) CryptoKeyImplementer {
	//log.Println("Calling Crypto.GenerateRsa()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(size)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Crypto", "generate_rsa")

	// Call the parent method.
	// CryptoKey
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newCryptoKeyFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(CryptoKeyImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "CryptoKey" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(CryptoKeyImplementer)
	}

	return &ret
}

/*
        Generates a self-signed [X509Certificate] from the given [CryptoKey] and [code]issuer_name[/code]. The certificate validity will be defined by [code]not_before[/code] and [code]not_after[/code] (first valid date and last valid date). The [code]issuer_name[/code] must contain at least "CN=" (common name, i.e. the domain name), "O=" (organization, i.e. your company name), "C=" (country, i.e. 2 lettered ISO-3166 code of the country the organization is based in). A small example to generate an RSA key and a X509 self-signed certificate. [codeblock] var crypto = Crypto.new() # Generate 4096 bits RSA key. var key = crypto.generate_rsa(4096) # Generate self-signed certificate using the given key. var cert = crypto.generate_self_signed_certificate(key, "CN=example.com,O=A Game Company,C=IT") [/codeblock]
	Args: [{ false key CryptoKey} {CN=myserver,O=myorganisation,C=IT true issuer_name String} {20140101000000 true not_before String} {20340101000000 true not_after String}], Returns: X509Certificate
*/
func (o *Crypto) GenerateSelfSignedCertificate(key CryptoKeyImplementer, issuerName gdnative.String, notBefore gdnative.String, notAfter gdnative.String) X509CertificateImplementer {
	//log.Println("Calling Crypto.GenerateSelfSignedCertificate()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 4, 4)
	ptrArguments[0] = gdnative.NewPointerFromObject(key.GetBaseObject())
	ptrArguments[1] = gdnative.NewPointerFromString(issuerName)
	ptrArguments[2] = gdnative.NewPointerFromString(notBefore)
	ptrArguments[3] = gdnative.NewPointerFromString(notAfter)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("Crypto", "generate_self_signed_certificate")

	// Call the parent method.
	// X509Certificate
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := newX509CertificateFromPointer(retPtr)

	// Check to see if we already have an instance of this object in our Go instance registry.
	if instance, ok := InstanceRegistry.Get(ret.GetBaseObject().ID()); ok {
		return instance.(X509CertificateImplementer)
	}

	// Check to see what kind of class this is and create it. This is generally used with
	// GetNode().
	className := ret.GetClass()
	if className != "X509Certificate" {
		actualRet := getActualClass(className, ret.GetBaseObject())
		return actualRet.(X509CertificateImplementer)
	}

	return &ret
}

// CryptoImplementer is an interface that implements the methods
// of the Crypto class.
type CryptoImplementer interface {
	ReferenceImplementer
	GenerateRandomBytes(size gdnative.Int) gdnative.PoolByteArray
	GenerateRsa(size gdnative.Int) CryptoKeyImplementer
	GenerateSelfSignedCertificate(key CryptoKeyImplementer, issuerName gdnative.String, notBefore gdnative.String, notAfter gdnative.String) X509CertificateImplementer
}
