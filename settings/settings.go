package settings

const SETTINGSVERSION = "v1"

type general struct {
	OsuDir string //localappdata
}

type graphics struct {
	Width, Height int64
	WindowWidth, WindowHeight int64
	Fullscreen bool //true
	VSync bool //false
	FPSCap int64 //1000
	MSAA int32 //16
}

func (gr graphics) GetSize() (int64, int64) {
	if gr.Fullscreen {
		return gr.Width, gr.Height
	}
	return gr.WindowWidth, gr.WindowHeight
}

func (gr graphics) GetSizeF() (float64, float64) {
	if gr.Fullscreen {
		return float64(gr.Width), float64(gr.Height)
	}
	return float64(gr.WindowWidth), float64(gr.WindowHeight)
}

func (gr graphics) GetWidth() int64 {
	if gr.Fullscreen {
		return gr.Width
	}
	return gr.WindowWidth
}

func (gr graphics) GetWidthF() float64 {
	if gr.Fullscreen {
		return float64(gr.Width)
	}
	return float64(gr.WindowWidth)
}

func (gr graphics) GetHeight() int64 {
	if gr.Fullscreen {
		return gr.Height
	}
	return gr.WindowHeight
}

func (gr graphics) GetHeightF() float64 {
	if gr.Fullscreen {
		return float64(gr.Height)
	}
	return float64(gr.WindowHeight)
}

func (gr graphics) GetAspectRatio() float64 {
	if gr.Fullscreen {
		return float64(gr.Width)/float64(gr.Height)
	}
	return float64(gr.WindowWidth)/float64(gr.WindowHeight)
}

type audio struct {
	GeneralVolume float64 //0.5
	MusicVolume float64 //=0.5
	SampleVolume float64 //=0.5
	EnableBeatmapSampleVolume bool //= false
}

type beat struct {
	BeatScale float64 //1.4
}

type color struct {
	EnableRainbow bool //true
	RainbowSpeed float64 //8, degrees per second
	BaseHue float64 //0..360, if EnableRainbow is disabled then this value will be used to calculate base color
	Saturation float64 //1.0
	Value float64 //1.0
	EnableCustomHueOffset bool //false, false means that every iteration has an offset of i*360/n
	HueOffset float64 //0, custom hue offset for mirror collages
	FlashToTheBeat bool //true, objects size is changing with music peak amplitude
	FlashAmplitude float64 //50, hue offset for flashes
}

type cursor struct {
	UseObjectColors bool //true, overrides lower color settings
	Colors *color
	EnableCustomTrailGlowOffset bool //true, if enabled, value set below will be used, if not, HueOffset of previous iteration will be used (or offset of 180° for single cursor)
	TrailGlowOffset float64 //-36, offset of the cursor trail glow
	ScaleToCS bool //false, if enabled, cursor will scale to beatmap CS value
	CursorSize float64 //18, cursor radius in osu!pixels
	ScaleToTheBeat bool //true, cursor size is changing with music peak amplitude
	ShowCursorsOnBreaks bool //true
	BounceOnEdges bool //false
}

type objects struct {
	MandalaTexturesTrigger int //5, minimum value of cursors needed to use more translucent textures
	DrawApproachCircles bool //true
	UseCursorColors bool //true, overrides lower color settings
	Colors *color
	ObjectsSize float64 //-1, objects radius in osu!pixels. If value is less than 0, beatmap's CS will be used
	ScaleToTheBeat bool //true, objects size is changing with music peak amplitude
	SliderLOD int64 //30, number of triangles in a circle
	SliderPathLOD int64 //0.5, int(pixelLength*(PathLOD/100)) => number of slider path points
	DrawFollowPoints bool //true
	WhiteFollowPoints bool //true
	FollowPointColorOffset float64 //0.0, hue offset of the followpoint
}

type playfield struct {
	LeadInTime float64 //5, time to the beginning of music
	BackgroundInDim float64 //0, background dim at the start of app
	BackgroundDim float64 // 0.95, background dim at the beatmap start
	BackgroundDimBreaks float64 // 0.95, background dim at the breaks
	Scale float64 //1, scale the playfield (1 means that 512 will be rescaled to 800 on FullHD monitor)
	FlashToTheBeat bool //true, background dim varies accoriding to music power
	KiaiFactor float64 //1.2, scale and flash factor during Kiai
}

type fileformat struct {
	Version *string
	General *general
	Graphics *graphics
	Audio *audio
	Beat *beat
	Cursor *cursor
	Objects *objects
	Playfield *playfield
}

var Version string
var General *general
var Graphics *graphics
var Audio *audio
var Beat *beat
var Cursor *cursor
var Objects *objects
var Playfield *playfield

var DIVIDES = 2