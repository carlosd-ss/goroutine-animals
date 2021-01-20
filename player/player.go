package player

import (
	"goroutines-music/colors"
	"log"
	"os"
	"sync"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/speaker"
	"github.com/faiface/beep/wav"
	"github.com/fatih/color"
)

var (
	lionSound = sound("drums/lion.wav")
	eagleSound  = sound("drums/eagle.wav")
	dolphinSound = sound("drums/dolphin.wav")
)

type Drum struct {
	Tempo int
}

func (d Drum) Lion(rhythm string, wg *sync.WaitGroup) {
	defer wg.Done()
	d.playbeat("Leão", rhythm, lionSound, colors.Blue)
}

func (d Drum) Eagle(rhythm string, wg *sync.WaitGroup) {
	defer wg.Done()
	d.playbeat("Águia", rhythm, eagleSound, colors.Green)
}

func (d Drum) Dolphin(rhythm string, wg *sync.WaitGroup) {
	defer wg.Done()
	d.playbeat("Golfinho", rhythm, dolphinSound, colors.Yellow)
}

func (d Drum) playbeat(name string, beats string, sound beep.StreamSeeker, color *color.Color) {
	ticker := time.NewTicker(time.Duration(d.Tempo) * time.Millisecond)
	defer ticker.Stop()

	runes := []rune(beats)
	count := 0
	loop := 1
	for count < len(runes) {
		select {
		case _ = <-ticker.C:
			r := runes[count]
			if r == 'x' {
				go play(sound)
				color.Printf("[T%d]%s\n", loop, name)
			}
			count++
			loop++
		}
	}
	//Give a little time to the last note to echo
	time.Sleep(time.Duration(d.Tempo) * time.Millisecond)
}

func play(sound beep.StreamSeeker) {
	playbackDone := make(chan bool)
	speaker.Play(beep.Seq(sound, beep.Callback(func() {
		playbackDone <- true
		sound.Seek(0)
	})))
	<-playbackDone
}

func sound(filename string) beep.StreamSeeker {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	streamer, format, err := wav.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	buffer := beep.NewBuffer(format)
	buffer.Append(streamer)
	streamer.Close()
	sound := buffer.Streamer(0, buffer.Len())

	return sound
}
