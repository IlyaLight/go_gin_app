package models

import (
	"github.com/gin-gonic/gin"
	"os"
	"testing"
)

// This function is used to do setup before executing the test functions
func TestMain(m *testing.M) {
	//Set Gin to Test Mode
	gin.SetMode(gin.TestMode)

	// Run the other tests
	os.Exit(m.Run())
}

// Helper function to create a router during testing
func getRouter(withTemplates bool) *gin.Engine {
	r := gin.Default()
	if withTemplates {
		r.LoadHTMLGlob("templates/*")
		//r.Use(setUserStatus())
	}
	return r
}

//// Test the function that fetches all articles
//func TestGetAllArticles(t *testing.T) {
//	alist := models.GetAllArticles()
//
//	// Check that the length of the list of articles returned is the
//	// same as the length of the global variable holding the list
//	if len(alist) != len(models.GetAllArticles()) {
//		t.Fail()
//	}
//
//	// Check that each member is identical
//	//for i, v := range alist {
//	//	if v.Content != articleList[i].Content ||
//	//		v.ID != articleList[i].ID ||
//	//		v.Title != articleList[i].Title {
//	//
//	//		t.Fail()
//	//		break
//	//	}
//	//}
//}
//
//// Test the function that fetche an Article by its ID
//func TestGetArticleByID(t *testing.T) {
//	a, err := getArticleByID(1)
//
//	if err != nil || a.ID != 1 || a.Title != "Article 1" || a.Content != "Article 1 body" {
//		t.Fail()
//	}
//}
