package port

import (
	"bytes"
	"encoding/json"
	"image"
	"image/png"
	"math"

	"github.com/swim-services/swim_porter/port/internal"
	"github.com/swim-services/swim_porter/port/utils"
)

func (p *porter) xp() error {
	if err := p.xpBar(); err != nil {
		return err
	}
	if err := p.xpJson(); err != nil {
		return err
	}
	return nil
}

func (p *porter) xpBar() error {
	if iconsFile, err := p.out.Read("textures/gui/icons.png"); err == nil {
		icons, err := png.Decode(bytes.NewReader(iconsFile))
		iconsSub := icons.(interface {
			SubImage(r image.Rectangle) image.Image
		})
		if err != nil {
			return err
		}
		bounds := icons.Bounds()
		x := 0
		y := bounds.Dy() / 4
		padding := float64(bounds.Dx()) / 3.45945945946
		extra := int(math.Round(padding))
		xpBarLength := bounds.Dx() - extra
		xpBarWidth := 5 * (bounds.Dx() / 256)
		emptyBar := iconsSub.SubImage(image.Rect(x, y, x+xpBarLength, y+xpBarWidth))
		fullBar := iconsSub.SubImage(image.Rect(x, y+xpBarWidth, x+xpBarLength, y+(xpBarWidth*2)))
		if err := internal.WritePng(emptyBar, "textures/ui/experiencebarempty.png", p.out); err != nil {
			return err
		}
		if err := internal.WritePng(fullBar, "textures/ui/experiencebarfull.png", p.out); err != nil {
			return err
		}
		p.out.Copy("textures/ui/experiencebarempty.png", "textures/gui/achievements/hotdogempty.png")
		p.out.Copy("textures/ui/experiencebarfull.png", "textures/gui/achievements/hotdogfull.png")
		p.out.Copy("textures/ui/experiencebarempty.png", "textures/ui/empty_progress_bar.png")
		p.out.Copy("textures/ui/experiencebarfull.png", "textures/ui/filled_progress_bar.png")

		if nub, err := assetsMapFS.Read("nub/nub.png"); err == nil {
			p.out.Write(nub, "textures/gui/achievements/nub.png")
		}
		if experienceNub, err := assetsMapFS.Read("nub/experiencenub.png"); err == nil {
			p.out.Write(experienceNub, "textures/ui/experiencenub.png")
		}
		if experienceBlueNub, err := assetsMapFS.Read("nub/experience_bar_nub_blue.png"); err == nil {
			p.out.Write(experienceBlueNub, "textures/ui/experience_bar_nub_blue.png")
		}
	}
	return nil
}

func (p *porter) xpJson() error {
	empty := utils.XpBar{
		NinesliceSize: [4]int{6, 1, 6, 1},
		BaseSize:      [2]int{182, 5},
	}
	emptyBytes, err := json.Marshal(empty)
	if err != nil {
		return err
	}
	p.out.Write(emptyBytes, "textures/ui/experiencebarempty.json")
	p.out.Write(emptyBytes, "textures/ui/empty_progress_bar.json")

	full := utils.XpBar{
		NinesliceSize: [4]int{1, 0, 1, 0},
		BaseSize:      [2]int{182, 5},
	}
	fullBytes, err := json.Marshal(full)
	if err != nil {
		return err
	}
	p.out.Write(fullBytes, "textures/ui/experiencebarfull.json")
	p.out.Write(fullBytes, "textures/ui/filled_progress_bar.json")

	return nil
}
