package dto

import (
	"net/http"
	"strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/naufalfmm/aquafarm-management-service/consts"
	"github.com/naufalfmm/aquafarm-management-service/model/dao"
)

type (
	UpsertEndpointRequest struct {
		Method string
		Path   string
	}

	UpsertEndpointRequests []UpsertEndpointRequest

	BulkUpsertEndpointsRequest struct {
		Endpoints UpsertEndpointRequests
	}
)

func isValidMethod(method string) bool {
	validMethods := []string{http.MethodPost, http.MethodPatch, http.MethodGet, http.MethodPut, http.MethodDelete}

	for _, validMethod := range validMethods {
		if validMethod == method {
			return true
		}
	}

	return false
}

func isValidEchoRoute(ecRoute *echo.Route) bool {
	if !isValidMethod(ecRoute.Method) {
		return false
	}

	if strings.Contains(ecRoute.Name, consts.EndpointNameEchoLibrary) ||
		strings.Contains(ecRoute.Name, consts.EndpointNameApp) {
		return false
	}

	return true
}

func cleanPath(path string) string {
	var asciiCodeSlash byte = 47

	if path[len(path)-1] == asciiCodeSlash {
		path = path[:len(path)-1]
	}

	return path
}

func (req *BulkUpsertEndpointsRequest) NewBulkUpsertEndpointsRequestFromEcho(ec echo.Echo) error {
	for _, ecRoute := range ec.Routes() {
		if !isValidEchoRoute(ecRoute) {
			continue
		}

		req.Endpoints = append(req.Endpoints, UpsertEndpointRequest{
			Method: ecRoute.Method,
			Path:   cleanPath(ecRoute.Path),
		})
	}

	return nil
}

func (req UpsertEndpointRequest) ToEndpoint() dao.Endpoint {
	now := time.Now()
	return dao.Endpoint{
		Method:    req.Method,
		Path:      req.Path,
		CreatedAt: now,
		UpdatedAt: now,
		CreatedBy: consts.SystemCreatedBy,
		UpdatedBy: consts.SystemCreatedBy,
	}
}

func (reqs UpsertEndpointRequests) FindByMethodPath(method, path string) UpsertEndpointRequest {
	for _, req := range reqs {
		if req.Method == method && req.Path == path {
			return req
		}
	}

	return UpsertEndpointRequest{}
}
