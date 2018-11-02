package entity

type Meeting struct {
	title string
	sponsor string
	participators []string
	startTime, endTime Time
}

func (meeting Meeting) isParticipator(username string) bool {
	for i := 0; i < len(meeting.participators); i++ {
		if meeting.participators[i] == username {
	    	return true
		}
	}
	return false
}

func (meeting *Meeting) addParticipator(username string) bool {
	if meeting.sponsor == username || meeting.isParticipator(username) {
		return false
	}
	meeting.participators = append(meeting.participators, username)
	return true
}

func (meeting *Meeting) deleteParticipator(username string) {
	for i := 0; i < len(meeting.participators); i++ {
		if meeting.participators[i] == username {
			meeting.participators = append(meeting.participators[:i], meeting.participators[i + 1 :]...)
		}
	}
}