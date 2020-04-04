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

// NetworkedMultiplayerPeerConnectionStatus is an enum for ConnectionStatus values.
type NetworkedMultiplayerPeerConnectionStatus int

const (
	NetworkedMultiplayerPeerConnectionConnected    NetworkedMultiplayerPeerConnectionStatus = 2
	NetworkedMultiplayerPeerConnectionConnecting   NetworkedMultiplayerPeerConnectionStatus = 1
	NetworkedMultiplayerPeerConnectionDisconnected NetworkedMultiplayerPeerConnectionStatus = 0
)

// NetworkedMultiplayerPeerTransferMode is an enum for TransferMode values.
type NetworkedMultiplayerPeerTransferMode int

const (
	NetworkedMultiplayerPeerTransferModeReliable          NetworkedMultiplayerPeerTransferMode = 2
	NetworkedMultiplayerPeerTransferModeUnreliable        NetworkedMultiplayerPeerTransferMode = 0
	NetworkedMultiplayerPeerTransferModeUnreliableOrdered NetworkedMultiplayerPeerTransferMode = 1
)

//func NewNetworkedMultiplayerPeerFromPointer(ptr gdnative.Pointer) NetworkedMultiplayerPeer {
func newNetworkedMultiplayerPeerFromPointer(ptr gdnative.Pointer) NetworkedMultiplayerPeer {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := NetworkedMultiplayerPeer{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Manages the connection to network peers. Assigns unique IDs to each client connected to the server.
*/
type NetworkedMultiplayerPeer struct {
	PacketPeer
	owner gdnative.Object
}

func (o *NetworkedMultiplayerPeer) BaseClass() string {
	return "NetworkedMultiplayerPeer"
}

/*
        Returns the current state of the connection. See [enum ConnectionStatus].
	Args: [], Returns: enum.NetworkedMultiplayerPeer::ConnectionStatus
*/
func (o *NetworkedMultiplayerPeer) GetConnectionStatus() NetworkedMultiplayerPeerConnectionStatus {
	//log.Println("Calling NetworkedMultiplayerPeer.GetConnectionStatus()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerPeer", "get_connection_status")

	// Call the parent method.
	// enum.NetworkedMultiplayerPeer::ConnectionStatus
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return NetworkedMultiplayerPeerConnectionStatus(ret)
}

/*
        Returns the ID of the [NetworkedMultiplayerPeer] who sent the most recent packet.
	Args: [], Returns: int
*/
func (o *NetworkedMultiplayerPeer) GetPacketPeer() gdnative.Int {
	//log.Println("Calling NetworkedMultiplayerPeer.GetPacketPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerPeer", "get_packet_peer")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: enum.NetworkedMultiplayerPeer::TransferMode
*/
func (o *NetworkedMultiplayerPeer) GetTransferMode() NetworkedMultiplayerPeerTransferMode {
	//log.Println("Calling NetworkedMultiplayerPeer.GetTransferMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerPeer", "get_transfer_mode")

	// Call the parent method.
	// enum.NetworkedMultiplayerPeer::TransferMode
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return NetworkedMultiplayerPeerTransferMode(ret)
}

/*
        Returns the ID of this [NetworkedMultiplayerPeer].
	Args: [], Returns: int
*/
func (o *NetworkedMultiplayerPeer) GetUniqueId() gdnative.Int {
	//log.Println("Calling NetworkedMultiplayerPeer.GetUniqueId()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerPeer", "get_unique_id")

	// Call the parent method.
	// int
	retPtr := gdnative.NewEmptyInt()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewIntFromPointer(retPtr)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *NetworkedMultiplayerPeer) IsRefusingNewConnections() gdnative.Bool {
	//log.Println("Calling NetworkedMultiplayerPeer.IsRefusingNewConnections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerPeer", "is_refusing_new_connections")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	return ret
}

/*
        Waits up to 1 second to receive a new network event.
	Args: [], Returns: void
*/
func (o *NetworkedMultiplayerPeer) Poll() {
	//log.Println("Calling NetworkedMultiplayerPeer.Poll()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerPeer", "poll")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false enable bool}], Returns: void
*/
func (o *NetworkedMultiplayerPeer) SetRefuseNewConnections(enable gdnative.Bool) {
	//log.Println("Calling NetworkedMultiplayerPeer.SetRefuseNewConnections()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enable)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerPeer", "set_refuse_new_connections")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Sets the peer to which packets will be sent. The [code]id[/code] can be one of: [constant TARGET_PEER_BROADCAST] to send to all connected peers, [constant TARGET_PEER_SERVER] to send to the peer acting as server, a valid peer ID to send to that specific peer, a negative peer ID to send to all peers except that one. By default, the target peer is [constant TARGET_PEER_BROADCAST].
	Args: [{ false id int}], Returns: void
*/
func (o *NetworkedMultiplayerPeer) SetTargetPeer(id gdnative.Int) {
	//log.Println("Calling NetworkedMultiplayerPeer.SetTargetPeer()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(id)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerPeer", "set_target_peer")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false mode int}], Returns: void
*/
func (o *NetworkedMultiplayerPeer) SetTransferMode(mode gdnative.Int) {
	//log.Println("Calling NetworkedMultiplayerPeer.SetTransferMode()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromInt(mode)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NetworkedMultiplayerPeer", "set_transfer_mode")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

// NetworkedMultiplayerPeerImplementer is an interface that implements the methods
// of the NetworkedMultiplayerPeer class.
type NetworkedMultiplayerPeerImplementer interface {
	PacketPeerImplementer
	GetPacketPeer() gdnative.Int
	GetUniqueId() gdnative.Int
	IsRefusingNewConnections() gdnative.Bool
	Poll()
	SetRefuseNewConnections(enable gdnative.Bool)
	SetTargetPeer(id gdnative.Int)
	SetTransferMode(mode gdnative.Int)
}
