package dto

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
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
)

func (req *CreatePondRequest) FromEchoContext(ec echo.Context) error {
	if err := ec.Bind(req); err != nil {
		return err
	}

	req.LoginData = ec.Get("x-user").(token.Data)

	if err := ec.Validate(req); err != nil {
		return err
	}

	return nil
}

func (req CreatePondRequest) ToPond() dao.Pond {
	now := time.Now()
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
