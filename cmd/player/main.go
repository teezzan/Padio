package main

import (
	"fmt"
	"os"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Queue struct {
	streamers []beep.Streamer
	playing   chan bool
}

func (q *Queue) Init() {
	q.playing = make(chan bool, 200)
}

func (q *Queue) Add(streamers ...beep.Streamer) {
	q.streamers = append(q.streamers, streamers...)
}

func (q *Queue) Stream(samples [][2]float64) (n int, ok bool) {
	// We use the filled variable to track how many samples we've
	// successfully filled already. We loop until all samples are filled.
	filled := 0
	for filled < len(samples) {
		// There are no streamers in the queue, so we stream silence.
		if len(q.streamers) == 0 {
			q.playing <- false
			for i := range samples[filled:] {
				samples[i][0] = 0
				samples[i][1] = 0
			}
			break
		}

		// We stream from the first streamer in the queue.
		n, ok := q.streamers[0].Stream(samples[filled:])
		// If it's drained, we pop it from the queue, thus continuing with
		// the next streamer.
		if !ok {
			q.streamers = q.streamers[1:]
			// q.playing <- false
		} else {
			q.playing <- true
		}
		// We update the number of filled samples.
		filled += n
	}
	return len(samples), true
}

func (q *Queue) Err() error {
	return nil
}

func main() {
	// done := make(chan bool, 5)
	var queue Queue
	queue.Init()

	sr := beep.SampleRate(44100)
	speaker.Init(sr, sr.N(time.Second/10))

	// A zero Queue is an empty Queue.

	speaker.Play(beep.Seq(&queue))
	GetNextAudio(&queue, sr, 3)

	for {

		select {
		case i := <-queue.playing:
			fmt.Println("Done ", i)
			if !i {
				GetNextAudio(&queue, sr, 2)
			}
		default:
			fmt.Print("")
		}
	}
}

func GetNextAudio(queue *Queue, sr beep.SampleRate, num int) {

	name := fmt.Sprintf("static/%d.mp3", num)

	// Open the file on the disk.
	f, err := os.Open(name)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Decode it.
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	// The speaker's sample rate is fixed at 44100. Therefore, we need to
	// resample the file in case it's in a different sample rate.
	resampled := beep.Resample(4, format.SampleRate, sr, streamer)

	// And finally, we add the song to the queue.
	speaker.Lock()
	queue.Add(resampled)
	speaker.Unlock()
	fmt.Print("Added Successfully!")
}
