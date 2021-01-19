package main

type createMetricsRequest struct {
	Metrics Metrics `json:"metrics"`
}

type createMetricsResponse struct {
	Metrics Metrics `json:"metrics"`
}

type getMetricsByIdRequest struct {
	Id string `json:"id"`
}

type getMetricsByIdResponse struct {
	Metrics *Metrics `json:"metrics"`
}

type ListMetricsRequest struct {

}

type ListMetricsResponse struct {
	Metrics []Metrics `json:"metrics"`
}

type updateMetricsRequest struct {
	Metrics Metrics `json:"metrics"`
}

type updateMetricsResponse struct {
	Metrics Metrics `json:"metrics"`
	isUpdated bool
}

type deleteMetricRequest struct {
	Id string `json:"id"`
}

type deleteMetricResponse struct {
	isDeleted bool
} 

