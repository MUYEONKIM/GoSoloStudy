package interfacetest

import "fmt"

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("playing", song)
}

type TapeRecoder struct {
	Microphones int
}

func (t TapeRecoder) Play(song string) {
	fmt.Println("recording", song)
}

func playlist(device TapePlayer, songs []string) {
	fmt.Println(device, "디바이스")
	for _, song := range songs {
		device.Play(song)
	}
}

func Interfacetest() {
	player := TapePlayer{}
	mixtape := []string{"first", "second", "third"}
	playlist(player, mixtape)
}

