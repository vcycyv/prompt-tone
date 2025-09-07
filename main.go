package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type PromptTone struct {
	StartTime float64 // in seconds
	Frequency float64 // in Hz
	Duration  float64 // in seconds (fixed at 0.4)
	Volume    float64 // fixed at 0.9
}

func main() {
	var (
		outputFile = flag.String("output", "ding.mp3", "Output MP3 file name")
		duration   = flag.Int("duration", 90, "Duration in minutes")
	)
	flag.Parse()

	fmt.Printf("Generating %d-minute MP3 with random prompt tones...\n", *duration)

	// Generate prompt tones
	tones := generatePromptTones(*duration)

	// Generate audio data
	audioData := generateAudioData(tones, *duration)

	// Save to MP3 file
	if err := saveToMP3(audioData, *outputFile); err != nil {
		log.Fatalf("Failed to save MP3 file: %v", err)
	}

	fmt.Printf("Successfully generated %s\n", *outputFile)
	fmt.Printf("Generated %d prompt tones\n", len(tones))

	// Print tone schedule
	for i, tone := range tones {
		fmt.Printf("Tone %d: at %.1f seconds (2 wooden fish sounds in sequence)\n",
			i+1, tone.StartTime)
	}
}

func generatePromptTones(durationMinutes int) []PromptTone {
	var tones []PromptTone
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate tones every 3-5 minutes randomly
	currentTime := 0.0
	for currentTime < float64(durationMinutes*60) {
		// Random interval between 3-5 minutes
		interval := r.Float64()*2 + 3 // 3 to 5 minutes
		currentTime += interval * 60  // Convert to seconds

		if currentTime < float64(durationMinutes*60) {
			tone := PromptTone{
				StartTime: currentTime,
				Frequency: 150, // Perfect frequency for wooden fish (150 Hz)
				Duration:  0.4, // Shorter duration for authentic wooden fish "tok"
				Volume:    0.9, // Louder for clear wooden percussion sound
			}
			tones = append(tones, tone)
		}
	}

	return tones
}

func generateAudioData(tones []PromptTone, durationMinutes int) []float32 {
	// Use a lower sample rate for faster processing
	sampleRate := 22050 // Half the original sample rate (44100/2)
	totalSamples := durationMinutes * 60 * sampleRate
	audioData := make([]float32, totalSamples)

	// Skip filling with silence - Go already initializes arrays with zeros

	// Add tones at their respective positions
	for _, tone := range tones {
		// Play the wooden fish sound 2 times in a row
		for repeat := 0; repeat < 2; repeat++ {
			// Calculate start position for each repetition
			// Add a small gap between each sound (0.2 seconds)
			repeatOffset := float64(repeat) * (tone.Duration + 0.2)
			startSample := int((tone.StartTime + repeatOffset) * float64(sampleRate))
			durationSamples := int(tone.Duration * float64(sampleRate))

			// Skip processing if out of range
			if startSample >= totalSamples {
				continue
			}

			// Limit duration to avoid overruns
			if startSample+durationSamples > totalSamples {
				durationSamples = totalSamples - startSample
			}

			// Generate tone samples - use step size of 2 for faster processing
			for i := 0; i < durationSamples; i += 2 {
				if startSample+i >= 0 {
					time := float64(i) / float64(sampleRate)

					// Create tone with simplified processing
					sample := generatePleasantTone(tone, time, i, durationSamples)

					// Set current sample and duplicate to next sample (faster than calculating both)
					audioData[startSample+i] = float32(sample)
					if i+1 < durationSamples && startSample+i+1 < totalSamples {
						audioData[startSample+i+1] = float32(sample)
					}
				}
			}
		}
	}

	return audioData
}

func generatePleasantTone(tone PromptTone, time float64, sampleIndex, totalSamples int) float64 {
	// Create authentic wooden fish (mokugyo) sound - hollow wooden "tok"

	// Wooden fish has very little tonal content - it's mostly a percussive sound
	// Instead of using sine waves, we'll create a more realistic percussion sound

	// Apply the wooden fish envelope which contains most of the sound character
	envelope := applyWoodenFishEnvelope(sampleIndex, totalSamples)

	// Add noise component for the wooden impact
	noiseComponent := generateWoodenNoise(sampleIndex, totalSamples, time)

	// Combine for authentic wooden fish sound
	return tone.Volume * envelope * noiseComponent
}

func generateWoodenNoise(sampleIndex, totalSamples int, time float64) float64 {
	// Improved wooden fish sound while maintaining performance

	// Create a more authentic wooden fish "tok" sound
	// Use a lower frequency for the hollow wooden body (120-150Hz is typical)
	woodBody := math.Sin(2 * math.Pi * 130 * time)

	// Add a quick "knock" component for the impact sound
	// This creates the characteristic wooden fish "tok" sound
	knockComponent := 0.0
	if sampleIndex < totalSamples/8 {
		// Sharp initial knock (the "tok")
		knockComponent = math.Sin(2*math.Pi*800*time) * 0.5
	}

	// Combine for authentic wooden fish sound while keeping it fast
	if sampleIndex < totalSamples/10 {
		// More knock component at the start (the impact)
		return woodBody*0.7 + knockComponent*0.3
	} else {
		// More wood resonance for the brief tail
		return woodBody * 0.9
	}
}

