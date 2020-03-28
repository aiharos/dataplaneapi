package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	client_native "github.com/haproxytech/client-native"
	"github.com/haproxytech/dataplaneapi/haproxy"
	"github.com/haproxytech/dataplaneapi/misc"
	"github.com/haproxytech/dataplaneapi/operations/resolver"
	"github.com/haproxytech/models"
)

//CreateResolverHandlerImpl implementation of the CreateResolverHandler interface using client-native client
type CreateResolverHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

//DeleteResolverHandlerImpl implementation of the DeleteResolverHandler interface using client-native client
type DeleteResolverHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

//GetResolverHandlerImpl implementation of the GetResolverHandler interface using client-native client
type GetResolverHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//GetResolversHandlerImpl implementation of the GetResolversHandler interface using client-native client
type GetResolversHandlerImpl struct {
	Client *client_native.HAProxyClient
}

//ReplaceResolverHandlerImpl implementation of the ReplaceResolverHandler interface using client-native client
type ReplaceResolverHandlerImpl struct {
	Client      *client_native.HAProxyClient
	ReloadAgent haproxy.IReloadAgent
}

//Handle executing the request and returning a response
func (h *CreateResolverHandlerImpl) Handle(params resolver.CreateResolverParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return resolver.NewCreateResolverDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.CreateResolver(params.Data, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return resolver.NewCreateResolverDefault(int(*e.Code)).WithPayload(e)
	}

	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return resolver.NewCreateResolverDefault(int(*e.Code)).WithPayload(e)
			}
			return resolver.NewCreateResolverCreated().WithPayload(params.Data)
		}
		rID := h.ReloadAgent.Reload()
		return resolver.NewCreateResolverAccepted().WithReloadID(rID).WithPayload(params.Data)
	}
	return resolver.NewCreateResolverAccepted().WithPayload(params.Data)
}

//Handle executing the request and returning a response
func (h *DeleteResolverHandlerImpl) Handle(params resolver.DeleteResolverParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return resolver.NewDeleteResolverDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.DeleteResolver(params.Name, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return resolver.NewDeleteResolverDefault(int(*e.Code)).WithPayload(e)
	}
	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return resolver.NewDeleteResolverDefault(int(*e.Code)).WithPayload(e)
			}
			return resolver.NewDeleteResolverNoContent()
		}
		rID := h.ReloadAgent.Reload()
		return resolver.NewDeleteResolverAccepted().WithReloadID(rID)
	}
	return resolver.NewDeleteResolverAccepted()
}

//Handle executing the request and returning a response
func (h *GetResolverHandlerImpl) Handle(params resolver.GetResolverParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	v, r, err := h.Client.Configuration.GetResolver(params.Name, t)
	if err != nil {
		e := misc.HandleError(err)
		return resolver.NewGetResolverDefault(int(*e.Code)).WithPayload(e).WithConfigurationVersion(v)
	}
	return resolver.NewGetResolverOK().WithPayload(&resolver.GetResolverOKBody{Version: v, Data: r}).WithConfigurationVersion(v)
}

//Handle executing the request and returning a response
func (h *GetResolversHandlerImpl) Handle(params resolver.GetResolversParams, principal interface{}) middleware.Responder {
	t := ""
	if params.TransactionID != nil {
		t = *params.TransactionID
	}

	v, rs, err := h.Client.Configuration.GetResolvers(t)
	if err != nil {
		e := misc.HandleError(err)
		return resolver.NewGetResolversDefault(int(*e.Code)).WithPayload(e).WithConfigurationVersion(v)
	}
	return resolver.NewGetResolversOK().WithPayload(&resolver.GetResolversOKBody{Version: v, Data: rs}).WithConfigurationVersion(v)
}

//Handle executing the request and returning a response
func (h *ReplaceResolverHandlerImpl) Handle(params resolver.ReplaceResolverParams, principal interface{}) middleware.Responder {
	t := ""
	v := int64(0)
	if params.TransactionID != nil {
		t = *params.TransactionID
	}
	if params.Version != nil {
		v = *params.Version
	}

	if t != "" && *params.ForceReload {
		msg := "Both force_reload and transaction specified, specify only one"
		c := misc.ErrHTTPBadRequest
		e := &models.Error{
			Message: &msg,
			Code:    &c,
		}
		return resolver.NewReplaceResolverDefault(int(*e.Code)).WithPayload(e)
	}

	err := h.Client.Configuration.EditResolver(params.Name, params.Data, t, v)
	if err != nil {
		e := misc.HandleError(err)
		return resolver.NewReplaceResolverDefault(int(*e.Code)).WithPayload(e)
	}

	if params.TransactionID == nil {
		if *params.ForceReload {
			err := h.ReloadAgent.ForceReload()
			if err != nil {
				e := misc.HandleError(err)
				return resolver.NewReplaceResolverDefault(int(*e.Code)).WithPayload(e)
			}
			return resolver.NewReplaceResolverOK().WithPayload(params.Data)
		}
		rID := h.ReloadAgent.Reload()
		return resolver.NewReplaceResolverAccepted().WithReloadID(rID).WithPayload(params.Data)
	}
	return resolver.NewReplaceResolverAccepted().WithPayload(params.Data)
}