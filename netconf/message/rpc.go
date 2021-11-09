package message

import (
	"encoding/xml"
	"fmt"
)

// RPC is used as a wrapper for any sent RPC
type RPC struct {
	XMLName   xml.Name `xml:"urn:ietf:params:xml:ns:netconf:base:1.0 rpc"`
	MessageID string   `xml:"message-id,attr"`
}

// RPCError defines an error reply to a RPC request
type RPCError struct {
	Type     string `xml:"error-type"`
	Tag      string `xml:"error-tag"`
	Severity string `xml:"error-severity"`
	Path     string `xml:"error-path"`
	Message  string `xml:"error-message"`
	Info     string `xml:",innerxml"`
}

// Error generates a string representation of the provided RPC error
func (re *RPCError) Error() string {
	return fmt.Sprintf("netconf rpc [%s] '%s'", re.Severity, re.Message)
}

// RPCReply defines a reply to a RPC request
type RPCReply struct {
	XMLName   xml.Name   `xml:"rpc-reply"`
	MessageID string     `xml:"message-id,attr"`
	Errors    []RPCError `xml:"rpc-error,omitempty"`
	Data      string     `xml:",innerxml"`
	Ok        bool       `xml:"ok,omitempty"`
	RawReply  string     `xml:"-"`
}

// NewRPCReply creates an instance of an RPCReply based on what was received
func NewRPCReply(rawXML []byte) (*RPCReply, error) {
	reply := &RPCReply{}
	reply.RawReply = string(rawXML)

	if err := xml.Unmarshal(rawXML, reply); err != nil {
		return nil, err
	}

	return reply, nil
}
