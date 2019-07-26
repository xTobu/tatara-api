package user

import (
	modelUser "tatara-api/lib/database/models/user"
	"tatara-api/lib/log"
)

// StructAllUsers ...
type StructAllUsers struct {
	UserList1 []*modelUser.User `json:"list1"`
	UserList2 []*modelUser.User `json:"list2"`
}

// AllUsers 取得所有 User
func AllUsers() (result *StructAllUsers, err error) {

	u, _ := modelUser.FindUsers()
	if err != nil {
		log.Error("AllUsers", err)
		return
	}
	result = &StructAllUsers{
		UserList1: u,
		UserList2: u,
	}

	return
}
