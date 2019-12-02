package main

import "testing"

func TestFraudActivityNotifications1(t *testing.T) {
	d := int32(5)
	expenditure := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}

	daysNotified := activityNotifications(expenditure, d)

	if daysNotified != 2 {
		t.Errorf("got %d instead of 2", daysNotified)
	}
}

func TestFraudActivityNotifications2(t *testing.T) {
	d := int32(6)
	expenditure := []int32{2, 3, 4, 2, 3, 6, 8, 4, 5}

	daysNotified := activityNotifications(expenditure, d)

	if daysNotified != 1 {
		t.Errorf("got %d instead of 1", daysNotified)
	}
}

func TestFraudActivityNotifications3(t *testing.T) {
	d := int32(4)
	expenditure := []int32{1, 2, 3, 4, 4}

	daysNotified := activityNotifications(expenditure, d)

	if daysNotified != 0 {
		t.Errorf("got %d instead of 0", daysNotified)
	}
}

// 2 2 3 3 4 6
