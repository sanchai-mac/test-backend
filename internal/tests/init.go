package tests

import (
	"log"
	"reflect"
	"test-backend/internal/config"
	container_mock "test-backend/internal/container/tests/mocks"
	"test-backend/internal/infrastructure/database"
	"test-backend/internal/infrastructure/server"
	"testing"
)

type Mock struct {
	MockFilePaths  []string
	MockCsvPath    []CsvMock
	Topic          *string
	MockFilesKafka *string
}

type CsvMock struct {
	Type     reflect.Type
	FilePath string
}

type Calls [][]interface{}
type AssertCalled struct {
	FnName        string
	IType         reflect.Type
	Calls         Calls
	NumberOfCalls *int
}

type Expected struct {
	Error     error
	Called    []AssertCalled
	NotCalled []AssertCalled
}

type Testcase map[string]struct {
	Mock     *Mock
	Input    interface{}
	Expected interface{}
	Skipped  bool
}

type Client struct {
}

type SetupOptions struct {
	MockSuiteDB func()
	MockClient  bool
	Mock        *Mock
}

type TestTools struct {
	t        testing.TB
	Teardown func(tb testing.TB)
	C        *container_mock.TestContainer
	Client   *Client
}

func SetupSuite(tb testing.TB, opts SetupOptions) TestTools {
	log.Println("setup suite")
	// Setup testsuite base data
	c := container_mock.NewTestContainer(tb)

	c.RequireInvoke(func(db *database.DB, s *server.Server) {
		cleanMockData(db, nil)
		if opts.MockSuiteDB != nil {

		}

		if opts.Mock != nil {
			mockDbWithJson(tb, c, db, opts.Mock.MockFilePaths)
		}
	})
	return TestTools{
		t: tb,
		Teardown: func(tb testing.TB) {
			log.Println("teardown suite")

			// Clean mock setup testsuite
			c.RequireInvoke(func(db *database.DB, s *server.Server) {

			})
		},
		C: c,
	}
}

func SetupTest(tb testing.TB, opts SetupOptions) TestTools {
	log.Println("setup test")
	client := Client{}
	// Setup testcase
	c := container_mock.NewTestContainer(tb)
	// Clean DB
	c.RequireInvoke(func(db *database.DB) {
		cleanMockData(db, nil)
	})

	m := opts.Mock

	// Mock testcase functions
	decorateMock(c, m)

	// Mock initialize data
	c.RequireInvoke(func(db *database.DB, s *server.Server, cfg *config.Configuration) {
		if m != nil {
			//Mock Producer
			// Mock with json file
			mockDbWithJson(tb, c, db, m.MockFilePaths)

			// Mock with CSV file
			for _, mock := range m.MockCsvPath {
				mockDbWithCsv(tb, c, db, mock)
			}
		}

		if opts.MockClient {
			go func() {
				if err := s.StartRestful(); err != nil {
					log.Println("[Container: Run] Start is panic: ", err)
					panic(err)
				}

			}()
		}
	})

	return TestTools{
		t: tb,
		Teardown: func(tb testing.TB) {
			log.Println("teardown test")
			c.RequireInvoke(func(db *database.DB, s *server.Server) {
			})
		},
		C:      c,
		Client: &client,
	}
}

func mockDbWithJson(tb testing.TB, c *container_mock.TestContainer, db *database.DB, mockFilePaths []string) {
	for _, path := range mockFilePaths {
		c.JsonToStruct(path, db)
	}
}

func mockDbWithCsv(tb testing.TB, c *container_mock.TestContainer, db *database.DB, mock CsvMock) {

}

func decorateMock(c *container_mock.TestContainer, m *Mock) {
	if m == nil {
		return
	}
}

func cleanMockData(db *database.DB, ignore []string) {
}

// Assertion
func (t TestTools) ExpectedCalls(called []AssertCalled) {
}
