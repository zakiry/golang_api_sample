package entity

type User struct {
	id   int64
	name string
	age  int32
}

type UserRecord struct {
	Id   int64
	Name string
	Age  int32
}

func (record UserRecord) Build() User {
	return User{
		id:   record.Id,
		name: record.Name,
		age:  record.Age,
	}
}

func (user User) Id() int64 {
	return user.id
}

func (user User) Name() string {
	return user.name
}

func (user User) Age() int32 {
	return user.age
}
