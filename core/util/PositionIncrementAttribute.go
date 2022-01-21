package util

type PositionIncrementAttribute interface {

	// SetPositionIncrement Set the position increment. The default value is one.
	SetPositionIncrement(positionIncrement int)

	// GetPositionIncrement Returns the position increment of this Token.
	GetPositionIncrement() int
}
