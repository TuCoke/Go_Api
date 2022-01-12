package dbops

import "testing"

var tempvid string

func clearTables() {
	dbConn.Exec("truncate users")
	dbConn.Exec("truncate video_info")
	dbConn.Exec("truncate comments")
	dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {

}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUser)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testRegetUser)
}

// region user
func testRegetUser(t *testing.T) {
	pwd, err := GetUserCredential("ces")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

func testGetUser(t *testing.T) {
	pwd, err := GetUserCredential("ces")
	if pwd != "123" || err != nil {
		t.Errorf("Error of GetUser")
	}
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("ces", "123")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("ces", "123")
	if err != nil {
		t.Errorf("error of DeleteUser: %v", err)
	}
}

// endregion
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUser)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetUser)
}

// region Video
func testReagetVideoInfo(t *testing.T) {
	video, err := GetVideoInfo("1")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
	if video != nil {
		t.Errorf("Deleting user test failed")
	}
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo("1")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of add viedeoInfo, %v", err)
	}
	tempvid = vi.Id
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo("1234")
	if err != nil {
		t.Errorf("error of DeleteVideoInfo: %v", err)
	}
}

func testAddUser(t *testing.T) {
	err := AddUserCredential("ces", "123")
	if err != nil {
		t.Errorf("Error of RegetUser: %v", err)
	}
}

// endregion

func TestComments(t *testing.T) {
	clearTables()
	//t.Run("AddUser")
}

func testAddComments(t *testing.T) {
	vid := "123"
	aid := 1
	content := "test 测试"
	err := AddNewComments(vid, aid, content)
	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func TestListComments(t *testing.T) {

}