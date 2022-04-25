package process

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/teezzan/padio/internal/player"
)

var LoadingInProgress bool = false
var StaticDir = "static"
var Queue player.Queue

func Init() {
	Queue.Init()

	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))

	speaker.Play(beep.Seq(&Queue))

	for {
		select {
		case i := <-Queue.Playing:

			load := !i && !LoadingInProgress
			if load {
				LoadingInProgress = true
				QueueAndPlay(&Queue, sr)
				LoadingInProgress = false
			}
		default:
			fmt.Print("")
		}
	}
}

func QueueAndPlay(queue *player.Queue, sr beep.SampleRate) error {
	f, err := os.Open(StaticDir)

	if err != nil {
		fmt.Println(err)
		return err
	}

	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return err
	}
	rand.Seed(time.Now().UnixNano())
	randomFile := files[rand.Intn(len(files))]
	fmt.Println(randomFile.Name())
	return PlayNextAudio(queue, sr, fmt.Sprintf("%s/%s", StaticDir, randomFile.Name()))
}

func PlayNextAudio(queue *player.Queue, sr beep.SampleRate, name string) error {

	// Open the file on the disk.
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// Decode it.
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println(err)
		return err
	}

	// The speaker's sample rate is fixed at 44100. Therefore, we need to
	// resample the file in case it's in a different sample rate.
	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	// And finally, we add the song to the queue.
	speaker.Lock()
	queue.Add(resampled)
	speaker.Unlock()
	fmt.Print("Added Successfully!")
	return nil
}
