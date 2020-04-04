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

//func NewConfigFileFromPointer(ptr gdnative.Pointer) ConfigFile {
func newConfigFileFromPointer(ptr gdnative.Pointer) ConfigFile {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := ConfigFile{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This helper class can be used to store [Variant] values on the filesystem using INI-style formatting. The stored values are identified by a section and a key: [codeblock] [section] some_key=42 string_example="Hello World3D!" a_vector=Vector3( 1, 0, 2 ) [/codeblock] The stored data can be saved to or parsed from a file, though ConfigFile objects can also be used directly without accessing the filesystem. The following example shows how to parse an INI-style file from the system, read its contents and store new values in it: [codeblock] var config = ConfigFile.new() var err = config.load("user://settings.cfg") if err == OK: # If not, something went wrong with the file loading # Look for the display/width pair, and default to 1024 if missing var screen_width = config.get_value("display", "width", 1024) # Store a variable if and only if it hasn't been defined yet if not config.has_section_key("audio", "mute"): config.set_value("audio", "mute", false) # Save the changes by overwriting the previous file config.save("user://settings.cfg") [/codeblock] Keep in mind that section and property names can't contain spaces. Anything after a space will be ignored on save and on load. ConfigFiles can also contain manually written comment lines starting with a semicolon ([code];[/code]). Those lines will be ignored when parsing the file. Note that comments will be lost when saving the ConfigFile. This can still be useful for dedicated server configuration files, which are typically never overwritten without explicit user action.
*/
type ConfigFile struct {
	Reference
	owner gdnative.Object
}

func (o *ConfigFile) BaseClass() string {
	return "ConfigFile"
}

/*
        Deletes the specified section along with all the key-value pairs inside. Raises an error if the section does not exist.
	Args: [{ false section String}], Returns: void
*/
func (o *ConfigFile) EraseSection(section gdnative.String) {
	//log.Println("Calling ConfigFile.EraseSection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(section)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "erase_section")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Deletes the specified key in a section. Raises an error if either the section or the key do not exist.
	Args: [{ false section String} { false key String}], Returns: void
*/
func (o *ConfigFile) EraseSectionKey(section gdnative.String, key gdnative.String) {
	//log.Println("Calling ConfigFile.EraseSectionKey()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(section)
	ptrArguments[1] = gdnative.NewPointerFromString(key)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "erase_section_key")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Returns an array of all defined key identifiers in the specified section. Raises an error and returns an empty array if the section does not exist.
	Args: [{ false section String}], Returns: PoolStringArray
*/
func (o *ConfigFile) GetSectionKeys(section gdnative.String) gdnative.PoolStringArray {
	//log.Println("Calling ConfigFile.GetSectionKeys()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(section)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "get_section_keys")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Returns an array of all defined section identifiers.
	Args: [], Returns: PoolStringArray
*/
func (o *ConfigFile) GetSections() gdnative.PoolStringArray {
	//log.Println("Calling ConfigFile.GetSections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "get_sections")

	// Call the parent method.
	// PoolStringArray
	retPtr := gdnative.NewEmptyPoolStringArray()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewPoolStringArrayFromPointer(retPtr)
	return ret
}

/*
        Returns the current value for the specified section and key. If either the section or the key do not exist, the method returns the fallback [code]default[/code] value. If [code]default[/code] is not specified or set to [code]null[/code], an error is also raised.
	Args: [{ false section String} { false key String} {Null true default Variant}], Returns: Variant
*/
func (o *ConfigFile) GetValue(section gdnative.String, key gdnative.String, aDefault gdnative.Variant) gdnative.Variant {
	//log.Println("Calling ConfigFile.GetValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(section)
	ptrArguments[1] = gdnative.NewPointerFromString(key)
	ptrArguments[2] = gdnative.NewPointerFromVariant(aDefault)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "get_value")

	// Call the parent method.
	// Variant
	retPtr := gdnative.NewEmptyVariant()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewVariantFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the specified section exists.
	Args: [{ false section String}], Returns: bool
*/
func (o *ConfigFile) HasSection(section gdnative.String) gdnative.Bool {
	//log.Println("Calling ConfigFile.HasSection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(section)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "has_section")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Returns [code]true[/code] if the specified section-key pair exists.
	Args: [{ false section String} { false key String}], Returns: bool
*/
func (o *ConfigFile) HasSectionKey(section gdnative.String, key gdnative.String) gdnative.Bool {
	//log.Println("Calling ConfigFile.HasSectionKey()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(section)
	ptrArguments[1] = gdnative.NewPointerFromString(key)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "has_section_key")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Loads the config file specified as a parameter. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on. Returns one of the [enum Error] code constants ([code]OK[/code] on success).
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *ConfigFile) Load(path gdnative.String) gdnative.Error {
	//log.Println("Calling ConfigFile.Load()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "load")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Loads the encrypted config file specified as a parameter, using the provided [code]key[/code] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on. Returns one of the [enum Error] code constants ([code]OK[/code] on success).
	Args: [{ false path String} { false key PoolByteArray}], Returns: enum.Error
*/
func (o *ConfigFile) LoadEncrypted(path gdnative.String, key gdnative.PoolByteArray) gdnative.Error {
	//log.Println("Calling ConfigFile.LoadEncrypted()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromPoolByteArray(key)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "load_encrypted")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Loads the encrypted config file specified as a parameter, using the provided [code]password[/code] to decrypt it. The file's contents are parsed and loaded in the [ConfigFile] object which the method was called on. Returns one of the [enum Error] code constants ([code]OK[/code] on success).
	Args: [{ false path String} { false pass String}], Returns: enum.Error
*/
func (o *ConfigFile) LoadEncryptedPass(path gdnative.String, pass gdnative.String) gdnative.Error {
	//log.Println("Calling ConfigFile.LoadEncryptedPass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromString(pass)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "load_encrypted_pass")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Saves the contents of the [ConfigFile] object to the file specified as a parameter. The output file uses an INI-style structure. Returns one of the [enum Error] code constants ([code]OK[/code] on success).
	Args: [{ false path String}], Returns: enum.Error
*/
func (o *ConfigFile) Save(path gdnative.String) gdnative.Error {
	//log.Println("Calling ConfigFile.Save()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(path)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "save")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [code]key[/code] to encrypt it. The output file uses an INI-style structure. Returns one of the [enum Error] code constants ([code]OK[/code] on success).
	Args: [{ false path String} { false key PoolByteArray}], Returns: enum.Error
*/
func (o *ConfigFile) SaveEncrypted(path gdnative.String, key gdnative.PoolByteArray) gdnative.Error {
	//log.Println("Calling ConfigFile.SaveEncrypted()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromPoolByteArray(key)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "save_encrypted")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Saves the contents of the [ConfigFile] object to the AES-256 encrypted file specified as a parameter, using the provided [code]password[/code] to encrypt it. The output file uses an INI-style structure. Returns one of the [enum Error] code constants ([code]OK[/code] on success).
	Args: [{ false path String} { false pass String}], Returns: enum.Error
*/
func (o *ConfigFile) SaveEncryptedPass(path gdnative.String, pass gdnative.String) gdnative.Error {
	//log.Println("Calling ConfigFile.SaveEncryptedPass()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 2, 2)
	ptrArguments[0] = gdnative.NewPointerFromString(path)
	ptrArguments[1] = gdnative.NewPointerFromString(pass)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "save_encrypted_pass")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Assigns a value to the specified key of the specified section. If either the section or the key do not exist, they are created. Passing a [code]null[/code] value deletes the specified key if it exists, and deletes the section if it ends up empty once the key has been removed.
	Args: [{ false section String} { false key String} { false value Variant}], Returns: void
*/
func (o *ConfigFile) SetValue(section gdnative.String, key gdnative.String, value gdnative.Variant) {
	//log.Println("Calling ConfigFile.SetValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 3, 3)
	ptrArguments[0] = gdnative.NewPointerFromString(section)
	ptrArguments[1] = gdnative.NewPointerFromString(key)
	ptrArguments[2] = gdnative.NewPointerFromVariant(value)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("ConfigFile", "set_value")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// ConfigFileImplementer is an interface that implements the methods
// of the ConfigFile class.
type ConfigFileImplementer interface {
	ReferenceImplementer
	EraseSection(section gdnative.String)
	EraseSectionKey(section gdnative.String, key gdnative.String)
	GetSectionKeys(section gdnative.String) gdnative.PoolStringArray
	GetSections() gdnative.PoolStringArray
	GetValue(section gdnative.String, key gdnative.String, aDefault gdnative.Variant) gdnative.Variant
	HasSection(section gdnative.String) gdnative.Bool
	HasSectionKey(section gdnative.String, key gdnative.String) gdnative.Bool
	SetValue(section gdnative.String, key gdnative.String, value gdnative.Variant)
}
