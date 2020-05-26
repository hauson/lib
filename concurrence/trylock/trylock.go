package trylock

type TryLock chan int

func New() TryLock {
	ch := make(chan int, 1)
	return TryLock(ch)
}

// lock without block
func (l TryLock) TryLock() bool {
	select {
	case l <- 1:
		return true
	default:
		return false
	}
	return false
}

// unlock without block
func (l TryLock) TryUnLock() bool {
	select {
	case <-l:
		return true
	default:
		return false
	}
	return false
}

// lock with block
func (l TryLock) Lock() {
	l <- 1
}

// unlock with block
func (l TryLock) UnLock() {
	<-l
}
