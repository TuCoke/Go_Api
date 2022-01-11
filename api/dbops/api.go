package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"main/api/defs"
	"main/api/utils"
	"time"
)

func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare(`insert into users (login_name,pwd) values(?,?)`)
	if err != nil {
		return err
	}
	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil
}

func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name=?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string

	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil
}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM USERS WHERE login_name=? and pwd =?")
	if err != nil {
		log.Printf("DeleteUser: %s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}

//region vieDeo
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006,15:04:05") // m d y , hh:mm:ss
	stmtIns, err := dbConn.Prepare("insert into video_info (id,author_id,name,display_ctime) values(?,?,?,?)")
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmtIns.Close()
	return res, err
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("select author_id,name,display_ctime from video_info where id =?")
	var aid int
	var dct string
	var name string
	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	defer stmtOut.Close()
	res := &defs.VideoInfo{vid, aid, name, dct}
	return res, err
}

func DeleteVideoInfo(vid string)  error{
	stmtDel,err:=dbConn.Prepare("Delete FROM video_info WHERE id=?")
	_,err = stmtDel.Exec(vid)
	if err !=nil{
		return  err
	}
	defer stmtDel.Close()
	return nil
}


//endregion
