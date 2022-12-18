package person

import (
	"bytes"
	"encoding/json"
	api "github.com/Hooneats/go-gin-pr4/common"
	"github.com/Hooneats/go-gin-pr4/model"
	"github.com/Hooneats/go-gin-pr4/model/person"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func GetMockPerson() *person.Person {
	return &person.Person{
		Name: "test",
		Age:  10,
		Pnum: "01011112222",
	}
}

func GetMockPersons() []*person.Person {
	return []*person.Person{
		GetMockPerson(),
		GetMockPerson(),
	}
}

func GetMockWebPerson() *WebPerson {
	return &WebPerson{
		Name: "test",
		Age:  10,
		Pnum: "01011112222",
	}
}

func GetMockWebPersons() []*WebPerson {
	return []*WebPerson{
		GetMockWebPerson(),
		GetMockWebPerson(),
	}
}

type MockModel struct {
	collection *mongo.Collection
}

func (m *MockModel) GetCollection(collection string) *mongo.Collection {
	return nil
}
func (m *MockModel) CreateIndex(colName string, indexName ...string) {
	return
}

type MockPersonModel struct {
}

func GetPersonModel(mod model.Modeler) *MockPersonModel {
	return nil
}

func (m *MockPersonModel) FindByName(name string) (*person.Person, error) {
	return GetMockPerson(), nil
}

func (m *MockPersonModel) FindByPnum(pnum string) (*person.Person, error) {
	return GetMockPerson(), nil
}
func (m *MockPersonModel) FindAll() ([]*person.Person, error) {
	return GetMockPersons(), nil
}
func (m *MockPersonModel) InsertOne(person *person.Person) (*person.Person, error) {
	return GetMockPerson(), nil
}
func (m *MockPersonModel) DeleteByPnum(pnum string) error {
	return nil
}
func (m *MockPersonModel) UpdateAgeByPnum(age int, pnum string) error {
	return nil
}
func setRouterAndPCon() (*gin.Engine, *PersonControl) {
	r := gin.Default()
	mkmodel := &MockModel{}
	mkPersonmodel := GetPersonModel(mkmodel)
	pControl := GetPersonControl(mkPersonmodel)
	return r, pControl
}

func TestPersonControl_UNIT(t *testing.T) {
	r, pControl := setRouterAndPCon()

	t.Run("name 으로 Person 데이터 가져오기", func(t *testing.T) {
		//c, _ := gin.CreateTestContext(httptest.NewRecorder())
		r.GET("/v1/persons/name", pControl.GetByName)
		mockWebPerson := GetMockWebPerson()
		mockResponse := api.SuccessData(mockWebPerson)
		expectedData, _ := json.Marshal(mockResponse) // Json 으로

		req, _ := http.NewRequest("GET", "/v1/persons/name", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		rep, _ := io.ReadAll(w.Body)
		log.Println("expected : " + string(expectedData))
		log.Println("actual : " + string(rep))
		assert.Equal(t, expectedData, rep)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("pnum 으로 Person 데이터 가져오기", func(t *testing.T) {
		r.GET("/v1/persons/pnum", pControl.GetByPnum)
		mockWebPerson := GetMockWebPerson()
		mockResponse := api.SuccessData(mockWebPerson)
		expectedData, _ := json.Marshal(mockResponse)

		req, _ := http.NewRequest("GET", "/v1/persons/pnum", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		rep, _ := io.ReadAll(w.Body)
		log.Println("expected : " + string(expectedData))
		log.Println("actual : " + string(rep))
		assert.Equal(t, expectedData, rep)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Person 데이터 전부 가져오기", func(t *testing.T) {
		r.GET("/v1/persons", pControl.GetAll)
		mockWebPerson := GetMockWebPersons()
		mockResponse := api.SuccessData(mockWebPerson)
		expectedData, _ := json.Marshal(mockResponse)

		req, _ := http.NewRequest("GET", "/v1/persons", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		rep, _ := io.ReadAll(w.Body)
		log.Println("expected : " + string(expectedData))
		log.Println("actual : " + string(rep))
		assert.Equal(t, expectedData, rep)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Person 데이터 Post One 요청", func(t *testing.T) {
		r.POST("/v1/persons/person", pControl.PostOne)
		mockWebPerson := GetMockWebPerson()
		mockResponse := api.SuccessData(mockWebPerson)
		expectedData, _ := json.Marshal(mockResponse)

		marshal, _ := json.Marshal(mockWebPerson)
		req, _ := http.NewRequest("POST", "/v1/persons/person", bytes.NewBuffer(marshal))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		rep, _ := io.ReadAll(w.Body)
		log.Println("expected : " + string(expectedData))
		log.Println("actual : " + string(rep))
		assert.Equal(t, expectedData, rep)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Person pnum 으로 삭제 요청: 본 테스트는 Database 의존성이 없는 Mock 활용 Controller 단위 테스트로 항상 통과", func(t *testing.T) {
		r.DELETE("/v1/persons/person", pControl.DeleteByPnum)
		mockWebPerson := GetMockWebPerson()
		mockResponse := api.Success()
		expectedData, _ := json.Marshal(mockResponse)
		req, _ := http.NewRequest("DELETE", "/v1/persons/person?pnum="+mockWebPerson.Pnum, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		rep, _ := io.ReadAll(w.Body)
		log.Println("expected : " + string(expectedData))
		log.Println("actual : " + string(rep))
		assert.Equal(t, expectedData, rep)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("pnum 으로 person 데이타 age 수정", func(t *testing.T) {
		r.PUT("/v1/persons/person", pControl.PutAgeByPnum)
		mockWebPerson := GetMockWebPerson()
		mockResponse := api.Success()
		expectedData, _ := json.Marshal(mockResponse)
		req, _ := http.NewRequest("PUT", "/v1/persons/person?age=20&pnum="+mockWebPerson.Pnum, nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)
		rep, _ := io.ReadAll(w.Body)
		log.Println("expected : " + string(expectedData))
		log.Println("actual : " + string(rep))
		assert.Equal(t, expectedData, rep)
		assert.Equal(t, http.StatusOK, w.Code)
	})

}
