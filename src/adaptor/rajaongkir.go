package adaptor

import (
	"encoding/json"
	"mini_e_commerce/src/pkg/httpclient"
)

const (
	RAJAONGKIR_API_KEY = "173c413e3e7e9dd49d7bb48af15a2175"
)

type RajaOngkirAdaptor struct {
	client *httpclient.Client
}

func NewRajaOngkirAdaptor(baseUrl string) *RajaOngkirAdaptor {
	client := httpclient.NewHttpClient(baseUrl)

	var header = make(map[string]string)
	header["key"] = RAJAONGKIR_API_KEY
	client.SetHeader(header)
	return &RajaOngkirAdaptor{
		client: client,
	}
}

func (t *RajaOngkirAdaptor) GetAllCity() (*[]JSONRajaOngkirCities, error) {
	data, err := t.client.Get("/starter/city")
	if err != nil {
		return nil, err
	}
	var datas []JSONRajaOngkirCities

	err = json.Unmarshal(data, &datas)
	if err != nil {
		return nil, err
	}

	return &datas, err

}

func (t *RajaOngkirAdaptor) GetCity(id string) (*[]JSONRajaOngkirCity, error) {
	data, err := t.client.Get("/starter/city?id=" + id)
	if err != nil {
		return nil, err
	}
	var datas []JSONRajaOngkirCity

	err = json.Unmarshal(data, &datas)
	if err != nil {
		return nil, err
	}

	return &datas, err

}

func (t *RajaOngkirAdaptor) GetAllProvince() (*[]JSONRajaOngkirProvinces, error) {
	data, err := t.client.Get("/starter/province")
	if err != nil {
		return nil, err
	}
	var datas []JSONRajaOngkirProvinces

	err = json.Unmarshal(data, &datas)
	if err != nil {
		return nil, err
	}

	return &datas, err

}

func (t *RajaOngkirAdaptor) GetProvince(id string) (*[]JSONRajaOngkirProvince, error) {
	data, err := t.client.Get("/starter/province?id=" + id)
	if err != nil {
		return nil, err
	}
	var datas []JSONRajaOngkirProvince

	err = json.Unmarshal(data, &datas)
	if err != nil {
		return nil, err
	}

	return &datas, err

}

func (t *RajaOngkirAdaptor) GetCost(payload interface{}) (*[]JSONRajaOngkirCost, error) {
	data, err := t.client.Post("/starter/cost", payload)
	if err != nil {
		return nil, err
	}
	var datas []JSONRajaOngkirCost

	err = json.Unmarshal(data, &datas)
	if err != nil {
		return nil, err
	}

	return &datas, err

}
