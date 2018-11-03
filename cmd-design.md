```
au          add a user to a meeting
cm          create a meeting
dam         delete all meetings
dm          delete a meeting
du          delete a user
fm          find meetings that are bewteen starttime and endtime
help        Help about any command
l           login
lo          logout
qm          quit a meeting
qu          query user
r           register
ru          remove user(s) from a meeting
```
---
`au`, add a user to a meeting
```
-t, --title string   the title of meeting
-u, --user strings   user(s) you want to add, input like "name1, name2"
```
---
`cm`, create a meeting
```
-e, --endtime string     the endTime of the meeting
-s, --starttime string   the startTime of the meeting
-t, --title string       the title of meeting
-u, --user strings       the user(s) of the meeting, input like "name1, name2"
```
---
`dam`, delete all meetings

---
`dm`, delete a meeting
```
-t, --title string   the title of meeting
```
---
`du`, delete a user

---
`fm`, find a meeting that are bewteen starttime and endtime
```
-e, --endtime string     the end time of the meeting
-s, --starttime string   the start time of the meeting
```
---
`l`, login
```
-p, --password string   agenda password
-u, --username string   agenda username
```
---
`lo`, logout

---
`qm`, quit a meeting
```
-t, --title string   the title of meeting
```
---
`qu`, query users

---
`r`, register
```
-c, --cellphone string   cellphone
-e, --email string       email
-p, --password string    password
-u, --username string    username
```
---
`ru`, remove user(s) from a meeting
```
-t, --title string   the title of the meeting
-u, --user strings   the user(s) in a meeting, input like "name1, name2"
```
---