package enum

type CholesterolLevel int

const (
	Good = iota
	Warning
	Danger
)

const (
	GoodString    = "Good"
	WarningString = "Warning"
	DangerString  = "Danger"
)

func (cl CholesterolLevel) String() string {
	return [...]string{GoodString, WarningString, DangerString}[cl]
}

func (cl CholesterolLevel) Int() int {
	return [...]int{Good, Warning, Danger}[cl]
}
