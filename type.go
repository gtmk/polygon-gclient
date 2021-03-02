package polygonio

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

type CommonResponse struct {
	Ticker      string `json:"ticker"`
	Status      string `json:"status"`
	Adjusted    bool   `json:"adjusted"`
	QueryCount  int32  `json:"query_count"`
	ResultCount int32  `json:"results_count"`
	Count       int32  `json:"count"`
	Page        int32  `json:"page"`
	PerPage     int32  `json:"perPage"`
}

type Bar struct {
	Ticker string  `json:"T"`
	Time   int64   `json:"t"`
	Volume float32 `json:"v"`
	Open   float32 `json:"o"`
	Close  float32 `json:"c"`
	High   float32 `json:"h"`
	Low    float32 `json:"l"`
	Trades int32   `json:"n"`
	VW     float32 `json:"vw"`
	AV     int64   `json:"av"`
}

//easyjson:json
type Bars []Bar

//easyjson:json
type StockBarsResponse struct {
	Results Bars `json:"results"`
}

type Trade struct {
	ID         int64   `json:"I"`
	Exchange   int32   `json:"x"` // exchange id
	Price      float64 `json:"p"` // price
	TradeID    string  `json:"i"` // trade id, uniquely identify the trade
	CorrID     int32   `json:"e"` // trade correction indicator
	ReportID   int32   `json:"r"` // report id
	ExTime     int64   `json:"y"` // participant timestamp
	SIPTime    int64   `json:"t"` // sip timestamp
	TRFTime    int64   `json:"f"` // trade report timestamp
	Conditions []int32 `json:"c"` // conditions
	Sequence   int32   `json:"q"` // sequence number in order
	Size       int32   `json:"s"` // size of the trade
	ListedEx   int32   `json:"z"` // listed exchange
}

type LastTrade struct {
	Condition1 int32   `json:"cond1"`
	Exchange   int32   `json:"exchange"`
	Price      float64 `json:"float64"`
	Size       int32   `json:"size"`
	Timestamp  int64   `json:"timestamp"`
}

type Trades []Trade

//easyjson:json
type StockTradesResponse struct {
	Results Trades `json:"results"`
}

type Quote struct {
	ExTime      int64   `json:"y"` // participant timestamp
	SIPTime     int64   `json:"t"` // sip timestamp
	TRFTime     int64   `json:"f"` // trade report timestamp
	Sequence    int32   `json:"q"` // sequence number in order
	Conditions  []int32 `json:"c"` // conditions
	Indicators  []int32 `json:"i"` // indicators
	BidPrice    float64 `json:"p"`
	BidExchange int32   `json:"x"`
	BidSize     int32   `json:"s"`
	AskPrice    float64 `json:"P"`
	AskExchange int32   `json:"X"`
	AskSize     int32   `json:"S"`
	ListedEx    int32   `json:"z"` // listed exchange
}

type LastQuote struct {
	BidPrice    float64 `json:"bidprice"`
	BidExchange int32   `json:"bidexchange"`
	BidSize     int32   `json:"bidsize"`
	AskPrice    float64 `json:"askprice"`
	AskExchange int32   `json:"askexchange"`
	AskSize     int32   `json:"asksize"`
	Timestamp   int64   `json:"timestamp"`
}

type Quotes []Quote

//easyjson:json
type StockQuotesResponse struct {
	Results Quotes `json:"results"`
}

type Reverse string

const (
	ReserveTrue  Reverse = "true"
	ReserveFalse Reverse = "false"
)

type Snapshot struct {
	Ticker         string    `json:"ticker"`
	TodayChange    float32   `json:"todaysChange"`
	TodayChangePct float32   `json:"todaysChangePerc"`
	Day            Bar       `json:"day"`
	PrevDay        Bar       `json:"prevDay"`
	LastQuote      LastQuote `json:"lastQuote"`
	LastTrade      LastTrade `json:"lastTrade"`
	Min            Bar       `json:"min"`
	Updated        int64     `json:"updated"`
}

type Snapshots []Snapshot

//easyjson:json
type StockSnapshotsResponse struct {
	Results Snapshots `json:"tickers"`
}

type Sort string

const (
	Asc  Sort = "asc"
	Desc Sort = "desc"
)

type Unadjusted string

const (
	UnadjustedTrue  Unadjusted = "true"
	UnadjustedFalse Unadjusted = "false"
)

