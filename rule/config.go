package rule

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type ConfigRules struct {
	Rules struct {
		LowestPriceRule struct {
			Packages []string `json:"packages"`
		} `json:"lowestPriceRule"`
		FreeShippingPerMonthRule struct {
			Providers struct {
				Lp struct {
					L int `json:"L"`
				} `json:"LP"`
			} `json:"providers"`
		} `json:"freeShippingPerMonthRule"`
		MaxDiscountPerMonthRule struct {
			MaximumDiscount float64 `json:"maximumDiscount"`
		} `json:"maxDiscountPerMonthRule"`
	} `json:"rules"`
}

type Config struct {
	ConfigDataFileLocation string
	ConfigRules            *ConfigRules
}

func NewConfig() (*Config, error) {
	c := &Config{}
	c = c.SetConfigFileLocation()

	file, err := ioutil.ReadFile(c.ConfigDataFileLocation)
	if err != nil {
		return nil, err
	}

	rules := new(ConfigRules)
	err = json.Unmarshal(file, rules)
	if err != nil {
		return nil, err
	}

	c.ConfigRules = rules

	return c, nil
}

func (c *Config) SetConfigFileLocation() *Config {
	c.ConfigDataFileLocation = os.Getenv("CONFIG_PATH")
	return c
}

func (c *Config) GetRules() *ConfigRules {
	return c.ConfigRules
}
