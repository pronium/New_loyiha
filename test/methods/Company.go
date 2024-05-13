package methods

import (
	"Job_Uz/models"
	"fmt"
	"os"
	"os/exec"
)

type Companies struct {
	Companies []*models.Company
}

// type CompanyInterface interface {
// 	PrintVacanciesOfCompany()
// 	PrintUsersOfCompany()
// 	AddUserToCompany()
// 	DeleteUserFromCompany()
// 	PrintCompaniesInfo()
// }

// Kompaniya vakansiyalarini chop etuvchi metod
func (c Companies) PrintVacanciesOfCompany(Vacancies Vacancies) {
	c.PrintCompaniesInfo()

	var CompanyID int
	fmt.Print("\nCompanya ID sini kiriting: ")
	fmt.Scan(&CompanyID)
	ClearTerminal()
	for ind, vacancy := range Vacancies.Vacancies {
		if vacancy.CompanyID == CompanyID {
			Vacancies.PrintVacancyInfo(ind, c)
		}
	}
}

// Kompaniya foydalanuvchilarini chop etuvchi metod
func (c Companies) PrintUsersOfCompany(CompanyID int) error {
	ind, err := c.FindIndexOfCompany(CompanyID)
	if err != nil {
		return fmt.Errorf("Companiya Id si xato berilgan: %v\n", CompanyID)
	}
	ClearTerminal()
	for _, User := range c.Companies[ind].Users {
		fmt.Println("UserID:", User.Id)
		fmt.Println("Ismi:", User.Name)
		fmt.Println("Familiyasi:", User.Surname)
		fmt.Println("Yoshi:", User.Age)
		fmt.Println("Telefon raqami:", User.Phone)
		fmt.Println("Email:", User.Email)

		if User.Status == 1 {
			fmt.Println("Status: Recruter")
		} else if User.Status == 2 {
			fmt.Println("Status: Employee")
		} else if User.Status == 3 {
			fmt.Println("Status: Candidate")
		}
		fmt.Println("Tajribasi:", User.Experince)
		fmt.Println("---------------------------------------")
	}
	return nil
}

// Foydalanuvchini kompaniyaga qo'shuvchi metod
func (c *Companies) AddUserToCompany(users *Users) {
	var err error
	for i := 0; i < 3; i++ {
		c.PrintCompaniesInfo()

		if err != nil {
			fmt.Println("\n", err, "\n")
		}

		fmt.Print("\nUser qo'shmoqchi bo'lgan kampaniya ID sini kiriting: ")
		var chooseCompanyID int
		fmt.Scan(&chooseCompanyID)
		ClearTerminal()

		var isExistedUser int // Existed(1) | nonExisted(2)
		fmt.Println("Dasturga kiritilgan User ni qo'shish uchun 1 ni bosing")
		fmt.Println("Dasturga kiritilmagan User ni qo'shish uchun 2 ni bosing")
		fmt.Scan(&isExistedUser)
		if isExistedUser == 1 {
			ClearTerminal()
			users.PrintAllUsersInfo()
			var choose int
			fmt.Print("\nKiritmoqchi bo'lgan User ni ID sini kiriting: ")
			fmt.Scan(&choose)
			ind, err := users.FindIndexOfUser(choose)
			if err != nil {
				continue
			}

			CompanyIndex, err := c.FindIndexOfCompany(chooseCompanyID)
			if err == nil {
				UserToAdd := users.Users[ind]
				c.Companies[CompanyIndex].Users = append(c.Companies[CompanyIndex].Users, &UserToAdd)
				break
			}
		} else if isExistedUser == 2 {
			err = users.AddUser()
			if err != nil {
				continue
			}

			CompanyIndex, err := c.FindIndexOfCompany(chooseCompanyID)
			if err == nil {
				UserToAdd := users.Users[len(users.Users)-1]
				c.Companies[CompanyIndex].Users = append(c.Companies[CompanyIndex].Users, &UserToAdd)
				break
			}
		} else {
			err = fmt.Errorf("Xato categoriyani tanladingiz")
		}
	}
	ClearTerminal()
	fmt.Println("User muvaffaqiyatli qo'shildi!\n")
}

// Foydalanuvchini kompaniyadan o'chiruvchi metod
func (c *Companies) DeleteUserFromCompany() error {
	var chooseCompanyID int
	fmt.Print("\nUser o'chirmoqchi bo'lgan kampaniya ID sini kiriting: ")
	fmt.Scan(&chooseCompanyID)
	ClearTerminal()
	err := c.PrintUsersOfCompany(chooseCompanyID)
	if err != nil {
		panic(err)
	}
	var chooseUserID int
	fmt.Print("\nO'chirmoqchi bo'lgan User ID sini kiriting: ")
	fmt.Scan(&chooseUserID)

	IndexOfCurrentCompany, err := c.FindIndexOfCompany(chooseCompanyID)
	if err != nil {
		return err
	}

	currentCompany := c.Companies[IndexOfCurrentCompany]
	ind, err := c.FindIndexOfCompanyUser(currentCompany.Users, chooseUserID)
	if err != nil {
		return err
	}
	currentCompany.Users = append(currentCompany.Users[:ind], currentCompany.Users[ind+1:]...)

	return nil
}

// Kompaniyalar haqida ma'lumotlarni chop etuvchi metod
func (c Companies) PrintCompaniesInfo() {
	for _, company := range c.Companies {
		fmt.Println("Kampaniya Id:", company.ID)
		fmt.Println("Kampaniya nomi:", company.Name)
		fmt.Println("Kampaniya turi:", company.Industry)
		fmt.Println("Kampaniya Userlar soni:", len(company.Users))
		fmt.Println("--------------------------------------------")
	}
}

func (c Companies) FindIndexOfCompanyUser(Users []*models.User, UserId int) (int, error) {
	start := 0
	end := len(Users) - 1

	for start <= end {
		mid := (start + end) / 2

		if Users[mid].Id == UserId {
			return mid, nil
		} else if Users[mid].Id > UserId {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0, fmt.Errorf("Bunday Id li User yo'q")
}

func (c Companies) FindCompanyName(CompanyID int) (string, error) {
	ind, err := c.FindIndexOfCompany(CompanyID)
	if err != nil {
		return "", err
	}
	return c.Companies[ind].Name, nil
}

func (c Companies) FindIndexOfCompany(CompanyID int) (int, error) {
	start := 0
	end := len(c.Companies) - 1

	for start <= end {
		mid := (start + end) / 2

		if c.Companies[mid].ID == CompanyID {
			return mid, nil
		} else if c.Companies[mid].ID < CompanyID {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}

	return 0, fmt.Errorf("kampaniya Id xato berilgan")
}

func ClearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
