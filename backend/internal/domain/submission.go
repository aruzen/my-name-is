package domain

// HueResultSubmission はユーザーデータを複製した状態で HueRecord と結び付ける。
type HueResultSubmission struct {
	userData UserData
	record   HueRecord
}

// NewHueResultSubmission は可変な map をその場で Clone してカプセル化する。
func NewHueResultSubmission(userData UserData, record HueRecord) HueResultSubmission {
	return HueResultSubmission{
		userData: userData.Clone(),
		record:   record,
	}
}

// UserData は内部状態を守るため常に Clone を返す。
func (s HueResultSubmission) UserData() UserData {
	return s.userData.Clone()
}

func (s HueResultSubmission) Record() HueRecord {
	return s.record
}
