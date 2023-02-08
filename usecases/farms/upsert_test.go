package farmsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/token/jwt"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func Test_usecase_Upsert(t *testing.T) {
	var (
		req = dto.UpsertFarmRequest{
			Code: "A1",
			LoginData: &jwt.UserLogin{
				Email: "engineer.test@test.com",
			},
		}

		farm = dao.Farm{
			ID:   1,
			Code: req.Code,
		}
	)

	t.Run("If the data is exist, it will update the data and return it", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetByCode(mock.ctx, req.Code).Return(farm, nil)
		mock.farmRepositories.EXPECT().UpdateByID(mock.ctx, farm.ID, req.ToUpdateMap()).Return(nil)
		mock.farmRepositories.EXPECT().GetByID(mock.ctx, farm.ID).Return(farm, nil)

		res, err := mock.usecases.Upsert(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, farm, res)
	})

	t.Run("If the data is not exist, it will create the data and return it", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetByCode(mock.ctx, req.Code).Return(dao.Farm{}, gorm.ErrRecordNotFound)
		mock.farmRepositories.EXPECT().Create(mock.ctx, req.ToFarm()).Return(farm, nil)

		res, err := mock.usecases.Upsert(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, farm, res)
	})

	t.Run("If update farm by id error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetByCode(mock.ctx, req.Code).Return(farm, nil)
		mock.farmRepositories.EXPECT().UpdateByID(mock.ctx, farm.ID, req.ToUpdateMap()).Return(errAny)

		res, err := mock.usecases.Upsert(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Farm{}, res)
	})

	t.Run("If get farm by code for checking existence error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().GetByCode(mock.ctx, req.Code).Return(dao.Farm{}, errAny)

		res, err := mock.usecases.Upsert(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Farm{}, res)
	})
}
