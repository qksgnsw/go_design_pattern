// 4. 구상 반복자

package main

type UserIterator struct {
	index int
	users []*User
}

func (u *UserIterator) hasNext() bool {
	// if u.index < len(u.users) {
	// 	return true
	// }
	// return false
	return u.index < len(u.users)

}
func (u *UserIterator) getNext() *User {
	if u.hasNext() {
		user := u.users[u.index]
		u.index++
		return user
	}
	return nil
}
