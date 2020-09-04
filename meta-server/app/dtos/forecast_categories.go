package dtos

type ForecastCategories string

const(
	Omitted ForecastCategories = "Omitted"
	Pipeline ForecastCategories = "Pipeline"
	BestCase ForecastCategories = "BestCase"
	Forecast ForecastCategories = "Forecast"
	Closed ForecastCategories = "Closed"
)