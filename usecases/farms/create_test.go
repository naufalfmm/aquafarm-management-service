package farmsUsecases

import (
	"testing"

	"github.com/naufalfmm/aquafarm-management-service/model/dao"
	"github.com/naufalfmm/aquafarm-management-service/model/dto"
	"github.com/naufalfmm/aquafarm-management-service/utils/token/jwt"
	"github.com/stretchr/testify/assert"
)

func Test_usecase_Create(t *testing.T) {
	var (
		req = dto.CreateFarmRequest{
			Code: "FA-1",

			LoginData: &jwt.UserLogin{
				Email: "engineer.test@test.com",
			},
		}

		farm = dao.Farm{
			ID:   1,
			Code: req.Code,
		}
	)

	t.Run("If no error, it will return the data", func(t *testing.T) {
		mock := setupMock(t)
		defer mock.Finish()

		mock.farmRepositories.EXPECT().Create(mock.ctx, req.ToFarm()).Return(farm, nil)

		res, err := mock.usecases.Create(mock.ctx, req)

		assert.Nil(t, err)
		assert.Equal(t, farm, res)
	})
}
