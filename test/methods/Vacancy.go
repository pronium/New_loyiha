package methods

import (
	"Job_Uz/models"
	"errors"
	"fmt"
	"strings"
)

type Vacancies struct {
	Vacancies []*models.Vacancy
}

// type VacancyInterface interface {
// 	AddVacancy()
// 	CloseVacancy()
// 	UpdateVacanyInfo()
// 	PrintAllVacanciesInfo()
// }

func (vacancies *Vacancies) AddVacancy(Companies Companies) error {
	newVacancy := &models.Vacancy{} // Yangi vakansiya uchun yangi struct o'rnatiyapmiz

	// Ish nomini kiritish
	fmt.Println("Ish nomi:")
	if _, err := fmt.Scan(&newVacancy.Title); err != nil {
		return err
	}

	// Ish haqida ma'lumot kiritish
	fmt.Println("Ish haqida ma'lumot kiriting:")
	var description string
	if _, err := fmt.Scan(&description); err != nil {
		return err
	}
	newVacancy.Description = description

	// Ish talablarini kiritish
	fmt.Println("Ish talablari (elementlarni vergul bilan ajrating):")
	var requirements string
	if _, err := fmt.Scan(&requirements); err != nil {
		return err
	}
	newVacancy.Requirement = strings.Split(requirements, ",")

	// Oylik maosh oralig'ini kiritish
	fmt.Println("Oylik maosh oralig'i:")
	if _, err := fmt.Scan(&newVacancy.SalaryRange); err != nil {
		return err
	}

	// Ish turi (online yoki offline)
	fmt.Println("Ish turi (true - Online, false - Offline):")
	if _, err := fmt.Scan(&newVacancy.WorkType); err != nil {
		return err
	}

	// Takliflar
	fmt.Println("Takliflar (elementlarni vergul bilan ajrating):")
	var offers string
	if _, err := fmt.Scan(&offers); err != nil {
		return err
	}
	newVacancy.Offer = strings.Split(offers, ",")

	// Vakansiya holati
	fmt.Println("Vakansiya holati (true - Ochiq, false - Yopiq):")
	if _, err := fmt.Scan(&newVacancy.Status); err != nil {
		return err
	}

	ClearTerminal()
	Companies.PrintCompaniesInfo()
	fmt.Println("\nVakansiya beradigan Kampaniya ID sini kiriting")
	if _, err := fmt.Scan(&newVacancy.CompanyID); err != nil {
		return err
	}

	newVacancy.VacancyID = vacancies.Vacancies[len(vacancies.Vacancies)-1].VacancyID + 1
	vacancies.Vacancies = append(vacancies.Vacancies, newVacancy)
	return nil
}

