package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Meeting struct {
	StartTime int
	EndTime   int
}

var meetingsCount int = 10

func main() {

	rand.Seed(time.Now().UnixNano())

	meets := generateMeetings(meetingsCount)

	roomsWithMeetings := assignMeetingsToRooms(meets)

	for i, m := range meets {
		fmt.Printf("Meeting %d: \n\tstart: %d\tend: %d\n", i+1, m.StartTime, m.EndTime)
	}

	fmt.Println()
	fmt.Println()
	fmt.Println()

	for i, room := range roomsWithMeetings {
		fmt.Printf("Room %d: %v\n", i+1, room)
	}

}

func generateMeetings(count int) []Meeting {

	meetings := make([]Meeting, count)

	for k := range meetings {
		startTime := rand.Intn(24)
		meetings[k].StartTime = startTime
		meetings[k].EndTime = getEndTimeFromStartTime(startTime)
	}

	sort.Slice(meetings, func(i, j int) bool {

		if meetings[i].StartTime < meetings[j].StartTime {
			return true
		} else if meetings[i].StartTime == meetings[j].StartTime {
			return meetings[i].EndTime < meetings[j].EndTime
		}

		return false
	})

	return meetings
}

func getEndTimeFromStartTime(start int) int {
	return start + rand.Intn(3) + 1
}

// Input Data Set: [1, 2], [1, 3], [2, 5], [3, 6], [4, 7]
// Output:
// Room 1: [1, 2], [2, 5]
// Room 2: [1, 3], [3, 6]
// Room 3: [4, 7]
func assignMeetingsToRooms(meets []Meeting) [][]Meeting {
	var rooms [][]Meeting

meetingsRoom:
	for _, m := range meets {

		if len(rooms) == 0 {
			room := []Meeting{
				m,
			}
			rooms = append(rooms, room)
			continue
		}

		if len(rooms) > 0 {

		roomsLoop:
			for k, r := range rooms {
				lastMeetinginRoom := r[len(r)-1]

				if m.StartTime >= lastMeetinginRoom.EndTime {
					rooms[k] = append(rooms[k], m)
					continue meetingsRoom
				} else {
					continue roomsLoop
				}
			}
		}

		room := []Meeting{
			m,
		}
		rooms = append(rooms, room)
		continue

	}

	return rooms
}
