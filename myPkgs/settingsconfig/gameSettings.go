package settings

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type GameSettings struct {
	VersionID   string `json:"versionID,omitempty"`
	WindowSizeX int    `json:"window_size_x"`
	WindowSizeY int    `json:"window_size_y"`
	ScreenResX  int    `json:"screen_res_x,"`
	ScreenResY  int    `json:"screen_res_y,"`
	//-------
	UIAudioVolume          int `json:"ui_audio_volume,"`
	GameBoardX             int `json:"game_board_x,"`
	GameBoardY             int `json:"game_board_y,"`
	GameBoardTileX         int `json:"game_board_tile_x,"`
	GameBoardTileY         int `json:"game_board_tile_y,"`
	GameBoardTile_Margin_X int `json:"game_board_tile_margin_x,"`
	GameBoardTile_Margin_Y int `json:"game_board_tile_margin_y,"`
}

func (sets *GameSettings) ToString() string {
	outstrng := fmt.Sprintf("SETTINGS:\n%12s: %s\n", "VERSION", sets.VersionID)
	outstrng += fmt.Sprintf("%12s: %3d %3d\n", "Window Size", sets.WindowSizeX, sets.WindowSizeY)
	outstrng += fmt.Sprintf("%12s: %3d %3d\n", "Screen Res", sets.ScreenResX, sets.ScreenResY)
	return outstrng
}

/*
This will load from a JSON file;
*/
func GetBytesFromJSON(filePath string) ([]byte, error) {
	fmt.Print("INIT JSON HELLO!\n\n")
	jSonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := jSonFile.Close(); err != nil {
			panic(err)
		}
	}()
	var rdal []byte
	rdr := bufio.NewReader(jSonFile)
	rdal, err = io.ReadAll(rdr)
	if err != nil {
		return nil, err
		// panic(err)
	}
	return rdal, nil
}

//	func(gSets *GameSettings) GetSettingsFromJSON(){
//		get
//	}
func GetSettingsFromJSON() GameSettings {
	var gSets GameSettings
	bee, err0 := GetBytesFromJSON("init.JSON")
	if err0 != nil {

		return GetSettingsFromBakedIn()
	}
	err2 := json.Unmarshal(bee, &gSets)
	if err2 != nil {
		log.Fatal(err2)
	}
	return gSets
}

func GetSettingsFromBakedIn() GameSettings {
	var gSets GameSettings = GameSettings{
		VersionID:              "0.0.00",
		WindowSizeX:            960, //860//892
		WindowSizeY:            640, //660 //720
		ScreenResX:             960, //860 //892
		ScreenResY:             640,
		UIAudioVolume:          100,
		GameBoardX:             32,
		GameBoardY:             32,
		GameBoardTileX:         16,
		GameBoardTileY:         16,
		GameBoardTile_Margin_X: 2,
		GameBoardTile_Margin_Y: 2,
	}
	return gSets
}
