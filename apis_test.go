package polygonio

import (
	"fmt"
	"testing"
)

func TestAPICalls(t *testing.T) {
	client := NewClient("api_key")

	//Reference Endpoints
	cms, tks, err := client.ReferenceTickers(&TickerOptions{Sort: ZATicker, Market: Stocks})
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", tks))
	fmt.Println(fmt.Sprintf("%+v", err))

	_, st, it, err := client.ReferenceTickerTypes()
	fmt.Println(fmt.Sprintf("%+v", st))
	fmt.Println(fmt.Sprintf("%+v", it))
	fmt.Println(fmt.Sprintf("%+v", err))

	td, err := client.ReferenceTickerDetail("AAPL")
	fmt.Println(fmt.Sprintf("%+v", td))
	fmt.Println(fmt.Sprintf("%+v", err))

	tn, err := client.ReferenceTickerNews("AAPL", &NewsOptions{PerPage: 10})
	fmt.Println(fmt.Sprintf("%+v", tn))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, mks, err := client.ReferenceMarkets()
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", mks))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, lcs, err := client.ReferenceLocales()
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", lcs))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, sps, err := client.ReferenceStockSplits("AAPL")
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", sps))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, fns, err := client.ReferenceFinancials("AAPL", &FinancialOptions{Limit: 10, Type: Y})
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", fns))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, dvs, err := client.ReferenceDividends("AAPL")
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", dvs))
	fmt.Println(fmt.Sprintf("%+v", err))

	ms, err := client.ReferenceMarketStatus()
	fmt.Println(fmt.Sprintf("%+v", ms))
	fmt.Println(fmt.Sprintf("%+v", err))

	mh, err := client.ReferenceMarketHolidays()
	fmt.Println(fmt.Sprintf("%+v", mh))
	fmt.Println(fmt.Sprintf("%+v", err))

	// Stock Endpoints
	exs, err := client.StockExchanges()
	fmt.Println(fmt.Sprintf("%+v", exs))
	fmt.Println(fmt.Sprintf("%+v", err))

	_, bars, err := client.StockPreviousClose("AAPL", &RequestOptions{Unadjusted: UnadjustedFalse})
	fmt.Println(fmt.Sprintf("%+v", bars))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, aggs, err := client.StockAggregates("AAPL", 1, Minute, "2020-10-05", "2020-10-06", &RequestOptions{Unadjusted: UnadjustedFalse, Sort: Asc})
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", aggs))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, grps, err := client.StockGroupedDaily(US, Stocks, "2020-10-05", nil)
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", grps))
	fmt.Println(fmt.Sprintf("%+v", err))

	dls, err := client.StockDaily("AAPL", "2020-10-05")
	fmt.Println(fmt.Sprintf("%+v", dls))
	fmt.Println(fmt.Sprintf("%+v", err))

	rs, err := client.StockConditionMappings(Trades)
	fmt.Println(fmt.Sprintf("%+v", rs))
	fmt.Println(fmt.Sprintf("%+v", err))

	// Forex endpoints
	cms, fpc, err := client.ForexPreviousClose("C:EURUSD", nil)
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", fpc))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, faggs, err := client.ForexAggregates("C:EURUSD", 1, Minute, "2020-10-05", "2020-10-06", &RequestOptions{Sort: Asc})
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", faggs))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, fd, err := client.ForexGroupedDaily(US, "2020-10-05", nil)
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", fd))
	fmt.Println(fmt.Sprintf("%+v", err))

	//Crypto Endpoints
	exs, err = client.CryptoExchanges()
	fmt.Println(fmt.Sprintf("%+v", exs))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, cpc, err := client.CryptoPreviousClose("X:ETH", nil)
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", cpc))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, caggs, err := client.CryptoAggregates("X:ETHUSDT", 1, Minute, "2020-10-05", "2020-10-06", &RequestOptions{Sort: Asc})
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", caggs))
	fmt.Println(fmt.Sprintf("%+v", err))

	cms, cd, err := client.CryptoGroupedDaily(US, "2020-10-05", nil)
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", cd))
	fmt.Println(fmt.Sprintf("%+v", err))

}
