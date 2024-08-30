package storage

import "main/bins"

type DB interface {
	AddBin(bin bins.Bin)
	Print()
}
