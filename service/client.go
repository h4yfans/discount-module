package service

import (
	"fmt"
	"os"

	"github.com/h4yfans/discount-module/entity"
	"github.com/h4yfans/discount-module/rule"
)

type Client struct {
	EntityListService     *EntityListService
	RuleService           *rule.Rules
	InputDataFileLocation string
}

func NewClient(e *EntityListService, r *rule.Rules) *Client {
	return &Client{EntityListService: e, RuleService: r}
}

func (c *Client) Output() error {
	c.SetInputDataFileLocation()

	providerList, err := c.EntityListService.GetProviderList()
	if err != nil {
		return err
	}

	fileService := NewFileService(c.InputDataFileLocation)
	file, err := fileService.Open()
	if err != nil {
		return err
	}

	for _, line := range file.Lines {
		shipment := entity.NewShipment()

		shipment, err = shipment.FromString(line, fileService.Delimiter, providerList)
		if err != nil {
			return err
		}

		if shipment != nil {
			ruleSet := rule.RuleSet{
				ProviderList: providerList,
				Shipment:     shipment,
			}
			shipment = c.RuleService.ApplyAllRules(ruleSet)
			price := fmt.Sprintf("%.2f", shipment.GetPrice())
			var discount string
			if shipment.GetDiscount() > 0 {
				discount = fmt.Sprintf("%.2f", shipment.GetDiscount())
			} else {
				discount = "-"
			}
			fmt.Printf("%s %s %s \n", line, price, discount)
		} else {
			fmt.Printf("%s %s \n", line, "Ignored")
		}
	}

	return nil
}

func (c *Client) SetInputDataFileLocation() *Client {
	c.InputDataFileLocation = os.Getenv("INPUT_PATH")
	return c
}
