package domain

// RecordRange はレコード取得の閉区間 [begin, end] を保持する。
type RecordRange struct {
	begin int
	end   int
}

// NewRecordRange は begin>=0 かつ begin<=end を満たさない場合 ErrInvalidRange を返す。
func NewRecordRange(begin, end int) (RecordRange, error) {
	if begin < 0 || end < begin {
		return RecordRange{}, ErrInvalidRange
	}

	return RecordRange{
		begin: begin,
		end:   end,
	}, nil
}

func (r RecordRange) Begin() int {
	return r.begin
}

func (r RecordRange) End() int {
	return r.end
}

func (r RecordRange) Count() int {
	return r.end - r.begin + 1
}
