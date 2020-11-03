package polygonio

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
)

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////               Client Configuration                     ////////////
////////                                                        ////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////

const apiURL = "https://api.polygon.io"

type Client struct {
	baseURL    string
	token      string
	httpClient *http.Client
}

type Error struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

func (e Error) Error() string {
	return fmt.Sprintf("%d %s: %s", e.StatusCode, http.StatusText(e.StatusCode), e.Message)
}

func NewClient(token string, options ...func(*Client)) *Client {
	client := &Client{
		token:      token,
		httpClient: &http.Client{},
	}

	// apply options
	for _, option := range options {
		option(client)
	}

	// set default values
	if client.baseURL == "" {
		client.baseURL = apiURL
	}
	return client
}

func WithHTTPClient(httpClient *http.Client) func(*Client) {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

func WithBaseURL(baseURL string) func(*Client) {
	return func(client *Client) {
		client.baseURL = baseURL
	}
}

func (c *Client) GetJSON(ctx context.Context, endpoint string, v interface{}) error {
	address, err := c.addToken(endpoint)
	if err != nil {
		return err
	}
	data, err := c.getBytes(ctx, address)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func (c *Client) GetBytes(ctx context.Context, endpoint string) ([]byte, error) {
	address, err := c.addToken(endpoint)
	if err != nil {
		return []byte{}, err
	}
	return c.getBytes(ctx, address)
}

func (c *Client) getBytes(ctx context.Context, address string) ([]byte, error) {
	req, err := http.NewRequest("GET", address, nil)
	if err != nil {
		return []byte{}, err
	}
	resp, err := c.httpClient.Do(req.WithContext(ctx))
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, err := ioutil.ReadAll(resp.Body)
		msg := ""

		if err == nil {
			msg = string(b)
		}

		return []byte{}, Error{StatusCode: resp.StatusCode, Message: msg}
	}
	return ioutil.ReadAll(resp.Body)
}

func (c *Client) addToken(endpoint string) (string, error) {
	u, err := url.Parse(c.baseURL + endpoint)
	if err != nil {
		return "", err
	}
	v := u.Query()
	v.Add("apiKey", c.token)
	u.RawQuery = v.Encode()
	return u.String(), nil
}

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////               Reference Endpoints                      ////////////
////////                                                        ////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
func (c *Client) ReferenceTickers(opts *TickerOptions) (CommonResponse, Tickers, error) {
	out := struct {
		Common  CommonResponse
		Tickers Tickers `json:"tickers"`
	}{}
	endpoint := fmt.Sprintf("/v2/reference/tickers")
	endpoint, err := c.referenceTickersWithOpts(endpoint, opts)
	if err != nil {
		return CommonResponse{}, nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	return out.Common, out.Tickers, err
}

func (c *Client) referenceTickersWithOpts(endpoint string, opts *TickerOptions) (string, error) {
	if opts == nil {
		return endpoint, nil
	}
	v, err := query.Values(opts)
	if err != nil {
		return "", err
	}
	optParams := v.Encode()
	if optParams != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, optParams)
	}
	return endpoint, nil
}

