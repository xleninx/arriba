package constants

type AssetID string

const (
	ETH AssetID = "ETH"
	BTC AssetID = "BTC"
	USD AssetID = "USD"
)

func (asset AssetID) IsValid() bool {
	switch asset {
	case ETH, BTC, USD:
		return true
	}
	return false
}
