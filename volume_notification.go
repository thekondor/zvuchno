// This file is a part of git.thekondor.net/zvuchno.git (mirror: github.com/thekondor/zvuchno)

package main

import (
	"github.com/godbus/dbus"
)

type VolumeNotification struct {
	conn                      dbus.BusObject
	notificationReplacementID uint32
}

func NewVolumeNotification(bus *dbus.Conn) VolumeNotification {
	obj := bus.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	return VolumeNotification{conn: obj}
}

func (vn *VolumeNotification) Show(title, message string, timeout uint32) {
	hints := map[string]interface{}{}
	out := vn.conn.Call(
		"org.freedesktop.Notifications.Notify", 0, "Volume",
		vn.notificationReplacementID, "volume", title, message, []string{}, hints, int(timeout))
	out.Store(&vn.notificationReplacementID)
}
