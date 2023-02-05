package dto

import (
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
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
)

func (req *CreateFarmRequest) FromEchoContext(ec echo.Context) error {
	if err := ec.Bind(req); err != nil {
		return err
	}

	req.LoginData = ec.Get("x-user").(token.Data)

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
