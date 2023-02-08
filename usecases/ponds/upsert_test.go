package pondsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/token/jwt"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_Upsert(t *testing.T) {
	var (
		req = dto.UpsertPondRequest{
			Code:     "A1",
			FarmID:   1,
			FarmCode: "FA-1",

			LoginData: &jwt.UserLogin{
				Email: "engineer.test@test.com",
			},
		}

		ponds = dao.Ponds{
			{
				ID:     1,
				Code:   "A1",
				FarmID: 1,
			},
		}

		farm = dao.Farm{
			ID: 1,
		}
	)

	t.Run("If the data is not exist, it will create the data and return it", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.pondRepositories.EXPECT().GetAll(mock.ctx, dto.PondListRequest{
			ListPondFilterRequest: dto.ListPondFilterRequest{
				Code:     req.Code,
				FarmID:   req.FarmID,
				FarmCode: req.FarmCode,
			},
		}).Return(nil, nil)

		mock.farmRepositories.EXPECT().GetByID(mock.ctx, req.FarmID).Return(farm, nil)
		mock.pondRepositories.EXPECT().Create(mock.ctx, req.ToCreatePondRequest().ToPond()).Return(ponds[0], nil)

		res, err := mock.usecases.Upsert(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, ponds[0], res)
	})

	t.Run("If the data is exist, it will update the data and return it", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.pondRepositories.EXPECT().GetAll(mock.ctx, dto.PondListRequest{
			ListPondFilterRequest: dto.ListPondFilterRequest{
				Code:     req.Code,
				FarmID:   req.FarmID,
				FarmCode: req.FarmCode,
			},
		}).Return(ponds, nil)

		mock.pondRepositories.EXPECT().UpdateByID(mock.ctx, ponds[0].ID, req.ToUpdateMap()).Return(nil)
		mock.pondRepositories.EXPECT().GetByID(mock.ctx, ponds[0].ID).Return(ponds[0], nil)

		res, err := mock.usecases.Upsert(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, ponds[0], res)
	})

	t.Run("If update the data by id error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.pondRepositories.EXPECT().GetAll(mock.ctx, dto.PondListRequest{
			ListPondFilterRequest: dto.ListPondFilterRequest{
				Code:     req.Code,
				FarmID:   req.FarmID,
				FarmCode: req.FarmCode,
			},
		}).Return(ponds, nil)

		mock.pondRepositories.EXPECT().UpdateByID(mock.ctx, ponds[0].ID, req.ToUpdateMap()).Return(errAny)

		res, err := mock.usecases.Upsert(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Pond{}, res)
	})

	t.Run("If get all endpoints error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.pondRepositories.EXPECT().GetAll(mock.ctx, dto.PondListRequest{
			ListPondFilterRequest: dto.ListPondFilterRequest{
				Code:     req.Code,
				FarmID:   req.FarmID,
				FarmCode: req.FarmCode,
			},
		}).Return(nil, errAny)

		res, err := mock.usecases.Upsert(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Pond{}, res)
	})
}