type RequestOptions struct {
	Unadjusted     Unadjusted `url:"unadjusted,omitempty"`
	Sort           Sort       `url:"sort,omitempty"`
	Timestamp      int64      `url:"timestamp,omitempty"`
	TimestampLimit int64      `url:"timestampLimit,omitempty"`
	Reverse        Reverse    `url:"reverse,omitempty"`
	Limit          int64      `url:"limit,omitempty"`
}

type TickerSort string

const (
	AZTicker TickerSort = "ticker"
	ZATicker TickerSort = "-ticker"
)

type TickerOptions struct {
	Sort    TickerSort `url:"sort,omitempty"`
	Type    string     `url:"type,omitempty"`
	Market  Market     `url:"market,omitempty"`
	Locale  Locale     `url:"locale,omitempty"`
	Search  string     `url:"search,omitempty"`
	PerPage int32      `url:"perpage,omitempty"`
	Page    int32      `url:"page,omitempty"`
	Active  bool       `url:"active,omitempty"`
}

type Ticker struct {
	Ticker      string    `json:"ticker"`
	Name        string    `json:"name"`
	Market      string    `json:"market"`
	Locale      Locale    `json:"locale"`
	Type        string    `json:"type"`
	Currency    string    `json:"currency"`
	Active      bool      `json:"active"`
	PrimaryExch string    `json:"primaryExch"`
	Updated     string    `json:"updated"`
	Codes       *CodesMap `json:"codes,omitempty"`
	Attrs       *AttrsMap `json:"attrs,omitempty"`
	URL         string    `json:"url"`
}

type Tickers []Ticker

type CodesMap map[string]string
type AttrsMap map[string]interface{}

func (cm CodesMap) Value() (driver.Value, error) {
	return json.Marshal(cm)
}

func (am AttrsMap) Value() (driver.Value, error) {
	return json.Marshal(am)
}

func (cm *CodesMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &cm)
}

func (am *AttrsMap) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(b, &am)
}

type NewsOptions struct {
	PerPage int32 `url:"perpage,omitempty"`
	Page    int32 `url:"page,omitempty"`
}

type FinancialOptionSort string

const (
	ReportPeriod        FinancialOptionSort = "reportPeriod"
	ReverseReportPeriod FinancialOptionSort = "-reportPeriod"
	CalendarDate        FinancialOptionSort = "calendarDate"
	ReverseCalendarDate FinancialOptionSort = "-calendarDate"
)

type FinancialOptionType string

const (
	Y  FinancialOptionType = "Y"
	YA FinancialOptionType = "YA"
	Q  FinancialOptionType = "Q"
	QA FinancialOptionType = "QA"
	T  FinancialOptionType = "T"
	TA FinancialOptionType = "TA"
)

type FinancialOptions struct {
	Limit int32               `url:"limit,omitempty"`
	Type  FinancialOptionType `url:"type,omitempty"`
	Sort  FinancialOptionSort `url:"sort,omitempty"`
}

type Exchange struct {
	ID     int32  `json:"id"`
	Type   string `json:"type"`
	Market string `json:"market"`
	Mic    string `json:"mic"`
	Name   string `json:"name"`
	Tape   string `json:"tape"`
}

type Exchanges []Exchange

type Timespan string

const (
	Minute Timespan = "minute"
	Hour   Timespan = "hour"
	Day    Timespan = "day"
	Week   Timespan = "week"
	Month  Timespan = "month"
	Quater Timespan = "quater"
	Year   Timespan = "year"
)

type Market string

const (
	Stocks  Market = "stocks"
	Crypto  Market = "crypto"
	Bond    Market = "mf"
	MMF     Market = "mmf"
	Indices Market = "indicies"
	FX      Market = "fx"
)

type Locale string

const (
	G  Locale = "global"
	US Locale = "us"
	GB Locale = "gb"
	CA Locale = "ca"
	NL Locale = "nl"
	GR Locale = "gr"
	SP Locale = "sp"
	DE Locale = "de"
	PE Locale = "pe"
	DK Locale = "dk"
	FI Locale = "fi"
	IE Locale = "ie"
	PT Locale = "pt"
	IN Locale = "in"
	MX Locale = "mx"
	FR Locale = "fr"
	CN Locale = "cn"
	CH Locale = "ch"
	SE Locale = "se"
)

type Tick string

