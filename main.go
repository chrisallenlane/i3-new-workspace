package main

import (
	"encoding/json"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
)

// Workspace represents an i3 workspace reported by `i3-msg`
type Workspace struct {
	Num int `json:"num"`
}

func main() {

	// get the workspaces
	jsonStr, err := exec.Command("i3-msg", "-t", "get_workspaces").Output()
	if err != nil {
		log.Fatalf("failed to get i3 workspaces: %v", err)
	}

	// initialize a struct
	var workspaces []Workspace

	// unmarshal the string into the struct
	if err := json.Unmarshal(jsonStr, &workspaces); err != nil {
		log.Fatalf("failed to unmarshal json: %v", err)
	}

	// sort workspaces
	sort.Slice(workspaces, func(i, j int) bool {
		return workspaces[i].Num < workspaces[j].Num
	})

	// determine the workspace start point
	num := 1
	if len(os.Args) > 1 {
		i, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatalf("failed to convert to integer: %v", err)
		}
		num = i
	}

	// find the lowest available workspace number
	for _, wksp := range workspaces {
		if num < wksp.Num {
			break
		}

		num = wksp.Num + 1
	}

	// initialize a new workspace
	cmd := exec.Command("i3-msg", "workspace", strconv.Itoa(num))
	if err := cmd.Run(); err != nil {
		log.Fatalf("failed to create workspace %d: %v", num, err)
	}
}
