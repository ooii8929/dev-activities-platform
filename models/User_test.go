package models

import (
	"testing"
)



	// Setup
func Test_UserGet_1(t *testing.T) {
			creds := &User{
				Email:"1@gmail.com",
			}

			actual := &User{
				Email:"1@gmail.com",
				Name:"123",
				Password:"123123",
			}

			// if i, e := creds.UserGetByEmail(creds.Email); i.Name != actual.Name || e != nil { //try a unit test on function
			// 		t.Error("除法函式測試沒透過") // 如果不是如預期的那麼就報錯
			// } else {
			// 		t.Log("第一個測試通過了") //記錄一些你期望記錄的資訊
			// }

			actualString, e := creds.UserGetByEmail(creds.Email)
			expectedString := actual.Password


			if actualString.Password != expectedString{
					t.Error("Expected String", e)
			}else{
				t.Log("第一個測試通過了")
			}
	}