const (
	TradesName Tick = "trades"
	QuotesName Tick = "quotes"
)

type Direction string

const (
	Gainers Direction = "gainers"
	Losers  Direction = "losers"
)

type Daily struct {
	Status     string  `json:"status"`
	From       string  `json:"from"`
	Ticker     string  `json:"symbol"`
	Volume     float32 `json:"volume"`
	Open       float32 `json:"open"`
	Close      float32 `json:"close"`
	High       float32 `json:"high"`
	Low        float32 `json:"low"`
	PreMarket  float32 `json:"preMarket"`
	AfterHours float32 `json:"afterHours"`
}

type TickerDetails struct {
	Logo           string   `json:"logo"`
	ListDate       string   `json:"listdate"`
	CIK            string   `json:"cik"`
	Bloomberg      string   `json:"bloomberg"`
	FIGI           string   `json:"figi"`
	LEI            string   `json:"lei"`
	SIC            int32    `json:"sic"`
	Country        string   `json:"country"`
	Industry       string   `json:"industry"`
	Sector         string   `json:"sector"`
	MarketCap      int64    `json:"marketcap"`
	Employees      int64    `json:"employees"`
	Phone          string   `json:"phone"`
	CEO            string   `json:"ceo"`
	URL            string   `json:"url"`
	Description    string   `json:"description"`
	Exchange       string   `json:"exchange"`
	Name           string   `json:"name"`
	Symbol         string   `json:"symbol"`
	ExchangeSymbol string   `json:"exchangeSymbol"`
	HQAddress      string   `json:"hq_address"`
	HQState        string   `json:"hq_state"`
	HQCountry      string   `json:"hq_country"`
	Type           string   `json:"type"`
	Updated        string   `json:"updated"`
	Tags           []string `json:"tags"`
	Similar        []string `json:"similar"`
	Active         bool     `json:"active"`
}

type TickerNews struct {
	Symbols   []string  `json:"symbols"`
	Title     string    `json:"title"`
	URL       string    `json:"url"`
	Source    string    `json:"source"`
	Summary   string    `json:"summary"`
	Image     string    `json:"image"`
	Timestamp time.Time `json:"timestamp"`
	Keywords  []string  `json:"keywords"`
}

type MarketDescription struct {
	Name        Market `json:"market"`
	Description string `json:"desc"`
}

type MarketDescriptions []MarketDescription

type LocaleName struct {
	Locale Locale `json:"locale"`
	Name   string `json:"name"`
}

type LocaleNames []LocaleName

type Split struct {
	Ticker        string  `json:"ticker"`
	ExDate        string  `json:"exDate"`
	PaymentDate   string  `json:"paymentDate"`
	RecorDate     string  `json:"recordDate"`
	DeclearedDate string  `json:"declaredDate"`
	Ratio         float32 `json:"ratio"`
	ToFactor      int32   `json:"tofactor"`
	ForFactor     int32   `json:"forfactor"`
}

type Splits []Split

type Dividend struct {
	Ticker        string  `json:"ticker"`
	Type          string  `json:"type"`
	ExDate        string  `json:"exDate"`
	PaymentDate   string  `json:"paymentDate"`
	RecorDate     string  `json:"recordDate"`
	DeclearedDate string  `json:"declaredDate"`
	Amount        float32 `json:"amount"`
	Qualified     string  `json:"qualified"`
	Flag          string  `json:"flag"`
}

type Dividends []Dividend

type MarketStatus struct {
	Market     string            `json:"market"`
	ServerTime string            `json:"serverTime"`
	Exchanges  map[string]string `json:"exchanges"`
	Currencies map[string]string `json:"currencies"`
}

type MarketHoliday struct {
	Exchange string `json:"exchange"`
	Name     string `json:"name"`
	Status   string `json:"status"`
	Date     string `json:"date"`
	Open     string `json:"open"`
	Close    string `json:"close"`
}

type MarketHolidays []MarketHoliday

