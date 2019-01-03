package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"

	"github.com/gotk3/gotk3/gtk"
	cron "gopkg.in/robfig/cron.v2"
)

type boss struct {
	short   string
	name    string
	enabled bool
}

var table map[string][]int

var bosses = []*boss{
	&boss{"admiral", "Admiral Taidha Covington", true},
	&boss{"claw", "Claw of Jormag", true},
	&boss{"jungleWorm", "Evolved Jungle Wurm", true},
	&boss{"fireElemental", "Fire Elemental", true},
	&boss{"golem", "Golem Mark II", true},
	&boss{"greatWorm", "Great Jungle Wurm", true},
	&boss{"karka", "Karka Queen", true},
	&boss{"megadestroyer", "Megadestroyer", true},
	&boss{"modniir", "Modniir Ulgoth", true},
	&boss{"behemoth", "Shadow Behemoth", true},
	&boss{"svanir", "Svanir Shaman Chief", true},
	&boss{"tequatl", "Tequatl the Sunless", true},
	&boss{"shatterer", "The Shatterer", true},
}

func main() {

	bossData, err := Asset("assets/data.json")
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bossData, &table)
	if err != nil {
		fmt.Print(err)
	}

	c := cron.New()
	c.AddFunc("*/15 * * * *", tick)
	c.Start()

	buildUI()
}

func buildUI() {
	data, err := Asset("assets/ui.glade")
	if err != nil {
		panic(err)
	}

	gtk.Init(nil)
	builder, err := gtk.BuilderNew()
	builder.AddFromString(string(data))
	if err != nil {
		panic(err)
	}
	obj, err := builder.GetObject("MainWindow")
	if err != nil {
		panic(err)
	}
	if w, ok := obj.(*gtk.Window); ok {
		w.Connect("destroy", func() {
			gtk.MainQuit()
		})
		w.ShowAll()
	} else {
		fmt.Println("not a *gtk.Window")
	}

	// create an event handler for every boss switch.
	// switches are named after the short name of the respective boss
	for _, b := range bosses {
		obj, err = builder.GetObject(b.short)

		if bossSwitch, ok := obj.(*gtk.Switch); ok {
			buildHandler(bossSwitch, b.short)
		}
	}

	gtk.Main()
}

// buildHandler creates an state-set event handler that (de)activate boss notifications
func buildHandler(sw *gtk.Switch, bossShort string) {
	sw.Connect("state-set", func(gtkSwitch *gtk.Switch) {
		getBoss(bossShort).enabled = gtkSwitch.GetActive()
	})
}

func tick() {
	time := getNextTime()
	nextBosses := table[getNextTime()]
	for _, bossID := range nextBosses {
		if bosses[bossID].enabled {
			notify(bosses[bossID].name, time)
		}
	}
}

func getBoss(short string) *boss {
	for _, b := range bosses {
		if b.short == short {
			return b
		}
	}
	return nil
}

func notify(bossName string, time string) {

	args := fmt.Sprintf("%s Starting at %s", bossName, time)

	cmd := exec.Command("notify-send", args)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func getNextTime() string {
	now := time.Now().Add(time.Duration(15 * time.Minute))

	hh := zeroPrefix(now.Hour())
	mm := zeroPrefix(now.Minute())

	return fmt.Sprintf("%s:%s", hh, mm)
}

func zeroPrefix(i int) string {
	s := strconv.Itoa(i)
	if i < 10 {
		s = "0" + s
	}
	return s
}
