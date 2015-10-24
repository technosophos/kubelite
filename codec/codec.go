package codec

type Oner interface {
	One() *Manifest
}

type Aller interface {
	All() []*Manifest
}

type Forer interface {
	For() string
}

type Musketeer interface {
	Oner
	Forer
	Aller
	// And aller forer oner
}

type Decoder interface {
	Encode() Musketeer
	Decode() Musketeer
}
