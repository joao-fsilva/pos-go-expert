package entity

type LimiterError struct {
	Message string
	Err     string
}

func (le *LimiterError) Error() string {
	return le.Message
}

func NewExpiredLimiterError() *LimiterError {
	return &LimiterError{
		Message: "Your limit has expired",
		Err:     "expired_limiter",
	}
}

func NewIncrementBlockedError() *LimiterError {
	return &LimiterError{
		Message: "You have reached the maximum number of requests or actions allowed within a certain time frame.",
		Err:     "is_blocked",
	}
}
