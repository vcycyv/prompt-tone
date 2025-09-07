# Wooden Fish Prompt Tone Generator

A Go command-line tool that generates audio files with wooden fish prompt tones at random intervals of 3-5 minutes.

## Features

- Generates MP3/WAV audio files of specified duration (default: 90 minutes)
- Places wooden fish prompt tones at random intervals between 3-5 minutes
- Each prompt plays 2 wooden fish sounds in sequence for better attention-getting
- Authentic wooden fish "tok" sound with proper resonance
- Optimized for fast generation
- Command-line interface with customizable parameters

## Requirements

- Go 1.21 or later
- Windows, macOS, or Linux
- FFmpeg (optional, for MP3 output - falls back to WAV if not available)

## Installation

1. Clone or download this repository
2. Navigate to the project directory
3. Build the executable:

```bash
go build -o ding.exe main.go
```

### Optional: Install FFmpeg for MP3 Output

To get true MP3 output, install FFmpeg:

**Windows**: Download from [ffmpeg.org](https://ffmpeg.org/download.html) and add to PATH  
**macOS**: `brew install ffmpeg`  
**Linux**: `sudo apt install ffmpeg` or equivalent for your distribution  

## Usage

### Basic Usage

Generate a 90-minute audio file with default settings:

```bash
./ding
```

This will create `ding.mp3` in the current directory.

### Custom Parameters

Generate a custom duration audio file:

```bash
./ding -duration 60
```

Specify a custom output filename:

```bash
./ding -output meditation.mp3
```

Combine both parameters:

```bash
./ding -duration 120 -output session.mp3
```

### Command Line Options

- `-output`: Output filename (default: `ding.mp3`)
- `-duration`: Duration in minutes (default: 90)

## Sound Characteristics

- **Sound Type**: Authentic wooden fish percussion ("tok")
- **Pattern**: Two wooden fish sounds in sequence for each prompt
- **Frequency**: 150 Hz (wooden body resonance)
- **Duration**: 0.4 seconds per sound with 0.2 second gap
- **Volume**: Optimized for clear presence

## Technical Details

- **Sample Rate**: 22,050 Hz (optimized for faster processing)
- **Bit Depth**: 16-bit
- **Channels**: Mono
- **Audio Format**: MP3 (128kbps) with WAV fallback
- **Optimization**: Efficient buffer handling and processing

## Notes

- The tool attempts to generate MP3 files using FFmpeg
- If FFmpeg is not available, it falls back to WAV format
- Each run produces different random tone patterns
- Tones are placed at random intervals between 3-5 minutes
- All tones have consistent wooden fish sound characteristics