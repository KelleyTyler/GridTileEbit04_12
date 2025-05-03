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
	VersionID   string `json:"versionID,omitempty,"`
	WindowSizeX int    `json:"window_size_x,"`
	WindowSizeY int    `json:"window_size_y,"`
	ScreenResX  int    `json:"screen_res_x,"`
	ScreenResY  int    `json:"screen_res_y,"`
	//-------
	SavePath string `json:"save_path,"`
	//-------
	UIAudioVolume          int `json:"ui_audio_volume,"`
	GameBoardX             int `json:"game_board_x,"`
	GameBoardY             int `json:"game_board_y,"`
	GameBoardXMax          int `json:"game_board_x_max,"`
	GameBoardYMax          int `json:"game_board_y_max,"`
	GameBoardTileX         int `json:"game_board_tile_x,"`
	GameBoardTileY         int `json:"game_board_tile_y,"`
	GameBoardTile_Margin_X int `json:"game_board_tile_margin_x,"`
	GameBoardTile_Margin_Y int `json:"game_board_tile_margin_y,"`
	//--------
}

/**/
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
	log.Print("INIT JSON HELLO!\n\n")
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

/**/
func Write_Byes_To_File(init_bytes []byte) error {
	file, err := os.Create("init.JSON")
	if err != nil {
		return err
	}

	defer func() {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}()
	wrtr := bufio.NewWriter(file)
	for _, b := range init_bytes {
		err := wrtr.WriteByte(b)
		if err != nil {
			return err
		}
	}
	wrtr.Flush()

	return nil
}

func MakingDirectory(file_path string) error {

	curr_dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return err
	}
	file_path_all := fmt.Sprintf("%s\\%s", curr_dir, file_path)
	// fpath := "/bin/Output/"
	if !directoryExists(file_path_all) {
		err = os.MkdirAll(file_path_all, 0700)
		// err := os.Mkdir(fpath, 0700)
		if err != nil {
			log.Println(err)

			return err
		}
	} else {
		log.Printf("Directory Exists! %s\n", file_path_all)
	}
	return nil
}

/**/
func directoryExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		// Path exists, check if it is a directory
		if stat, err := os.Stat(path); err == nil && stat.IsDir() {
			return true
		}
	}
	return false
}

/**/
func GetSettingsFromJSON() GameSettings {
	var gSets GameSettings
	bee, err0 := GetBytesFromJSON("init.JSON")
	if err0 != nil {
		gSets = GetSettingsFromBakedIn()
		// init_bytes, err := json.Marshal(gSets)
		// if err != nil {
		// 	log.Printf("ERROR ERROR ERROR!!!\n")
		// }
		// err = Write_Byes_To_File(init_bytes)
		// if err != nil {
		// 	panic(err)
		// }
		// MakingDirectory(gSets.SavePath)
		return gSets
	}
	err2 := json.Unmarshal(bee, &gSets)
	if err2 != nil {
		log.Fatal(err2)
	}

	//MakingDirectory(gSets.SavePath)
	return gSets
}

/**/
func GetSettingsFromBakedIn() GameSettings {
	var gSets GameSettings = GameSettings{
		VersionID:              "0.0.00",
		WindowSizeX:            960, //860//892
		WindowSizeY:            640, //660 //720
		ScreenResX:             960, //860 //892
		ScreenResY:             640,
		SavePath:               "bin\\Output",
		UIAudioVolume:          100,
		GameBoardX:             64,
		GameBoardY:             64,
		GameBoardXMax:          128,
		GameBoardYMax:          128,
		GameBoardTileX:         16,
		GameBoardTileY:         16,
		GameBoardTile_Margin_X: 2,
		GameBoardTile_Margin_Y: 2,
	}
	return gSets
}
