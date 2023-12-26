package quorum

func NewReadRepairFunc[Query, Data any](_ Discovery[Query, Data]) ReadRepairFunc[Query, Data] {
	return nil
}
