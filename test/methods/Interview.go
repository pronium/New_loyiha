package methods

import (
	"Job_Uz/models"
	"fmt"
)

type Interviews struct {
	Interviews []*models.Interview
}

// type InterviewInterface interface {
// 	AddInterview()
// 	PrintAllInterviewInfo()
// 	ChangeInterviewStatus()
// }

func (i *Interviews) AddInterview(Users *Users, Companies *Companies, Vacancies *Vacancies) error {
	var cID, vID, uID int
	var format bool
	var d, t string

	Companies.PrintCompaniesInfo()
	fmt.Print("Kompaniya id kiriting: ")
	fmt.Scan(&cID)
	Vacancies.PrintAllVacanciesInfo()
	fmt.Print("Vakansiya id kiriting: ")
	fmt.Scan(&vID)
	Users.PrintAllUsersInfo()
	fmt.Print("User id kiriting: ")
	fmt.Scan(&uID)
	fmt.Print("Intervyu formatini kiriting(online:true/offline:false): ")
	fmt.Scan(&format)
	fmt.Print("Intervyu sanasini kiriting(dd.mm.yyyy): ")
	fmt.Scan(&d)
	fmt.Print("Intervyu vaqtini kiriting(hh:mm): ")
	fmt.Scan(&t)

	Count := i.Interviews[len(i.Interviews)-1].ID + 1
	newInt := models.Interview{
		ID:        Count,
		CompanyID: cID,
		VacancyID: vID,
		UserID:    uID,
		IsOnline:  format,
		Date:      d,
		Time:      t,
		Status:    2,
	}

	if format {
		fmt.Print("Intervyu bo'lib o'tadigan platformani/saytni kiriting: ")
		fmt.Scan(&newInt.WebSite)
	} else {
		fmt.Println("Intervyu bo'lib o'tadigan address ma'lumotlarini kiriting: ")
		fmt.Print("Ko'cha: ")
		fmt.Scan(&newInt.Address.Street)
		fmt.Print("Suit: ")
		fmt.Scan(&newInt.Address.Suite)
		fmt.Print("Shahar: ")
		fmt.Scan(&newInt.Address.City)
		fmt.Print("Po'chta indeksi: ")
		fmt.Scan(&newInt.Address.Zipcode)
		fmt.Println("Lokatsiya")
		fmt.Print("Shimol-Janub koordinatisi: ")
		fmt.Scan(&newInt.Address.Geo.Lat)
		fmt.Print("Sharq-G'arb koordinatisi: ")
		fmt.Scan(&newInt.Address.Geo.Lng)
	}

	i.Interviews = append(i.Interviews, &newInt)

	ind, err := Users.FindIndexOfUser(uID)
	if err != nil {
		return err
	}

	Users.Users[ind].Interviews = append(Users.Users[ind].Interviews, &newInt)
	ClearTerminal()
	fmt.Println("Interview muvaffaqiyatli qo'shildi!\n")
	return nil
}

func (i *Interviews) PrintAllInterviewInfo() {
	if len(i.Interviews) < 1 {
		fmt.Println("Empty list!")
		return
	}

	for _, v := range i.Interviews {
		i.PrintInterviewInfo(v)
	}
}

func (i *Interviews) PrintInterviewInfo(currentInterview *models.Interview) {
	fmt.Printf("Intervyu ID: %d\n", currentInterview.ID)
	fmt.Printf("Kompaniya ID: %d\n", currentInterview.CompanyID)
	fmt.Printf("Vakansiya ID: %d\n", currentInterview.VacancyID)
	fmt.Printf("User ID: %d\n", currentInterview.UserID)
	fmt.Printf("Intervyu onlayn shaklda: %t\n", currentInterview.IsOnline)
	if currentInterview.IsOnline {
		fmt.Printf("Veb sayt: %s\n", currentInterview.WebSite)
	} else {
		fmt.Printf("Address: %+v\n", currentInterview.Address)
	}
	fmt.Printf("Sana: %s\n", currentInterview.Date)
	fmt.Printf("Vaqt: %s\n", currentInterview.Time)
	switch currentInterview.Status {
	case 1:
		fmt.Println("Intervyu holati: Muvaffaqiyatli")
	case 2:
		fmt.Println("Intervyu holati: Jarayonda")
	case 3:
		fmt.Println("Intervyu holati: Muvaffaqiyatsiz")
	default:
		fmt.Println("Intervyu holatida xatolik!")
	}
	fmt.Println()
}

func (i *Interviews) ChangeInterviewStatus() error {
	i.PrintAllInterviewInfo()

	var targetID, newStatus int
	fmt.Print("Intervyu ID kiriting: ")
	fmt.Scan(&targetID)
	fmt.Print("Intervyu holatini kiriting(1:Success | 2:Possess | 3:Reject): ")
	fmt.Scan(&newStatus)

	ind, err := i.FindIndexOfInterview(targetID)
	if err != nil {
		return err
	}

	if !(newStatus > 0 && newStatus < 4) {
		return fmt.Errorf("Yangi Status xato kiritildi")
	}
	i.Interviews[ind].Status = newStatus

	return nil
}

func (i *Interviews) FindIndexOfInterview(InterviewId int) (int, error) {
	start := 0
	end := len(i.Interviews) - 1
	for start <= end {
		mid := (start + end) / 2
		if i.Interviews[mid].ID < InterviewId {
			start = mid + 1
		} else if i.Interviews[mid].ID > InterviewId {
			end = mid - 1
		} else {
			return mid, nil
		}
	}

	return 0, fmt.Errorf("%d ID'lik intervyu topilmadi!", InterviewId)
}
