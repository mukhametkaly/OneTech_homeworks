package main

import (
	"context"
	"github.com/go-kit/kit/endpoint"
)

func makeCreateMetricsEndpoint(s *metricsService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(createMetricsRequest)

		resp = s.CreateMetric(req)

		return resp, nil
	}
}

func GetMetricsByIdEndpoint(s *metricsService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(getMetricsByIdRequest)

		resp = s.GetMetricById(req)

		return resp, nil
	}
}

func UpdateMetricEndpoint(s *metricsService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(updateMetricsRequest)

		resp = s.UpdateMetric(req)

		return resp, nil
	}
}

func ListMetricsEndpoint(s *metricsService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(ListMetricsRequest)

		resp = s.ListMetrics(req)

		return resp, nil
	}
}

func DeleteMetricByIdEndpoint(s *metricsService) endpoint.Endpoint {
	return func(ctx context.Context, r interface{}) (resp interface{}, err error) {
		req := r.(deleteMetricRequest)

		resp = s.DeleteMetricById(req)

		return resp, nil
	}
}
