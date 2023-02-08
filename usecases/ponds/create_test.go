package pondsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/token/jwt"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_Create(t *testing.T) {
	var (
		farmID   uint64 = 1
		farmCode string = "FA-1"

		farm dao.Farm = dao.Farm{
			ID:   farmID,
			Code: farmCode,
		}

		userLogin = jwt.UserLogin{
			Email: "engineer.test@test.com",
		}
	)

	t.Run("If create ponds with farm id, it will return the created pond", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.CreatePondRequest{
			FarmID: farmID,

			Code: "A1",

			LoginData: &userLogin,
		}

		pond := dao.Pond{
			ID:     1,
			FarmID: farmID,
			Code:   req.Code,
		}

		mock.farmRepositories.EXPECT().GetByID(mock.ctx, farmID).Return(farm, nil)
		mock.pondRepositories.EXPECT().Create(mock.ctx, req.ToPond()).Return(pond, nil)

		res, err := mock.usecases.Create(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, pond, res)
	})

	t.Run("If create ponds with farm code, it will return the created pond", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.CreatePondRequest{
			FarmCode: farmCode,

			Code: "A1",

			LoginData: &userLogin,
		}

		updatedReq := dto.CreatePondRequest{
			FarmID:   farmID,
			FarmCode: farmCode,

			Code: "A1",

			LoginData: &userLogin,
		}

		pond := dao.Pond{
			ID:     1,
			FarmID: farmID,
			Code:   req.Code,
		}

		mock.farmRepositories.EXPECT().GetByCode(mock.ctx, farmCode).Return(farm, nil)
		mock.pondRepositories.EXPECT().Create(mock.ctx, updatedReq.ToPond()).Return(pond, nil)

		res, err := mock.usecases.Create(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, pond, res)
	})

	t.Run("If get farm by id error, it will return error", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		req := dto.CreatePondRequest{
			FarmID: farmID,

			Code: "A1",

			LoginData: &userLogin,
		}

		mock.farmRepositories.EXPECT().GetByID(mock.ctx, farmID).Return(dao.Farm{}, errAny)

		res, err := mock.usecases.Create(mock.ctx, req)

		assert.Equal(t, errAny, err)
		assert.Equal(t, dao.Pond{}, res)
	})
}
