package main

import (
	readwrite "Job_Uz/ReadWrite"
	"Job_Uz/methods"
	"fmt"
	"os"
	"os/exec"
)

var Users methods.Users
var Companies methods.Companies
var Vacancies methods.Vacancies
var Interviews methods.Interviews

func main() {
	Users = readwrite.ReadUsers("user.json")               // for Users
	Companies = readwrite.ReadCompanies("kompaniya.json")  // for Companies
	Vacancies = readwrite.ReadVacancies("vakasiya.json")   // for Vacancies
	Interviews = readwrite.ReadInterviews("intervyu.json") // for Interviews

	runner := "y"
	for runner != "yes" {
		ClearTerminal()
		fmt.Println("Userler ustida ishlash uchun 1 ni bosing")
		fmt.Println("Companiyalar ustida ishlash uchun 2 ni bosing")
		fmt.Println("Vakansiyalar ustida ishlash uchun 3 ni bosing")
		fmt.Println("Suhbatlar ustida ishlash uchun 4 ni bosing")
		category := -1
		fmt.Scan(&category)
		ClearTerminal()
		if category == 1 {
			MainOfUsers()
		} else if category == 2 {
			MainOfCompanies()
		} else if category == 3 {
			MainOfVacancies()
		} else if category == 4 {
			MainOfInterviews()
		} else {
			fmt.Println("Xato raqam kiritidingiz!!!")
			return
		}
		fmt.Println("\n---------------------------------\n")
		fmt.Println("Chiqich uchun 'yes' deb yozing")
		fmt.Println("Aks holda hohlagan harf ni bosing")
		fmt.Scan(&runner)
	}

	readwrite.WriteUsers("user.json", Users)               // for Users
	readwrite.WriteCompanies("kompaniya.json", Companies)  // for Companies
	readwrite.WriteVacancies("vakasiya.json", Vacancies)   // for Vacancies
	readwrite.WriteInterviews("intervyu.json", Interviews) // for Interviews                      // for Interviews
}

func MainOfUsers() {
	fmt.Println("Hamma user lar haqida malumot kerak bo'lsa 1 ni bosing")
	fmt.Println("Biror bir userni interviewlari haqida bilmoqchi bo'lsangiz 2 ni bosing")
	fmt.Println("User qo'shish uchun 3 ni bosing")
	fmt.Println("User o'chirish uchun 4 ni bosing")
	fmt.Println("User malumotini taxrirlash uchun 5 ni bosing")
	fmt.Println("Userni statusini o'zgartirish uchun 6 ni bosing")
	command := 0
	fmt.Scan(&command)
	ClearTerminal()
	switch command {
	case 1:
		Users.PrintAllUsersInfo()
	case 2:
		Users.PrintAllInterviewsOfUser(&Interviews)
	case 3:
		err := Users.AddUser()
		if err != nil {
			panic(err)
		} else {
			ClearTerminal()
			fmt.Println("User muvaffaqiyatli qo'shildi!\n")
		}
	case 4:
		Users.PrintAllUsersInfo()
		var chooseID int
		fmt.Print("\nO'chirmoqchi bo'lgan user ni ID sini kiriting: ")
		fmt.Scan(&chooseID)
		ClearTerminal()
		err := Users.DeleteUser(chooseID)
		if err != nil {
			panic(err)
		} else {
			ClearTerminal()
			fmt.Println("User muvaffaqiyatli o'chirildi!\n")
		}
	case 5:
		Users.PrintAllUsersInfo()
		var chooseID int
		fmt.Print("\nTaxrirlamoqchi bo'lgan user ni ID sini kiriting: ")
		fmt.Scan(&chooseID)
		ClearTerminal()
		err := Users.UpdateUserInfo(chooseID)
		if err != nil {
			panic(err)
		} else {
			ClearTerminal()
			fmt.Println("User malumotlari muvaffaqiyatli o'zgartilirdi!\n")
		}
	case 6:
		Users.PrintAllUsersInfo()
		var chooseID int
		fmt.Print("\nTaxrirlamoqchi bo'lgan user ni ID sini kiriting: ")
		fmt.Scan(&chooseID)
		ClearTerminal()
		fmt.Println("Recruter qilmoqchi bo'lsangiz 1 bosing")
		fmt.Println("Employee qilmoqchi bo'lsangiz 2 bosing")
		fmt.Println("Condidate qilmoqchi bo'lsangiz 3 bosing")
		var chooseStatus int
		fmt.Scan(&chooseStatus)
		err := Users.ChangeUserStatus(chooseID, chooseStatus)
		if err != nil {
			panic(err)
		} else {
			ClearTerminal()
			fmt.Println("User Statusi muvaffaqiyatli o'zgartirildi!\n")
		}
	default:
		panic(fmt.Errorf("Xato son tanladingiz!"))
	}
}

