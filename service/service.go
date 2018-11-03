package service

import (
	"agenda/entity"
	"agenda/logger"
	"log"
)


var curuserinfoPath = "/src/agenda/data/curuser.txt"
var ErrLog *log.Logger
type User entity.User
type Meeting entity.Meeting

func init() {
	ErrLog = logger.Error
}

func UserLogout() bool {
	if err := entity.Logout(); err != nil {
		return false
	} else {
		return true
	}
}

func GetCurUser() (entity.User,bool) {
	if cu,err := entity.GetCurUser(); err != nil {
		return cu, false
	} else {
		return cu, true
	}
}

func UserLogin(username string, password string) bool {
	user := entity.QueryUser(func (u *entity.User) bool {
		if u.Name == username && u.Password == password {
			return true
		}
		return false
	})
	if len(user) == 0 {
		ErrLog.Println("Login: User not Exist")
		return false
	}
	entity.SetCurUser(&user[0])
	if err := entity.Sync(); err != nil {
		ErrLog.Println("Login: error occurred when set curuser")
		return false
	}
	return true
}

func UserRegister(username string, password string, email string, phone string) (bool, error) {
	user := entity.QueryUser(func (u *entity.User) bool {
		return u.Name == username
	})
	if len(user) == 1 {
		ErrLog.Println("User Register: Already exist username")
		return false, nil
	}
	entity.CreateUser(&entity.User{username, password, email, phone})
	if err := entity.Sync(); err != nil {
		return true, err
	}
	return true, nil
}


func DeleteUser(username string) bool {
	entity.DeleteUser(func (u *entity.User) bool {
		return u.Name == username
	})
	entity.UpdateMeeting(
		func(m *entity.Meeting) bool {
			return m.IsParticipator(username)
		},
		func(m *entity.Meeting) {
			m.DeleteParticipator(username)
		})
	entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == username || len(m.Participators) == 0
	})
	if err := entity.Sync(); err != nil {
		return false
	}
	return UserLogout()
}

func ListAllUser() []entity.User {
	return entity.QueryUser(func (u *entity.User) bool {
		return true
	})
}

func CreateMeeting(username string, title string, startTime string, endTime string, participator []string) bool {
	for _, i := range participator {
		if username == i {
			ErrLog.Println("Create Meeting: sponsor can't be participator")
			return false
		}
		l := entity.QueryUser(func (u *entity.User) bool{
			return u.Name == i
		})
		if (len(l) == 0) {
			ErrLog.Println("Create Meeting: no such a user : ", i)
			return false
		}
		dc := 0
		for _, j := range participator {
			if j == i {
				dc++
				if dc == 2 {
					ErrLog.Println("Create Meeting: duplicate participator")
					return false
				}
			}
		}
	}
	sTime := entity.StringToTime(startTime)
	eTime := entity.StringToTime(endTime)
	if sTime.IsValidTime() && eTime.IsValidTime() != true {
		ErrLog.Println("Create Meeting: Invalid time")
		return false
	}
	if eTime.LessThan(sTime) == true {
		ErrLog.Println("Create Meeting: Start Time greater than end time")
		return false
	}
	for _, p := range participator {
		l := entity.QueryMeeting(func (m *entity.Meeting) bool {
			if m.Sponsor == p || m.IsParticipator(p) {
				if (m.StartTime.LessThan(sTime) || m.StartTime.Equal(sTime)) && m.EndTime.MoreThan(sTime) {
					return true
				}
				if m.StartTime.LessThan(eTime) && (m.EndTime.MoreThan(eTime) || m.EndTime.Equal(eTime)) {
					return true
				}
				if (m.StartTime.MoreThan(sTime) || m.StartTime.Equal(sTime)) && (m.EndTime.LessThan(eTime) || m.EndTime.Equal(eTime)) {
					return true
				}
			}
			return false
		})
		if len(l) > 0 {
			ErrLog.Println("Create Meeting: ", p, " time conflict")
			return false
		}
	}
	tu := entity.QueryUser(func (u *entity.User) bool {
		return u.Name == username
	})
	if len(tu) == 0 {
		ErrLog.Println("Create Meeting: Sponsor ", username, " not exist")
		return false
	}
	l := entity.QueryMeeting(func (m *entity.Meeting) bool {
		if m.Sponsor == username || m.IsParticipator(username) {
			if (m.StartTime.LessThan(sTime) || m.StartTime.Equal(sTime)) && m.EndTime.MoreThan(sTime) {
				return true
			}
			if m.StartTime.LessThan(eTime) && (m.EndTime.MoreThan(eTime) || m.EndTime.Equal(eTime)) {
				return true
			}
			if (m.StartTime.MoreThan(sTime) || m.StartTime.Equal(sTime)) && (m.EndTime.LessThan(eTime) || m.EndTime.Equal(eTime)) {
				return true
			}
		}
		return false
	})

	if len(l) > 0 {
		ErrLog.Println("Create Meeting: ", username, " time conflict")
		return false;
	}
	entity.CreateMeeting(&entity.Meeting{Title: title, Sponsor: username, Participators: participator, StartTime: sTime, EndTime: eTime})
	if err := entity.Sync(); err != nil {
		return false
	}
	return true
}

