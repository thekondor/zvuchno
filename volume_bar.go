// This file is a part of git.thekondor.net/zvuchno.git (mirror: github.com/thekondor/zvuchno)

package main

import (
	"bytes"
	"fmt"
	"github.com/cheggaaa/pb"
	"log"
	"text/template"
)

type VolumeBar struct {
	progressBar *pb.ProgressBar
	tpl         *template.Template
}

func NewVolumeBar(width byte, formatTpl, barFormat string) VolumeBar {
	bar := pb.New(100)
	bar.ManualUpdate = true
	bar.ShowCounters = false
	bar.ShowTimeLeft = false
	bar.ShowFinalTime = false
	bar.ShowPercent = false
	bar.SetWidth(int(width))
	bar.Format(barFormat)

	return VolumeBar{bar, template.Must(template.New("full format").Parse(formatTpl))}
}

func (vb VolumeBar) Update(value int) string {
	percent := value / 655
	vb.progressBar.Set(percent)
	vb.progressBar.Update()

	return vb.toString(percent)
}

func (vb VolumeBar) toString(percent int) string {
	var output bytes.Buffer

	progress := vb.progressBar.String()
	if err := vb.tpl.Execute(&output, struct {
		Percent int
		Bar     string
	}{percent, progress}); nil != err {
		log.Printf("W: template formatting error = %s", err)
		return fmt.Sprintf("(fallback): %d%% %s", percent, progress)
	}

	return output.String()
}
