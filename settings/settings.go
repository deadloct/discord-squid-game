package settings

import (
	"io/ioutil"
	"path"
	"text/template"
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	// Seconds
	DefaultStartDelay = 60
	MinimumStartDelay = 5
	MaximumStartDelay = 60 * 60 // 60 minutes in seconds

	EnableClone  = true
	DefaultClone = 1
	MinimumClone = 1
	MaximumClone = 100

	DefaultDayDelay    = 5 * time.Second
	DefaultVictorCount = 1
	MinimumVictorCount = 0

	JokeInterval = 10 * time.Second

	PhrasesFile = "phrases.en.json"
	JokesFile   = "jokes.en.json"
	IntroFile   = "intro.en.template"
	HelpFile    = "help.en.template"

	MaxMsgLen            = 1500
	DiscordMaxMessages   = 100
	DiscordMaxBulkDelete = 100

	WhiteSpaceChar = "\u200d"

	DataLocation = "data"
	DayEmoji     = "skull_crossbones"
)

var (
	Intro *template.Template
	Help  string // not currently a template
)

type IntroValues struct {
	Delay       time.Duration
	EntryEmoji  string
	EffieEmoji  string
	Sponsor     string
	VictorCount int
}

func ImportData() {
	importIntro()
	importHelp()
}

func importIntro() {
	introPath := path.Join(DataLocation, IntroFile)
	v, err := ioutil.ReadFile(introPath)
	if err != nil {
		log.Panicf("unable to open intro template file %v: %v", introPath, err)
	}

	tmplStr := string(v[:])
	Intro, err = template.New("intro-template").Parse(tmplStr)
	if err != nil {
		log.Panicf("unable to parse intro template '%v': %v", tmplStr, err)
	}

	log.Info("imported intro template")
}

func importHelp() {
	helpPath := path.Join(DataLocation, HelpFile)
	v, err := ioutil.ReadFile(helpPath)
	if err != nil {
		log.Panicf("unable to open help file %v: %v", helpPath, err)
	}

	Help = string(v[:])
	log.Info("imported help file")
}