func (c *Client) ReferenceTickerTypes() (CommonResponse, map[string]string, map[string]string, error) {
	out := struct {
		Common  CommonResponse
		Results struct {
			Types      map[string]string `json:"types"`
			IndexTypes map[string]string `json:"indexTypes"`
		} `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/reference/types")
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out.Common, out.Results.Types, out.Results.IndexTypes, err
}

func (c *Client) ReferenceTickerDetail(ticker string) (TickerDetail, error) {
	var out TickerDetail
	endpoint := fmt.Sprintf("/v1/meta/symbols/%s/company", url.PathEscape(ticker))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

func (c *Client) ReferenceTickerNews(ticker string, opts *NewsOptions) ([]TickerNews, error) {
	var out []TickerNews
	endpoint := fmt.Sprintf("/v1/meta/symbols/%s/news", url.PathEscape(ticker))
	endpoint, err := c.referenceTickerNewsWithOpts(endpoint, opts)
	if err != nil {
		return out, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

func (c *Client) referenceTickerNewsWithOpts(endpoint string, opts *NewsOptions) (string, error) {
	if opts == nil {
		return endpoint, nil
	}
	v, err := query.Values(opts)
	if err != nil {
		return "", err
	}
	optParams := v.Encode()
	if optParams != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, optParams)
	}
	return endpoint, nil
}

func (c *Client) ReferenceMarkets() (CommonResponse, MarketDescriptions, error) {
	out := struct {
		Common  CommonResponse
		Results MarketDescriptions `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/reference/markets")
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out.Common, out.Results, err
}

func (c *Client) ReferenceLocales() (CommonResponse, LocaleNames, error) {
	out := struct {
		Common  CommonResponse
		Results LocaleNames `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/reference/locales")
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out.Common, out.Results, err
}

func (c *Client) ReferenceStockSplits(ticker string) (CommonResponse, Splits, error) {
	out := struct {
		Common  CommonResponse
		Results Splits `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/reference/splits/%s", url.PathEscape(ticker))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out.Common, out.Results, err
}

func (c *Client) ReferenceDividends(ticker string) (CommonResponse, Dividends, error) {
	out := struct {
		Common  CommonResponse
		Results Dividends `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/reference/dividends/%s", url.PathEscape(ticker))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out.Common, out.Results, err
}

func (c *Client) ReferenceFinancials(ticker string, opts *FinancialOptions) (CommonResponse, Financials, error) {
	out := struct {
		Common  CommonResponse
		Results Financials `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/reference/financials/%s", url.PathEscape(ticker))
	endpoint, err := c.referenceFinancialsWithOpts(endpoint, opts)
	if err != nil {
		return CommonResponse{}, nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	return out.Common, out.Results, err
}

func (c *Client) referenceFinancialsWithOpts(endpoint string, opts *FinancialOptions) (string, error) {
	if opts == nil {
		return endpoint, nil
	}
	v, err := query.Values(opts)
	if err != nil {
		return "", err
	}
	optParams := v.Encode()
	if optParams != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, optParams)
	}
	return endpoint, nil
}

func (c *Client) ReferenceMarketStatus() (MarketStatus, error) {
	var out MarketStatus
	endpoint := fmt.Sprintf("/v1/marketstatus/now")
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

func (c *Client) ReferenceMarketHolidays() (MarketHolidays, error) {
	var out MarketHolidays
	endpoint := fmt.Sprintf("/v1/marketstatus/upcoming")
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////               Stock Endpoints                          ////////////
////////                                                        ////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////

func (c *Client) StockExchanges() (Exchanges, error) {
	var out Exchanges
	endpoint := fmt.Sprintf("/v1/meta/exchanges")
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

func (c *Client) StockPreviousClose(ticker string, opts *RequestOptions) (CommonResponse, Bars, error) {
	out := struct {
		Common  CommonResponse
		Results Bars `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/aggs/ticker/%s/prev", url.PathEscape(ticker))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return CommonResponse{}, nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	if err != nil {
		return CommonResponse{}, nil, err
	}
	return out.Common, out.Results, err
}

func (c *Client) endpointWithOpts(endpoint string, opts *RequestOptions) (string, error) {
	if opts == nil {
		return endpoint, nil
	}
	v, err := query.Values(opts)
	if err != nil {
		return "", err
	}
	optParams := v.Encode()
	if optParams != "" {
		endpoint = fmt.Sprintf("%s?%s", endpoint, optParams)
	}
	return endpoint, nil
}

func (c *Client) StockAggregates(ticker string, multiplier int32, timespan Timespan, from, to string, opts *RequestOptions) (CommonResponse, Bars, error) {
	out := struct {
		Common  CommonResponse
		Results Bars `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/aggs/ticker/%s/range/%s/%s/%s/%s", url.PathEscape(ticker), url.PathEscape(strconv.Itoa(int(multiplier))), url.PathEscape(string(timespan)), url.PathEscape(from), url.PathEscape(to))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return CommonResponse{}, nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	return out.Common, out.Results, err
}

func (c *Client) StockGroupedDaily(locale Locale, market Market, date string, opts *RequestOptions) (CommonResponse, Bars, error) {
	out := struct {
		Common  CommonResponse
		Results Bars `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/aggs/grouped/locale/%s/market/%s/%s", url.PathEscape(string(locale)), url.PathEscape(string(market)), url.PathEscape(date))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return CommonResponse{}, nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	return out.Common, out.Results, err
}

func (c *Client) StockHistoricTrades() {}
func (c *Client) StockHistoricQuotes() {}
func (c *Client) StockLastTrade()      {}
func (c *Client) StockLastQuote()      {}

func (c *Client) StockDaily(ticker, date string) (Daily, error) {
	var out Daily
	endpoint := fmt.Sprintf("/v1/open-close/%s/%s", url.PathEscape(ticker), url.PathEscape(date))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

func (c *Client) StockConditionMappings(tick Tick) (map[string]string, error) {
	out := make(map[string]string)
	endpoint := fmt.Sprintf("/v1/meta/conditions/%s", url.PathEscape(string(tick)))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

func (c *Client) StockSnapshotAll()                 {}
func (c *Client) StockSnapshotSingle(ticker string) {}
func (c *Client) StockSnapshotTopGainersLosers()    {}

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////               Forex Endpoints                          ////////////
////////                                                        ////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
func (c *Client) ForexPreviousClose(ticker string, opts *RequestOptions) (CommonResponse, Bars, error) {
	return c.StockPreviousClose(ticker, nil)
}

func (c *Client) ForexAggregates(ticker string, multiplier int32, timespan Timespan, from, to string, opts *RequestOptions) (CommonResponse, Bars, error) {
	return c.StockAggregates(ticker, multiplier, timespan, from, to, nil)
}

func (c *Client) ForexGroupedDaily(locale Locale, date string, opts *RequestOptions) (CommonResponse, Bars, error) {
	return c.StockGroupedDaily(locale, FX, date, opts)
}

func (c *Client) ForexHistoricTicks()             {}
func (c *Client) ForexRealTimeConversion()        {}
func (c *Client) ForexLastQuotesForCurrencyPair() {}
func (c *Client) ForexSnapshotAll()               {}
func (c *Client) ForexSnapshotTopGainersLosers()  {}

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////               Crypto Endpoints                         ////////////
////////                                                        ////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
func (c *Client) CryptoExchanges() (Exchanges, error) {
	var out Exchanges
	endpoint := fmt.Sprintf("/v1/meta/crypto-exchanges")
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

func (c *Client) CryptoPreviousClose(ticker string, opts *RequestOptions) (CommonResponse, Bars, error) {
	return c.StockPreviousClose(ticker, opts)
}

func (c *Client) CryptoAggregates(ticker string, multiplier int32, timespan Timespan, from, to string, opts *RequestOptions) (CommonResponse, Bars, error) {
	return c.StockAggregates(ticker, multiplier, timespan, from, to, opts)
}

func (c *Client) CryptoGroupedDaily(locale Locale, date string, opts *RequestOptions) (CommonResponse, Bars, error) {
	return c.StockGroupedDaily(locale, Crypto, date, opts)
}

func (c *Client) CryptoLastTradeForCryptoPair() {}

func (c *Client) CryptoDaily(from, to, date string) (CryptoDaily, error) {
	var out CryptoDaily
	endpoint := fmt.Sprintf("/v1/open-close/crypto/%s/%s/%s", url.PathEscape(from), url.PathEscape(to), url.PathEscape(date))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

func (c *Client) CryptoHistoricTrades()   {}
func (c *Client) CryptoSnapshotAll()      {}
func (c *Client) CryptoSnapshotFullBook() {}
