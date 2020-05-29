package controllers

import (
	"github.com/jjamieson1/eden-frontend/models"
	"github.com/jjamieson1/eden-tenant-service/app/services"
	"github.com/revel/revel"
)

type Api struct {
	*revel.Controller
}

func (c Api) SetTenantUserServiceProvider(tenantId string, provider string) revel.Result  {
	c.Response.SetStatus(200)
	return c.Result
}

/*
POST    /api/tenant                                                 Api.AddNewTenant
POST    /api/tenant/:tenantId                                       Api.UpdateTenant
DELETE  /api/tenant/:tenantId                                       Api.DeleteTenant

*/

func (c Api) AddNewTenant() revel.Result {
	var tenant models.Tenant
	err := c.Params.BindJSON(&tenant)
	if err != nil {
		return c.RenderJSON(err)
	}
	tenant, err = services.AddUpdateTenantDetails("", tenant)
	if err != nil {
		c.Response.SetStatus(400)
		return c.RenderJSON(err)
	}
	return c.RenderJSON(tenant)
}

func (c Api) GetTenantById(tenantId string) revel.Result {

	revel.AppLog.Debugf("requesting tenant details for: %v", tenantId)

	tenant, err := services.GetTenantDetails(tenantId)
	if err != nil {
		c.Response.SetStatus(400)
		return c.RenderJSON(err)
	}
	return c.RenderJSON(tenant)
}


func (c Api) GetTenantByUrl(url string) revel.Result {

	revel.AppLog.Debugf("requesting tenant details for url: %v", url)

	tenant, err := services.GetTenantByUrl(url)
	if err != nil {
		c.Response.SetStatus(400)
		return c.RenderJSON(err)
	}
	return c.RenderJSON(tenant)
}


func (c Api)UpdateTenant(tenantId string) revel.Result {
	var tenant models.Tenant
	err := c.Params.BindJSON(&tenant)
	if err != nil {
		c.Response.SetStatus(400)
		return c.RenderJSON(err)
	}
	tenant, err = services.AddUpdateTenantDetails(tenantId, tenant)
	if err != nil {
		c.Response.SetStatus(400)
		return c.RenderJSON(err)
	}
	return c.RenderJSON(tenant)
}

func (c Api) DeleteTenant(tenantId string)  revel.Result {
	err := services.DeleteTenant(tenantId)
	if err != nil {
		c.Response.SetStatus(400)
		return c.RenderJSON(err)
	}
	return c.Result
}