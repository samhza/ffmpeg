package ffmpeg

import (
	"encoding/json"
	"io"
	"os/exec"
)

type CodecType string

const (
	CodecTypeAudio CodecType = "audio"
	CodecTypeVideo CodecType = "video"
)

type ProbeResults struct {
	Streams []ProbedStream `json:"streams"`
	Format  Format         `json:"format"`
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
type ProbedStream struct {
	Tags             StreamTags  `json:"tags"`
	NALLengthSize    string      `json:"nal_length_size"`
	CodecLongName    string      `json:"codec_long_name"`
	Profile          string      `json:"profile"`
	CodecType        CodecType   `json:"codec_type"`
	CodecTagString   string      `json:"codec_tag_string"`
	CodecTag         string      `json:"codec_tag"`
	NbFrames         string      `json:"nb_frames"`
	BitsPerRawSample string      `json:"bits_per_raw_sample"`
	BitRate          string      `json:"bit_rate"`
	Duration         string      `json:"duration"`
	StartTime        string      `json:"start_time"`
	TimeBase         string      `json:"time_base"`
	PixFmt           string      `json:"pix_fmt"`
	AvgFrameRate     string      `json:"avg_frame_rate"`
	CodecName        string      `json:"codec_name"`
	RFrameRate       string      `json:"r_frame_rate"`
	IsAVC            string      `json:"is_avc"`
	ChromaLocation   string      `json:"chroma_location"`
	Disposition      Disposition `json:"disposition"`
	Refs             int         `json:"refs"`
	Level            int         `json:"level"`
	HasBFrames       int         `json:"has_b_frames"`
	StartPTS         int         `json:"start_pts"`
	DurationTS       int         `json:"duration_ts"`
	CodedHeight      int         `json:"coded_height"`
	CodedWidth       int         `json:"coded_width"`
	Height           int         `json:"height"`
	Width            int         `json:"width"`
	ClosedCaptions   int         `json:"closed_captions"`
	Index            int         `json:"index"`
}
type StreamTags struct {
	Language    string `json:"language"`
	HandlerName string `json:"handler_name"`
	VendorID    string `json:"vendor_id"`
}

type Format struct {
	Tags           FormatTags `json:"tags"`
	Filename       string     `json:"filename"`
	FormatName     string     `json:"format_name"`
	FormatLongName string     `json:"format_long_name"`
	StartTime      string     `json:"start_time"`
	Duration       string     `json:"duration"`
	NStreams       int        `json:"nb_streams"`
	NPrograms      int        `json:"nb_programs"`
	ProbeScore     int        `json:"probe_score"`
}

type FormatTags struct {
	MajorBrand       string `json:"major_brand"`
	MinorVersion     string `json:"minor_version"`
	CompatibleBrands string `json:"compatible_brands"`
	Encoder          string `json:"encoder"`
}

func Probe(f string) (*ProbeResults, error) {
	return probe(f, nil)
}

func ProbeReader(r io.Reader) (*ProbeResults, error) {
	return probe("-", r)
}

func probe(f string, r io.Reader) (*ProbeResults, error) {
	cmd := exec.Command("ffprobe", "-v", "quiet", "-print_format", "json", "-show_format", "-show_streams", f)
	cmd.Stdin = r
	out, err := cmd.Output()
	if err != nil {
		return nil, err
	}
	var probeResults ProbeResults
	return &probeResults, json.Unmarshal(out, &probeResults)
}
