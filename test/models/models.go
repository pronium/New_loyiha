package models

type Geo struct {
	Lat string
	Lng string
}

type Address struct {
	Street  string
	Suite   string
	City    string
	Zipcode string
	Geo     Geo
}

type Company struct {
	ID       int
	Name     string
	Industry string
	Address  Address
	Users    []*User
}

type User struct {
	Id           int
	Name         string
	Surname      string
	Age          int
	Phone        string
	Email        string
	Status       int // Recruter(1) | Employee(2) | Candidate(3)
	Experince    int
	Applications [][]int      // CompanyId | Vacancy ID
	Interviews   []*Interview //Interview_ID
}

type Interview struct {
	ID        int
	CompanyID int
	VacancyID int
	UserID    int
	IsOnline  bool `json:"InterviewType"`
	Address   Address
	WebSite   string `json:"Platform"`
	Date      string
	Time      string
	Status    int // Success(1) | Possess(2) | Reject(3)
}

type Vacancy struct {
	VacancyID   int //* Vakansiya raqami
	CompanyID   int
	Title       string   //* Ish nomi
	Description string   //* Ish vazifalarining batafsil tavsifi
	Requirement []string //* Job requirements
	SalaryRange string   //* Kutilayotgan ish haqi oralig'i
	WorkType    bool     //* true = Online or false = Offline
	Offer       []string //* Takliflar
	Status      bool     //! *Open or Closed
}
