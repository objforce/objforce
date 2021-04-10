package models

type ForecastCategories string

const (
	ForecastCategoriesOmitted  ForecastCategories = "Omitted"
	ForecastCategoriesPipeline ForecastCategories = "Pipeline"
	ForecastCategoriesBestCase ForecastCategories = "BestCase"
	ForecastCategoriesForecast ForecastCategories = "Forecast"
	ForecastCategoriesClosed   ForecastCategories = "Closed"
)
