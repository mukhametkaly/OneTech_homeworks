package main

import (
	"github.com/rs/xid"
	"time"
)

type MetricsInterface interface {
	CreateMetric(req createMetricsRequest) *createMetricsResponse
	UpdateMetric(req updateMetricsRequest) *updateMetricsResponse
	GetMetricById(req getMetricsByIdRequest) *getMetricsByIdResponse
	ListMetrics(req ListMetricsRequest) ListMetricsResponse
	DeleteMetricById(req deleteMetricRequest) *deleteMetricResponse
}

type metricsService struct {
	metricsSlice []Metrics
}

func (s *metricsService) Init() {
	s.metricsSlice = []Metrics{}
}

func (s *metricsService) CreateMetric(req createMetricsRequest) *createMetricsResponse {
	metrics := req.Metrics

	id := xid.New()
	metrics.Id = id.String()
	metrics.Time = time.Now()

	s.metricsSlice = append(s.metricsSlice, metrics)

	resp := createMetricsResponse{Metrics: metrics}

	return &resp
}

func (s *metricsService) GetMetricById(req getMetricsByIdRequest) *getMetricsByIdResponse {
	resp := getMetricsByIdResponse{}

	for _, metrics := range s.metricsSlice {
		if metrics.Id == req.Id {
			resp.Metrics = &metrics
			return &resp
		}
	}

	resp.Metrics = nil
	return &resp
}

func (s *metricsService) ListMetrics(_ ListMetricsRequest) ListMetricsResponse{

	resp := ListMetricsResponse{
		Metrics: s.metricsSlice,
	}
	return resp
}

func (s *metricsService) UpdateMetric(req updateMetricsRequest) *updateMetricsResponse  {
	resp := updateMetricsResponse{}
	for w, i := range s.metricsSlice {
		if req.Metrics.Id == i.Id {
			(*s).metricsSlice[w] = req.Metrics
			(*s).metricsSlice[w].Time = time.Now()
			resp.isUpdated = true
			resp.Metrics = i
			return &resp
		}
	}
	resp.isUpdated = false
	return &resp

}

func (s *metricsService) DeleteMetricById(req deleteMetricRequest) *deleteMetricResponse   {
	resp := deleteMetricResponse{}
	for q, i := range s.metricsSlice {
		if req.Id == i.Id {
			s.metricsSlice[q] = s.metricsSlice[len(s.metricsSlice) - 1]
			s.metricsSlice = s.metricsSlice[: len(s.metricsSlice) - 1]
			resp.isDeleted = true
			return  &resp
		}
	}

	resp.isDeleted = false
	return &resp
}