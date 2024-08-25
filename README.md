
# VideoCropAnalyzer

VideoCropAnalyzer is a command-line tool that analyzes video files to detect optimal crop values for different frames. It supports multiple operating systems and architectures, including Linux, Windows, and macOS (both Intel and ARM).

## Features

- **Cross-Platform**: Works on Linux, Windows, and macOS.
- **Multiple Architectures**: Supports both Intel (amd64) and ARM (arm64) architectures.
- **Automatic Crop Value Detection**: Automatically detects the optimal crop values (top, bottom, left, right) for video frames.
- **Lightweight and Fast**: Quickly processes videos to provide crop recommendations.

## Prerequisites

### FFmpeg

VideoCropAnalyzer requires FFmpeg to be installed on your system, as it uses FFmpeg for video processing.

- **Install FFmpeg on macOS**:
  
  ```bash
  brew install ffmpeg
  ```
  
- **Install FFmpeg on Ubuntu**:
  
  ```bash
  sudo apt-get install ffmpeg
  ```

- **Install FFmpeg on Windows**:  
  Download and install FFmpeg from [FFmpeg's official website](https://ffmpeg.org/download.html), and ensure that the `ffmpeg` executable is available in your system's `PATH`.

## Installation

### Pre-built Binaries

You can download pre-built binaries for your operating system from the [Releases](https://github.com/Kianrad/VideoCropAnalyzer/releases) page.

1. Download the appropriate binary for your operating system and architecture.
2. Extract the binary to a directory in your `PATH`.
3. Give the binary executable permissions (if needed):

   ```bash
   chmod +x videocropanalyzer
   ```

### Homebrew (macOS)

If you're using macOS, you can install the tool using Homebrew:

```bash
brew tap Kianrad/videocropanalyzer
brew install videocropanalyzer
```

## Usage

### Basic Command

To analyze a video file and detect crop values:

```bash
videocropanalyzer /path/to/your/video.mp4
```

This command will output the suggested crop values in JSON format.

### Example Output

```json
{
  "top": 40,
  "bottom": 40,
  "left": 10,
  "right": 10
}
```

This output indicates that the tool suggests cropping 40 pixels from the top and bottom, and 10 pixels from the left and right of the video.

### Sample Output

Here’s a sample of what you might see when running the tool:

```bash
$ videocropanalyzer /path/to/your/video.mp4
{
  "top": 40,
  "bottom": 40,
  "left": 10,
  "right": 10
}
```

This output shows the analysis process and the resulting crop values in JSON format.

### Command-line Options

- `--help`: Display help information.
- `--version`: Show the version of the tool.

## Building from Source

To build the tool from source, you'll need Go installed on your machine.

1. Clone the repository:

   ```bash
   git clone https://github.com/Kianrad/VideoCropAnalyzer.git
   cd VideoCropAnalyzer
   ```

2. Build the project:

   ```bash
   go build -o videocropanalyzer
   ```

3. Run the tool:

   ```bash
   ./videocropanalyzer /path/to/your/video.mp4
   ```

### Cross-Compiling

To compile the project for multiple operating systems and architectures, use the following commands:

```bash
GOOS=linux GOARCH=amd64 go build -o videocropanalyzer-linux-amd64
GOOS=windows GOARCH=amd64 go build -o videocropanalyzer-windows-amd64.exe
GOOS=darwin GOARCH=amd64 go build -o videocropanalyzer-darwin-amd64
GOOS=darwin GOARCH=arm64 go build -o videocropanalyzer-darwin-arm64
```

## Contributing

Contributions are welcome! Please follow these steps:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature/your-feature`).
3. Commit your changes (`git commit -am 'Add some feature'`).
4. Push to the branch (`git push origin feature/your-feature`).
5. Open a pull request.

Please make sure your code passes the existing tests and includes new tests where appropriate.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [FFmpeg](https://ffmpeg.org/) - The tool uses FFmpeg for video processing.
- [Go](https://golang.org/) - Built with Go.
