package id

type AccountID string

func (a AccountID) String() string {
	return string(a)
}

type TripID string

func (t TripID) String() string {
	return string(t)
}

type IdentiTyID string

func (t IdentiTyID) String() string {
	return string(t)
}

type CarID string

func (t CarID) String() string {
	return string(t)
}
