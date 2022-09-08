package main

import (
	"fmt"
	"testing"
)

func TestGenerateMeetings(t *testing.T) {
	counts := []int{4, 7, 78, 12}

	for k, count := range counts {

		t.Run(fmt.Sprintf("Testing generating %d meetings that are ordered by start datetime in ascending order", k), func(t *testing.T) {

			meetings := generateMeetings(count)

			for x, m := range meetings {
				if x == 0 {
					continue
				}
				if m.StartTime < meetings[x-1].StartTime {
					t.Errorf("Expected startTime %d to be below or equal to %d but it's not the case", m.StartTime, meetings[x-1].StartTime)
				}
			}
		})
	}
}

// Input Data Set: [1, 2], [1, 3], [2, 5], [3, 6], [4, 7]
// Output:
// Room 1: [1, 2], [2, 5]
// Room 2: [1, 3], [3, 6]
// Room 3: [4, 7]

func TestAssignMeetingsToRooms(t *testing.T) {

	type MeetingWithAssignment []struct {
		TestName          string
		Meetings          []Meeting
		RoomsWithMeetings [][]Meeting
	}

	data := MeetingWithAssignment{
		{
			TestName: "default meetings data set provided in interview",
			Meetings: []Meeting{
				{StartTime: 1, EndTime: 2},
				{StartTime: 1, EndTime: 3},
				{StartTime: 2, EndTime: 5},
				{StartTime: 3, EndTime: 6},
				{StartTime: 4, EndTime: 7},
			},
			RoomsWithMeetings: [][]Meeting{
				{{StartTime: 1, EndTime: 2},
					{StartTime: 2, EndTime: 5}},
				{{StartTime: 1, EndTime: 3},
					{StartTime: 3, EndTime: 6}},
				{{StartTime: 4, EndTime: 7}},
			},
		},
		{
			TestName: "random meetings data set I",
			Meetings: []Meeting{
				{StartTime: 0, EndTime: 2},
				{StartTime: 1, EndTime: 3},
				{StartTime: 12, EndTime: 14},
			},
			RoomsWithMeetings: [][]Meeting{
				{{StartTime: 0, EndTime: 2},
					{StartTime: 12, EndTime: 14}},
				{{StartTime: 1, EndTime: 3}},
			},
		},
		{
			TestName: "random meetings data set II",
			Meetings: []Meeting{
				{StartTime: 0, EndTime: 3},
				{StartTime: 1, EndTime: 3},
				{StartTime: 1, EndTime: 4},
				{StartTime: 7, EndTime: 9},
			},
			RoomsWithMeetings: [][]Meeting{
				{{StartTime: 0, EndTime: 3},
					{StartTime: 7, EndTime: 9}},
				{{StartTime: 1, EndTime: 3}},
				{{StartTime: 1, EndTime: 4}},
			},
		},
	}

	for _, dataSet := range data {

		t.Run(dataSet.TestName, func(t *testing.T) {

			roomsWithMeetings := assignMeetingsToRooms(dataSet.Meetings)

			for x, v := range roomsWithMeetings {

				receivedStartime := dataSet.RoomsWithMeetings[x][0].StartTime
				expectedStartime := v[0].StartTime
				if receivedStartime != expectedStartime {
					t.Errorf("Expected starttime %d but got %d", expectedStartime, receivedStartime)
				}

				receivedEndtime := dataSet.RoomsWithMeetings[x][0].EndTime
				expectedEndtime := v[0].EndTime
				if receivedStartime != expectedStartime {
					t.Errorf("Expected endtime %d but got %d", expectedEndtime, receivedEndtime)
				}

			}

		})
	}

}
