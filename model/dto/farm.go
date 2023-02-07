package dto

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

type (
	CreateFarmRequest struct {
		Code        string `json:"code" validate:"required"`
		Description string `json:"description" validate:"required"`
		Address     string `json:"address" validate:"required"`
		Village     string `json:"village" validate:"required"`
		District    string `json:"district" validate:"required"`
		City        string `json:"city" validate:"required"`
		Province    string `json:"province" validate:"required"`
		PostalCode  string `json:"postalCode" validate:"required"`

		Latitude  *float64 `json:"latitude" validate:"required_with=Longitude"`
		Longitude *float64 `json:"longitude" validate:"required_with=Latitude"`

		LoginData token.Data `validate:"dive,required"`
	}

	UpsertFarmRequest struct {
		Code string `param:"code" validate:"required"`

		Description string `json:"description" validate:"required"`
		Address     string `json:"address" validate:"required"`
		Village     string `json:"village" validate:"required"`
		District    string `json:"district" validate:"required"`
		City        string `json:"city" validate:"required"`
		Province    string `json:"province" validate:"required"`
		PostalCode  string `json:"postalCode" validate:"required"`

		Latitude  *float64 `json:"latitude" validate:"required_with=Longitude"`
		Longitude *float64 `json:"longitude" validate:"required_with=Latitude"`

		LoginData token.Data `validate:"dive,required"`
	}

	ListFarmFilterRequest struct {
		Code             string     `query:"code"`
		Village          string     `query:"village"`
		District         string     `query:"district"`
		City             string     `query:"city"`
		Province         string     `query:"province"`
		PostalCode       string     `query:"postalCode"`
		CreatedDateStart *time.Time `query:"createdDateStart"`
		CreatedDateEnd   *time.Time `query:"createdDateEnd"`
	}

	FarmPagingRequest struct {
		orm.PagingRequest
		Filter ListFarmFilterRequest
	}

	FarmResponse struct {
		ID          uint64 `json:"id"`
		Code        string `json:"code"`
		Description string `json:"description"`

		Address    string `json:"address"`
		Village    string `json:"village"`
		District   string `json:"district"`
		City       string `json:"city"`
		Province   string `json:"province"`
		PostalCode string `json:"postalCode"`

		Latitude  *float64 `json:"latitude"`
		Longitude *float64 `json:"longitude"`

		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		CreatedBy string    `json:"createdBy"`
		UpdatedBy string    `json:"updatedBy"`

		Ponds PondResponses `json:"ponds,omitempty"`
	}

	FarmResponses []FarmResponse

	FarmPagingResponse struct {
		orm.BasePagingResponse
		Items FarmResponses `json:"items"`
	}
)

func (req *CreateFarmRequest) FromEchoContext(ec echo.Context) error {
	if err := ec.Bind(req); err != nil {
		return err
	}

	req.LoginData = ec.Get(consts.XUserHeader).(token.Data)

	if err := ec.Validate(req); err != nil {
		return err
	}

	return nil
}

func (req CreateFarmRequest) ToFarm() dao.Farm {
	now := time.Now()
	return dao.Farm{
		Code:        req.Code,
		Description: req.Description,
		Address:     req.Address,
		Village:     req.Village,
		District:    req.District,
		City:        req.City,
		Province:    req.Province,
		PostalCode:  req.PostalCode,

		Latitude:  req.Latitude,
		Longitude: req.Longitude,

		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.LoginData.CreatedBy(),
		UpdatedBy: req.LoginData.CreatedBy(),
	}
}

func (req *UpsertFarmRequest) FromEchoContext(ec echo.Context) error {
	if err := ec.Bind(req); err != nil {
		return err
	}

	req.LoginData = ec.Get(consts.XUserHeader).(token.Data)

	if err := ec.Validate(req); err != nil {
		return err
	}

	return nil
}

func (req UpsertFarmRequest) ToFarm() dao.Farm {
	now := time.Now()
	return dao.Farm{
		Code:        req.Code,
		Description: req.Description,
		Address:     req.Address,
		Village:     req.Village,
		District:    req.District,
		City:        req.City,
		Province:    req.Province,
		PostalCode:  req.PostalCode,

		Latitude:  req.Latitude,
		Longitude: req.Longitude,

		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.LoginData.CreatedBy(),
		UpdatedBy: req.LoginData.CreatedBy(),
	}
}

func (req UpsertFarmRequest) ToUpdateMap() map[string]interface{} {
	return map[string]interface{}{
		"description": req.Description,
		"address":     req.Address,
		"village":     req.Village,
		"district":    req.District,
		"city":        req.City,
		"province":    req.Province,
		"postal_code": req.PostalCode,
		"latitude":    req.Latitude,
		"longitude":   req.Longitude,
		"updated_at":  time.Now(),
		"updated_by":  req.LoginData.CreatedBy(),
	}
}

func (filter ListFarmFilterRequest) Apply(o orm.Orm) orm.Orm {
	if filter.Code != "" {
		o = o.Where("farms.code LIKE ?", "%"+filter.Code+"%")
	}

	if filter.Village != "" {
		o = o.Where("farms.village LIKE ?", "%"+filter.Village+"%")
	}

	if filter.District != "" {
		o = o.Where("farms.district LIKE ?", "%"+filter.District+"%")
	}

	if filter.City != "" {
		o = o.Where("farms.city LIKE ?", "%"+filter.City+"%")
	}

	if filter.Province != "" {
		o = o.Where("farms.province LIKE ?", "%"+filter.Province+"%")
	}

	if filter.PostalCode != "" {
		o = o.Where("farms.postal_code LIKE ?", "%"+filter.PostalCode+"%")
	}

	if filter.CreatedDateStart != nil && filter.CreatedDateEnd != nil {
		o = o.Where("farms.created_at BETWEEN ? AND ?", *filter.CreatedDateStart, *filter.CreatedDateEnd)
	}

	return o
}

func (req *FarmPagingRequest) FromEchoContext(ec echo.Context) error {
	req.PagingRequest = orm.NewPagingRequest(ec, []string{"createdDate"})

	if err := ec.Bind(&req.Filter); err != nil {
		return err
	}

	return nil
}

func NewFarmResponse(data dao.Farm) FarmResponse {
	return FarmResponse{
		ID:          data.ID,
		Code:        data.Code,
		Description: data.Description,

		Address:    data.Address,
		Village:    data.Village,
		District:   data.District,
		City:       data.City,
		Province:   data.Province,
		PostalCode: data.PostalCode,

		Latitude:  data.Latitude,
		Longitude: data.Longitude,

		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		CreatedBy: data.CreatedBy,
		UpdatedBy: data.UpdatedBy,

		Ponds: NewPondResponses(data.Ponds),
	}
}

func NewFarmResponses(farms dao.Farms) FarmResponses {
	resps := make(FarmResponses, len(farms))
	for i, farm := range farms {
		resps[i] = NewFarmResponse(farm)
	}

	return resps
}

func NewFarmPagingResponse(farmPaging dao.FarmsPagingResponse) FarmPagingResponse {
	return FarmPagingResponse{
		BasePagingResponse: farmPaging.BasePagingResponse,
		Items:              NewFarmResponses(farmPaging.Items),
	}
}
