// 2. 구상 컬렉션

package main

type UserCollection struct {
	users []*User
}

func (u *UserCollection) createIterator() Iterator {
	return &UserIterator{
		users: u.users,
	}
}
