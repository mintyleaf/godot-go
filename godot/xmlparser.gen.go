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

// XMLParserNodeType is an enum for NodeType values.
type XMLParserNodeType int

const (
	XMLParserNodeCdata      XMLParserNodeType = 5
	XMLParserNodeComment    XMLParserNodeType = 4
	XMLParserNodeElement    XMLParserNodeType = 1
	XMLParserNodeElementEnd XMLParserNodeType = 2
	XMLParserNodeNone       XMLParserNodeType = 0
	XMLParserNodeText       XMLParserNodeType = 3
	XMLParserNodeUnknown    XMLParserNodeType = 6
)

//func NewXMLParserFromPointer(ptr gdnative.Pointer) XMLParser {
func newXMLParserFromPointer(ptr gdnative.Pointer) XMLParser {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := XMLParser{}
	obj.SetBaseObject(owner)

	return obj
}

/*
This class can serve as base to make custom XML parsers. Since XML is a very flexible standard, this interface is low-level so it can be applied to any possible schema.
*/
type XMLParser struct {
	Reference
	owner gdnative.Object
}

func (o *XMLParser) BaseClass() string {
	return "XMLParser"
}

/*
        Gets the amount of attributes in the current element.
	Args: [], Returns: int
*/
func (o *XMLParser) GetAttributeCount() gdnative.Int {
	//log.Println("Calling XMLParser.GetAttributeCount()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_attribute_count")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Gets the name of the attribute specified by the index in [code]idx[/code] argument.
	Args: [{ false idx int}], Returns: String
*/
func (o *XMLParser) GetAttributeName(idx gdnative.Int) gdnative.String {
	//log.Println("Calling XMLParser.GetAttributeName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_attribute_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Gets the value of the attribute specified by the index in [code]idx[/code] argument.
	Args: [{ false idx int}], Returns: String
*/
func (o *XMLParser) GetAttributeValue(idx gdnative.Int) gdnative.String {
	//log.Println("Calling XMLParser.GetAttributeValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(idx)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_attribute_value")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Gets the current line in the parsed file (currently not implemented).
	Args: [], Returns: int
*/
func (o *XMLParser) GetCurrentLine() gdnative.Int {
	//log.Println("Calling XMLParser.GetCurrentLine()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_current_line")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Gets the value of a certain attribute of the current element by name. This will raise an error if the element has no such attribute.
	Args: [{ false name String}], Returns: String
*/
func (o *XMLParser) GetNamedAttributeValue(name gdnative.String) gdnative.String {
	//log.Println("Calling XMLParser.GetNamedAttributeValue()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_named_attribute_value")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Gets the value of a certain attribute of the current element by name. This will return an empty [String] if the attribute is not found.
	Args: [{ false name String}], Returns: String
*/
func (o *XMLParser) GetNamedAttributeValueSafe(name gdnative.String) gdnative.String {
	//log.Println("Calling XMLParser.GetNamedAttributeValueSafe()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_named_attribute_value_safe")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Gets the contents of a text node. This will raise an error in any other type of node.
	Args: [], Returns: String
*/
func (o *XMLParser) GetNodeData() gdnative.String {
	//log.Println("Calling XMLParser.GetNodeData()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_node_data")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Gets the name of the current element node. This will raise an error if the current node type is neither [constant NODE_ELEMENT] nor [constant NODE_ELEMENT_END].
	Args: [], Returns: String
*/
func (o *XMLParser) GetNodeName() gdnative.String {
	//log.Println("Calling XMLParser.GetNodeName()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_node_name")

	// Call the parent method.
	// String
	retPtr := gdnative.NewEmptyString()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewStringFromPointer(retPtr)
	return ret
}

/*
        Gets the byte offset of the current node since the beginning of the file or buffer.
	Args: [], Returns: int
*/
func (o *XMLParser) GetNodeOffset() gdnative.Int {
	//log.Println("Calling XMLParser.GetNodeOffset()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_node_offset")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Gets the type of the current node. Compare with [enum NodeType] constants.
	Args: [], Returns: enum.XMLParser::NodeType
*/
func (o *XMLParser) GetNodeType() XMLParserNodeType {
	//log.Println("Calling XMLParser.GetNodeType()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "get_node_type")

	// Call the parent method.
	// enum.XMLParser::NodeType
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return XMLParserNodeType(ret)
}

/*
        Check whether the current element has a certain attribute.
	Args: [{ false name String}], Returns: bool
*/
func (o *XMLParser) HasAttribute(name gdnative.String) gdnative.Bool {
	//log.Println("Calling XMLParser.HasAttribute()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(name)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "has_attribute")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Check whether the current element is empty (this only works for completely empty tags, e.g. [code]<element \>[/code]).
	Args: [], Returns: bool
*/
func (o *XMLParser) IsEmpty() gdnative.Bool {
	//log.Println("Calling XMLParser.IsEmpty()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "is_empty")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Opens an XML file for parsing. This returns an error code.
	Args: [{ false file String}], Returns: enum.Error
*/
func (o *XMLParser) Open(file gdnative.String) gdnative.Error {
	//log.Println("Calling XMLParser.Open()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromString(file)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "open")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Opens an XML raw buffer for parsing. This returns an error code.
	Args: [{ false buffer PoolByteArray}], Returns: enum.Error
*/
func (o *XMLParser) OpenBuffer(buffer gdnative.PoolByteArray) gdnative.Error {
	//log.Println("Calling XMLParser.OpenBuffer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromPoolByteArray(buffer)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "open_buffer")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Reads the next node of the file. This returns an error code.
	Args: [], Returns: enum.Error
*/
func (o *XMLParser) Read() gdnative.Error {
	//log.Println("Calling XMLParser.Read()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "read")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Moves the buffer cursor to a certain offset (since the beginning) and read the next node there. This returns an error code.
	Args: [{ false position int}], Returns: enum.Error
*/
func (o *XMLParser) Seek(position gdnative.Int) gdnative.Error {
	//log.Println("Calling XMLParser.Seek()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(position)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "seek")

	// Call the parent method.
	// enum.Error
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return gdnative.Error(ret)
}

/*
        Skips the current section. If the node contains other elements, they will be ignored and the cursor will go to the closing of the current element.
	Args: [], Returns: void
*/
func (o *XMLParser) SkipSection() {
	//log.Println("Calling XMLParser.SkipSection()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("XMLParser", "skip_section")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// XMLParserImplementer is an interface that implements the methods
// of the XMLParser class.
type XMLParserImplementer interface {
	ReferenceImplementer
	GetAttributeCount() gdnative.Int
	GetAttributeName(idx gdnative.Int) gdnative.String
	GetAttributeValue(idx gdnative.Int) gdnative.String
	GetCurrentLine() gdnative.Int
	GetNamedAttributeValue(name gdnative.String) gdnative.String
	GetNamedAttributeValueSafe(name gdnative.String) gdnative.String
	GetNodeData() gdnative.String
	GetNodeName() gdnative.String
	GetNodeOffset() gdnative.Int
	HasAttribute(name gdnative.String) gdnative.Bool
	IsEmpty() gdnative.Bool
	SkipSection()
}
