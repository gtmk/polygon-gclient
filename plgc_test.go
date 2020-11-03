package plgc

import (
	"fmt"
	"testing"
)

func TestAPICalls(t *testing.T) {
	client := NewClient("dpQp_57HKzynXlG4crynl2_T5KccWp2A")

	//_, bars, err := client.PreviousClose("AAPL")
	//fmt.Println(fmt.Sprintf("%+v", bars))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//exs, err := client.Exchanges()
	//fmt.Println(fmt.Sprintf("%+v", exs))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//cms, aggs, err := client.Aggregates("AAPL", 1, Minute, "2020-10-05", "2020-10-06")
	//fmt.Println(fmt.Sprintf("%+v", cms))
	//fmt.Println(fmt.Sprintf("%+v", aggs))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//cms, grps, err := client.StockGroupedDaily(US, Stocks, "2020-10-05")
	//fmt.Println(fmt.Sprintf("%+v", cms))
	//fmt.Println(fmt.Sprintf("%+v", grps))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//dls, err := client.Daily("AAPL", "2020-10-05")
	//fmt.Println(fmt.Sprintf("%+v", dls))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//rs, err := client.StockConditionMappings(Trades)
	//fmt.Println(fmt.Sprintf("%+v", rs))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//_, st, it, err := client.ReferenceTickerTypes()
	//fmt.Println(fmt.Sprintf("%+v", st))
	//fmt.Println(fmt.Sprintf("%+v", it))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//td, err := client.ReferenceTickerDetail("AAPL")
	//fmt.Println(fmt.Sprintf("%+v", td))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//tn, err := client.ReferenceTickerNews("AAPL")
	//fmt.Println(fmt.Sprintf("%+v", tn))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//cms, mks, err := client.ReferenceMarkets()
	//fmt.Println(fmt.Sprintf("%+v", cms))
	//fmt.Println(fmt.Sprintf("%+v", mks))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//cms, lcs, err := client.ReferenceLocales()
	//fmt.Println(fmt.Sprintf("%+v", cms))
	//fmt.Println(fmt.Sprintf("%+v", lcs))
	//fmt.Println(fmt.Sprintf("%+v", err))

	//cms, sps, err := client.ReferenceStockSplits("AAPL")
	//fmt.Println(fmt.Sprintf("%+v", cms))
	//fmt.Println(fmt.Sprintf("%+v", sps))
	//fmt.Println(fmt.Sprintf("%+v", err))

	cms, dvs, err := client.ReferenceDividends("AAPL")
	fmt.Println(fmt.Sprintf("%+v", cms))
	fmt.Println(fmt.Sprintf("%+v", dvs))
	fmt.Println(fmt.Sprintf("%+v", err))

}
