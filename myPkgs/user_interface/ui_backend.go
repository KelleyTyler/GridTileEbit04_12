package user_interface

import (
	"bytes"
	"log"

	gensound "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/generated_sound"
	settings "github.com/KelleyTyler/GridTileEbit04_12/myPkgs/settingsconfig"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/gomono"
	"golang.org/x/image/font/gofont/goregular"
)

type UI_Backend struct {
	Settings                    *settings.GameSettings
	SoundSystem                 *gensound.Basic_SoundSystem
	Btn_Sounds                  [][]byte
	Textsrcs                    []*text.GoTextFaceSource
	Btn_Text_Mono, Btn_Text_Reg text.Face
}

func GetUIBackend(settings *settings.GameSettings, gsounds *gensound.Basic_SoundSystem) UI_Backend {
	// gsounds := gensound.InitSoundSet(settings, 3200, 480)
	// gsounds.Init01(settings,3200,480)
	bckend := UI_Backend{
		Settings:    settings,
		SoundSystem: gsounds,
	}

	if gsounds == nil {
		soundsys := gensound.Get_Basic_SoundSystem(bckend.Settings, 3200, 400)
		bckend.SoundSystem = &soundsys
	}
	bckend.Textsrcs = make([]*text.GoTextFaceSource, 0)
	var err error
	var tempTextSrc *text.GoTextFaceSource
	tempTextSrc, err = text.NewGoTextFaceSource(bytes.NewReader(gomono.TTF))
	if err != nil {
		log.Fatal("err: ", err)
	}
	bckend.Textsrcs = append(bckend.Textsrcs, tempTextSrc)
	bckend.Btn_Text_Mono = &text.GoTextFace{
		Source: bckend.Textsrcs[0],
		Size:   20,
	}
	tempTextSrc, err = text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal("err: ", err)
	}
	bckend.Textsrcs = append(bckend.Textsrcs, tempTextSrc)
	bckend.Btn_Text_Reg = &text.GoTextFace{
		Source: bckend.Textsrcs[1],
		Size:   10,
	}

	// bckend.Btn_Text_Mono
	bckend.InitSounds()
	return bckend
}

func (uiBack *UI_Backend) init() {

}

func (uiBack *UI_Backend) PlaySound(sound_num int) {
	if sound_num < int(len(uiBack.Btn_Sounds)) {
		uiBack.SoundSystem.PlayByte(uiBack.Btn_Sounds[sound_num])
	}
}
func (uiBack *UI_Backend) InitSounds() {
	uiBack.Btn_Sounds = append(uiBack.Btn_Sounds, gensound.Soundwave_CreateSound(3200, 200, 0, 110, []float32{1.0}, []float32{0.0750000}))
	uiBack.Btn_Sounds = append(uiBack.Btn_Sounds, gensound.Soundwave_CreateSound(3200, 200, 10, 110, []float32{1.0}, []float32{0.0750000}))
	uiBack.Btn_Sounds = append(uiBack.Btn_Sounds, gensound.Soundwave_CreateSound(3200, 200, 15, 110, []float32{1.0}, []float32{0.0750000}))
	uiBack.Btn_Sounds = append(uiBack.Btn_Sounds, gensound.Soundwave_CreateSound(3200, 200, 20, 110, []float32{1.0}, []float32{0.0750000}))
	uiBack.Btn_Sounds = append(uiBack.Btn_Sounds, gensound.Soundwave_CreateSound(3200, 200, 25, 110, []float32{1.0}, []float32{0.0750000}))
	uiBack.Btn_Sounds = append(uiBack.Btn_Sounds, gensound.Soundwave_CreateSound(3200, 200, 25, 110, []float32{1.0}, []float32{0.0750000}))
}

/*
func (ui_Helper *UI_Helper) InitSounds() { //1200

		ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 0, 110, []float32{1.0}, []float32{0.0750000}))
		ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 10, 110, []float32{1.0}, []float32{0.0750000}))
		ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 15, 110, []float32{1.0}, []float32{0.0750000}))
		ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 20, 110, []float32{1.0}, []float32{0.0750000}))
		ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 25, 110, []float32{1.0}, []float32{0.0750000}))
		ui_Helper.Btn_Sounds = append(ui_Helper.Btn_Sounds, Soundwave_CreateSound(3200, 220, 25, 110, []float32{1.0}, []float32{0.0750000}))
	}
*/
func (uiBack *UI_Backend) GetTextFace(textnum, size int) *text.Face {
	var textOut text.Face
	textOut = &text.GoTextFace{
		Source: uiBack.Textsrcs[textnum],
		Size:   float64(size),
	}
	return &textOut
}
