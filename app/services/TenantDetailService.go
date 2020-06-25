package services

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/jjamieson1/eden-frontend/models"
	"github.com/revel/revel"
)

func AddUpdateTenantDetails(tenantId string, tenant models.Tenant) (models.Tenant, error) {
	databaseString := revel.Config.StringDefault("connectionString", "root:root@tcp(localhost:3306)/eden_tenant?parseTime=true")

	db, err := sql.Open("mysql", databaseString)
	if err != nil {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return tenant, errors.New(error)
	}
	defer db.Close()


	if tenantId == "" {
		tenant.TenantId = uuid.New().String()

		query := `INSERT INTO tenant_details (
					tenant_id,
					parent_tenant_id,
 					url, 
 					common_name,
 					primary_logo_url,
 					secondary_logo_url,
 					mission, 
 					primary_phone, 
 					primary_email, 
 					street,
 					city,
 					state,
 					postal,
 					hours_monday,
 					hours_tuesday,
 					hours_wednesday,
 					hours_thursday,
 					hours_friday,
 					hours_saturday,
 					hours_sunday,
 					promotional,
 					tenant_type) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`

		revel.AppLog.Infof("adding new tenant  (tenantId: %v)", tenant.TenantId)

		stmt, err := db.Prepare(query)
		if err != nil {
			error := fmt.Sprintf("error performing query: %v, error: %v",query, err.Error())
			revel.AppLog.Errorf(error)
			return tenant, errors.New(error)
		}

		_, err = stmt.Exec(
			tenant.TenantId,
			tenant.ParentTenantId,
			tenant.Url,
			tenant.CommonName,
			tenant.LogoUrl,
			tenant.LogoSecondaryUrl,
			tenant.Mission,
			tenant.Phone,
			tenant.Email,
			tenant.Street,
			tenant.City,
			tenant.State,
			tenant.Postal,
			tenant.Hours.Monday,
			tenant.Hours.Tuesday,
			tenant.Hours.Wednesday,
			tenant.Hours.Thursday,
			tenant.Hours.Friday,
			tenant.Hours.Saturday,
			tenant.Hours.Sunday,
			tenant.Promotional,
			tenant.TenantType)
		if err != nil {
			error := fmt.Sprintf("error performing query: %v, error: %v",query, err.Error())
			revel.AppLog.Errorf(error)
			return tenant, errors.New(error)
		}

	} else {
		query :=  `UPDATE tenant_details SET 
						tenant_parent_id,
						url=?, 
						common_name=?, 
						primary_logo_url=?,
						secondary_logo_url=?,
						mission=?, 
						primary_phone=?,
						primary_email=?,
						street=?,
						city=?,
						state=?,
						postal=?,
						hours_monday=?,
						hours_tuesday=?,
						hours_wednesday=?,
						hours_thursday=?,
						hours_friday=?,
						hours_saturday=?,
						hours_sunday=?, 
						promotional=?,
						tenant_type=?
						WHERE tenant_id=?`

		revel.AppLog.Infof("updating CMS article cmsId: %v for tenant: %v", tenantId)

		stmt, err := db.Prepare(query)
		if err != nil {
			error := fmt.Sprintf("error performing query: %v, error: %v",query, err.Error())
			revel.AppLog.Errorf(error)
			return tenant, errors.New(error)
		}

		_, err = stmt.Exec(
			tenant.ParentTenantId,
			tenant.Url,
			tenant.CommonName,
			tenant.Url,
			tenant.LogoSecondaryUrl,
			tenant.Mission,
			tenant.Phone,
			tenant.Email,
			tenant.Street,
			tenant.City,
			tenant.State,
			tenant.Postal,
			tenant.Hours.Monday,
			tenant.Hours.Tuesday,
			tenant.Hours.Wednesday,
			tenant.Hours.Thursday,
			tenant.Hours.Friday,
			tenant.Hours.Saturday,
			tenant.Hours.Sunday,
			tenant.Promotional,
			tenant.TenantType,
			tenant.TenantId,
			)
		if err != nil && err != sql.ErrNoRows {
			error := fmt.Sprint( err.Error())
			revel.AppLog.Errorf(error)
			return tenant, errors.New(error)
		} else if err == sql.ErrNoRows {
			error := fmt.Sprint( err.Error())
			revel.AppLog.Info(error)
			return tenant, sql.ErrNoRows
		} else {
			revel.AppLog.Debugf("updated/added tenant_details for tenantId: %v ", tenantId)
		}
	}

	return tenant, err

}

