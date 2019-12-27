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

//func NewWebRTCPeerConnectionGDNativeFromPointer(ptr gdnative.Pointer) WebRTCPeerConnectionGDNative {
func newWebRTCPeerConnectionGDNativeFromPointer(ptr gdnative.Pointer) WebRTCPeerConnectionGDNative {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := WebRTCPeerConnectionGDNative{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Undocumented
*/
type WebRTCPeerConnectionGDNative struct {
	WebRTCPeerConnection
	owner gdnative.Object
}

func (o *WebRTCPeerConnectionGDNative) BaseClass() string {
	return "WebRTCPeerConnectionGDNative"
}

// WebRTCPeerConnectionGDNativeImplementer is an interface that implements the methods
// of the WebRTCPeerConnectionGDNative class.
type WebRTCPeerConnectionGDNativeImplementer interface {
	WebRTCPeerConnectionImplementer
}
