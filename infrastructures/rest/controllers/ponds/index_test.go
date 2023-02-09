package pondsControllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	mock_endpointLogsUsecases "github.com/naufalfmm/aquafarm-management-service/mocks/usecases/endpointLogs"
	mock_endpointsUsecases "github.com/naufalfmm/aquafarm-management-service/mocks/usecases/endpoints"
	mock_farmsUsecases "github.com/naufalfmm/aquafarm-management-service/mocks/usecases/farms"
	mock_pondsUsecases "github.com/naufalfmm/aquafarm-management-service/mocks/usecases/ponds"
	mock_logger "github.com/naufalfmm/aquafarm-management-service/mocks/utils/logger"
	mock_orm "github.com/naufalfmm/aquafarm-management-service/mocks/utils/orm"
	mock_validator "github.com/naufalfmm/aquafarm-management-service/mocks/utils/validator"
	"github.com/naufalfmm/aquafarm-management-service/resources"
	"github.com/naufalfmm/aquafarm-management-service/resources/db"
	"github.com/naufalfmm/aquafarm-management-service/usecases"
	"github.com/naufalfmm/aquafarm-management-service/utils/frTime"
	"github.com/naufalfmm/aquafarm-management-service/utils/token/jwt"
	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	t.Run("If no error, it will return the controllers", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		orm := mock_orm.NewMockOrm(ctrl)
		logger := mock_logger.NewMockLogger(ctrl)
		validator := mock_validator.NewMockValidator(ctrl)

		reso := resources.Resources{
			MySql: &db.DB{
				Orm: orm,
			},
			Logger:    logger,
			Validator: validator,
		}

		farmUsec := mock_farmsUsecases.NewMockUsecases(ctrl)
		pondUsec := mock_pondsUsecases.NewMockUsecases(ctrl)
		endpointUsec := mock_endpointsUsecases.NewMockUsecases(ctrl)
		endpointLogUsec := mock_endpointLogsUsecases.NewMockUsecases(ctrl)

		usecs := usecases.Usecases{
			Farms:        farmUsec,
			Ponds:        pondUsec,
			Endpoints:    endpointUsec,
			EndpointLogs: endpointLogUsec,
		}

		res, err := Init(usecs, reso)

		assert.Nil(t, err)
		assert.Equal(t, Controllers{Usecases: usecs, Resources: reso}, res)
	})
}

var (
	errAny = errors.New("any error")
)

type mock struct {
	ctrl        *gomock.Controller
	ctx         context.Context
	eCtx        echo.Context
	resRecorder *httptest.ResponseRecorder

	orm       *mock_orm.MockOrm
	logger    *mock_logger.MockLogger
	validator *mock_validator.MockValidator

	farmUsecases        *mock_farmsUsecases.MockUsecases
	pondUsecases        *mock_pondsUsecases.MockUsecases
	endpointUsecases    *mock_endpointsUsecases.MockUsecases
	endpointLogUsecases *mock_endpointLogsUsecases.MockUsecases

	controllers Controllers

	now       time.Time
	loginData *jwt.UserLogin
}

func setupMock(t *testing.T) mock {
	mock := mock{}
	mock.ctrl = gomock.NewController(t)

	mock.orm = mock_orm.NewMockOrm(mock.ctrl)
	mock.logger = mock_logger.NewMockLogger(mock.ctrl)
	mock.validator = mock_validator.NewMockValidator(mock.ctrl)

	mock.farmUsecases = mock_farmsUsecases.NewMockUsecases(mock.ctrl)
	mock.pondUsecases = mock_pondsUsecases.NewMockUsecases(mock.ctrl)
	mock.endpointUsecases = mock_endpointsUsecases.NewMockUsecases(mock.ctrl)
	mock.endpointLogUsecases = mock_endpointLogsUsecases.NewMockUsecases(mock.ctrl)

	mock.controllers = Controllers{
		Usecases: usecases.Usecases{
			Farms:        mock.farmUsecases,
			Ponds:        mock.pondUsecases,
			Endpoints:    mock.endpointUsecases,
			EndpointLogs: mock.endpointLogUsecases,
		},
		Resources: resources.Resources{
			MySql: &db.DB{
				Orm: mock.orm,
			},
			Logger:    mock.logger,
			Validator: mock.validator,
		},
	}

	mock.loginData = &jwt.UserLogin{Email: "engineer.test@test.com"}

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	mock.resRecorder = httptest.NewRecorder()

	mock.eCtx = echo.New().NewContext(req, mock.resRecorder)
	mock.eCtx.Set(consts.XUserHeader, mock.loginData)

	mock.eCtx.Echo().Validator = mock.validator
	mock.ctx = mock.eCtx.Request().Context()

	mock.now = time.Now()
	frTime.Mock(mock.now)

	return mock
}

func (m mock) Finish() {
	m.ctrl.Finish()
	frTime.ResetMock()
}

func (m mock) SetRequestBody(body interface{}) {
	requestByte, _ := json.Marshal(body)

	req := httptest.NewRequest(http.MethodGet, "/", strings.NewReader(string(requestByte)))

	m.eCtx.SetRequest(req)
	m.eCtx.Request().Header.Set(consts.ContentTypeHeader, consts.ApplicationJsonContentType)
}

func (m mock) SetParam(key, val string) {
	m.eCtx.SetParamNames(key)
	m.eCtx.SetParamValues(val)
}

func (m mock) MakeSuccessResponse(message string, expectedContent interface{}) string {
	expectedResponse := struct {
		Ok      bool        `json:"ok"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
	}{Ok: true, Message: message, Data: expectedContent}

	return m.makeDataResponse(expectedResponse)
}

func (m mock) MakeErrorResponse(message string, err error) string {
	expectedResponse := struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
		Error   string `json:"error"`
	}{Ok: false, Message: message, Error: err.Error()}

	return m.makeDataResponse(expectedResponse)
}

func (m mock) makeDataResponse(expectedContent interface{}) string {
	expectedByte, _ := json.Marshal(expectedContent)
	expectedResult := string(expectedByte) + "\n"
	return expectedResult
}