func GetTenantDetails(tenantId string) (models.Tenant, error) {
	databaseString := revel.Config.StringDefault("connectionString", "root:root@tcp(localhost:3306)/eden_tenant?parseTime=true")
	var tenant models.Tenant
	revel.AppLog.Infof("retrieving tenant  (tenantId: %v)", tenantId)

	db, err := sql.Open("mysql", databaseString)
	if err != nil {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return tenant, errors.New(error)
	}
	defer db.Close()

	query := `SELECT * FROM  tenant_details WHERE tenant_id = ?`


	err = db.QueryRow(query, tenantId).Scan(
		&tenant.TenantId,
		&tenant.ParentTenantId,
		&tenant.Url,
		&tenant.CommonName,
		&tenant.LogoUrl,
		&tenant.LogoSecondaryUrl,
		&tenant.Mission,
		&tenant.Phone,
		&tenant.Email,
		&tenant.Street,
		&tenant.City,
		&tenant.State,
		&tenant.Postal,
		&tenant.Hours.Monday,
		&tenant.Hours.Tuesday,
		&tenant.Hours.Wednesday,
		&tenant.Hours.Thursday,
		&tenant.Hours.Friday,
		&tenant.Hours.Saturday,
		&tenant.Hours.Sunday,
		&tenant.Promotional,
		&tenant.TenantType,
		)


	if err != nil && err != sql.ErrNoRows {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return tenant, errors.New(error)
	} else if err == sql.ErrNoRows {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Info(error)
		return tenant, sql.ErrNoRows
	} else {
		revel.AppLog.Debugf("retrieved tenant_details for tenantId: %v ", tenantId)
	}


	return tenant, err

}

