package services

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/jjamieson1/eden-frontend/models"
	"github.com/revel/revel"
	"github.com/revel/revel/cache"
	"time"
)

func GetProvidersForTenantByType(tenantId, providerType string) ([]models.TenantProvider, error) {
	databaseString := revel.Config.StringDefault("connectionString", "root:root@tcp(localhost:3306)/eden_tenant?parseTime=true")
	var tenantProvider models.TenantProvider
	var tenantProviders []models.TenantProvider
	var err error

	if err := cache.Get("tenant-" + providerType + "-" + tenantId, &tenantProviders); err != nil {
		revel.AppLog.Infof("cache miss, getting new data, error: %v", err.Error())

		db, err := sql.Open("mysql", databaseString)
		if err != nil {
			error := fmt.Sprint(err.Error())
			revel.AppLog.Errorf(error)
			return tenantProviders, errors.New(error)
		}
		defer db.Close()

		query := `SELECT tenant_provider.*, eden_adapter.*, eden_provider_type.*, auth_strategy.*
				FROM tenant_provider
				    JOIN eden_adapter ON
				        tenant_provider.eden_adapter_id = eden_adapter.id
                    JOIN eden_provider_type ON
                        eden_adapter.eden_provider_type_id = eden_provider_type.id
                    JOIN auth_strategy ON
                        eden_adapter.auth_strategy_id = auth_strategy.id
					WHERE 
						tenant_provider.tenant_id = ? 
					AND eden_provider_type.eden_provider_type_name = ?
`

		revel.AppLog.Infof("getting provider information for tenantId: %v for provider_type: %v)", tenantId, providerType)

		stmt, err := db.Prepare(query)
		if err != nil {
			error := fmt.Sprintf("error preparing query: %v, error: %v", query, err.Error())
			revel.AppLog.Errorf(error)
			return tenantProviders, errors.New(error)
		}

		results, err := stmt.Query(tenantId, providerType)
		if err != nil {
			error := fmt.Sprintf("error performing query: %v, error: %v", query, err.Error())
			revel.AppLog.Errorf(error)
			return tenantProviders, errors.New(error)
		}
		for results.Next() {
			err := results.Scan(
				&tenantProvider.Id,
				&tenantProvider.EdenAdapter.Id,
				&tenantProvider.TenantId,
				&tenantProvider.CalloutUrl,
				&tenantProvider.UserName,
				&tenantProvider.Password,
				&tenantProvider.ApiKey,
				&tenantProvider.AppKey,
				&tenantProvider.Token,
				&tenantProvider.RefreshToken,
				&tenantProvider.EdenAdapter.Id,
				&tenantProvider.EdenAdapter.Name,
				&tenantProvider.EdenAdapter.PluginName,
				&tenantProvider.EdenAdapter.ProviderType.Id,
				&tenantProvider.EdenAdapter.AuthStrategy.Id,
				&tenantProvider.EdenAdapter.AdapterUrl,
				&tenantProvider.EdenAdapter.ProviderType.Id,
				&tenantProvider.EdenAdapter.ProviderType.Name,
				&tenantProvider.EdenAdapter.AuthStrategy.Id,
				&tenantProvider.EdenAdapter.AuthStrategy.Name,
				&tenantProvider.EdenAdapter.AuthStrategy.Parameters,
			)
			if err != nil {
				error := fmt.Sprintf("error mapping query to model: %v, error: %v", query, err.Error())
				revel.AppLog.Errorf(error)
				return tenantProviders, errors.New(error)
			}
			tenantProviders = append(tenantProviders, tenantProvider)
		}
		go cache.Set("tenant-" + tenantId, tenantProviders, 30*time.Minute)
	} else {
		revel.AppLog.Debugf("cache hit, returning tenantId: %v", tenantId)
	}
	return tenantProviders, err
}

func GetAllProvidersForTenant(tenantId string) ([]models.TenantProvider, error) {
	databaseString := revel.Config.StringDefault("connectionString", "root:root@tcp(localhost:3306)/eden_tenant?parseTime=true")
	var tenantProvider models.TenantProvider
	var tenantProviders []models.TenantProvider

	db, err := sql.Open("mysql", databaseString)
	if err != nil {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return tenantProviders, errors.New(error)
	}
	defer db.Close()

	query := `SELECT tenant_provider.*, eden_adapter.*, eden_provider_type.*, auth_strategy.*
				FROM tenant_provider
				    JOIN eden_adapter ON
				        tenant_provider.eden_adapter_id = eden_adapter.id
                    JOIN eden_provider_type ON
                        eden_adapter.eden_provider_type_id = eden_provider_type.id
                    JOIN auth_strategy ON
                        eden_adapter.auth_strategy_id = auth_strategy.id
					WHERE 
						tenant_provider.tenant_id = ?
`

	revel.AppLog.Infof("getting provider information for tenantId: %v for provider_type: %v)", tenantId)

	stmt, err := db.Prepare(query)
	if err != nil {
		error := fmt.Sprintf("error preparing query: %v, error: %v",query, err.Error())
		revel.AppLog.Errorf(error)
		return tenantProviders, errors.New(error)
	}

	results, err := stmt.Query(tenantId)
	if err != nil {
		error := fmt.Sprintf("error performing query: %v, error: %v",query, err.Error())
		revel.AppLog.Errorf(error)
		return tenantProviders, errors.New(error)
	}
	for results.Next() {
		err := results.Scan(
			&tenantProvider.Id,
			&tenantProvider.EdenAdapter.Id,
			&tenantProvider.TenantId,
			&tenantProvider.CalloutUrl,
			&tenantProvider.UserName,
			&tenantProvider.Password,
			&tenantProvider.ApiKey,
			&tenantProvider.AppKey,
			&tenantProvider.Token,
			&tenantProvider.RefreshToken,
			&tenantProvider.EdenAdapter.Id,
			&tenantProvider.EdenAdapter.Name,
			&tenantProvider.EdenAdapter.PluginName,
			&tenantProvider.EdenAdapter.ProviderType.Id,
			&tenantProvider.EdenAdapter.AuthStrategy.Id,
			&tenantProvider.EdenAdapter.AdapterUrl,
			&tenantProvider.EdenAdapter.ProviderType.Id,
			&tenantProvider.EdenAdapter.ProviderType.Name,
			&tenantProvider.EdenAdapter.AuthStrategy.Id,
			&tenantProvider.EdenAdapter.AuthStrategy.Name,
			&tenantProvider.EdenAdapter.AuthStrategy.Parameters,
		)
		if err != nil {
			error := fmt.Sprintf("error mapping query to model: %v, error: %v" ,query, err.Error())
			revel.AppLog.Errorf(error)
			return tenantProviders, errors.New(error)
		}
		tenantProviders = append(tenantProviders, tenantProvider)
	}
	return tenantProviders, err

}