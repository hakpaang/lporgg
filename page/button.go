package page

import (
	"log/slog"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"github.com/blacktop/lporg/internal/command"
)

func DefaultBtn(win fyne.Window) *widget.Button {
	return widget.NewButton("default", func() {
		useICloud := false
		backup := false
		items := []*widget.FormItem{
			widget.NewFormItem("use iCloud for config", widget.NewCheck("", func(checked bool) {
				useICloud = checked
			})),
			widget.NewFormItem("Backup your current Launchpad/Dock settings?", widget.NewCheck("", func(checked bool) {
				backup = checked
			})),
		}
		dialog.ShowForm("Default...", "Yes", "Cancel", items, func(b bool) {
			if !b {
				return
			}
			conf := &command.Config{
				Cmd:    "default",
				Cloud:  useICloud,
				Backup: backup,
			}

			if err := conf.Verify(); err != nil {
				return
			}
			if conf.Backup {
				slog.Info("Backing up current launchpad settings")
				if err := command.SaveConfig(conf); err != nil {
					return
				}
			}
			slog.Info("Apply default launchpad settings")
			if err := command.DefaultOrg(conf); err != nil {
				dialog.ShowError(err, win)
				return
			}
			dialog.ShowInformation("Default", "Launchpad settings defaulted", win)
		}, win)
	})
}

func SaveBtn(win fyne.Window) *widget.Button {
	return widget.NewButton("save", func() {
		iCloud := false
		items := []*widget.FormItem{
			widget.NewFormItem("use iCloud for config", widget.NewCheck("", func(checked bool) {
				iCloud = checked
			})),
		}
		dialog.ShowForm("Save...", "Save", "Cancel", items, func(b bool) {
			if !b {
				return
			}
			conf := &command.Config{
				Cmd:   "save",
				Cloud: iCloud,
			}
			if err := conf.Verify(); err != nil {
				return
			}
			slog.Info("Saving launchpad settings")
			if err := command.SaveConfig(conf); err != nil {
				dialog.ShowError(err, win)
				return
			}
			dialog.ShowInformation("Saved", "Launchpad settings saved:"+conf.File, win)
		}, win)
	})
}

func LoadBtn(win fyne.Window) *widget.Button {
	return widget.NewButton("load", func() {
		useICloud := false
		backup := false
		items := []*widget.FormItem{
			widget.NewFormItem("use iCloud for config", widget.NewCheck("", func(checked bool) {
				useICloud = checked
			})),
			widget.NewFormItem("Backup your current Launchpad/Dock settings?", widget.NewCheck("", func(checked bool) {
				backup = checked
			})),
		}
		dialog.ShowForm("Load...", "Yes", "Cancel", items, func(b bool) {
			if !b {
				return
			}
			conf := &command.Config{
				Cmd:    "load",
				Cloud:  useICloud,
				Backup: backup,
			}

			if err := conf.Verify(); err != nil {
				return
			}
			if conf.Backup {
				slog.Info("Backing up current launchpad settings")
				if err := command.SaveConfig(conf); err != nil {
					return
				}
			}
			if err := command.LoadConfig(conf); err != nil {
				dialog.ShowError(err, win)
				return
			}
			dialog.ShowInformation("Default", "Launchpad settings defaulted", win)
		}, win)
	})
}

func RevertBtn(win fyne.Window) *widget.Button {
	return widget.NewButton("revert", func() {
		useICloud := false
		items := []*widget.FormItem{
			widget.NewFormItem("use iCloud for config", widget.NewCheck("", func(checked bool) {
				useICloud = checked
			})),
		}
		dialog.ShowForm("Revert...", "Yes", "Cancel", items,
			func(b bool) {
				if !b {
					return
				}
				conf := &command.Config{
					Cmd:   "revert",
					Cloud: useICloud,
				}

				if err := conf.Verify(); err != nil {
					return
				}
				if err := command.LoadConfig(conf); err != nil {
					dialog.ShowError(err, win)
					return
				}
				dialog.ShowInformation("Reverted", "Launchpad settings reverted", win)
			}, win)
	})
}
