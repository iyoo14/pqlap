package pqlap

import (
    "testing"
)

func TestGetInstance(t *testing.T) {
    _, err := GetInstance()
    if err != nil {
        t.Error("got error")
    }
}

func TestSum(t *testing.T) {
    actual := Sum(1,2)
    expected := 3
    if actual != expected {
        t.Errorf("got %v\nwant %v", actual, expected)
    }
}
