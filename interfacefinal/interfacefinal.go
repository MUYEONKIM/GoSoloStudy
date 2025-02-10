package interfacefinal

import (
	"fmt"
)

type TapeInerface interface {
	Play(string)
	Stop()
}

type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("playing : ", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stop!!")
}

type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Recording : ", song)
}

func (t TapeRecorder) Recording() {
	fmt.Println("삐빠라삐라빠")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stop!!")
}

func playlist(device TapeInerface, songs []string) {
	for _, arg := range songs {
		fmt.Println(arg)
	}

	recorder, ok := device.(TapeRecorder)
	if ok {
		recorder.Recording()
	}
	device.Stop()
}

func Interfacefinal() {
	player := TapePlayer{}
	record := TapeRecorder{}

	mixtape := []string{"first", "second", "third"}

	playlist(player, mixtape)
	playlist(record, mixtape)

}


