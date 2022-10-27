package adaptor

type RajaOngkirQueryProp struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Weight      uint32 `json:"weight"`
	Courier     uint16 `json:"courier"`
}

type RajaOngkirStatusProp struct {
	Code        uint16 `json:"code"`
	Description string `json:"description"`
}

type RajaOngkirOriginAndDestinationProp struct {
	CityId     string `json:"city_id"`
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
	Type       string `json:"type"`
	CityName   string `json:"city_name"`
	PostalCode string `json:"postal_code"`
}

type RajaOngkirCourierResultProp struct {
	Code  string                      `json:"code"`
	Name  string                      `json:"name"`
	Costs []RajaOngkirCourierCostProp `json:"costs"`
}

type RajaOngkirCourierCostProp struct {
	Service     string                            `json:"service"`
	Description string                            `json:"description"`
	Cost        []RajaOngkirCourierCostDetailProp `json:"cost"`
}

type RajaOngkirCourierCostDetailProp struct {
	Value uint32 `json:"value"`
	Etd   string `json:"etd"`
	Note  string `json:"note"`
}

type RajaOngkirProvinceResultProp struct {
	ProvinceId string `json:"province_id"`
	Province   string `json:"province"`
}

type RajaOngkirCityResultProp struct {
	CityId string `json:"city_id"`
	City   string `json:"city"`
}

type RajaOngkirProvinceCityQueryProp struct {
	Id string `json:"id"`
}

type JSONRajaOngkirCost struct {
	Query              RajaOngkirQueryProp                `json:"query"`
	Status             RajaOngkirStatusProp               `json:"status"`
	OriginDetails      RajaOngkirOriginAndDestinationProp `json:"origin_details"`
	DestinationDetails RajaOngkirOriginAndDestinationProp `json:"destination_details"`
	Results            []RajaOngkirCourierResultProp      `json:"results"`
}

type JSONRajaOngkirProvince struct {
	Query   RajaOngkirProvinceCityQueryProp `json:"query"`
	Status  RajaOngkirStatusProp            `json:"status"`
	Results RajaOngkirProvinceResultProp    `json:"results"`
}

type JSONRajaOngkirProvinces struct {
	Query   interface{}                    `json:"query"`
	Status  RajaOngkirStatusProp           `json:"status"`
	Results []RajaOngkirProvinceResultProp `json:"results"`
}

type JSONRajaOngkirCity struct {
	Query   RajaOngkirProvinceCityQueryProp `json:"query"`
	Status  RajaOngkirStatusProp            `json:"status"`
	Results RajaOngkirCityResultProp        `json:"results"`
}

type JSONRajaOngkirCities struct {
	Query   interface{}                `json:"query"`
	Status  RajaOngkirStatusProp       `json:"status"`
	Results []RajaOngkirCityResultProp `json:"results"`
}
