package currency_test

import (
	"testing"

	"bitbucket.org/syb-devs/currency"
)

var currencyEqualsTests = []struct {
	currencyA currency.Currency
	currencyB currency.Currency
	expected  bool
}{
	{currency.Currency{"USD", 840, 2, "United States dollar"}, currency.Currency{"USD", 840, 2, "United States dollar"}, true},
	{currency.Currency{"USD", 840, 2, "United States dollar"}, currency.Currency{"EUR", 840, 2, "United States dollar"}, false},
	{currency.Currency{"USD", 840, 2, "United States dollar"}, currency.Currency{"USD", 666, 2, "United States dollar"}, false},
	{currency.Currency{"USD", 840, 2, "United States dollar"}, currency.Currency{"USD", 840, 1, "United States dollar"}, false},
	{currency.Currency{"USD", 840, 2, "United States dollar"}, currency.Currency{"USD", 840, 2, "Unstated Unites dollar"}, false},
}

func TestCurrencyEquals(t *testing.T) {
	for _, test := range currencyEqualsTests {
		actual := test.currencyA.Equals(test.currencyB)
		if actual != test.expected {
			t.Errorf("Currency.Equals() expecting %v, got %v while comparing %+v with %+v", test.expected, actual, test.currencyA, test.currencyB)
		}
	}
}

var getByCodeTests = []struct {
	code     string
	expected currency.Currency
}{
	{"USD", currency.Currency{"USD", 840, 2, "United States dollar"}},
	{"KGS", currency.Currency{"KGS", 417, 2, "Kyrgyzstani som"}},
}

func TestGetByCode(t *testing.T) {
	for _, test := range getByCodeTests {
		actual, _ := currency.GetByCode(test.code)
		if actual != test.expected {
			t.Errorf("GetByCode(%s) expecting %+v, got %+v", test.code, test.expected, actual)
		}
	}
}

func TestGetByCodeNotFound(t *testing.T) {
	expected := currency.ErrCurrencyNotFound
	_, err := currency.GetByCode("XWX")
	if err != expected {
		t.Errorf("expecting '%v', got '%v'", expected, err)
	}
}

var getByIDTests = []struct {
	ID       int
	expected currency.Currency
}{
	{840, currency.Currency{"USD", 840, 2, "United States dollar"}},
	{417, currency.Currency{"KGS", 417, 2, "Kyrgyzstani som"}},
}

func TestGetByID(t *testing.T) {
	for _, test := range getByIDTests {
		actual, _ := currency.GetByID(test.ID)
		if actual != test.expected {
			t.Errorf("GetByID(%v) expecting %+v, got %+v", test.ID, test.expected, actual)
		}
	}
}

func TestGetByIDNotFound(t *testing.T) {
	expected := currency.ErrCurrencyNotFound
	_, err := currency.GetByID(8765)
	if err != expected {
		t.Errorf("expecting '%v', got '%v'", expected, err)
	}
}

func TestGetList(t *testing.T) {
	list := currency.GetList("USD", "XWX", "KGS", "GBP")
	if len(list) != 3 {
		t.Errorf("expected a list with 3 elements, but got %v. List: %+v", len(list), list)
	}

	c := list[1]
	if c.Code != "KGS" {
		t.Errorf("expected the second currency in list to be KGS, but got: %+v", c.Code)
	}

	list = currency.GetList()
	if len(list) != 181 {
		t.Errorf("expecting the full currency list to have 181 elements, but got %v", len(list))
	}
}
