package pondsRepositories

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	mock_logger "github.com/naufalfmm/aquafarm-management-service/mocks/utils/logger"
	mock_orm "github.com/naufalfmm/aquafarm-management-service/mocks/utils/orm"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/resources/db"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Run("If no error, it will return the repositories", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orm := mock_orm.NewMockOrm(ctrl)
		logger := mock_logger.NewMockLogger(ctrl)

		reso := resources.Resources{
			MySql: &db.DB{
				Orm: orm,
			},
			Logger: logger,
		}

		expRepo := repositories{
			resources: reso,
		}

		res, err := Init(reso)

		assert.Nil(t, err)
		assert.Equal(t, &expRepo, res)
	})
}

var (
	errAny = errors.New("any error")
)

type mock struct {
	ctrl         *gomock.Controller
	ctx          context.Context
	orm          *mock_orm.MockOrm
	logger       *mock_logger.MockLogger
	repositories repositories
	now          time.Time
}

func setupMock(t *testing.T) mock {
	mock := mock{}
	mock.ctrl = gomock.NewController(t)

	mock.orm = mock_orm.NewMockOrm(mock.ctrl)
	mock.logger = mock_logger.NewMockLogger(mock.ctrl)

	mock.repositories = repositories{
		resources: resources.Resources{
			MySql: &db.DB{
				Orm: mock.orm,
			},
			Logger: mock.logger,
		},
	}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	ecCtx := echo.New().NewContext(req, httptest.NewRecorder())
	mock.ctx = ecCtx.Request().Context()

	mock.now = time.Now()
	timeNow = func() time.Time {
		return mock.now
	}

	return mock
}

func (m mock) Before() {
	m.orm.EXPECT().WithContext(m.ctx).Return(m.orm)
}

func (m mock) Finish() {
	m.ctrl.Finish()

	timeNow = time.Now
}
