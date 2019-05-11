package asset

type Store interface {
	Fetch(assetToFetch Asset, assetsToPreserve ...WritableAsset) error
	Destroy(Asset) error
	DestroyState() error
}
