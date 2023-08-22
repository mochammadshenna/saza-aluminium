package state

var (
	_userPromotionTypes userPromotionTypes
	_userStatuses       userStatuses
)

func init() {
	_userPromotionTypes = newUserPromotionTypes()
	_userStatuses = newUserStatuses()
}

type UserPromotionType int64

func (s UserPromotionType) Int64() int64 {
	return int64(s)
}

func (s UserPromotionType) Name() string {
	switch s {
	case UserPromotionTypes().Birthday:
		return "BIRTHDAY"
	case UserPromotionTypes().Voucher:
		return "VOUCHER"
	default:
		return ""
	}
}

type userPromotionTypes struct {
	Birthday UserPromotionType
	Voucher  UserPromotionType
}

func newUserPromotionTypes() userPromotionTypes {
	return userPromotionTypes{
		Birthday: 1,
		Voucher:  2,
	}
}

func UserPromotionTypes() userPromotionTypes {
	return _userPromotionTypes
}

type UserStatus int64

func (s UserStatus) Int64() int64 {
	return int64(s)
}

func (s UserStatus) Name() string {
	switch s {
	case UserStatuses().Pending:
		return "PENDING"
	case UserStatuses().Verified:
		return "VERIFIED"
	case UserStatuses().Rejected:
		return "REJECTED"
	default:
		return ""
	}
}

type userStatuses struct {
	Pending  UserStatus
	Verified UserStatus
	Rejected UserStatus
}

func newUserStatuses() userStatuses {
	return userStatuses{
		Pending:  1,
		Verified: 2,
		Rejected: 3,
	}
}

func UserStatuses() userStatuses {
	return _userStatuses
}
