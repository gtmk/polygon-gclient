package polygonio

import "time"

////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
////////               Forex Endpoints                          ////////////
////////                                                        ////////////
////////////////////////////////////////////////////////////////////////////
////////////////////////////////////////////////////////////////////////////
func (c *Client) ForexPreviousClose(ticker string, opts *RequestOptions) (*Bars, error) {
	return c.StockPreviousClose(ticker, nil)
}

func (c *Client) ForexAggregates(ticker string, multiplier int32, timespan Timespan, from, to time.Time, opts *RequestOptions) (*Bars, error) {
	return c.StockAggregates(ticker, multiplier, timespan, from, to, nil)
}

func (c *Client) ForexGroupedDaily(locale Locale, date time.Time, opts *RequestOptions) (*Bars, error) {
	return c.StockGroupedDaily(locale, FX, date, opts)
}

func (c *Client) ForexHistoricTicks()             {}
func (c *Client) ForexRealTimeConversion()        {}
func (c *Client) ForexLastQuotesForCurrencyPair() {}
func (c *Client) ForexSnapshotAll()               {}
func (c *Client) ForexSnapshotTopGainersLosers()  {}
