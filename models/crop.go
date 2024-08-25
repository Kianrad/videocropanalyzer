package models

type CropValues struct {
	Top    int `json:"top"`
	Bottom int `json:"bottom"`
	Left   int `json:"left"`
	Right  int `json:"right"`
}

type HorizontalCrop struct {
	Top    int `json:"top"`
	Bottom int `json:"bottom"`
}

type VerticalCrop struct {
	Left  int `json:"left"`
	Right int `json:"right"`
}
