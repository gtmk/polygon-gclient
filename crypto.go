package polygonio

import (
	"context"
	"fmt"
	"net/url"
)

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

func (c *Client) CryptoPreviousClose(ticker string, opts *RequestOptions) (*Bars, error) {
	return c.StockPreviousClose(ticker, opts)
}

func (c *Client) CryptoAggregates(ticker string, multiplier int32, timespan Timespan, from, to string, opts *RequestOptions) (*Bars, error) {
	return c.StockAggregates(ticker, multiplier, timespan, from, to, opts)
}

func (c *Client) CryptoGroupedDaily(locale Locale, date string, opts *RequestOptions) (*Bars, error) {
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
