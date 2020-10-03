package decimal

import (
	"math/rand"
	"time"

	"github.com/shopspring/decimal"
)

type Decimal struct {
	decimal.Decimal
}

func New(value int64) Decimal {
	return Decimal{
		Decimal: decimal.New(value, 0),
	}
}

func Seed() {
	rand.Seed(time.Now().UnixNano())
}

func NewRandom(min float64, max float64) Decimal {
	return NewFromFloat(min + rand.Float64()*(max-min))
}

func NewFromFloat(value float64) Decimal {
	return Decimal{
		Decimal: decimal.NewFromFloat(value),
	}
}

func NewFromDecimal(value decimal.Decimal) Decimal {
	return Decimal{
		Decimal: value,
	}
}

func NewFromString(str string) (Decimal, error) {
	d, err := decimal.NewFromString(str)
	return Decimal{
		Decimal: d,
	}, err
}

func (d Decimal) Add(add Decimal) Decimal {
	return NewFromDecimal(d.Decimal.Add(add.Decimal))
}

func (d Decimal) Sub(sub Decimal) Decimal {
	return NewFromDecimal(d.Decimal.Sub(sub.Decimal))
}

func (d Decimal) Mul(mul Decimal) Decimal {
	return NewFromDecimal(d.Decimal.Mul(mul.Decimal))
}

func (d Decimal) Div(div Decimal) Decimal {
	return NewFromDecimal(d.Decimal.Div(div.Decimal))
}

func (d Decimal) Shift(shift int32) Decimal {
	return NewFromDecimal(d.Decimal.Shift(shift))
}

func (d Decimal) Miser(precision int32) Decimal {
	return New(1).Shift(-precision)
}

func (d Decimal) AddMiser(precision int32) Decimal {
	return d.Add(New(1).Shift(-precision))
}

func (d Decimal) AddCountMiser(precision int32, count int64) Decimal {
	return d.Add(New(count).Shift(-precision))
}

func (d Decimal) SubMiser(precision int32) Decimal {
	return d.Sub(New(1).Shift(-precision))
}

func (d Decimal) SubCountMiser(precision int32, count int64) Decimal {
	return d.Sub(New(count).Shift(-precision))
}

func (d Decimal) Equal(val Decimal) bool {
	return d.Decimal.Equal(val.Decimal)
}

func (d Decimal) GreaterThan(val Decimal) bool {
	return d.Decimal.GreaterThan(val.Decimal)
}

func (d Decimal) GreaterThanOrEqual(val Decimal) bool {
	return d.Decimal.GreaterThanOrEqual(val.Decimal)
}

func (d Decimal) LessThan(val Decimal) bool {
	return d.Decimal.LessThan(val.Decimal)
}

func (d Decimal) LessThanOrEqual(val Decimal) bool {
	return d.Decimal.LessThanOrEqual(val.Decimal)
}

func (d Decimal) Median(v Decimal) Decimal {
	return d.Add(v).Div(New(2))
}

func (d Decimal) SpreadPercent(low Decimal, high Decimal) Decimal {
	return NewFromDecimal(high.Sub(low).Mul(New(100)).Div(d).Abs())
}

func (d Decimal) Round(places int32) Decimal {
	return NewFromDecimal(d.Decimal.Round(places))
}

func (d Decimal) AddFloat(value float64) Decimal {
	return d.Add(NewFromFloat(value))
}

func (d Decimal) SubFloat(value float64) Decimal {
	return d.Sub(NewFromFloat(value))
}

func (d Decimal) AddPercent(percent float64) Decimal {
	return d.Add(d.Div(New(100)).Mul(NewFromFloat(percent)))
}

func (d Decimal) SubPercent(percent float64) Decimal {
	return d.Sub(d.Div(New(100)).Mul(NewFromFloat(percent)))
}

func (d Decimal) AddRandom(min float64, max float64) Decimal {
	return d.AddFloat(rand.Float64() * (max - min))
}

func (d Decimal) SubRandom(min float64, max float64) Decimal {
	return d.SubFloat(rand.Float64() * (max - min))
}

func (d Decimal) AddRandomPercent(min float64, max float64) Decimal {
	return d.AddPercent(min + rand.Float64()*(max-min))
}

func (d Decimal) SubRandomPercent(min float64, max float64) Decimal {
	return d.SubPercent(min + rand.Float64()*(max-min))
}

func (d Decimal) Distance(to Decimal) Decimal {
	return NewFromDecimal(d.Sub(to).Abs())
}

func (from Decimal) DistancePercent(to Decimal) Decimal {
	return NewFromDecimal(from.Sub(to).Mul(New(100)).Div(from).Abs())
}

func (d Decimal) DistancePercentFromLow(to Decimal) Decimal {
	if d.GreaterThan(to) {
		return to.DistancePercent(d)
	}

	return d.DistancePercent(to)
}

func (d Decimal) DistancePercentFromBigger(to Decimal) Decimal {
	if d.GreaterThan(to) {
		return d.DistancePercent(to)
	}

	return to.DistancePercent(d)
}
