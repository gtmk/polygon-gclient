package polygonio

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/google/go-querystring/query"
)

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

func (c *Client) StockPreviousClose(ticker string, opts *RequestOptions) (*Bars, error) {
	out := struct {
		Results Bars `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/aggs/ticker/%s/prev", url.PathEscape(ticker))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	if err != nil {
		return nil, err
	}
	return &out.Results, err
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

func (c *Client) StockAggregates(ticker string, multiplier int32, timespan Timespan, from, to string, opts *RequestOptions) (*Bars, error) {
	out := struct {
		Results Bars `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/aggs/ticker/%s/range/%s/%s/%s/%s", url.PathEscape(ticker), url.PathEscape(strconv.Itoa(int(multiplier))), url.PathEscape(string(timespan)), url.PathEscape(from), url.PathEscape(to))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	return &out.Results, err
}

func (c *Client) StockGroupedDaily(locale Locale, market Market, date string, opts *RequestOptions) (*Bars, error) {
	out := struct {
		Results Bars `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/aggs/grouped/locale/%s/market/%s/%s", url.PathEscape(string(locale)), url.PathEscape(string(market)), url.PathEscape(date))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	return &out.Results, err
}

func (c *Client) StockTrades(ticker, date string, opts *RequestOptions) (*Trades, error) {
	out := struct {
		Results Trades `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/ticks/stocks/trades/%s/%s", url.PathEscape(ticker), url.PathEscape(date))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	return &out.Results, err
}

func (c *Client) StockDailyTrades(ticker, date string) ([]*Trades, error) {
	opts := RequestOptions{Limit: 50000}
	var out []*Trades
	for {
		trades, err := c.StockTrades(ticker, date, &opts)
		if err != nil {
			return nil, err
		}
		if len(*trades) <= 1 {
			out = append(out, trades)
			break
		}
		out = append(out, trades)
		opts.Timestamp = (*trades)[len(*trades)-1].SIPTime
	}
	return out, nil
}

func (c *Client) StockQuotes(ticker, date string, opts *RequestOptions) (*Quotes, error) {
	out := struct {
		Results Quotes `json:"results"`
	}{}
	endpoint := fmt.Sprintf("/v2/ticks/stocks/nbbo/%s/%s", url.PathEscape(ticker), url.PathEscape(date))
	endpoint, err := c.endpointWithOpts(endpoint, opts)
	if err != nil {
		return nil, err
	}
	err = c.GetJSON(context.Background(), endpoint, &out)
	return &out.Results, err
}

func (c *Client) StockDailyQuotes(ticker, date string) ([]*Quotes, error) {
	opts := RequestOptions{Limit: 50000}
	var out []*Quotes
	for {
		quotes, err := c.StockQuotes(ticker, date, &opts)
		if err != nil {
			return nil, err
		}
		if len(*quotes) <= 1 {
			out = append(out, quotes)
			break
		}
		fmt.Println(len(*quotes))
		out = append(out, quotes)
		opts.Timestamp = (*quotes)[len(*quotes)-1].SIPTime
	}
	return out, nil
}

func (c *Client) StockLastTrade(ticker string) (LastTrade, error) {
	out := struct {
		Last LastTrade `json:"last"`
	}{}
	endpoint := fmt.Sprintf("/v1/last/stocks/%s", url.PathEscape(ticker))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out.Last, err
}

func (c *Client) StockLastQuote(ticker string) (LastQuote, error) {
	out := struct {
		Last LastQuote `json:"last"`
	}{}
	endpoint := fmt.Sprintf("/v1/last_quote/stocks/%s", url.PathEscape(ticker))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out.Last, err
}

func (c *Client) StockDaily(ticker, date string) (*Daily, error) {
	var out Daily
	endpoint := fmt.Sprintf("/v1/open-close/%s/%s", url.PathEscape(ticker), url.PathEscape(date))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return &out, err
}

func (c *Client) StockConditionMappings(tick Tick) (map[string]string, error) {
	out := make(map[string]string)
	endpoint := fmt.Sprintf("/v1/meta/conditions/%s", url.PathEscape(string(tick)))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return out, err
}

func (c *Client) StockSnapshotAll() (*Snapshots, error) {
	out := struct {
		Tickers Snapshots `json:"tickers"`
	}{}
	endpoint := fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/tickers")
	err := c.GetJSON(context.Background(), endpoint, &out)
	return &out.Tickers, err

}
func (c *Client) StockSnapshotSingle(ticker string) (*Snapshot, error) {
	out := struct {
		Ticker Snapshot `json:"ticker"`
	}{}
	endpoint := fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/tickers/%s", url.PathEscape(ticker))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return &out.Ticker, err

}
func (c *Client) StockSnapshotTopGainersLosers(direction Direction) (*Snapshots, error) {
	out := struct {
		Tickers Snapshots `json:"tickers"`
	}{}
	endpoint := fmt.Sprintf("/v2/snapshot/locale/us/markets/stocks/%s", url.PathEscape(string(direction)))
	err := c.GetJSON(context.Background(), endpoint, &out)
	return &out.Tickers, err
}