func (v *Vacancies) CloseVacancy() error {
	for {
		var vacancyId int
		v.PrintAllVacanciesInfo()
		fmt.Println("\nLoading...\nEnter Vacancy ID to close (or -1 to exit):")
		if _, err := fmt.Scan(&vacancyId); err != nil {
			return err
		}
		if vacancyId == -1 {
			return nil // Allow user to exit the loop
		}
		start, end := 0, len(v.Vacancies)-1
		for start <= end {
			mid := (start + end) / 2
			if v.Vacancies[mid].VacancyID == vacancyId {
				v.Vacancies[mid].Status = false
				return nil
			} else if v.Vacancies[mid].VacancyID > vacancyId {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
		return fmt.Errorf("No vacancy found with ID %v", vacancyId)
	}
}

func (v *Vacancies) UpdateVacanyInfo() error {
	for {
		var vacancyId int
		v.PrintAllVacanciesInfo()
		fmt.Println("Title(1)\nDescription(2)\nRequirement(3)\nSalaryRange(4)\nWorkType(5)\nOffer(6)\nStatus(7)\nChoose fields to update (e.g., 1_2_5):")
		var num, value, offers string
		fmt.Println("Enter Vacancy ID (or -1 to exit):")
		if _, err := fmt.Scan(&vacancyId); err != nil {
			return err
		}
		if vacancyId == -1 {
			return nil // Allow user to exit the loop
		}
		if _, err := fmt.Scan(&num); err != nil {
			return err
		}
		nums := strings.Split(num, "_")
		start, end := 0, len(v.Vacancies)-1
		found := false
		for start <= end {
			mid := (start + end) / 2
			if v.Vacancies[mid].VacancyID == vacancyId {
				found = true
				for _, x := range nums {
					switch x {
					case "1":
						fmt.Println("Enter new Title:")
						if _, err := fmt.Scan(&v.Vacancies[mid].Title); err != nil {
							return err
						}
					case "2":
						fmt.Println("Enter new Description:")
						if _, err := fmt.Scan(&value); err != nil {
							return err
						}
						v.Vacancies[mid].Description = value
					case "3":
						fmt.Println("Enter new Requirements with ',':")
						if _, err := fmt.Scan(&value); err != nil {
							return err
						}
						v.Vacancies[mid].Requirement = strings.Split(value, ",")
					case "4":
						fmt.Println("Enter new Salary Range:")
						if _, err := fmt.Scan(&v.Vacancies[mid].SalaryRange); err != nil {
							return err
						}
					case "5":
						fmt.Println("Enter Work Type (true for Online, false for Offline):")
						if _, err := fmt.Scan(&v.Vacancies[mid].WorkType); err != nil {
							return err
						}
					case "6":
						fmt.Println("Enter new Offers (with ','):")
						if _, err := fmt.Scan(&offers); err != nil {
							return err
						}
						v.Vacancies[mid].Offer = strings.Split(offers, ",")
					case "7":
						fmt.Println("Enter Status (true for Open, false for Closed):")
						if _, err := fmt.Scan(&v.Vacancies[mid].Status); err != nil {
							return err
						}
					default:
						fmt.Printf("Invalid choice %s\n", x)
					}
				}
				break
			} else if v.Vacancies[mid].VacancyID > vacancyId {
				end = mid - 1
			} else {
				start = mid + 1
			}
		}
		if found {
			return nil
		}
		fmt.Println("No vacancy found with ID", vacancyId)
	}
}

func (v *Vacancies) PrintAllVacanciesInfo() error {
	if len(v.Vacancies) == 0 {
		return errors.New("no vacancies available")
	}
	for _, vacancy := range v.Vacancies {
		fmt.Printf("Vacancy ID: %d\nCompany ID: %d\nTitle: %s\n", vacancy.VacancyID, vacancy.CompanyID, vacancy.Title)
		fmt.Println("Description:", vacancy.Description)
		fmt.Println("Requirements:", strings.Join(vacancy.Requirement, ", "))
		fmt.Println("Salary Range:", vacancy.SalaryRange)
		fmt.Println("Work Type:", workType(vacancy.WorkType))
		fmt.Println("Offers:", strings.Join(vacancy.Offer, ", "))
		if vacancy.Status == true {
			fmt.Println("Status: Active")
		} else {
			fmt.Println("Status: Yopiq")
		}
		fmt.Println("\n|---------=---------=---------=---------=---------=---------|\n")
	}
	return nil
}

func workType(workType bool) string {
	if workType {
		return "Online"
	}
	return "Offline"
}

func (Vacancies Vacancies) PrintVacancyInfo(ind int, Companies Companies) {
	currentVacancy := Vacancies.Vacancies[ind]
	companyName, err := Companies.FindCompanyName(currentVacancy.CompanyID)
	if err != nil {
		fmt.Println("Kampaniya: XXX")
	}
	fmt.Println("Vakansiya: ", currentVacancy.Title)
	fmt.Println("Kampaniya: ", companyName)
	fmt.Println("Izoh: ", currentVacancy.Description)

	fmt.Println("Talablar: ")
	for _, requirment := range currentVacancy.Requirement {
		fmt.Println("\t", requirment)
	}

	fmt.Println("Maosh: ", currentVacancy.SalaryRange)

	if currentVacancy.WorkType {
		fmt.Println("Ish turi: Office")
	} else {
		fmt.Println("Ish turi: Masofaviy")
	}

	fmt.Println("Biz beramiz: ")
	for _, Offer := range currentVacancy.Offer {
		fmt.Println("\t", Offer)
	}

	if currentVacancy.Status {
		fmt.Println("Vakansiya ochiq")
	} else {
		fmt.Println("Vakansiya yopilgan")
	}
	fmt.Println("--------------------------------------------")
}
