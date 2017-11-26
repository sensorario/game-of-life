package main

import "fmt"
import "time"
import "os/exec"
import "os"

func main() {
	size := 25
	steps := 120

	b := Board{
		[]Cell{

			// gilde
			Cell{0, 4},
			Cell{1, 4},
			Cell{2, 4},
			Cell{2, 3},
			Cell{1, 2},

			// warm
			Cell{10, 14},
			Cell{11, 14},
			Cell{12, 14},
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
