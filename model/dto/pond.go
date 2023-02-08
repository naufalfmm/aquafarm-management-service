package dto

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/utils/frTime"
	"github.com/naufalfmm/aquafarm-management-service/utils/orm"
	"github.com/naufalfmm/aquafarm-management-service/utils/token"
)

type (
	CreatePondRequest struct {
		FarmID   uint64 `json:"farmId" param:"id" validate:"required_without=FarmCode"`
		FarmCode string `json:"farmCode" validate:"required_without=FarmID"`

		Code        string `json:"code" validate:"required"`
		Description string `json:"description" validate:"required"`

		Wide  float64 `json:"wide" validate:"required"`
		Long  float64 `json:"long" validate:"required"`
		Depth float64 `json:"depth" validate:"required"`

		LoginData token.Data `validate:"dive,required"`
	}

	UpsertPondRequest struct {
		FarmID   uint64 `json:"farmId" param:"id" validate:"required_without=FarmCode"`
		FarmCode string `json:"farmCode" validate:"required_without=FarmID"`

		Code        string `param:"code" validate:"required"`
		Description string `json:"description" validate:"required"`

		Wide  float64 `json:"wide" validate:"required"`
		Long  float64 `json:"long" validate:"required"`
		Depth float64 `json:"depth" validate:"required"`

		LoginData token.Data `validate:"dive,required"`
	}

	ListPondFilterRequest struct {
		Code             string     `query:"code"`
		VolumeStart      float64    `query:"volumeStart"`
		VolumeEnd        float64    `query:"volumeEnd"`
		AreaStart        float64    `query:"areaStart"`
		AreaEnd          float64    `query:"areaEnd"`
		CreatedDateStart *time.Time `query:"createdDateStart"`
		CreatedDateEnd   *time.Time `query:"createdDateEnd"`

		FarmID   uint64
		FarmCode string
	}

	PondPagingRequest struct {
		orm.PagingRequest
		Filter ListPondFilterRequest
	}

	PondListRequest struct {
		ListPondFilterRequest
	}

	PondResponse struct {
		ID          uint64 `json:"id"`
		FarmID      uint64 `json:"farmId"`
		Code        string `json:"code"`
		Description string `json:"description"`

		Wide  float64 `json:"wide"`
		Long  float64 `json:"long"`
		Depth float64 `json:"depth"`

		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
		CreatedBy string    `json:"createdBy"`
		UpdatedBy string    `json:"updatedBy"`

		Farm *FarmResponse `json:"farm,omitempty"`
	}

	PondResponses []PondResponse

	PondPagingResponse struct {
		orm.BasePagingResponse
		Items PondResponses `json:"items"`
	}
)

func (req *CreatePondRequest) FromEchoContext(ec echo.Context) error {
	if err := ec.Bind(req); err != nil {
		return err
	}

	req.LoginData = ec.Get(consts.XUserHeader).(token.Data)

	if err := ec.Validate(req); err != nil {
		return err
	}

	return nil
}

func (req CreatePondRequest) ToPond() dao.Pond {
	now := frTime.Now()
	return dao.Pond{
		FarmID:      req.FarmID,
		Code:        req.Code,
		Description: req.Description,

		Wide:  req.Wide,
		Long:  req.Long,
		Depth: req.Depth,

		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.LoginData.CreatedBy(),
		UpdatedBy: req.LoginData.CreatedBy(),
	}
}

func (req *UpsertPondRequest) FromEchoContext(ec echo.Context) error {
	if err := ec.Bind(req); err != nil {
		return err
	}

	req.LoginData = ec.Get(consts.XUserHeader).(token.Data)

	if err := ec.Validate(req); err != nil {
		return err
	}

	return nil
}

func (req UpsertPondRequest) ToPond() dao.Pond {
	now := time.Now()
	return dao.Pond{
		Code:        req.Code,
		Description: req.Description,

		Wide:  req.Wide,
		Long:  req.Long,
		Depth: req.Depth,

		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: req.LoginData.CreatedBy(),
		UpdatedBy: req.LoginData.CreatedBy(),
	}
}

