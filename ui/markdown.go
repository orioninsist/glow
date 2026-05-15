package ui

import (
	"fmt"
	"math"
	"time"
	"unicode"

	"github.com/charmbracelet/log"
	"github.com/dustin/go-humanize"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type markdown struct {
	// Full path of a local markdown file. Only relevant to local documents and
	// those that have been stashed in this session.
	localPath string

	// Value we filter against. This exists so that we can maintain positions
	// of filtered items if notes are edited while a filter is active. This
	// field is ephemeral, and should only be referenced during filtering.
	filterValue string

	Body    string
	Note    string
	Modtime time.Time
}

const projectCommandAndConfigGuideTR = `
# Glow Proje Komut ve Config Kontrolu

Bu not, bu checkout icindeki kaynak kod taranarak hazirlandi. Komutlar Cobra
baglantilarindan ` + "`main.go`" + `, config anahtarlari Viper baglantilarindan ve
` + "`glow.yml`" + ` dosyasindan gelir.

## Kaynak Analizi

Onemli dosyalar:

` + "```text" + `
main.go              CLI/TUI girisi, flag'ler, config okuma, URL/dosya/stdin render akisi
config_cmd.go        glow config komutu ve varsayilan glow.yml icerigi
github.go            GitHub repo/blob linklerini raw markdown kaynagina cevirir
gitlab.go            GitLab repo/blob linklerini raw markdown kaynagina cevirir
url.go               URL yardimcilari
utils/utils.go       Stil secimi, frontmatter temizleme, code block sarmalama
ui/                  Bubble Tea TUI, pager, liste, filtre ve editor akisi
glow.yml             Bu checkout icin onerilen config dosyasi
myconfig/goz-yormayan.json  Goz yormayan ozel Glamour stili
` + "```" + `

Config yukleme sirasi:

` + "```text" + `
GLOW_CONFIG_HOME verilirse once o dizindeki glow.yml okunur.
XDG_CONFIG_HOME verilirse XDG config dizini de aranir.
Hiç config bulunmazsa sistemin varsayilan glow config dizininde glow.yml olusturulur.
Bu checkout icin onerilen kullanim:
GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow
` + "```" + `

## Tum CLI Komutlari

Derleme ve test:

` + "```bash" + `
cd /mnt/samsung/orion-backup-local/projects/glow
go test ./...
go build -o glow .
` + "```" + `

Repo kokundeki config ile calistirma:

` + "```bash" + `
export GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow
export PAGER="less -R"
./glow README.md
` + "```" + `

Temel kullanim:

` + "```bash" + `
glow                         # mevcut dizinde TUI acar
glow .                       # dizin README'sini/TUI akisina uygun kaynagi acar
glow README.md               # lokal markdown dosyasini render eder
glow -                       # stdin'den markdown okur
curl -L "URL" | glow -p -
glow "https://github.com/KULLANICI/REPO"
glow "https://github.com/KULLANICI/REPO/blob/main/DOSYA.md"
glow "https://raw.githubusercontent.com/KULLANICI/REPO/refs/heads/main/DOSYA.md"
glow "https://gitlab.com/KULLANICI/REPO"
` + "```" + `

Okuma modlari:

` + "```bash" + `
glow -p -w 100 README.md        # pager ile oku
glow -t -w 100 README.md        # TUI ile oku
glow -w 100 README.md           # direkt terminale bas
glow -p -w 80 README.md         # dar metin alani
glow -p -w 120 README.md        # genis kod bloklari icin
` + "```" + `

Stil ve gorsel komutlari:

` + "```bash" + `
glow -s auto README.md
glow -s dark README.md
glow -s light README.md
glow -s /mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json -p -w 100 README.md
glow --kitty-images -p -w 100 README.md
glow --kitty-images -p -w 100 "https://raw.githubusercontent.com/KULLANICI/REPO/refs/heads/main/README.md"
` + "```" + `

TUI dosya listeleme komutlari:

` + "```bash" + `
glow -t .
glow -t -a .
glow -t -l .
glow -t -m .
glow -t -a -l -w 100 .
` + "```" + `

Config ve yardim komutlari:

` + "```bash" + `
glow config
glow config --config /path/to/glow.yml
glow --config /mnt/samsung/orion-backup-local/projects/glow/glow.yml README.md
glow --help
glow --version
glow completion zsh
glow man
` + "```" + `

Notlar:

` + "```text" + `
-p ve -t ayni komutta birlikte kullanilmaz.
--mouse flag'i vardir ama gizlidir; TUI icin config'ten mouse: true onerilir.
--kitty-images sadece Kitty terminalde ve CLI output akisi icin anlamlidir.
GITHUB/GITLAB repo veya blob linkleri raw markdown kaynagina cevrilebilir.
` + "```" + `

## glow.yml Ayarlari

Bu checkout icin en iyi genel ayar:

` + "```yaml" + `
style: "/mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json"
pager: false
tui: false
width: 100
mouse: true
all: false
showLineNumbers: true
preserveNewLines: true
kittyImages: true
` + "```" + `

Ayar aciklamalari:

` + "```text" + `
style:
  Markdown renk/stil temasi. "auto", "dark", "light", "pink", "dracula",
  "tokyo-night" gibi Glamour stilleri veya JSON stil dosyasi yolu olabilir.
  Bu proje icin goz yormayan ozel JSON dosyasi onerilir.

pager:
  true olursa CLI render sonucu $PAGER ile acilir. Uzun README ve GitHub
  dokumanlari icin komutta -p kullanmak daha kontrolludur.

tui:
  true olursa cikti interaktif Bubble Tea TUI icinde acilir. pager ile ayni
  anda true olamaz.

width:
  Kelime sarma genisligi. 100 bu proje icin dengeli; 80 daha dar, 120 kod
  bloklari icin daha genistir. 0 terminal genisligini otomatik kullanir.

mouse:
  TUI'da mouse/tekerlek destegini acar. CLI render akisini etkilemez.

all:
  TUI dosya listesinde gizli, ignored ve sistem dosyalarini gosterir. false
  gunluk kullanim icin daha temizdir; gerekirse -a ile acilir.

showLineNumbers:
  TUI pager'da satir numaralarini gosterir. Proje dokumani okurken referans
  vermeyi kolaylastirir.

preserveNewLines:
  Markdown icindeki yazilmis satir sonlarini korur. Teknik dokumanlarda
  yazarın satir duzenini saklamak icin yararlidir.

kittyImages:
  Kitty terminalde markdown image satirlarini kitten icat ile inline gostermeyi
  dener. Terminal Kitty degilse veya kaynak kod render ediliyorsa normal
  markdown render'a geri doner.
` + "```" + `
`

// Generate the value we're doing to filter against.
func (m *markdown) buildFilterValue() {
	note, err := normalize(m.Note)
	if err != nil {
		log.Error("error normalizing", "note", m.Note, "error", err)
		m.filterValue = m.Note
	}

	m.filterValue = note
}

func (m markdown) relativeTime() string {
	return relativeTime(m.Modtime)
}

// Normalize text to aid in the filtering process. In particular, we remove
// diacritics, "ö" becomes "o". Note that Mn is the unicode key for nonspacing
// marks.
func normalize(in string) (string, error) {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	out, _, err := transform.String(t, in)
	if err != nil {
		return "", fmt.Errorf("error normalizing: %w", err)
	}
	return out, nil
}

// Return the time in a human-readable format relative to the current time.
func relativeTime(then time.Time) string {
	now := time.Now()
	if ago := now.Sub(then); ago < time.Minute {
		return "just now"
	} else if ago < humanize.Week {
		return humanize.CustomRelTime(then, now, "ago", "from now", magnitudes)
	}
	return then.Format("02 Jan 2006 15:04 MST")
}

// Magnitudes for relative time.
var magnitudes = []humanize.RelTimeMagnitude{
	{D: time.Second, Format: "now", DivBy: time.Second},
	{D: 2 * time.Second, Format: "1 second %s", DivBy: 1},
	{D: time.Minute, Format: "%d seconds %s", DivBy: time.Second},
	{D: 2 * time.Minute, Format: "1 minute %s", DivBy: 1},
	{D: time.Hour, Format: "%d minutes %s", DivBy: time.Minute},
	{D: 2 * time.Hour, Format: "1 hour %s", DivBy: 1},
	{D: humanize.Day, Format: "%d hours %s", DivBy: time.Hour},
	{D: 2 * humanize.Day, Format: "1 day %s", DivBy: 1},
	{D: humanize.Week, Format: "%d days %s", DivBy: humanize.Day},
	{D: 2 * humanize.Week, Format: "1 week %s", DivBy: 1},
	{D: humanize.Month, Format: "%d weeks %s", DivBy: humanize.Week},
	{D: 2 * humanize.Month, Format: "1 month %s", DivBy: 1},
	{D: humanize.Year, Format: "%d months %s", DivBy: humanize.Month},
	{D: 18 * humanize.Month, Format: "1 year %s", DivBy: 1},
	{D: 2 * humanize.Year, Format: "2 years %s", DivBy: 1},
	{D: humanize.LongTime, Format: "%d years %s", DivBy: humanize.Year},
	{D: math.MaxInt64, Format: "a long while %s", DivBy: 1},
}