func MainOfCompanies() {
	fmt.Println("kampaniyaga tegishli vakansiyalarni ko'rish uchun 1 ni bosing")
	fmt.Println("Kampaniyaning userlarini ko'rish uchun 2 ni bosing")
	fmt.Println("User qo'shish uchun 3 ni bosing")
	fmt.Println("User o'chirish uchun 4 ni bosing")
	fmt.Println("Hamma kampaniyalar haqidagi malumotni ko'rish uchun 5 ni bosing")
	command := 0

	fmt.Scan(&command)
	ClearTerminal()
	switch command {
	case 1:
		Companies.PrintVacanciesOfCompany(Vacancies)
	case 2:
		var err error
		for i := 0; i < 3; i++ {
			ClearTerminal()

			Companies.PrintCompaniesInfo()

			if err != fmt.Errorf("Xato") {
				fmt.Println(err)
			}

			var chooseID int
			fmt.Print("\nUserlarini ko'rmoqchi bo'lgan kampaniya ID sini kiriting: ")
			fmt.Scan(&chooseID)

			err = Companies.PrintUsersOfCompany(chooseID)
			if err == nil {
				break
			}
		}
	case 3:
		Companies.AddUserToCompany(&Users)
	case 4:
		var err error
		for i := 0; i < 3; i++ {
			ClearTerminal()
			Companies.PrintCompaniesInfo()

			if err != fmt.Errorf("Xato") {
				fmt.Println(err)
			}

			err = Companies.DeleteUserFromCompany()
			if err == nil {
				ClearTerminal()
				fmt.Println("User muvaffaqiyatli o'chirildi!\n")
				break
			}
		}
	case 5:
		Companies.PrintCompaniesInfo()
	default:
		panic(fmt.Errorf("Xato son tanladingiz!"))
	}
}

func MainOfVacancies() {
	fmt.Println("Vakansiya qo'shish uchun 1 ni bosing")
	fmt.Println("Vakansiyani yopish uchun 2 ni bosing")
	fmt.Println("Vakansiyaga o'zgarishlar kiritish uchun 3 ni bosing")
	fmt.Println("Vakansiyalarni hammasini ko'rish uchun 4 ni bosing")
	command := 0
	fmt.Scan(&command)
	ClearTerminal()
	switch command {
	case 1:
		// fmt.Println("Vacansiya qo'shish uchun 1 ni bosing")
		err := Vacancies.AddVacancy(Companies)
		if err != nil {
			panic(err)
		} else {
			ClearTerminal()
			fmt.Println("Vakansiya muvaffaqiyatli qo'shildi!\n")
		}
	case 2:
		// fmt.Println("Vakansiyani yopish uchun 2 ni bosing")
		err := Vacancies.CloseVacancy()
		if err != nil {
			panic(err)
		} else {
			ClearTerminal()
			fmt.Println("Vakansiya muvaffaqiyatli yopildi!\n")
		}
	case 3:
		// fmt.Println("Vakansiyaga o'zgarishlar kiritish uchun 3 ni bosing")
		err := Vacancies.UpdateVacanyInfo()
		if err != nil {
			panic(err)
		} else {
			ClearTerminal()
			fmt.Println("Vakansiya muvaffaqiyatli o'zgartirildi!\n")
		}
	case 4:
		// fmt.Println("Vakansiyalarni hammasini ko'rish uchun 4 ni bosing")
		Vacancies.PrintAllVacanciesInfo()
	default:
		panic(fmt.Errorf("Xato son tanladingiz!"))
	}
}

func MainOfInterviews() {
	fmt.Println("Suhbat belgilash uchun 1 ni bosing")
	fmt.Println("Suhbatlar haqida malumot olish uchun 2 ni bosing")
	fmt.Println("Suhbatni natijasini yozish uchun 3 ni bosing")
	command := 0
	fmt.Scan(&command)
	ClearTerminal()
	switch command {
	case 1:
		// fmt.Println("Suhbat belgilash uchun 1 ni bosing")
		var err error
		for i := 0; i < 3; i++ {
			ClearTerminal()
			if err != fmt.Errorf("Xato") {
				fmt.Println(err)
			}
			err = Interviews.AddInterview(&Users, &Companies, &Vacancies)
			if err == nil {
				break
			}
		}
	case 2:
		// fmt.Println("Suhbatlar haqida malumot olish uchun 2 ni bosing")
		Interviews.PrintAllInterviewInfo()
	case 3:
		// fmt.Println("Suhbatni natijasini yozish uchun 3 ni bosing")
		var err error
		for i := 0; i < 3; i++ {
			ClearTerminal()
			if err != nil {
				fmt.Println(err)
			}
			err = Interviews.ChangeInterviewStatus()
			if err == nil {
				break
			}
		}
	default:
		panic(fmt.Errorf("Xato son tanladingiz!"))
	}
}

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
