package polygonio

import (
	"fmt"
	"testing"
	"time"
)

func TestAPICalls(t *testing.T) {

	//stream test
	{
		stream, err := GetStream("API_KEY", "wss://socket.polygon.io/stocks")
		if err != nil {
			fmt.Println("error on getting the stream: ", err)
		}
		go func() {
			for msg := range stream.messageC {
				fmt.Println("message: ", string(msg))
			}
		}()
		go func() {
			for msg := range stream.errorC {
				fmt.Println("message: ", msg.Error())
			}
		}()

		if err := stream.Subscribe(`A.*`); err != nil {
			fmt.Println("error on subscribing: ", err)
		}
		if err := stream.Subscribe(`T.*`); err != nil {
			fmt.Println("error on subscribing: ", err)
		}

		if err := stream.Unsubscribe(`A.*`); err != nil {
			fmt.Println("error on unsubscribing: ", err)
		}
		if err := stream.Unsubscribe(`T.*`); err != nil {
			fmt.Println("error on unsubscribing: ", err)
		}

		time.Sleep(time.Duration(5 * time.Second))

		if err := stream.Close(); err != nil {
			fmt.Println("error on closing: ", err)
		}
		fmt.Println("close done")

		if err := stream.Subscribe(`T.*`); err != nil {
			fmt.Println("error on subscribing: ", err)
		}
		for {
			time.Sleep(time.Duration(5) * time.Second)
		}
	}

	//rest test
	{
		client := NewClient("API_KEY")
		//Reference Endpoints
		tks, err := client.ReferenceTickers(&TickerOptions{Sort: ZATicker, Market: Stocks})
		fmt.Println(fmt.Sprintf("%+v", tks))
		fmt.Println(fmt.Sprintf("%+v", err))

		st, it, err := client.ReferenceTickerTypes()
		fmt.Println(fmt.Sprintf("%+v", st))
		fmt.Println(fmt.Sprintf("%+v", it))
		fmt.Println(fmt.Sprintf("%+v", err))

		td, err := client.ReferenceTickerDetail("AAPL")
		fmt.Println(fmt.Sprintf("%+v", td))
		fmt.Println(fmt.Sprintf("%+v", err))

		tn, err := client.ReferenceTickerNews("AAPL", &NewsOptions{PerPage: 10})
		fmt.Println(fmt.Sprintf("%+v", tn))
		fmt.Println(fmt.Sprintf("%+v", err))

		mks, err := client.ReferenceMarkets()
		fmt.Println(fmt.Sprintf("%+v", mks))
		fmt.Println(fmt.Sprintf("%+v", err))

		lcs, err := client.ReferenceLocales()
		fmt.Println(fmt.Sprintf("%+v", lcs))
		fmt.Println(fmt.Sprintf("%+v", err))

		sps, err := client.ReferenceStockSplits("AAPL")
		fmt.Println(fmt.Sprintf("%+v", sps))
		fmt.Println(fmt.Sprintf("%+v", err))

		fns, err := client.ReferenceFinancials("AAPL", &FinancialOptions{Limit: 10, Type: Y})
		fmt.Println(fmt.Sprintf("%+v", fns))
		fmt.Println(fmt.Sprintf("%+v", err))

		dvs, err := client.ReferenceDividends("AAPL")
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

		bars, err := client.StockPreviousClose("AAPL", &RequestOptions{Unadjusted: UnadjustedFalse})
		fmt.Println(fmt.Sprintf("%+v", bars))
		fmt.Println(fmt.Sprintf("%+v", err))

		aggs, err := client.StockAggregates("AAPL", 1, Minute, "2021-01-04", "2021-01-05", &RequestOptions{Unadjusted: UnadjustedFalse, Sort: Asc})
		fmt.Println(fmt.Sprintf("%+v", aggs))
		fmt.Println(fmt.Sprintf("%+v", err))
		for _, agg := range *aggs {
			fmt.Println(fmt.Sprintf("%+v", agg))
		}

		trades, err := client.StockDailyQuotes("AMD", "2020-10-05", nil)
		fmt.Println(fmt.Sprintf("%+v", len(trades)))
		fmt.Println(fmt.Sprintf("%+v", err))

		grps, err := client.StockGroupedDaily(US, Stocks, "2020-10-05", nil)
		fmt.Println(fmt.Sprintf("%+v", grps))
		fmt.Println(fmt.Sprintf("%+v", err))

		dls, err := client.StockDaily("AAPL", "2020-10-05")
		fmt.Println(fmt.Sprintf("%+v", dls))
		fmt.Println(fmt.Sprintf("%+v", err))

		rs, err := client.StockConditionMappings(Trades)
		fmt.Println(fmt.Sprintf("%+v", rs))
		fmt.Println(fmt.Sprintf("%+v", err))

		opts := RequestOptions{Limit: 100}

		cms, tds, err := client.StockHistoricQuotes("AAPL", "2020-10-14", &opts)
		fmt.Println(fmt.Sprintf("%+v", cms))
		fmt.Println(fmt.Sprintf("%+v", len(tds)))
		fmt.Println(fmt.Sprintf("%+v", err))

		last, err := client.StockLastQuote("AAPL")
		fmt.Println(fmt.Sprintf("%+v", last))
		fmt.Println(fmt.Sprintf("%+v", err))

		sps, err := client.StockSnapshotAll()
		fmt.Println(fmt.Sprintf("%+v", sps))
		fmt.Println(fmt.Sprintf("%+v", err))
		for i, e := range *sps {
			fmt.Println(fmt.Sprintf("%s, %d, %+v, %+v", e.Ticker, i, e.PrevDay, e.Min))
		}
		sps, err := client.StockSnapshotSingle("AAPL")
		fmt.Println(fmt.Sprintf("%+v", sps))
		fmt.Println(fmt.Sprintf("%+v", err))

		gappers, err := client.StockSnapshotTopGainersLosers(Gainers)
		fmt.Println(fmt.Sprintf("%+v", gappers))
		fmt.Println(fmt.Sprintf("%+v", err))

		// Forex endpoints
		fpc, err := client.ForexPreviousClose("C:EURUSD", nil)
		fmt.Println(fmt.Sprintf("%+v", fpc))
		fmt.Println(fmt.Sprintf("%+v", err))

		faggs, err := client.ForexAggregates("C:EURUSD", 1, Minute, "2020-10-05", "2020-10-06", &RequestOptions{Sort: Asc})
		fmt.Println(fmt.Sprintf("%+v", faggs))
		fmt.Println(fmt.Sprintf("%+v", err))

		fd, err := client.ForexGroupedDaily(US, "2020-10-05", nil)
		fmt.Println(fmt.Sprintf("%+v", fd))
		fmt.Println(fmt.Sprintf("%+v", err))

		//Crypto Endpoints
		exs, err = client.CryptoExchanges()
		fmt.Println(fmt.Sprintf("%+v", exs))
		fmt.Println(fmt.Sprintf("%+v", err))

		cpc, err := client.CryptoPreviousClose("X:ETH", nil)
		fmt.Println(fmt.Sprintf("%+v", cpc))
		fmt.Println(fmt.Sprintf("%+v", err))

		caggs, err := client.CryptoAggregates("X:ETHUSDT", 1, Minute, "2020-10-05", "2020-10-06", &RequestOptions{Sort: Asc})
		fmt.Println(fmt.Sprintf("%+v", caggs))
		fmt.Println(fmt.Sprintf("%+v", err))

		cd, err := client.CryptoGroupedDaily(US, "2020-10-05", nil)
		fmt.Println(fmt.Sprintf("%+v", cd))
		fmt.Println(fmt.Sprintf("%+v", err))
	}
}
