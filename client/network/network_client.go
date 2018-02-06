// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new network API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for network API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
NetworkConnect connects a container to a network
*/
func (a *Client) NetworkConnect(params *NetworkConnectParams) (*NetworkConnectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkConnectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkConnect",
		Method:             "POST",
		PathPattern:        "/networks/{id}/connect",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/octet-stream"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NetworkConnectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkConnectOK), nil

}

/*
NetworkCreate creates a network
*/
func (a *Client) NetworkCreate(params *NetworkCreateParams) (*NetworkCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkCreate",
		Method:             "POST",
		PathPattern:        "/networks/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NetworkCreateReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkCreateCreated), nil

}

/*
NetworkDelete removes a network
*/
func (a *Client) NetworkDelete(params *NetworkDeleteParams) (*NetworkDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkDelete",
		Method:             "DELETE",
		PathPattern:        "/networks/{id}",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NetworkDeleteReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkDeleteNoContent), nil

}

/*
NetworkDisconnect disconnects a container from a network
*/
func (a *Client) NetworkDisconnect(params *NetworkDisconnectParams) (*NetworkDisconnectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkDisconnectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkDisconnect",
		Method:             "POST",
		PathPattern:        "/networks/{id}/disconnect",
		ProducesMediaTypes: []string{"application/json", "text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NetworkDisconnectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkDisconnectOK), nil

}

/*
NetworkInspect inspects a network
*/
func (a *Client) NetworkInspect(params *NetworkInspectParams) (*NetworkInspectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkInspectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkInspect",
		Method:             "GET",
		PathPattern:        "/networks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NetworkInspectReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkInspectOK), nil

}

/*
NetworkList lists networks

Returns a list of networks. For details on the format, see [the network inspect endpoint](#operation/NetworkInspect).

Note that it uses a different, smaller representation of a network than inspecting a single network. For example,
the list of containers attached to the network is not propagated in API versions 1.28 and up.

*/
func (a *Client) NetworkList(params *NetworkListParams) (*NetworkListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkList",
		Method:             "GET",
		PathPattern:        "/networks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NetworkListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkListOK), nil

}

/*
NetworkPrune deletes unused networks
*/
func (a *Client) NetworkPrune(params *NetworkPruneParams) (*NetworkPruneOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewNetworkPruneParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "NetworkPrune",
		Method:             "POST",
		PathPattern:        "/networks/prune",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json", "text/plain"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &NetworkPruneReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*NetworkPruneOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
