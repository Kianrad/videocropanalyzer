package models

type FFProbeOutput struct {
	Streams []Stream `json:"streams"`
	Format  Format   `json:"format"`
	// Other fields might include "programs", "chapters", etc.
}

type Stream struct {
	Index              int               `json:"index"`
	CodecName          string            `json:"codec_name"`
	CodecLongName      string            `json:"codec_long_name"`
	Profile            string            `json:"profile"`
	CodecType          string            `json:"codec_type"`
	CodecTimeBase      string            `json:"codec_time_base"`
	CodecTagString     string            `json:"codec_tag_string"`
	CodecTag           string            `json:"codec_tag"`
	Width              int               `json:"width,omitempty"`
	Height             int               `json:"height,omitempty"`
	CodedWidth         int               `json:"coded_width,omitempty"`
	CodedHeight        int               `json:"coded_height,omitempty"`
	ClosedCaptions     int               `json:"closed_captions,omitempty"`
	HasBFrames         int               `json:"has_b_frames,omitempty"`
	SampleAspectRatio  string            `json:"sample_aspect_ratio,omitempty"`
	DisplayAspectRatio string            `json:"display_aspect_ratio,omitempty"`
	PixFmt             string            `json:"pix_fmt,omitempty"`
	Level              int               `json:"level,omitempty"`
	ChromaLocation     string            `json:"chroma_location,omitempty"`
	Refs               int               `json:"refs,omitempty"`
	IsAvc              string            `json:"is_avc,omitempty"`
	NalLengthSize      string            `json:"nal_length_size,omitempty"`
	RFrameRate         string            `json:"r_frame_rate"`
	AvgFrameRate       string            `json:"avg_frame_rate"`
	TimeBase           string            `json:"time_base"`
	StartPts           int64             `json:"start_pts,omitempty"`
	StartTime          string            `json:"start_time,omitempty"`
	DurationTs         int64             `json:"duration_ts,omitempty"`
	Duration           string            `json:"duration,omitempty"`
	BitRate            string            `json:"bit_rate,omitempty"`
	MaxBitRate         string            `json:"max_bit_rate,omitempty"`
	BitsPerRawSample   string            `json:"bits_per_raw_sample,omitempty"`
	NbFrames           string            `json:"nb_frames,omitempty"`
	NbReadFrames       string            `json:"nb_read_frames,omitempty"`
	NbReadPackets      string            `json:"nb_read_packets,omitempty"`
	Disposition        Disposition       `json:"disposition"`
	Tags               map[string]string `json:"tags,omitempty"`
	// Add other fields as necessary depending on your specific ffprobe output
}

type Disposition struct {
	Default         int `json:"default"`
	Dub             int `json:"dub"`
	Original        int `json:"original"`
	Comment         int `json:"comment"`
	Lyrics          int `json:"lyrics"`
	Karaoke         int `json:"karaoke"`
	Forced          int `json:"forced"`
	HearingImpaired int `json:"hearing_impaired"`
	VisualImpaired  int `json:"visual_impaired"`
	CleanEffects    int `json:"clean_effects"`
	AttachedPic     int `json:"attached_pic"`
	TimedThumbnails int `json:"timed_thumbnails"`
}

type Format struct {
	Filename       string            `json:"filename"`
	NbStreams      int               `json:"nb_streams"`
	NbPrograms     int               `json:"nb_programs,omitempty"`
	FormatName     string            `json:"format_name"`
	FormatLongName string            `json:"format_long_name"`
	StartTime      string            `json:"start_time,omitempty"`
	Duration       string            `json:"duration,omitempty"`
	Size           string            `json:"size,omitempty"`
	BitRate        string            `json:"bit_rate,omitempty"`
	ProbeScore     int               `json:"probe_score,omitempty"`
	Tags           map[string]string `json:"tags,omitempty"`
}
