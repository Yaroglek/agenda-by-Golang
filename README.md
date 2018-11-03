# CLI 命令行实用程序开发实战 - Agenda
## 数据科学与计算机学院 2016级软件工程(数字媒体)
## 16340299 赵博然
## 16340300 赵佳乐
## 16340297 张子轩
## 16340312 周泽昊
---
agenda命令见`cmd-design.md`.

---
安装.
```
[yaroglek@centos-new agenda]$ go install agenda
```
---
测试. 包括全部命令的基本测试, 以及一些有逻辑冲突的错误测试(会议时间冲突, 用户重名, 未登录等问题). 
```
[yaroglek@centos-new agenda]$ agenda r -u yaroglek -p yaroglek -e king_yaroglek@163.com -c 1562613471
Successfully register!
[yaroglek@centos-new agenda]$ agenda r -u pudge -p pudge -e pudge@pudge.com -c 110
Successfully register!
[yaroglek@centos-new agenda]$ agenda r -u king -p king -e king@king.com -c 119
Successfully register!
[yaroglek@centos-new agenda]$ agenda qu
----------------
Username: yaroglek
Phone: 1562613471
Email: king_yaroglek@163.com
----------------
----------------
Username: pudge
Phone: 110
Email: pudge@pudge.com
----------------
----------------
Username: king
Phone: 119
Email: king@king.com
----------------
[yaroglek@centos-new agenda]$ agenda cm -t N1 -s 2018-12-02/14:00 -e 2018-12-02/17:00 -u pudge
Error: please login
[yaroglek@centos-new agenda]$ agenda l -u yaroglek -p yaroglek
Login Successfully. Current User: yaroglek
[yaroglek@centos-new agenda]$ agenda cm -t N1 -s 2018-12-02/14:00 -e 2018-12-02/17:00 -u pudge
Create meeting successfully!
[yaroglek@centos-new agenda]$ agenda cm -t N2 -s 2018-07-01/14:00 -e 2018-07-01/17:00 -u pudge,king
Create meeting successfully!
[yaroglek@centos-new agenda]$ agenda au -t N1 -u king
Unexpected error. Check error.log for detail
[yaroglek@centos-new agenda]$ agenda au -t N1 -u yaroglek
Unexpected error. Check error.log for detail
[yaroglek@centos-new agenda]$ agenda cm -t ogori -s 2018-12-02/16:00 -e 2018-12-02/16:30 -u pudge
Error: create Failed. Please check error.log for more detail
[yaroglek@centos-new agenda]$ agenda lo
Logout Successfully
[yaroglek@centos-new agenda]$ agenda l -u pudge -p pudge
Login Successfully. Current User: pudge
[yaroglek@centos-new agenda]$ agenda cm -t ogori -s 2018-12-02/16:00 -e 2018-12-02/16:30 -u yaroglek
Error: create Failed. Please check error.log for more detail
[yaroglek@centos-new agenda]$ agenda qm -t N2
Quit Successfully
[yaroglek@centos-new agenda]$ agenda qm -t N2
Error: Meeting not exist or you're not a participator of it
[yaroglek@centos-new agenda]$ agenda dm -t N1
Error: Meeting not exist or you're not a Sponsor of it
[yaroglek@centos-new agenda]$ agenda lo
Logout Successfully
[yaroglek@centos-new agenda]$ agenda l -u yaroglek -p yaroglek
Login Successfully. Current User: yaroglek
[yaroglek@centos-new agenda]$ agenda dm -t N1
Delete Successfully
[yaroglek@centos-new agenda]$ agenda dm -t N1
Error: Meeting not exist or you're not a Sponsor of it
[yaroglek@centos-new agenda]$ agenda lo
Logout Successfully
[yaroglek@centos-new agenda]$ agenda l -u king -p king
Login Successfully. Current User: king
[yaroglek@centos-new agenda]$ agenda fm -s 2018-01-01/00:00 -e 2018-12-31/23:59
----------------
Title: N2
Start Time: 2018-07-01/14:00
End Time: 2018-07-01/17:00
Participator(s): king
----------------
[yaroglek@centos-new agenda]$ agenda qm -t N2
Quit Successfully
[yaroglek@centos-new agenda]$ agenda lo
Logout Successfully
[yaroglek@centos-new agenda]$ agenda l -u yaroglek -p yaroglek
Login Successfully. Current User: yaroglek
[yaroglek@centos-new agenda]$ agenda fm -s 2018-01-01/00:00 -e 2018-12-31/23:59
[yaroglek@centos-new agenda]$ agenda dam
Successfully deleted 0 meeting(s)
[yaroglek@centos-new agenda]$ agenda cm -t ogori -s 2018-12-02/16:00 -e 2018-12-02/16:30 -u pudge,king
Create meeting successfully!
[yaroglek@centos-new agenda]$ agenda cm -t N5 -s 2018-12-02/14:00 -e 2018-12-02/16:00 -u pudge,king
Create meeting successfully!
[yaroglek@centos-new agenda]$ agenda dam
Successfully deleted 2 meeting(s)
[yaroglek@centos-new agenda]$ agenda du
Delete Sucessfully
[yaroglek@centos-new agenda]$ agenda du
Error: please login
[yaroglek@centos-new agenda]$ agenda qu
----------------
Username: king
Phone: 119
Email: king@king.com
----------------
----------------
Username: pudge
Phone: 110
Email: pudge@pudge.com
----------------
```