package main

import (
	"time"

	aw "github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	t := time.Now().UTC()
	it := wf.NewItem(t.Format("15:04"))
	it.Subtitle(t.Format("Mon, 02 Jan 2006 – Coordinated Universal Time (UTC)"))
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
