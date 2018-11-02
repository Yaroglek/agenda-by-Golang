package entity

import (
	"os"
	"io"
	"bufio"
	"path/filepath"
	"errors"
	"log"
	"encoding/json"
)

type UserFilter func (*User) bool
type MeetingFilter func (*Meeting) bool

var userinfoPath = "/src/agenda/data/user"
var meetinginfoPath = "/src/agenda/data/meeting"
var curUserPath = "/src/agenda/data/curUser.txt"

var curUsername *string;

var dirty bool

var userData []User
var meetingData []Meeting

var errLog *log.Logger

func init()  {
	errLog = loghelper.Error
	dirty = false
	userinfoPath = filepath.Join(loghelper.GoPath, userinfoPath)
	meetinginfoPath = filepath.Join(loghelper.GoPath, meetinginfoPath)
	curUserPath = filepath.Join(loghelper.GoPath, curUserPath)
	if err := readFromFile(); err != nil {
		errLog.Println("readFromFile fail: ", err)
	}
}

func logout() error {
	curUsername = nil
	return sync()
}

func sync() error {
	if err := writeToFile(); err != nil {
		errLog.Println("writeToFile fail:", err)
		return err
	}
	return nil
}

func createUser(v User) {
	userData = append(userData, v)
	dirty = true
}

func queryUser(filter UserFilter) []User {
	var user []User
	for _, v := range userData {
		if filter(&v) {
			user = append(user, v)
		}
	}
	return user
}

func updateUser(filter UserFilter, switcher func (*User)) int {
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

func deleteUser(filter UserFilter) int {
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

func createMeeting(v Meeting) {
	meetingData = append(meetingData, v)
	dirty = true
}

func queryMeeting(filter MeetingFilter) (meeting []Meeting) {
	for _, v := range meetingData {
		if filter(&v) {
			meeting = append(meeting, v)
		}
	}
	return meeting;
}

func updateMeeting(filter MeetingFilter, switcher func (*Meeting)) int {
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

func deleteMeeting(filter MeetingFilter) int {
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

func getCurUser() (User, error) {
	if curUsername == nil {
		return User{}, errors.New("Current user does not exist")
	}
	for _ , v := range userData {
		if v.name == *curUsername {
			return v, nil
		}
	}
	return User{}, errors.New("Current user does not exist")
}

func setCurUser(u *User) {
	if u == nil {
		curUsername = nil
		return
	}
	if (curUsername == nil) {
		p := u.name
		curUsername = &p
	} else {
		*curUsername = u.name
	}
}

func readFromFile() error {
	var e []error
	str, err1 := readString(curUserPath)
	if err1 != nil {
		e = append(e, err1)
	}
	curUsername = str
	if err := readUser(); err != nil {
		e = append(e, err)
	}
	if err := readMeeting(); err != nil {
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

func writeToFile() error {
	var e []error
	if err := writeString(curUserPath, curUsername); err != nil {
		e = append(e, err)
	}
	if dirty {
		if err := writeJSON(userinfoPath, userData); err != nil {
			e = append(e, err)
		}
		if err := writeJSON(meetinginfoPath, meetingData); err != nil {
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

func readUser() error {
	file, err := os.Open(userinfoPath);
	if err != nil {
		errLog.Println("Open File Fail: ", userinfoPath, err)
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	switch err := dec.Decode(&userData); err {
	case nil, io.EOF:
		return nil
	default:
		errLog.Println("Decode User Fail: ", err)
		return err
	}
}

func readMeeting() error {
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

func writeJSON(fpath string, data interface{}) error {
	file, err := os.Create(fpath);
	if err != nil {
		return err
	}
	defer file.Close()
	enc := json.NewEncoder(file)
	if err := enc.Encode(&data); err != nil {
		errLog.Println("writeJSON: ", err)
		return err
	}
	return nil
}

func writeString(path string, data *string) error {
	file, err := os.Create(path)
	if err != nil {
		loghelper.Error.Println("Create file error: ", path)
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	if data != nil {
		if _, err := writer.WriteString(*data); err != nil {
			loghelper.Error.Println("Write file fail:", path)
			return err
		}
	}
	if err := writer.Flush(); err != nil {
		loghelper.Error.Println("Flush file fail:", path)
		return err
	}
	return nil
}

func readString(path string) (*string, error) {
	file, err := os.Open(path)
	if err != nil {
		loghelper.Error.Println("Open file error:", path)
		return nil, err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	str, err := reader.ReadString('\n');
	if err != nil && err != io.EOF {
		loghelper.Error.Println("Read file fail:", path)
		return nil, err
	}
	return &str, nil
}