package controllers

import (
	"errors"
	"github.com/jjamieson1/eden-frontend/models"
	"github.com/jjamieson1/eden-tenant-service/app/services"
	"github.com/revel/revel"
	"strings"
)

type Api struct {
	*revel.Controller
}

func (c Api) SetTenantUserServiceProvider(tenantId string, provider string) revel.Result  {
	c.Response.SetStatus(200)
	return c.Result
}

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

func (c Api) GetAllChildrenOfTenant(tenantId string) revel.Result {
	revel.AppLog.Debugf("getting all  tenant details for children of tenantId: %v", tenantId)
	tenants, _ := services.GetAllTenantChildrenDetails(tenantId)

	return c.RenderJSON(tenants)
}


func (c Api) GetTenantByUrl(url string) revel.Result {

	revel.AppLog.Debugf("requesting tenant details for url: %v", url)

	// In case a port number is appended remove
	u := strings.Split(url, ":")

	tenant, err := services.GetTenantByUrl(u[0])
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


func (c Api) GetTenantType() revel.Result {
tenantId := c.Request.Header.Get("tenantId")
	tenantTypes, err := services.GetTenantType(tenantId)
	if err != nil {
		c.Response.SetStatus(400)
		return c.RenderJSON(err)
	}
	return c.RenderJSON(tenantTypes)
}

func (c Api) AddTenantType() revel.Result {
	var tenantType models.TenantType
	err := c.Params.BindJSON(&tenantType)
	if err != nil {
		return c.RenderJSON(err)
	}
	i, err := services.AddTenantType(tenantType)
	if err != nil {
		c.Response.SetStatus(400)
		return c.RenderJSON(err)
	}
	tenantType.Id = i
	return c.RenderJSON(tenantType)
}

func (c Api) DeleteTenantType()  revel.Result {
	return c.RenderError(errors.New("not implemented"))
}

func (c Api) UpdateTenantType()  revel.Result {
	return c.RenderError(errors.New("not implemented"))

}