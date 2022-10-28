package spot

import (
	"errors"
	"fmt"
	. "github.com/nntaoli-project/goex/v2"
	"net/url"
)

func (s spotImpl) GetName() string {
	return "huobi.com"
}

func (s spotImpl) GetDepth(pair CurrencyPair, limit int, opt ...OptionParameter) (*Depth, error) {
	//TODO implement me
	panic("implement me")
}

func (s spotImpl) GetTicker(pair CurrencyPair, opt ...OptionParameter) (*Ticker, error) {
	cli := GetHttpCli()
	params := url.Values{}
	params.Set("symbol", pair.Symbol)

	data, err := cli.DoRequest("GET", s.uriOpts.Endpoint+s.uriOpts.TickerUri, &params, nil)
	if err != nil {
		return nil, fmt.Errorf("%w%s", err, errors.New(string(data)))
	}

	tk, err := s.unmarshalerOpts.TickerUnmarshaler(data)
	if err != nil {
		return nil, err
	}

	tk.Pair = pair
	tk.Origin = data

	return tk, err
}

func (s spotImpl) GetKline(pair CurrencyPair, period KlinePeriod, opt ...OptionParameter) ([]Kline, error) {
	//TODO implement me
	panic("implement me")
}
