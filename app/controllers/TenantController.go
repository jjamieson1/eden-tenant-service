package controllers

import (
	"github.com/jjamieson1/eden-tenant-service/app/Models"
	"github.com/revel/revel"
)

type Api struct {
	*revel.Controller
}

func (c Api) GetTenantUserServiceProvider(tenantId string) revel.Result {
	var provider Models.UserProvider
	provider.Name = "Eden User Service"
	provider.PluginName = "eden"
	provider.CalloutUrl = "http://localhost:9100/api"
	provider.Id = "2d8ebb78-7bfb-4d33-8bd5-8fdd75a95127"
	return c.RenderJSON(provider)
}

func (c Api) SetTenantUserServiceProvider(tenantId string, provider string) revel.Result  {
	c.Response.SetStatus(200)
	return c.Result
}

func (c Api) GetUserServiceProviders() revel.Result {
	// This is the list of plugins connecting the platform to user providers.
	var provider Models.UserProvider
	var providers []Models.UserProvider

	provider.Name = "Ping Identity"
	provider.PluginName = "ping"
	provider.CalloutUrl = "http://localhost:9100/api"
	provider.Id = "4e0f5e2f-180c-47f2-a595-256fea217e1c"
	providers = append(providers, provider)

	provider.Name = "Eden User Service"
	provider.PluginName = "eden"
	provider.CalloutUrl = "http://localhost:9100/api"
	provider.Id = "2d8ebb78-7bfb-4d33-8bd5-8fdd75a95127"
	providers = append(providers, provider)

	provider.Name = "Vivvo Trust Cloud"
	provider.PluginName = "vivvo"
	provider.CalloutUrl = "http://localhost:9100/api"
	provider.Id = "75d56c9c-2b43-49f7-bdf6-6ac82c580173"
	providers = append(providers, provider)

	return c.RenderJSON(providers)
}

