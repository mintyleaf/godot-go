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

//func NewWebRTCDataChannelGDNativeFromPointer(ptr gdnative.Pointer) WebRTCDataChannelGDNative {
func newWebRTCDataChannelGDNativeFromPointer(ptr gdnative.Pointer) WebRTCDataChannelGDNative {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := WebRTCDataChannelGDNative{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type WebRTCDataChannelGDNative struct {
	WebRTCDataChannel
	owner gdnative.Object
}

func (o *WebRTCDataChannelGDNative) BaseClass() string {
	return "WebRTCDataChannelGDNative"
}

// WebRTCDataChannelGDNativeImplementer is an interface that implements the methods
// of the WebRTCDataChannelGDNative class.
type WebRTCDataChannelGDNativeImplementer interface {
	WebRTCDataChannelImplementer
}