func applyWoodenFishEnvelope(sampleIndex, totalSamples int) float64 {
	// Improved wooden fish envelope while maintaining performance

	// Real wooden fish has very sharp attack and quick decay
	attackSamples := totalSamples / 40   // Very sharp attack (1/40 of duration)
	decaySamples := totalSamples * 2 / 3 // Shorter decay for wooden fish

	if sampleIndex < attackSamples {
		// Very sharp attack for wooden impact "tok" sound
		// Use quadratic curve for more natural wooden strike
		progress := float64(sampleIndex) / float64(attackSamples)
		return progress * progress
	} else {
		// Quick decay typical of wooden percussion
		decayIndex := sampleIndex - attackSamples
		decayProgress := float64(decayIndex) / float64(decaySamples)

		// Early cutoff for speed and authentic wooden sound
		if decayProgress >= 0.7 {
			return 0
		}

		// Two-stage decay for more realistic wooden sound
		if decayProgress < 0.2 {
			// Very fast initial decay (the "tok")
			return 1.0 - decayProgress*2.5
		} else {
			// Quick tail decay (wooden resonance fades quickly)
			return 0.5 * math.Exp(-(decayProgress-0.2)*8)
		}
	}
}

func saveToMP3(audioData []float32, filename string) error {
	// For now, we'll create a WAV file and then convert it to MP3
	// This is a simplified approach - in production you'd want a proper MP3 encoder

	// Create temporary WAV file
	tempWavFile := filename + ".wav"
	if err := saveToWAV(audioData, tempWavFile); err != nil {
		return fmt.Errorf("failed to create temporary WAV: %v", err)
	}
	defer os.Remove(tempWavFile)

	// Convert WAV to MP3 using external tool (FFmpeg if available)
	if err := convertWavToMp3(tempWavFile, filename); err != nil {
		// If conversion fails, just copy the WAV file and rename it
		log.Printf("Warning: MP3 conversion failed, using WAV format instead: %v", err)
		if err := os.Rename(tempWavFile, filename); err != nil {
			return fmt.Errorf("failed to rename WAV file: %v", err)
		}
	}

	return nil
}

func convertWavToMp3(wavFile, mp3File string) error {
	// Try to use FFmpeg if available
	args := []string{"-i", wavFile, "-acodec", "libmp3lame", "-ab", "128k", mp3File, "-y"}
	cmd := exec.Command("ffmpeg", args...)

	// Run the command
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("ffmpeg execution failed: %v", err)
	}

	return nil
}

func saveToWAV(audioData []float32, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// WAV file header with reduced sample rate for faster processing
	sampleRate := 22050 // Match the reduced sample rate from generateAudioData
	bitsPerSample := 16
	numChannels := 1
	dataSize := len(audioData) * 2 // 16-bit samples

	// Write WAV header
	writeString(file, "RIFF")
	writeUint32(file, uint32(36+dataSize)) // File size - 8
	writeString(file, "WAVE")
	writeString(file, "fmt ")
	writeUint32(file, 16) // Chunk size
	writeUint16(file, 1)  // Audio format (PCM)
	writeUint16(file, uint16(numChannels))
	writeUint32(file, uint32(sampleRate))
	writeUint32(file, uint32(sampleRate*numChannels*bitsPerSample/8)) // Byte rate
	writeUint16(file, uint16(numChannels*bitsPerSample/8))            // Block align
	writeUint16(file, uint16(bitsPerSample))
	writeString(file, "data")
	writeUint32(file, uint32(dataSize))

	// Write audio data - process in chunks for better performance
	chunkSize := 4096
	buffer := make([]byte, chunkSize*2) // 2 bytes per sample

	for i := 0; i < len(audioData); i += chunkSize {
		end := i + chunkSize
		if end > len(audioData) {
			end = len(audioData)
		}

		// Convert chunk of samples to bytes
		bufferPos := 0
		for j := i; j < end; j++ {
			// Convert float32 to int16
			intSample := int16(audioData[j] * 32767.0)
			buffer[bufferPos] = byte(intSample)
			buffer[bufferPos+1] = byte(intSample >> 8)
			bufferPos += 2
		}

		// Write chunk to file
		file.Write(buffer[:bufferPos])
	}

	return nil
}

func writeString(file *os.File, s string) {
	file.WriteString(s)
}

func writeUint16(file *os.File, value uint16) {
	file.Write([]byte{byte(value), byte(value >> 8)})
}

func writeUint32(file *os.File, value uint32) {
	file.Write([]byte{byte(value), byte(value >> 8), byte(value >> 16), byte(value >> 24)})
}