type Financial struct {
	Ticker                                                 string  `json:"ticker"`
	Period                                                 string  `json:"period"`
	CalendarDate                                           string  `json:"calendarDate"`
	ReportPeriod                                           string  `json:"reportPeriod"`
	Updated                                                string  `json:"updated"`
	AccumulatedOtherComprehensiveIncom                     float32 `json:"accumulatedOtherComprehensiveIncome"`
	Asset                                                  float32 `json:"assets"`
	AssetAverage                                           float32 `json:"assetsAverage"`
	AssetCurrent                                           float32 `json:"assetsCurrent"`
	AssetTurnOver                                          float32 `json:"assetTurnover"`
	AssetNonCurrent                                        float32 `json:"assetsNonCurrent"`
	BookValuePerShare                                      float32 `json:"bookValuePerShare"`
	CapitalExpenditure                                     float32 `json:"capitalExpenditure"`
	CashAndEquivalents                                     float32 `json:"cashAndEquivalents"`
	CashAndEquivalentsUSD                                  float32 `json:"cashAndEquivalentsUSD"`
	CostOfRevenue                                          float32 `json:"costOfRevenue"`
	ConsolidatedIncome                                     float32 `json:"consolidatedIncome"`
	CurrentRatio                                           float32 `json:"currentRatio"`
	DebtToEquityRatio                                      float32 `json:"debtToEquityRatio"`
	Debt                                                   float32 `json:"debt"`
	DebtCurrent                                            float32 `json:"debtCurrent"`
	DebtNonCurrent                                         float32 `json:"debtNonCurrent"`
	DebtUSD                                                float32 `json:"debtUSD"`
	DeferredRevenue                                        float32 `json:"deferredRevenue"`
	DepreciationAmortizationAndAccretion                   float32 `json:"depreciationAmortizationAndAccretion"`
	Deposits                                               float32 `json:"deposits"`
	DividentdYield                                         float32 `json:"dividendYield"`
	DividendsPerBasicCommonShare                           float32 `json:"dividendsPerBasicCommonShare"`
	EarningBeforeInterestTaxes                             float32 `json:"earningBeforeInterestTaxes"`
	EarningsBeforeInterestTaxesDepreciationAmortization    float32 `json:"earningsBeforeInterestTaxesDepreciationAmortization"`
	EBITDAMargin                                           float32 `json:"EBITDAMargin"`
	EarningsBeforeInterestTaxesDepreciationAmortizationUSD float32 `json:"earningsBeforeInterestTaxesDepreciationAmortizationUSD"`
	EarningBeforeInterestTaxesUSD                          float32 `json:"earningBeforeInterestTaxesUSD"`
	EarningsBeforeTax                                      float32 `json:"earningsBeforeTax"`
	EarningsPerBasicShare                                  float32 `json:"earningsPerBasicShare"`
	EarningsPerDilutedShare                                float32 `json:"earningsPerDilutedShare"`
	EarningsPerBasicShareUSD                               float32 `json:"earningsPerBasicShareUSD"`
	ShareholdersEquity                                     float32 `json:"shareholdersEquity"`
	EverageEquity                                          float32 `json:"everageEquity"`
	ShareholdersEquityUSD                                  float32 `json:"shareholdersEquityUSD"`
	EnterpriseValue                                        float32 `json:"enterpriseValue"`
	EnterpriseValueOverEBIT                                float32 `json:"enterpriseValueOverEBIT"`
	EnterpriseValueOverEBITDA                              float32 `json:"enterpriseValueOverEBITDA"`
	FreeCashFlow                                           float32 `json:"freeCashFlow"`
	FreeCashFlowPerShare                                   float32 `json:"freeCashFlowPerShare"`
	ForeignCurrencyUSDExchangeRate                         float32 `json:"foreignCurrencyUSDExchangeRate"`
	GrossProfit                                            float32 `json:"grossProfit"`
	GrossMargin                                            float32 `json:"grossMargin"`
	GoodwillAndIntangibleAssets                            float32 `json:"goodwillAndIntangibleAssets"`
	InterestExpense                                        float32 `json:"interestExpense"`
	InvestedCapital                                        float32 `json:"investedCapital"`
	InvestedCapitalAverage                                 float32 `json:"investedCapitalAverage"`
	Inventory                                              float32 `json:"inventory"`
	Investments                                            float32 `json:"investments"`
	InvestmentsCurrent                                     float32 `json:"investmentsCurrent"`
	InvestmentsNonCurrent                                  float32 `json:"investmentsNonCurrent"`
	TotalLiabilities                                       float32 `json:"totalLiabilities"`
	CurrentLiabilities                                     float32 `json:"currentLiabilities"`
	LiabilitiesNonCurrent                                  float32 `json:"liabilitiesNonCurrent"`
	MarketCapitalization                                   float32 `json:"marketCapitalization"`
	NetCashFlow                                            float32 `json:"netCashFlow"`
	NetCashFlowBusinessAcquisitionsDisposals               float32 `json:"netCashFlowBusinessAcquisitionsDisposals"`
	IssuanceEquityShares                                   float32 `json:"issuanceEquityShares"`
	IssuanceDebtSecurities                                 float32 `json:"issuanceDebtSecurities"`
	PaymentDividendsOtherCashDistributions                 float32 `json:"paymentDividendsOtherCashDistributions"`
	NetCashFlowFromFinancing                               float32 `json:"netCashFlowFromFinancing"`
	NetCashFlowFromInvesting                               float32 `json:"netCashFlowFromInvesting"`
	NetCashFlowInvestmentAcquisitionsDisposals             float32 `json:"netCashFlowInvestmentAcquisitionsDisposals"`
	NetCashFlowFromOperations                              float32 `json:"netCashFlowFromOperations"`
	EffectOfExchangeRateChangesOnCash                      float32 `json:"effectOfExchangeRateChangesOnCash"`
	NetIncome                                              float32 `json:"netIncome"`
	NetIncomeCommonStock                                   float32 `json:"netIncomeCommonStock"`
	NetIncomeCommonStockUSD                                float32 `json:"netIncomeCommonStockUSD"`
	NetLossIncomeFromDiscontinuedOperations                float32 `json:"netLossIncomeFromDiscontinuedOperations"`
	NetIncomeToNonControllingInterests                     float32 `json:"netIncomeToNonControllingInterests"`
	ProfitMargin                                           float32 `json:"profitMargin"`
	OperatingExpenses                                      float32 `json:"operatingExpenses"`
	OperatingIncome                                        float32 `json:"operatingIncome"`
	TradeAndNonTradePayables                               float32 `json:"tradeAndNonTradePayables"`
	PayoutRatio                                            float32 `json:"payoutRatio"`
	PriceToBookValue                                       float32 `json:"priceToBookValue"`
	PriceEarnings                                          float32 `json:"priceEarnings"`
	PriceToEarningsRatio                                   float32 `json:"priceToEarningsRatio"`
	PropertyPlantEquipmentNet                              float32 `json:"propertyPlantEquipmentNet"`
	PreferredDividendsIncomeStatementImpact                float32 `json:"preferredDividendsIncomeStatementImpact"`
	SharePriceAdjustedClose                                float32 `json:"sharePriceAdjustedClose"`
	PriceSales                                             float32 `json:"priceSales"`
	PriceToSalesRatio                                      float32 `json:"priceToSalesRatio"`
	TradeAndNonTradeReceivables                            float32 `json:"tradeAndNonTradeReceivables"`
	AccumulatedRetainedEarningsDeficit                     float32 `json:"accumulatedRetainedEarningsDeficit"`
	Revenues                                               float32 `json:"revenues"`
	RevenuesUSD                                            float32 `json:"revenuesUSD"`
	ResearchAndDevelopmentExpense                          float32 `json:"researchAndDevelopmentExpense"`
	ReturnOnAverageAssets                                  float32 `json:"returnOnAverageAssets"`
	ReturnOnAverageEquity                                  float32 `json:"returnOnAverageEquity"`
	ReturnOnInvestedCapital                                float32 `json:"returnOnInvestedCapital"`
	ReturnOnSales                                          float32 `json:"returnOnSales"`
	ShareBasedCompensation                                 float32 `json:"shareBasedCompensation"`
	SellingGeneralAndAdministrativeExpense                 float32 `json:"sellingGeneralAndAdministrativeExpense"`
	ShareFactor                                            float32 `json:"shareFactor"`
	Shares                                                 float32 `json:"shares"`
	WeightedAverageShares                                  float32 `json:"weightedAverageShares"`
	WeightedAverageSharesDiluted                           float32 `json:"weightedAverageSharesDiluted"`
	SalesPerShare                                          float32 `json:"salesPerShare"`
	TangibleAssetValue                                     float32 `json:"tangibleAssetValue"`
	TaxAssets                                              float32 `json:"taxAssets"`
	IncomeTaxExpense                                       float32 `json:"incomeTaxExpense"`
	TaxLiabilities                                         float32 `json:"taxLiabilities"`
	TangibleAssetsBookValuePerShare                        float32 `json:"tangibleAssetsBookValuePerShare"`
	WorkingCapital                                         float32 `json:"workingCapital"`
}

