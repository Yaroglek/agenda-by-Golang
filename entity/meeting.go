package entity

type Meeting struct {
	Title string
	Sponsor string
	Participators []string
	StartTime, EndTime Time
}

func (meeting Meeting) IsParticipator(username string) bool {
	for i := 0; i < len(meeting.Participators); i++ {
		if meeting.Participators[i] == username {
	    	return true
		}
	}
	return false
}

func (meeting *Meeting) AddParticipator(username string) bool {
	if meeting.Sponsor == username || meeting.IsParticipator(username) {
		return false
	}
	meeting.Participators = append(meeting.Participators, username)
	return true
}

func (meeting *Meeting) DeleteParticipator(username string) {
	for i := 0; i < len(meeting.Participators); i++ {
		if meeting.Participators[i] == username {
			meeting.Participators = append(meeting.Participators[:i], meeting.Participators[i + 1 :]...)
			break
		}
	}
}