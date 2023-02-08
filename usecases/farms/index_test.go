package farmsUsecases

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	mock_endpointLogsRepositories "github.com/naufalfmm/aquafarm-management-service/mocks/persistents/repositories/endpointLogs"
	mock_endpointsRepositories "github.com/naufalfmm/aquafarm-management-service/mocks/persistents/repositories/endpoints"
	mock_farmsRepositories "github.com/naufalfmm/aquafarm-management-service/mocks/persistents/repositories/farms"
	mock_pondsRepositories "github.com/naufalfmm/aquafarm-management-service/mocks/persistents/repositories/ponds"
	mock_logger "github.com/naufalfmm/aquafarm-management-service/mocks/utils/logger"
	mock_orm "github.com/naufalfmm/aquafarm-management-service/mocks/utils/orm"
	"github.com/naufalfmm/aquafarm-management-service/persistents"
	"github.com/naufalfmm/aquafarm-management-service/persistents/repositories"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/resources/db"
	"github.com/naufalfmm/aquafarm-management-service/utils/frTime"
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

		farmRepo := mock_farmsRepositories.NewMockRepositories(ctrl)
		pondRepo := mock_pondsRepositories.NewMockRepositories(ctrl)
		endpointRepo := mock_endpointsRepositories.NewMockRepositories(ctrl)
		endpointLogRepo := mock_endpointLogsRepositories.NewMockRepositories(ctrl)

		persists := persistents.Persistents{
			Repositories: repositories.Repositories{
				Farms:        farmRepo,
				Ponds:        pondRepo,
				Endpoints:    endpointRepo,
				EndpointLogs: endpointLogRepo,
			},
		}

		expUsecases := usecases{
			persistents: persists,
			resources:   reso,
		}

		res, err := Init(persists, reso)

		assert.Nil(t, err)
		assert.Equal(t, &expUsecases, res)
	})
}

var (
	errAny = errors.New("any error")
)

type mock struct {
	ctrl *gomock.Controller
	ctx  context.Context

	orm    *mock_orm.MockOrm
	logger *mock_logger.MockLogger

	farmRepositories        *mock_farmsRepositories.MockRepositories
	pondRepositories        *mock_pondsRepositories.MockRepositories
	endpointRepositories    *mock_endpointsRepositories.MockRepositories
	endpointLogRepositories *mock_endpointLogsRepositories.MockRepositories

	usecases usecases

	now time.Time
}

func setupMock(t *testing.T) mock {
	mock := mock{}
	mock.ctrl = gomock.NewController(t)

	mock.orm = mock_orm.NewMockOrm(mock.ctrl)
	mock.logger = mock_logger.NewMockLogger(mock.ctrl)

	mock.farmRepositories = mock_farmsRepositories.NewMockRepositories(mock.ctrl)
	mock.pondRepositories = mock_pondsRepositories.NewMockRepositories(mock.ctrl)
	mock.endpointRepositories = mock_endpointsRepositories.NewMockRepositories(mock.ctrl)
	mock.endpointLogRepositories = mock_endpointLogsRepositories.NewMockRepositories(mock.ctrl)

	mock.usecases = usecases{
		persistents: persistents.Persistents{
			Repositories: repositories.Repositories{
				Farms:        mock.farmRepositories,
				Ponds:        mock.pondRepositories,
				Endpoints:    mock.endpointRepositories,
				EndpointLogs: mock.endpointLogRepositories,
			},
		},
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
	frTime.Mock(mock.now)

	return mock
}

func (m mock) Finish() {
	m.ctrl.Finish()
	frTime.ResetMock()
}
