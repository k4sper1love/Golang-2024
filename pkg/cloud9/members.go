package cloud9

type Member struct {
	Type string `json:"type"`
	Role string `json:"role"`
	Nickname string `json:"nickname"`
	FirstName string `json:"firstname"`
	SecondName string `json:"secondname"`
}

func GetMembers() []Member{
	var members []Member
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "HObbit", FirstName: "Abay", SecondName: "Khassenov"})
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "Ax1Le", FirstName: "Sergey", SecondName: "Rykhtorov"})
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "electroNic", FirstName: "Denis", SecondName: "Sharipov"})
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "Perfecto", FirstName: "Ilia", SecondName: "Zalutskii"})
	members = append(members, Member{Type: "Player", Role: "Rifler", Nickname: "Boombl4", FirstName: "Kirill", SecondName: "Mikhaylov"})
	members = append(members, Member{Type: "Staff", Role: "Coach", Nickname: "groove", FirstName: "Konstantin", SecondName: "Pikiner"})
	members = append(members, Member{Type: "Staff", Role: "Analyst", Nickname: "F_1N", FirstName: "Ivan", SecondName: "Kochugov"})
	members = append(members, Member{Type: "Staff", Role: "Manager", Nickname: "Sweetypotz", FirstName: "Aleksandr", SecondName: "Shcherbakov"})
	return members
}