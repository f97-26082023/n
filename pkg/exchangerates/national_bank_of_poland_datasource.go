package exchangerates

import (
	"bytes"
	"encoding/xml"
	"math"
	"time"

	"github.com/f97/n/pkg/core"
	"github.com/f97/n/pkg/errs"
	"github.com/f97/n/pkg/log"
	"github.com/f97/n/pkg/models"
	"github.com/f97/n/pkg/utils"
	"github.com/f97/n/pkg/validators"
)

const nationalBankOfPolandDailyExchangeRateUrl = "https://www.nbp.pl/kursy/xml/en/lastaen.xml"
const nationalBankOfPolandInconvertibleCurrencyExchangeRateUrl = "https://www.nbp.pl/kursy/xml/en/lastben.xml"
const nationalBankOfPolandExchangeRateReferenceUrl = "https://www.nbp.pl/homen.aspx?f=/kursy/kursyen.htm"
const nationalBankOfPolandDataSource = "Narodowy Bank Polski"
const nationalBankOfPolandBaseCurrency = "PLN"

const nationalBankOfPolandDataUpdateDateFormat = "2006-01-02 15:04"
const nationalBankOfPolandDataUpdateDateTimezone = "Europe/Warsaw"

// NationalBankOfPolandDataSource defines the structure of exchange rates data source of National Bank of Poland
type NationalBankOfPolandDataSource struct {
	ExchangeRatesDataSource
}

// NationalBankOfPolandExchangeRateData represents the whole data from National Bank of Poland
type NationalBankOfPolandExchangeRateData struct {
	XMLName          xml.Name                            `xml:"exchange_rates"`
	Date             string                              `xml:"date,attr"`
	AllExchangeRates []*NationalBankOfPolandExchangeRate `xml:"mid-rate"`
}

// NationalBankOfPolandExchangeRate represents the exchange rate data from National Bank of Poland
type NationalBankOfPolandExchangeRate struct {
	Currency string `xml:"code,attr"`
	Units    string `xml:"units,attr"`
	Rate     string `xml:",chardata"`
}

// ToLatestExchangeRateResponse returns a view-object according to original data from National Bank of Poland
func (e *NationalBankOfPolandExchangeRateData) ToLatestExchangeRateResponse(c *core.Context) *models.LatestExchangeRateResponse {
	if len(e.AllExchangeRates) < 1 {
		log.ErrorfWithRequestId(c, "[national_bank_of_poland_datasource.ToLatestExchangeRateResponse] all exchange rates is empty")
		return nil
	}

	exchangeRates := make(models.LatestExchangeRateSlice, 0, len(e.AllExchangeRates))

	for i := 0; i < len(e.AllExchangeRates); i++ {
		exchangeRate := e.AllExchangeRates[i]

		if _, exists := validators.AllCurrencyNames[exchangeRate.Currency]; !exists {
			continue
		}

		finalExchangeRate := exchangeRate.ToLatestExchangeRate(c)

		if finalExchangeRate == nil {
			continue
		}

		exchangeRates = append(exchangeRates, finalExchangeRate)
	}

	timezone, err := time.LoadLocation(nationalBankOfPolandDataUpdateDateTimezone)

	if err != nil {
		log.ErrorfWithRequestId(c, "[national_bank_of_poland_datasource.ToLatestExchangeRateResponse] failed to get timezone, timezone name is %s", nationalBankOfPolandDataUpdateDateTimezone)
		return nil
	}

	updateDateTime := e.Date + " 12:15" // Table A of the average foreign currency exchange rates is published (updated) on the NBP website on business days, between 11:45 a.m. and 12:15 p.m.
	updateTime, err := time.ParseInLocation(nationalBankOfPolandDataUpdateDateFormat, updateDateTime, timezone)

	if err != nil {
		log.ErrorfWithRequestId(c, "[national_bank_of_poland_datasource.ToLatestExchangeRateResponse] failed to parse update date, datetime is %s", updateDateTime)
		return nil
	}

	latestExchangeRateResp := &models.LatestExchangeRateResponse{
		DataSource:    nationalBankOfPolandDataSource,
		ReferenceUrl:  nationalBankOfPolandExchangeRateReferenceUrl,
		UpdateTime:    updateTime.Unix(),
		BaseCurrency:  nationalBankOfPolandBaseCurrency,
		ExchangeRates: exchangeRates,
	}

	return latestExchangeRateResp
}

// ToLatestExchangeRate returns a data pair according to original data from National Bank of Poland
func (e *NationalBankOfPolandExchangeRate) ToLatestExchangeRate(c *core.Context) *models.LatestExchangeRate {
	amount, err := utils.StringToInt64(e.Units)

	if err != nil {
		log.WarnfWithRequestId(c, "[national_bank_of_poland_datasource.ToLatestExchangeRate] failed to parse amount, currency is %s, amount is %s", e.Currency, e.Units)
		return nil
	}

	rate, err := utils.StringToFloat64(e.Rate)

	if err != nil {
		log.WarnfWithRequestId(c, "[national_bank_of_poland_datasource.ToLatestExchangeRate] failed to parse rate, currency is %s, rate is %s", e.Currency, e.Rate)
		return nil
	}

	if rate <= 0 {
		log.WarnfWithRequestId(c, "[national_bank_of_poland_datasource.ToLatestExchangeRate] rate is invalid, currency is %s, rate is %s", e.Currency, e.Rate)
		return nil
	}

	finalRate := float64(amount) / rate

	if math.IsInf(finalRate, 0) {
		return nil
	}

	return &models.LatestExchangeRate{
		Currency: e.Currency,
		Rate:     utils.Float64ToString(finalRate),
	}
}

// GetRequestUrls returns the National Bank of Poland data source urls
func (e *NationalBankOfPolandDataSource) GetRequestUrls() []string {
	return []string{nationalBankOfPolandInconvertibleCurrencyExchangeRateUrl, nationalBankOfPolandDailyExchangeRateUrl}
}

// Parse returns the common response entity according to the National Bank of Poland data source raw response
func (e *NationalBankOfPolandDataSource) Parse(c *core.Context, content []byte) (*models.LatestExchangeRateResponse, error) {
	nationalBankOfPolandData := &NationalBankOfPolandExchangeRateData{}

	xmlDecoder := xml.NewDecoder(bytes.NewReader(content))
	xmlDecoder.CharsetReader = utils.IdentReader
	err := xmlDecoder.Decode(&nationalBankOfPolandData)

	if err != nil {
		log.ErrorfWithRequestId(c, "[national_bank_of_poland_datasource.Parse] failed to parse xml data, content is %s, because %s", string(content), err.Error())
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	latestExchangeRateResponse := nationalBankOfPolandData.ToLatestExchangeRateResponse(c)

	if latestExchangeRateResponse == nil {
		log.ErrorfWithRequestId(c, "[national_bank_of_poland_datasource.Parse] failed to parse latest exchange rate data, content is %s", string(content))
		return nil, errs.ErrFailedToRequestRemoteApi
	}

	return latestExchangeRateResponse, nil
}
