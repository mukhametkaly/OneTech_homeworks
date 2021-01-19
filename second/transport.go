package main

import (
	"context"
	"encoding/json"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"net/http"
)

func MakeHandler(s metricsService) http.Handler {

	options := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}

	createMetrics := kithttp.NewServer(
		makeCreateMetricsEndpoint(&s),
		decodeCreateMetricsRequest,
		encodeResponse,
		options...)

	getMetrics := kithttp.NewServer(
		GetMetricsByIdEndpoint(&s),
		decodeGetMetricsRequest,
		encodeResponse,
		options...)

	listMetrics := kithttp.NewServer(
		ListMetricsEndpoint(&s),
		decodeListMetrics,
		encodeResponse,
		options...)

	updateMetric := kithttp.NewServer(
		UpdateMetricEndpoint(&s),
		decodeUpdateMetric,
		encodeResponse,
		options...)

	deleteMetric := kithttp.NewServer(
		DeleteMetricByIdEndpoint(&s),
		decodeDeleteMetricById,
		encodeResponse,
		options...)


	r := mux.NewRouter()

	//create metrics
	r.Handle("/metrics/create/", createMetrics).Methods("POST")

	//get metrics
	r.Handle("/metrics/get/{id}", getMetrics).Methods("GET")

	//list metrics
	r.Handle("/metrics/", listMetrics).Methods("GET")

	//update metrics
	r.Handle("/metrics/update", updateMetric).Methods("PUT")

	//delete metrics
	r.Handle("/metrics/delete/{id}", deleteMetric).Methods("DELETE")

	return r
}

func decodeCreateMetricsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	apiRequest := createMetricsRequest{}

	if err := json.NewDecoder(r.Body).Decode(&apiRequest); err != nil {
		return nil, errors.New("400")
	}else if apiRequest.Metrics.Value == "" {
		return nil, errors.New("No Value")
	} else if apiRequest.Metrics.Type == "" {
		return nil, errors.New("No metric type")
	}

	return apiRequest, nil
}

func decodeGetMetricsRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	apiRequest := getMetricsByIdRequest{}

	vars := mux.Vars(r)
	id := vars["id"]
	apiRequest.Id = id

	return apiRequest, nil
}

func decodeUpdateMetric(_ context.Context, r *http.Request) (request interface{}, err error) {
	apiRequest := updateMetricsRequest{}

	if err := json.NewDecoder(r.Body).Decode(&apiRequest); err != nil {
		return nil, errors.New("400")
	}else if apiRequest.Metrics.Value == "" {
		return nil, errors.New("No Value")
	} else if apiRequest.Metrics.Type == "" {
		return nil, errors.New("No metric type")
	}
	return apiRequest, nil
}

func decodeListMetrics(_ context.Context, _ *http.Request) (request interface{}, err error) {
	apiRequest := ListMetricsRequest{}
	return apiRequest, nil
}

func decodeDeleteMetricById(_ context.Context, r *http.Request) (request interface{}, err error) {
	apiRequest := deleteMetricRequest{}

	vars := mux.Vars(r)
	id := vars["id"]
	apiRequest.Id = id

	return apiRequest, nil
}


func encodeResponse(_ context.Context, w http.ResponseWriter, r interface{}) error {
	if e, ok := r.(errorInterface); ok && e.error() != nil {
		encodeError(context.Background(), e.error(), w)
		return nil
	}

	return json.NewEncoder(w).Encode(r)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	_ = json.NewEncoder(w).Encode(err)
}

type errorInterface interface {
	error() error
}