func (req UpsertPondRequest) ToUpdateMap() map[string]interface{} {
	return map[string]interface{}{
		"description": req.Description,
		"wide":        req.Wide,
		"long":        req.Long,
		"depth":       req.Depth,
		"updated_at":  frTime.Now(),
		"updated_by":  req.LoginData.CreatedBy(),
	}
}

func (req UpsertPondRequest) ToCreatePondRequest() CreatePondRequest {
	return CreatePondRequest{
		FarmID:      req.FarmID,
		FarmCode:    req.FarmCode,
		Code:        req.Code,
		Description: req.Description,
		Wide:        req.Wide,
		Long:        req.Long,
		Depth:       req.Depth,
		LoginData:   req.LoginData,
	}
}

func (filter ListPondFilterRequest) Apply(o orm.Orm) orm.Orm {
	if filter.Code != "" {
		o = o.Where("ponds.code LIKE ?", "%"+filter.Code+"%")
	}

	if filter.VolumeStart != 0 && filter.VolumeEnd == 0 {
		o = o.Where("(ponds.wide * ponds.long) >= ?", filter.VolumeStart)
	}

	if filter.VolumeStart == 0 && filter.VolumeEnd != 0 {
		o = o.Where("(ponds.wide * ponds.long) <= ?", filter.VolumeEnd)
	}

	if filter.VolumeStart != 0 && filter.VolumeEnd != 0 {
		o = o.Where("(ponds.wide * ponds.long) BETWEEN ? AND ?", filter.VolumeStart, filter.VolumeEnd)
	}

	if filter.AreaStart != 0 && filter.AreaEnd == 0 {
		o = o.Where("(ponds.wide * ponds.long * ponds.depth) >= ?", filter.AreaStart)
	}

	if filter.AreaStart == 0 && filter.AreaEnd != 0 {
		o = o.Where("(ponds.wide * ponds.long * ponds.depth) <= ?", filter.AreaEnd)
	}

	if filter.AreaStart != 0 && filter.AreaEnd != 0 {
		o = o.Where("(ponds.wide * ponds.long * ponds.depth) BETWEEN ? AND ?", filter.AreaStart, filter.AreaEnd)
	}

	if filter.CreatedDateStart != nil && filter.CreatedDateEnd != nil {
		o = o.Where("ponds.created_at BETWEEN ? AND ?", *filter.CreatedDateStart, *filter.CreatedDateEnd)
	}

	if filter.FarmID != 0 {
		o = o.Where("ponds.farm_id", filter.FarmID)
	}

	if filter.FarmCode != "" {
		o = o.Joins("farms").Where("farms.code", filter.FarmCode)
	}

	return o
}

func (req *PondPagingRequest) FromEchoContext(ec echo.Context) error {
	req.PagingRequest = orm.NewPagingRequest(ec, []string{"createdDate"})

	if err := ec.Bind(&req.Filter); err != nil {
		return err
	}

	return nil
}

func NewPondResponse(pond dao.Pond) PondResponse {
	resp := PondResponse{
		ID:          pond.ID,
		FarmID:      pond.FarmID,
		Code:        pond.Code,
		Description: pond.Description,

		Wide:  pond.Wide,
		Long:  pond.Long,
		Depth: pond.Depth,

		CreatedAt: pond.CreatedAt,
		UpdatedAt: pond.UpdatedAt,
		CreatedBy: pond.CreatedBy,
		UpdatedBy: pond.UpdatedBy,
	}

	if pond.Farm.ID != 0 {
		farm := NewFarmResponse(pond.Farm)
		resp.Farm = &farm
	}

	return resp
}

func NewPondResponses(ponds dao.Ponds) PondResponses {
	resps := make(PondResponses, len(ponds))
	for i, pond := range ponds {
		resps[i] = NewPondResponse(pond)
	}

	return resps
}

func NewPondPagingResponse(pondPaging dao.PondsPagingResponse) PondPagingResponse {
	return PondPagingResponse{
		BasePagingResponse: pondPaging.BasePagingResponse,
		Items:              NewPondResponses(pondPaging.Items),
	}
}
