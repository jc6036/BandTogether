package user_controller

import (
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
)

// queryDb
func TestQueryDB_NoRows_ReturnsMessage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer db.Close()

	query := "SELECT id, jret FROM t WHERE id=0"
	rows := sqlmock.NewRows([]string{"id", "jret"}) // no rows
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	got, err := queryDB(db, query)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "No rows found!" {
		t.Fatalf("got %q, want %q", got, "No rows found!")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestQueryDB_SingleRow_ReturnsValue(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer db.Close()

	query := "SELECT id, jret FROM t WHERE id=1"
	rows := sqlmock.NewRows([]string{"id", "jret"}).
		AddRow(int64(1), "only")
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	got, err := queryDB(db, query)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "only" {
		t.Fatalf("got %q, want %q", got, "only")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestQueryDB_MultipleRows_ReturnsLast(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer db.Close()

	query := "SELECT id, jret FROM t ORDER BY id"
	rows := sqlmock.NewRows([]string{"id", "jret"}).
		AddRow(int64(1), "first").
		AddRow(int64(2), "second").
		AddRow(int64(3), "third")
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	got, err := queryDB(db, query)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got != "third" {
		t.Fatalf("got %q, want %q", got, "third")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestQueryDB_QueryError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer db.Close()

	query := "SELECT id, jret FROM t WHERE bad=1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnError(assertErr("boom"))

	_, err = queryDB(db, query)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestQueryDB_ScanError_NullString(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer db.Close()

	query := "SELECT id, jret FROM t WHERE id=2"
	// NULL in jret; Scan into string should error.
	rows := sqlmock.NewRows([]string{"id", "jret"}).
		AddRow(int64(2), nil)
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	_, err = queryDB(db, query)
	if err == nil {
		t.Fatalf("expected scan error, got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestQueryDB_RowIterError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New: %v", err)
	}
	defer db.Close()

	query := "SELECT id, jret FROM t"
	// Inject an iteration error on the second returned row.
	rows := sqlmock.NewRows([]string{"id", "jret"}).
		AddRow(int64(1), "a").
		AddRow(int64(2), "b").
		RowError(1, assertErr("iter err")) // 0-based index among returned rows
	mock.ExpectQuery(regexp.QuoteMeta(query)).WillReturnRows(rows)

	_, err = queryDB(db, query)
	if err == nil {
		t.Fatalf("expected iteration error, got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

// GetUserById
func TestGetUserById_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New error: %v", err)
	}
	defer db.Close()

	userID := "123" // NOTE: function builds raw SQL without quotes, so use numeric to keep SQL valid
	qstr := "SELECT userId, json FROM btusers WHERE userId = " + userID + ";"

	// queryDB expects two columns: id (int) and json (string)
	jsonDoc := `{"id":"123","name":"Ada","avatar":"https://example.com/a.png"}`
	rows := sqlmock.NewRows([]string{"userId", "json"}).AddRow(int64(123), jsonDoc)
	mock.ExpectQuery(regexp.QuoteMeta(qstr)).WillReturnRows(rows)

	c := ctxWithQuery("userId", userID)

	got, err := GetUserById(c, db)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got["id"] != "123" || got["name"] != "Ada" || got["avatar"] != "https://example.com/a.png" {
		t.Fatalf("unexpected result: %#v", got)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestGetUserById_DBError(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New error: %v", err)
	}
	defer db.Close()

	userID := "99"
	qstr := "SELECT userId, json FROM btusers WHERE userId = " + userID + ";"

	mock.ExpectQuery(regexp.QuoteMeta(qstr)).WillReturnError(assertErr("boom"))

	c := ctxWithQuery("userId", userID)

	_, err = GetUserById(c, db)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestGetUserById_InvalidJSON(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New error: %v", err)
	}
	defer db.Close()

	userID := "7"
	qstr := "SELECT userId, json FROM btusers WHERE userId = " + userID + ";"

	// Return a row where the "json" column isn't valid JSON
	rows := sqlmock.NewRows([]string{"userId", "json"}).AddRow(int64(7), "{not valid json")
	mock.ExpectQuery(regexp.QuoteMeta(qstr)).WillReturnRows(rows)

	c := ctxWithQuery("userId", userID)

	_, err = GetUserById(c, db)
	if err == nil {
		t.Fatalf("expected json unmarshal error, got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

func TestGetUserById_NoRows(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("sqlmock.New error: %v", err)
	}
	defer db.Close()

	userID := "0"
	qstr := "SELECT userId, json FROM btusers WHERE userId = " + userID + ";"

	// No rows returned. queryDB returns "", nil -> then Unmarshal("") errors.
	rows := sqlmock.NewRows([]string{"userId", "json"})
	mock.ExpectQuery(regexp.QuoteMeta(qstr)).WillReturnRows(rows)

	c := ctxWithQuery("userId", userID)

	_, err = GetUserById(c, db)
	if err == nil {
		t.Fatalf("expected error due to empty result (unmarshal), got nil")
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatalf("unmet expectations: %v", err)
	}
}

// Helpers
func ctxWithQuery(key, val string) *gin.Context {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request = httptest.NewRequest("GET", "/?"+key+"="+val, nil)
	return c
}

type assertErr string

func (e assertErr) Error() string { return string(e) }
