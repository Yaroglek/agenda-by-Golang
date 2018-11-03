package entity

import (
	"os"
	"io"
	"bufio"
	"path/filepath"
	"errors"
	"log"
	"encoding/json"
	"agenda/logger"
)

type UserFilter func (*User) bool
type MeetingFilter func (*Meeting) bool

var userinfoPath = "/src/agenda/data/user"
var meetinginfoPath = "/src/agenda/data/meeting"
var curUserPath = "/src/agenda/data/curUser"

var curUsername *string;

var dirty bool

var userData []User
var meetingData []Meeting

var errLog *log.Logger

func init()  {
	errLog = logger.Error
	dirty = false
	userinfoPath = filepath.Join(logger.GoPath, userinfoPath)
	meetinginfoPath = filepath.Join(logger.GoPath, meetinginfoPath)
	curUserPath = filepath.Join(logger.GoPath, curUserPath)
	if err := ReadFromFile(); err != nil {
		errLog.Println("readFromFile fail:", err)
	}
}

func Logout() error {
	curUsername = nil
	return Sync()
}

func Sync() error {
	if err := WriteToFile(); err != nil {
		errLog.Println("writeToFile fail:", err)
		return err
	}
	return nil
}

func CreateUser(user *User) {
	userData = append(userData, *user)
	dirty = true
}

func QueryUser(filter UserFilter) []User {
	var user []User
	for _, v := range userData {
		if filter(&v) {
			user = append(user, v)
		}
	}
	return user
}

func UpdateUser(filter UserFilter, switcher func (*User)) int {
	count := 0
	for i := 0; i < len(userData); i++ {
		if v := &userData[i]; filter(v) {
			switcher(v)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

func DeleteUser(filter UserFilter) int {
	count := 0
	length := len(userData)
	for i := 0; i < length; {
		if filter(&userData[i]) {
			length--
			userData[i] = userData[length]
			userData = userData[:length]
			count++
		} else {
			i++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

func CreateMeeting(meeting *Meeting) {
	meetingData = append(meetingData, *meeting)
	dirty = true
}

func QueryMeeting(filter MeetingFilter) (meeting []Meeting) {
	for _, v := range meetingData {
		if filter(&v) {
			meeting = append(meeting, v)
		}
	}
	return meeting;
}

func UpdateMeeting(filter MeetingFilter, switcher func (*Meeting)) int {
	count := 0
	for i := 0; i < len(meetingData); i++ {
		if v := &meetingData[i]; filter(v) {
			switcher(v)
			count++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

func DeleteMeeting(filter MeetingFilter) int {
	count := 0
	length := len(meetingData)
	for i := 0; i < length; {
		if filter(&meetingData[i]) {
			length--
			meetingData[i] = meetingData[length]
			meetingData = meetingData[:length]
			count++
		} else {
			i++
		}
	}
	if count > 0 {
		dirty = true
	}
	return count
}

func GetCurUser() (User, error) {
	if curUsername == nil {
		return User{}, errors.New("Current user does not exist")
	}
	for _ , v := range userData {
		if v.Name == *curUsername {
			return v, nil
		}
	}
	return User{}, errors.New("Current user does not exist")
}

func SetCurUser(u *User) {
	if u == nil {
		curUsername = nil
		return
	}
	if (curUsername == nil) {
		p := u.Name
		curUsername = &p
	} else {
		*curUsername = u.Name
	}
}

func ReadFromFile() error {
	var e []error
	str, err1 := ReadString(curUserPath)
	if err1 != nil {
		e = append(e, err1)
	}
	curUsername = str
	if err := ReadUser(); err != nil {
		e = append(e, err)
	}
	if err := ReadMeeting(); err != nil {
		e = append(e, err)
	}
	if len(e) == 0 {
		return nil
	}
	result := e[0]
	for i := 1; i < len(e); i++ {
		result = errors.New(result.Error() + e[i].Error())
	}
	return result
}

func WriteToFile() error {
	var e []error
	if err := WriteString(curUserPath, curUsername); err != nil {
		e = append(e, err)
	}
	if dirty {
		if err := WriteJSON(userinfoPath, userData); err != nil {
			e = append(e, err)
		}
		if err := WriteJSON(meetinginfoPath, meetingData); err != nil {
			e = append(e, err)
		}
	}
	if len(e) == 0 {
		return nil
	}
	result := e[0]
	for i := 1; i < len(e); i++ {
		result = errors.New(result.Error() + e[i].Error())
	}
	return result
}

func ReadUser() error {
	file, err := os.Open(userinfoPath);
	if err != nil {
		errLog.Println("Open File Fail:", userinfoPath, err)
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	switch err := dec.Decode(&userData); err {
	case nil, io.EOF:
		return nil
	default:
		errLog.Println("Decode User Fail:", err)
		return err
	}
}

func ReadMeeting() error {
	file, err := os.Open(meetinginfoPath);
	if err != nil {
		errLog.Println("Open File Fail:", meetinginfoPath, err)
		return err
	}
	defer file.Close()
	dec := json.NewDecoder(file)
	switch err := dec.Decode(&meetingData); err {
	case nil, io.EOF:
		return nil
	default:
		errLog.Println("Decode Met Fail:", err)
		return err
	}
}

func WriteJSON(fpath string, data interface{}) error {
	file, err := os.Create(fpath);
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	if err := enc.Encode(&data); err != nil {
		errLog.Println("writeJSON:", err)
		return err
	}
	return nil
}

func WriteString(path string, data *string) error {
	file, err := os.Create(path)
	if err != nil {
		logger.Error.Println("Create file error:", path)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if data != nil {
		if _, err := writer.WriteString(*data); err != nil {
			logger.Error.Println("Write file fail:", path)
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		logger.Error.Println("Flush file fail:", path)
		return err
	}
	return nil
}

func ReadString(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		logger.Error.Println("Open file error:", path)
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n');
	if err != nil && err != io.EOF {
		logger.Error.Println("Read file fail:", path)
		return nil, err
	}
	return &str, nil
}