func GetAllTenantChildrenDetails(tenantId string) ([]models.Tenant, error) {
	databaseString := revel.Config.StringDefault("connectionString", "root:root@tcp(localhost:3306)/eden_tenant?parseTime=true")
	var tenant models.Tenant
	var tenants []models.Tenant

	db, err := sql.Open("mysql", databaseString)
	if err != nil {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return tenants, errors.New(error)
	}
	defer db.Close()

	query := `SELECT * FROM  tenant_details WHERE parent_tenant_id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		error := fmt.Sprintf("error preparing query: %v, error: %v", query, err.Error())
		revel.AppLog.Errorf(error)
		return tenants, errors.New(error)
	}

	results, err := stmt.Query(tenantId)
	if err != nil {
		error := fmt.Sprintf("error performing query: %v, error: %v", query, err.Error())
		revel.AppLog.Errorf(error)
		return tenants, errors.New(error)
	}
	for results.Next() {
		err := results.Scan(
			&tenant.TenantId,
			&tenant.ParentTenantId,
			&tenant.Url,
			&tenant.CommonName,
			&tenant.LogoUrl,
			&tenant.LogoSecondaryUrl,
			&tenant.Mission,
			&tenant.Phone,
			&tenant.Email,
			&tenant.Street,
			&tenant.City,
			&tenant.State,
			&tenant.Postal,
			&tenant.Hours.Monday,
			&tenant.Hours.Tuesday,
			&tenant.Hours.Wednesday,
			&tenant.Hours.Thursday,
			&tenant.Hours.Friday,
			&tenant.Hours.Saturday,
			&tenant.Hours.Sunday,
			&tenant.Promotional,
			&tenant.TenantType,
		)
		if err != nil {
			error := fmt.Sprintf("error mapping query to model: %v, error: %v", query, err.Error())
			revel.AppLog.Errorf(error)
			return tenants, errors.New(error)
		}
		tenants = append(tenants, tenant)
	}


return tenants, err


}

func GetTenantByUrl (url string) (models.Tenant, error) {

	databaseString := revel.Config.StringDefault("connectionString", "root:root@tcp(localhost:3306)/eden_tenant?parseTime=true")
	var tenant models.Tenant

	db, err := sql.Open("mysql", databaseString)
	if err != nil {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return tenant, errors.New(error)
	}
	defer db.Close()

	query := `SELECT * FROM  tenant_details WHERE url = ?`

	revel.AppLog.Infof("getting tenant  by url: %v", url)

	err = db.QueryRow(query, url).Scan(
		&tenant.TenantId,
		&tenant.Url,
		&tenant.CommonName,
		&tenant.LogoUrl,
		&tenant.LogoSecondaryUrl,
		&tenant.Mission,
		&tenant.Phone,
		&tenant.Email,
		&tenant.Street,
		&tenant.City,
		&tenant.State,
		&tenant.Postal,
		&tenant.Hours.Monday,
		&tenant.Hours.Tuesday,
		&tenant.Hours.Wednesday,
		&tenant.Hours.Thursday,
		&tenant.Hours.Friday,
		&tenant.Hours.Saturday,
		&tenant.Hours.Sunday,
		&tenant.Promotional,
	)


	if err != nil && err != sql.ErrNoRows {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return tenant, errors.New(error)
	} else if err == sql.ErrNoRows {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Info(error)
		return tenant, sql.ErrNoRows
	} else {
		revel.AppLog.Debugf("retrieved tenant_details for url: %v ", url)
	}


	return tenant, err

}

func DeleteTenant(tenantId string,)  error {
	databaseString := revel.Config.StringDefault("connectionString", "root:root@tcp(localhost:3306)/eden_tenant?parseTime=true")

	db, err := sql.Open("mysql", databaseString)
	if err != nil {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return  errors.New(error)
	}
	defer db.Close()

	query :=  `DELETE FROM tenant_details 
						WHERE tenant_id=?`

	revel.AppLog.Infof("updating CMS article cmsId: %v for tenant: %v", tenantId)

	stmt, err := db.Prepare(query)
	if err != nil {
		error := fmt.Sprintf("error performing query: %v, error: %v",query, err.Error())
		revel.AppLog.Errorf(error)
		return errors.New(error)
	}
	_, err = stmt.Exec(tenantId)

	if err != nil && err != sql.ErrNoRows {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return  errors.New(error)
	} else if err == sql.ErrNoRows {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Info(error)
		return sql.ErrNoRows
	} else {
		revel.AppLog.Debugf("deleted tenantId: %v ", tenantId)
	}

	return  err
}


// Tenant Type services

func AddTenantType(tenantType models.TenantType) (string, error) {
	databaseString := revel.Config.StringDefault("connectionString", "root:root@tcp(localhost:3306)/eden_tenant?parseTime=true")
	db, err := sql.Open("mysql", databaseString)
	if err != nil {
		error := fmt.Sprint( err.Error())
		revel.AppLog.Errorf(error)
		return "", errors.New(error)
	}
	defer db.Close()

	i := uuid.New().String()

		query := `INSERT INTO tenant_type (id, type_name, description, tenant_id) VALUES (?,?,?,?)`

		revel.AppLog.Infof("adding new tenant type: %v, for tenantId: %v", tenantType.Name, tenantType.Id)

		stmt, err := db.Prepare(query)
		if err != nil {
			error := fmt.Sprintf("error performing query: %v, error: %v",query, err.Error())
			revel.AppLog.Errorf(error)
			return  i, errors.New(error)
		}
		_, err = stmt.Exec(
			i,
			tenantType.Name,
			tenantType.Description,
			tenantType.TenantId,
			)
		if err != nil {
			error := fmt.Sprintf("error performing query: %v, error: %v",query, err.Error())
			revel.AppLog.Errorf(error)
			return  i, errors.New(error)
		}

return i, err
	}

func GetTenantType(tenantId string) ([]models.TenantType, error) {

	revel.AppLog.Debugf("getting tenant types for tenantId: %v", tenantId)

	databaseString := revel.Config.StringDefault("connectionString", "root:root@tcp(localhost:3306)/eden_tenant?parseTime=true")
	var tenantType models.TenantType
	var tenantTypes []models.TenantType
	var err error

		db, err := sql.Open("mysql", databaseString)
		if err != nil {
			error := fmt.Sprint(err.Error())
			revel.AppLog.Errorf(error)
			return tenantTypes, errors.New(error)
		}
		defer db.Close()

		query := `SELECT * FROM tenant_type WHERE tenant_id = ?`

		stmt, err := db.Prepare(query)
		if err != nil {
			error := fmt.Sprintf("error preparing query: %v, error: %v", query, err.Error())
			revel.AppLog.Errorf(error)
			return tenantTypes, errors.New(error)
		}

		results, err := stmt.Query(tenantId)
		if err != nil {
			error := fmt.Sprintf("error performing query: %v, error: %v", query, err.Error())
			revel.AppLog.Errorf(error)
			return tenantTypes, errors.New(error)
		}
		for results.Next() {
			err := results.Scan(
				&tenantType.Id,
				&tenantType.Name,
				&tenantType.Description,
				&tenantType.TenantId,
			)
			if err != nil {
				error := fmt.Sprintf("error mapping query to model: %v, error: %v", query, err.Error())
				revel.AppLog.Errorf(error)
				return tenantTypes, errors.New(error)
			}
			tenantTypes = append(tenantTypes, tenantType)
		}
	return tenantTypes, err
}