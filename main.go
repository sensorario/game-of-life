package main

import "fmt"
import "time"
import "os/exec"
import "os"

func main() {
	size := 18
	steps := 60

	b := Board{
		[]Cell{
			Cell{0, 4},
			Cell{1, 4},
			Cell{2, 4},
			Cell{2, 3},
			Cell{1, 2},
		},
	}

	for i := 1; i < steps; i++ {
		b = b.NewGen()

		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()

		for i := 0; i <= size; i++ {
			for j := 0; j <= size*2; j++ {
				if b.Contains(Cell{i, j}) {
					fmt.Print("●")
				} else {
					fmt.Print("○")
				}
			}
			fmt.Println("")
		}

		time.Sleep(time.Second / 10)
	}
}