func QueryMeeting(username, startTime, endTime string) ([]entity.Meeting,bool) {
	sTime := entity.StringToTime(startTime)
	eTime := entity.StringToTime(endTime)
	var m []entity.Meeting
	if sTime.IsValidTime() && eTime.IsValidTime() != true {
		ErrLog.Println("Query Meeting: Wrong StartDate")
		return m,false
	}
	if eTime.LessThan(sTime) == true {
		ErrLog.Println("Query Meeting: Start Time greater than end time")
		return m, false
	}

	tm := entity.QueryMeeting(func (m *entity.Meeting) bool {
		if m.Sponsor == username || m.IsParticipator(username) {
			if (m.StartTime.LessThan(sTime) || m.StartTime.Equal(sTime)) && m.EndTime.MoreThan(sTime) {
				return true
			}
			if m.StartTime.LessThan(eTime) && (m.EndTime.MoreThan(eTime) || m.EndTime.Equal(eTime)) {
				return true
			}
			if (m.StartTime.MoreThan(sTime) || m.StartTime.Equal(sTime)) && (m.EndTime.LessThan(eTime) || m.EndTime.Equal(eTime)) {
				return true
			}
		}
		return false
	})
	return tm,true
}

func DeleteMeeting(username, title string) int {
	return entity.DeleteMeeting(func (m *entity.Meeting) bool {
		return m.Sponsor == username && m.Title == title
	})
}

func QuitMeeting(username string, title string) bool {
	flag :=entity.QueryMeeting(func (m *entity.Meeting) bool {
		return m.Title == title && m.IsParticipator(username) == true
	})
	if len(flag) == 0 {
		return false
	}
	entity.UpdateMeeting(func (m *entity.Meeting) bool {
		return m.IsParticipator(username) == true && m.Title == title
	}, func (m *entity.Meeting) {
		m.DeleteParticipator(username)
	})
	entity.DeleteMeeting(func (m *entity.Meeting) bool {
		return len(m.Participators) == 0
	})
	return true
}

func ClearMeeting(username string) (int, bool) {
	cm := entity.DeleteMeeting(func (m *entity.Meeting) bool {
		return m.Sponsor == username
	})
	if err := entity.Sync(); err != nil {
		ErrLog.Println("Clear Meeting: Delete failed")
		return cm, false
	} else {
		return cm, true
	}
}

func AddMeetingParticipator(username string, title string, participators []string) bool {
	for _, p := range participators {
		uc := entity.QueryUser(func (u *entity.User) bool {
			return u.Name == p
		})
		if len(uc) == 0 {
			ErrLog.Println("Add Meeting Participator: No such a user: ", p)
			return false
		}
		qm := entity.QueryMeeting(func (m *entity.Meeting) bool {
			return m.Sponsor == username && m.Title == title && m.IsParticipator(p)
		})
		if len(qm) != 0 {
			ErrLog.Println("Add Meeting Participator: ",p, "Already in meeting")
			return false
		}
	}
	mt := entity.UpdateMeeting(func (m *entity.Meeting) bool {
		return m.Sponsor == username && m.Title == title
	}, func (m *entity.Meeting) {
		for _,p := range participators {
			m.AddParticipator(p)
		}
	})
	if mt == 0 {
		ErrLog.Println("Add Meeting Participator: no such meeting")
		return false
	}
	if err := entity.Sync(); err != nil {
		return false
	}
	return true
}

func RemoveMeetingParticipator(username string, title string, participators []string) bool {
	for _, p := range participators {
		uc := entity.QueryUser(func (u *entity.User) bool {
			return u.Name == p
		})
		if len(uc) == 0 {
			ErrLog.Println("Remove Meeting Participator: No such a user: ", p)
			return false
		}
		qm := entity.QueryMeeting(func (m *entity.Meeting) bool {
			return m.Sponsor == username && m.Title == title  && m.IsParticipator(p)
		})
		if len(qm) == 0 {
			ErrLog.Println("Remove Meeting Participator: Not in Meeting :", p)
			return false
		}
	}
	mt := entity.UpdateMeeting(func (m *entity.Meeting) bool {
		return m.Sponsor == username && m.Title == title
	}, func (m *entity.Meeting) {
		for _ , p := range participators {
			m.DeleteParticipator(p)
		}
	})
	if mt == 0 {
		ErrLog.Println("Remove Meeting Participator: no such a meeting: ", title)
		return false
	}
	entity.DeleteMeeting(func(m *entity.Meeting) bool {
		return m.Sponsor == username && len(m.Participators) == 0
	})
	if err := entity.Sync(); err != nil {
		return false
	}
	return true
}