package core


import "path/filepath"
import "os"
import "encoding/json"


type Theme struct {
	ImagesZipUri string `json:"imagesZipUri"`
	ImageFilename string `json:"imageFilename"`
	ImageCredits string `json:"imageCredits"`
	DayHighlight int `json:"dayHighlight"`
	NightHighlight int `json:"nightHighlight"`
	DayImageList []int `json:"dayImageList"`
	SunriseImageList []int `json:"sunriseImageList"`
	SunsetImageList []int `json:"sunsetImageList"`
	NightImageList []int `json:"nightImageList"`
}





func WriteTheme(basename string, steps int) {
	sunriseLen := steps - 2
	if sunriseLen < 0 {
		sunriseLen = 0
	}

	sunriseIndex := make([]int, sunriseLen)
	for idx := 0; idx < len(sunriseIndex); idx++ {
		sunriseIndex[idx] = idx+2
	}

	sunsetIndex := make([]int, len(sunriseIndex))
	for idx := 0; idx < len(sunriseIndex); idx++ {
		sunsetIndex[idx] = sunriseIndex[len(sunriseIndex) - idx - 1]
	}


	// generate json
	theme := Theme { ImagesZipUri: "", 
					 ImageFilename: "image_*.png",
					 ImageCredits: "",
					 DayHighlight: 1,
					 NightHighlight: steps,
					 DayImageList: []int { 1 },
					 SunriseImageList: sunriseIndex,
					 SunsetImageList: sunsetIndex,
					 NightImageList: []int { steps } }

	outFile := filepath.Join(basename, "theme.json")
	outputFile, err := os.Create(outFile)
	if err != nil {
		panic("failed to write json file")
	}
	defer outputFile.Close()	

	encoder := json.NewEncoder(outputFile)
	encoder.SetIndent("", " ")
	err = encoder.Encode(theme)
	if err != nil {
		panic("failed to encode to json")
	}
}
