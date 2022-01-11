package unitconv

//go:generate stringer -type=Unit -linecomment

type Unit uint

func (u Unit) ethExponent() int {
	return etherPowMap[u]
}

func (u Unit) weiExponent() int {
	return weiPowMap[u]
}

func (u Unit) gtEther() bool {
	return etherPowMap[u] > etherPowMap[Ether]
}

func (u Unit) ltEther() bool {
	return etherPowMap[u] < etherPowMap[Ether]
}

func (u Unit) valid() bool {
	_, ok := etherPowMap[u]

	return ok
}

const (
	// Wei is the base unit of currency in the Ethereum blockchain universe.
	// Wei converts to 10^-18 Ether.
	Wei Unit = iota
	// KWei represents 1000 Wei, or 10^-15 Ether.
	KWei
	// MWei represents 1000000 Wei, or 10^-12 Ether.
	MWei
	// GWei represents 1000000000 Wei, or 10^-9 Ether. GWei is usually used as the unit of measurement for transaction gas prices.
	GWei
	// Szabo represents 1000000000000 Wei, or 10^-6 Ether.
	Szabo
	// Finney represents 1000000000000000 Wei, or 10^-3 Ether.
	Finney
	// Ether is 10^18 Wei, and usually the unit used when displaying cryptocurrency amounts in "normal" contexts.
	Ether
	// KEther represents 1000 (10^3) Ether.
	KEther
	// MEther represents 1000000 (10^6) Ether.
	MEther
	// GEther represents 1000000000 (10^9) Ether.
	GEther
	// TEther represents 1000000000000 (10^12) Ether.
	TEther
)

var (
	etherPowMap = map[Unit]int{
		Wei:    -18,
		KWei:   -15,
		MWei:   -12,
		GWei:   -9,
		Szabo:  -6,
		Finney: -3,
		Ether:  1,
		KEther: 3,
		MEther: 6,
		GEther: 9,
		TEther: 12,
	}
	weiPowMap = map[Unit]int{
		Wei:    0,
		KWei:   3,
		MWei:   6,
		GWei:   9,
		Szabo:  12,
		Finney: 15,
		Ether:  18,
		KEther: 21,
		MEther: 24,
		GEther: 27,
		TEther: 30,
	}
)
