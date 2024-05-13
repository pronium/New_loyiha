package methods

import (
	"Job_Uz/models"
	"fmt"
	"strings"
)

type Users struct {
	Users []models.User
}

// type UserInterface interface {
// 	PrintAllUsersInfo()
// 	PrintAllVacancyOfUser()
// 	AddUser()
// 	DeleteUser()
// 	UpdateUserInfo()
// 	ChangeUserStatus()
// }

func (Users Users) FindIndexOfUser(UserId int) (int, error) {
	start := 0
	end := len(Users.Users) - 1

	for start <= end {
		mid := (start + end) / 2

		if Users.Users[mid].Id == UserId {
			return mid, nil
		} else if Users.Users[mid].Id > UserId {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0, fmt.Errorf("Bunday Indexli User yo'q ekan")
}

func (Users Users) PrintAllUsersInfo() {
	for _, v := range Users.Users {
		fmt.Printf("Id: %d\n", v.Id)
		fmt.Printf("Ismi: %s\n", v.Name)
		fmt.Printf("Familyasi: %s\n", v.Surname)
		fmt.Printf("Yoshi: %d\n", v.Age)
		fmt.Printf("Tel nomeri: %s\n", v.Phone)
		fmt.Printf("Email: %s\n", v.Email)
		fmt.Printf("Status: %d\n", v.Status)
		fmt.Printf("Daraja: %d\n", v.Id)
		fmt.Println("-------------------------------------")
	}
}

func (Users Users) PrintAllInterviewsOfUser(Interviews *Interviews) {
	var err error

	for i := 0; i < 3; i++ {
		ClearTerminal()

		Users.PrintAllUsersInfo()

		if err != nil {
			fmt.Println("\n", err)
		}

		fmt.Println("\nQaysi User ning Interview lari haqida bilmoqchisiz: ")
		fmt.Print("User ning Id sini kiriting: ")
		var UserId int
		fmt.Scan(&UserId)
		var ind int
		ind, err = Users.FindIndexOfUser(UserId)
		if err == nil {
			ClearTerminal()
			for _, v := range Users.Users[ind].Interviews {
				Interviews.PrintInterviewInfo(v)
			}
			break
		}
	}
}

func (Users *Users) AddUser() error {
	newId := Users.Users[len(Users.Users)-1].Id + 1
	var newName string
	var newSurname string
	var newAge int
	var newPhone string
	var newEmail string
	var newStatus int
	var newExperince int
	var nweAplication = make([][]int, 0)
	var nevIdInterview = make([]*models.Interview, 0)
	fmt.Println("Yangi user ismi: ")
	fmt.Scan(&newName)
	fmt.Println("Familiyasi: ")
	fmt.Scan(&newSurname)
	fmt.Println("Yoshi: ")
	fmt.Scan(&newAge)
	fmt.Println("Tel nomeri: ")
	fmt.Scan(&newPhone)
	fmt.Println("Emaili: ")
	fmt.Scan(&newEmail)
	fmt.Println("Statusi: Recruter(1) | Employee(2) | Candidate(3)")
	fmt.Scan(&newStatus)
	fmt.Println("Experince: ")
	fmt.Scan(&newExperince)
	Users.Users = append(Users.Users, models.User{
		Id:           newId,
		Name:         newName,
		Surname:      newSurname,
		Age:          newAge,
		Phone:        newPhone,
		Email:        newEmail,
		Status:       newStatus,
		Experince:    newExperince,
		Applications: nweAplication,
		Interviews:   nevIdInterview,
	})
	return nil

}

func (Users *Users) DeleteUser(UserId int) error {
	var closeId int
	closeId, err := Users.FindIndexOfUser(UserId)
	if err != nil {
		return fmt.Errorf("Bunday user yoq")
	}
	Users.Users = append(Users.Users[:closeId], Users.Users[closeId+1:]...)
	return nil
}

func (Users *Users) UpdateUserInfo(UserId int) error {
	fmt.Println("Userni qaysi malumotlarini o'zgartirmoqchisiz: ")
	fmt.Println("1-ismi\n2-familiyasi\n3-yosh\n4-tel nomeri\n5-emaili\n6-statusi\n7-tajribasi")
	fmt.Println("O'zgartirmoqchi bo'lgan malumotlaringizni quydagicha kiritasiz (5_6_2_3)! ")
	fmt.Println("O'zogartimoqchi bolgan malumotlaringizni ayting: ")
	var newInfo string
	fmt.Scan(&newInfo)
	var newName string
	var newSurname string
	var newAge int
	var newPhone string
	var newEmail string
	var newStatus int
	var newExperince int
	i, err := Users.FindIndexOfUser(UserId)
	if err != nil {
		return err
	}
	for _, value := range strings.Split(newInfo, "_") {
		if value == "1" {
			fmt.Println("Yani ismini kiriting: ")
			fmt.Scan(&newName)
			Users.Users[i].Name = newName
		} else if value == "2" {
			fmt.Println("Yani familiya kiriting: ")
			fmt.Scan(&newSurname)
			Users.Users[i].Surname = newSurname
		} else if value == "3" {
			fmt.Println("Yani yosh kiriting: ")
			fmt.Scan(&newAge)
			Users.Users[i].Age = newAge
		} else if value == "4" {
			fmt.Println("Yani tel nomer kiriting: ")
			fmt.Scan(&newPhone)
			Users.Users[i].Phone = newPhone
		} else if value == "5" {
			fmt.Println("Yani email kiriting: ")
			fmt.Scan(&newEmail)
			Users.Users[i].Email = newEmail
		} else if value == "6" {
			fmt.Println("Yani status kiriting: ")
			fmt.Scan(&newStatus)
			Users.Users[i].Status = newStatus
		} else if value == "7" {
			fmt.Println("yangilangan tajribani kiriting: ")
			fmt.Scan(&newExperince)
			Users.Users[i].Experince = newExperince
		} else {
			return fmt.Errorf("Userning Bunday alumotini ozgartirolmaysiz! ")
		}
	}
	return nil
}

func (Users *Users) ChangeUserStatus(UserID int, StatusChanger int) error {
	i, err := Users.FindIndexOfUser(UserID)
	if err != nil {
		return err
	}
	if !(StatusChanger > 0 && StatusChanger < 4) {
		return fmt.Errorf("Taxrirlashda raqam noto'g'ri kiritilgan")
	}
	Users.Users[i].Status = StatusChanger
	return nil
}

// vacancyIndex = v[1]
// 		companyIndex = v[0]

// 		fmt.Printf("Kompaniya nomi: %s", Companies[companyIndex].Name)
// 		fmt.Printf("Bajaradigan ishi: %s",Vacancies[vacancyIndex].Description)
// 		fmt.Printf("Talablar: ")
// 		for _, Val := range Vacancies[vacancyIndex].Requirement {
// 			fmt.Println("\t", Val)
// 		  }
// 		fmt.Printf("Ish haqi %d,",Vacancies[vacancyIndex].SalaryRange)
// 		fmt.Printf("")
// 		if Vacancies[vacancyIndex].WorkType {
// 			fmt.Println("ish turi: Oflain")
// 		} else {
// 			fmt.Println("ish turi: Onlain")
// 		}
// 		fmt.Printf("Takloiflar: %s", Vacancies[vacancyIndex].Offer)

// 		if Vacancies[vacancyIndex].Status{
// 			fmt.Println("Vakansiya aktive")
// 		}else{
// 			fmt.Println("Vakansiya aktive emas")
// 		}