type Financials []Financial

type CryptoTrade struct {
	Price      float32 `json:"p"`
	Size       float32 `json:"s"`
	Exchange   int32   `json:"x"`
	Time       int64   `json:"t"`
	Conditions []int32 `json:"c"`
}

type CryptoTrades []CryptoTrade
type CryptoDaily struct {
	Ticker        string       `json:"symbol"`
	IsUTC         bool         `json:"isUTC"`
	Day           string       `json:"day"`
	Open          float32      `json:"open"`
	Close         float32      `json:"close"`
	OpenTrades    CryptoTrades `json:"openTrades"`
	ClosingTrades CryptoTrades `json:"closingTrades"`
}

// PolygonClientMsg is the standard message sent by clients of the stream interface
type PolygonClientMsg struct {
	Action string `json:"action"`
	Params string `json:"params"`
}

type PolygonAuthMsg struct {
	Event   string `json:"ev"`
	Status  string `json:"status"`
	Message string `json:"message"`
}

type StreamingServerMsg struct {
	Event             string      `json:"ev"`
	Symbol            string      `json:"sym"`
	Exchange          int32       `json:"x"`
	TradeID           string      `json:"i"`
	Price             float32     `json:"p"`
	S                 int64       `json:"s"`
	C                 interface{} `json:"c"`
	Timestamp         int64       `json:"t"`
	Unknown           int64       `json:"z"`
	BidExchange       int32       `json:"bx"`
	AskExchange       int32       `json:"ax"`
	BidPrice          float32     `json:"bp"`
	AskPrice          float32     `json:"ap"`
	BidSize           int32       `json:"bs"`
	AskSize           int32       `json:"as"`
	Volume            int32       `json:"v"`
	AccumulatedVolume int64       `json:"av"`
	OpeningPrice      float32     `json:"op"`
	VWAP              float32     `json:"vw"`
	OpenPrice         float32     `json:"o"`
	HighPrice         float32     `json:"h"`
	LowPrice          float32     `json:"l"`
	Average           float32     `json:"a"`
	EndTimestamp      int64       `json:"e"`
}

