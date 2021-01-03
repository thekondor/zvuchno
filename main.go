// This file is a part of git.thekondor.net/zvuchno.git (mirror: github.com/thekondor/zvuchno)

package main

import (
	"github.com/godbus/dbus"
	"github.com/sqp/pulseaudio"
	"log"
)

type AppSettings struct {
	Timeout      uint32
	Title        string
	OnMuteText   string
	OnUnmuteText string
}

type App struct {
	deviceEventCh      chan struct{}
	volumeNotification VolumeNotification
	bar                VolumeBar
	settings           AppSettings
}

func (app *App) DeviceVolumeUpdated(path dbus.ObjectPath, values []uint32) {
	app.lockPulseAudio(true)
	defer app.lockPulseAudio(false)

	if values[0] == values[1] {
		app.volumeNotification.Show(app.settings.Title, app.bar.Update(int(values[0])), app.settings.Timeout)
	}
}

func (app *App) DeviceMuteUpdated(path dbus.ObjectPath, value bool) {
	app.lockPulseAudio(true)
	defer app.lockPulseAudio(false)

	var message string
	if value {
		message = app.settings.OnMuteText
	} else {
		message = app.settings.OnUnmuteText
	}

	app.volumeNotification.Show(app.settings.Title, message, app.settings.Timeout)
}

func (app *App) Loop(pa *pulseaudio.Client) {
	app.lockPulseAudio(false)
	pa.Register(app)
	pa.Listen()
}

func (app *App) lockPulseAudio(state bool) {
	if state {
		<-app.deviceEventCh
	} else {
		app.deviceEventCh <- struct{}{}
	}
}

func mustPulseAudio() *pulseaudio.Client {
	pa, err := pulseaudio.New()
	if err != nil {
		log.Panicf("Failed to connect to PulseAudio: %s", err)
	}

	return pa
}

func mustDbus() *dbus.Conn {
	sessionBus, err := dbus.SessionBus()
	if err != nil {
		log.Panicf("Failed to connect to DBUS session bus: %s", err)
	}

	return sessionBus
}

func main() {
	pa := mustPulseAudio()
	dbusConn := mustDbus()

	config := NewConfig()
	app := &App{
		deviceEventCh:      make(chan struct{}, 1),
		volumeNotification: NewVolumeNotification(dbusConn),
		bar: NewVolumeBar(config.Appearance.Width,
			config.Appearance.Format.Full,
			config.Appearance.Format.Bar),
		settings: AppSettings{config.Notification.Timeout,
			config.Appearance.Text.Title,
			config.Appearance.Text.OnMute,
			config.Appearance.Text.OnUnmute}}

	log.Printf("Serving PA events...")
	app.Loop(pa)
}
