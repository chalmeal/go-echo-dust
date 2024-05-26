package util

type Util struct {
	Check   Check
	Convert Convert
	Time    Time
	Id      Id
}

func NewUtil(u Util) *Util {
	return &Util{
		Check:   u.Check,
		Convert: u.Convert,
		Time:    u.Time,
		Id:      u.Id,
	}
}
