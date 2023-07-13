package server

import (
	"time"

	"github.com/TurboHsu/munager/cmd/sync/structure"
)

var UserBase []User

func (u *User) DeleteFile(filebase string) {
	for i, f := range u.AccessList {
		if f.PathBase == filebase {
			// Delete this file from access list
			u.AccessList = append(u.AccessList[:i], u.AccessList[i+1:]...)
			break
		}
	}
}

func registerUser(fingerprint string) {
	UserBase = append(UserBase, User{
		Fingerprint: fingerprint,
		AccessList:  []structure.FileInfo{},
		TimeCreated: time.Now(),
	})
}

func killUser(fingerprint string) {
	for i, u := range UserBase {
		if u.Fingerprint == fingerprint {
			UserBase = append(UserBase[:i], UserBase[i+1:]...)
			break
		}
	}
}

func getUser(fingerprint string) (ret *User) {
	for i := 0; i < len(UserBase); i++ {
		if UserBase[i].Fingerprint == fingerprint {
			ret = &UserBase[i]
			return
		}
	}
	return
}

type User struct {
	Fingerprint string
	AccessList  []structure.FileInfo
	TimeCreated time.Time
}
