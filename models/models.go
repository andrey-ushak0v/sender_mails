package models

type UserInDb struct {
	FirstName string
	LastName  string
	BirthYear int
}

//хранилище емейлов и данных о пользователе
func UsersData() map[string]UserInDb {
	users := map[string]UserInDb{
		"prigolovkina@mail.ru":      UserInDb{FirstName: "Alla", LastName: "Prigolovkina", BirthYear: 1989},
		"uandrey220@gmail.com":      UserInDb{FirstName: "Vas", LastName: "Pupkin", BirthYear: 1990},
		"andreyushakov09@yandex.ru": UserInDb{FirstName: "Vas", LastName: "Pkin", BirthYear: 1980},
		"foo@baiirhhh.com":          UserInDb{FirstName: "Vaspp", LastName: "Pkmnin", BirthYear: 1981},
	}
	return users
}

//получение данных о юзере по емейлу
func GetUserByEmail(mail string, users map[string]UserInDb) UserInDb {
	return users[mail]
}