//easyjson:json
type StreamingServerMsges []StreamingServerMsg

// StreamTrade is the structure that defines a trade that
// polygon transmits via websocket protocol.
type StreamTrade struct {
	Event      string  `json:"ev"`
	Symbol     string  `json:"sym"`
	Exchange   int32   `json:"x"`
	TradeID    string  `json:"i"`
	Price      float32 `json:"p"`
	Size       int32   `json:"s"`
	Timestamp  int64   `json:"t"`
	Conditions []int32 `json:"c"`
	Unknown    int32   `json:"z"`
}

//easyjson:json
type StreamTrades []StreamTrade

// StreamQuote is the structure that defines a quote that
// polygon transmits via websocket protocol.
type StreamQuote struct {
	Event       string  `json:"ev"`
	Symbol      string  `json:"sym"`
	Condition   int32   `json:"c"`
	BidExchange int32   `json:"bx"`
	AskExchange int32   `json:"ax"`
	BidPrice    float32 `json:"bp"`
	AskPrice    float32 `json:"ap"`
	BidSize     int32   `json:"bs"`
	AskSize     int32   `json:"as"`
	Timestamp   int64   `json:"t"`
	Unknown     int32   `json:"z"`
}

//easyjson:json
type StreamQuotes []StreamQuote

// StreamAggregate is the structure that defines an aggregate that
// polygon transmits via websocket protocol.
type StreamAggregate struct {
	Event             string  `json:"ev"`
	Symbol            string  `json:"sym"`
	Volume            int32   `json:"v"`
	AccumulatedVolume int64   `json:"av"`
	OpeningPrice      float32 `json:"op"`
	VWAP              float32 `json:"vw"`
	OpenPrice         float32 `json:"o"`
	ClosePrice        float32 `json:"c"`
	HighPrice         float32 `json:"h"`
	LowPrice          float32 `json:"l"`
	Average           float32 `json:"a"`
	TotalTrade        int32   `jdon:"z"`
	StartTimestamp    int64   `json:"s"`
	EndTimestamp      int64   `json:"e"`
}

//easyjson:json
type StreamAggregates []StreamAggregate
