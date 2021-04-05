package models

type ForecastCategory string

const(
	Omitted ForecastCategory = "Omitted" 
	Pipeline ForecastCategory = "Pipeline"
	BestCase ForecastCategory = "BestCase"
	Forecast ForecastCategory = "Forecast"
	Closed ForecastCategory = "Closed"
)