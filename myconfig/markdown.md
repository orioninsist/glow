# Glow Proje Analizi, Build ve Markdown Link Komutlari

Bu dosya bu checkout icin pratik kullanim notudur. Glow artik bu dizinden build edilecek:

```bash
cd /mnt/samsung/orion-backup-local/projects/glow
```

## Kisa Proje Analizi

Bu proje Go ile yazilmis `github.com/charmbracelet/glow/v2` CLI/TUI markdown okuyucusu.

Onemli dosyalar:

```text
main.go              CLI/TUI girisi, argumanlar, URL/GitHub/GitLab/local dosya okuma, render akisi
config_cmd.go        glow config komutu ve varsayilan config olusturma
github.go            GitHub repo/README linklerini raw markdown kaynagina cevirme
gitlab.go            GitLab repo/README linklerini raw markdown kaynagina cevirme
ui/                  Bubble Tea tabanli TUI ve pager
utils/utils.go       Glamour style secimi, frontmatter temizleme, code block yardimcilari
glow.yml             Bu checkout icin kullanilacak tek config
myconfig/goz-yormayan.json  Terminalde rahat okunan ozel Glamour stili
```

Config yukleme davranisi:

```text
GLOW_CONFIG_HOME varsa once orasi okunur.
Okunacak dosya adi glow.yml olmalidir.
Bu projede tek config dosyasi kokte tutulur:
/mnt/samsung/orion-backup-local/projects/glow/glow.yml
```

## Once Paket Yoneticisi Kurulumunu Kaldir

Ayni isimli sistem Glow'u karistirmamak icin once hangi binary calisiyor bak:

```bash
which -a glow
glow --version
```

Kurulum sekline gore sadece sende olan komutu calistir:

```bash
# Bu makinede pacman/paru paketindeki glow kaldirildi:
sudo pacman -Rns glow
```

## Buradan Build Et

```bash
cd /mnt/samsung/orion-backup-local/projects/glow
go test ./...
go build -o glow .
```

Build edilen binary:

```text
/mnt/samsung/orion-backup-local/projects/glow/glow
```

## Config Dosyasinin Yeri

Olusturulan ana config:

```text
/mnt/samsung/orion-backup-local/projects/glow/glow.yml
```

Bu config su stil dosyasini kullanir:

```text
/mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json
```

Her komutta bu config'i kullandirmak icin:

```bash
export GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow
```

Kalici yapmak istersen `~/.zshrc` icine ekle:

```bash
export GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow
export PAGER="less -R"
export PATH="/mnt/samsung/orion-backup-local/projects/glow:$PATH"
```

## 9:16 1080x1920 YouTube Shorts Icin En Iyi Genislik

9:16 dikey video icin onerilen ana deger:

```bash
glow -p -w 72 "GITHUB_MARKDOWN_LINKI"
```

Neden `72`?

```text
1080x1920 dikey kayitta terminal genisligi sinirlidir.
-w 100 normal masaustu okuma icin guzel olsa da Shorts kaydinda satirlar uzun,
yazi daha kucuk ve goz takibi daha yorucu olabilir.

-w 72 dikey videoda daha okunakli satir uzunlugu verir.
Basliklar, listeler ve paragraflar ekranda daha rahat taranir.
Telefon ekraninda izleyen kisi icin metin daha buyuk ve daha sakin gorunur.
```

Pratik secim:

```text
YouTube Shorts / 9:16 / 1080x1920:  -w 72
Daha buyuk font veya cok yakin crop:       -w 64
Kod bloklari cok genisse:                  -w 80
Normal masaustu terminal okuma:            -w 100
```

Bu proje icin Shorts kaydinda en guzel genel komut:

```bash
GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow glow -p -w 72 "GITHUB_MARKDOWN_LINKI"
```

Eger binary'yi direkt repo icinden calistiriyorsan:

```bash
GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow ./glow -p -w 72 "GITHUB_MARKDOWN_LINKI"
```

## GitHub Markdown Linki Icin En Iyi Komut

En guzel Shorts gorunumu icin secilen ayar: repo kokundeki tek `glow.yml`, ozel `goz-yormayan.json` stil dosyasi, `less -R` pager ve `72` kolon genislik.

Normal masaustu okuma icin `-w 100` hala iyi bir degerdir. Dikey YouTube Shorts kaydi icin `-w 72` daha okunakli ve daha az yorucudur.

`GITHUB_MARKDOWN_LINKI` sadece yer tutucudur; tirnaklarin icine gercek GitHub repo, GitHub `blob/.../*.md` ya da raw markdown linkini koy.

~~~bash
glow -p -w 72 "GITHUB_MARKDOWN_LINKI"
~~~

Pager acilinca en basta baslik gorunur; asagi inmek icin `Space`, `j`, `Down`, cikmak icin `q` kullan.

Senin README linkin icin:

~~~bash
glow -p -w 72 "https://raw.githubusercontent.com/orioninsist/orioninsist.github.io/refs/heads/main/README.md"
~~~

TUI gorunumuyle acmak istersen:

~~~bash
glow -t -w 72 "GITHUB_MARKDOWN_LINKI"
~~~

Senin README linkini TUI ile acmak icin:

~~~bash
glow -t -w 72 "https://raw.githubusercontent.com/orioninsist/orioninsist.github.io/refs/heads/main/README.md"
~~~

## goz-yormayan.json Stil Analizi

`myconfig/goz-yormayan.json` bu proje icin ozel Glamour stilidir. Amaci parlak,
sert ve neon renklerden kacip uzun markdown okurken gozu yormayan bir terminal
gorunumu vermektir.

Stilin karakteri:

```text
Ana metin:        #d6d3c7  kirik sicak beyaz; saf beyaz kadar sert degil
H1 baslik:        #f2e6b8  koyu yesil/mavi arka plan ustunde sicak baslik
H2 baslik:        #b8d98f  yumusak yesil
H3 baslik:        #d7b98e  yumusak amber
Link:             #8fbfbd  sakin camgobegi, underline ile seciliyor
Kod:              #f0b58f  sicak turuncu, koyu arka planla ayriliyor
Kod blogu bg:     #242824  cok koyu yesilimsi gri
Yorumlar:         #7d867f  geri planda kalan gri-yesil
```

Analiz sonucu:

```text
Bu JSON genel olarak goz yormayan bir stil.
YouTube Shorts icin degistirmekten cok -w degerini 72 yapmak daha onemli.
Renkler zaten parlak beyaz, neon mor veya sert mavi gibi yorucu tonlara dayanmiyor.
```

Eger videoda kontrast az kalirsa JSON'u degistirmek yerine once terminal fontunu
biraz buyut ve `-w 64` dene. Renk degistirme ancak kayitta metin soluk gorunurse
gerekir.

## -s Yazinca ve Yazmayinca Ne Olur?

Senin kullanmak istedigin komut:

```bash
glow -p -w 72 "URL"
```

Bu komutta `-s` yok. O zaman Glow stil ayarini config'ten alir. Yani su sartla
otomatik olarak `goz-yormayan.json` secilir:

```bash
export GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow
```

veya komutta config dosyasini acikca verirsen:

```bash
glow --config /mnt/samsung/orion-backup-local/projects/glow/glow.yml -p -w 72 "URL"
```

Bu sart yoksa Glow repo kokundeki `glow.yml` dosyasini kendiliginden okumaz.
Sistem config dizinindeki ayari veya varsayilan `auto` stilini kullanir.

Yani:

```text
GLOW_CONFIG_HOME dogruysa:
  glow -p -w 72 "URL"
  -> goz-yormayan.json otomatik kullanilir.

GLOW_CONFIG_HOME yoksa:
  glow -p -w 72 "URL"
  -> goz-yormayan.json garanti degil; sistem config veya auto kullanilir.
```

Eger sunu yazarsan:

```bash
glow -p -s dark -w 80 "URL"
```

`-s dark` komut satiri flag'i oldugu icin config'teki JSON stilini ezer. Bu
durumda `goz-yormayan.json` kullanilmaz; Glow'un hazir `dark` stili kullanilir.

Kisa karar:

```text
En iyi Shorts gorunumu:
  GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow glow -p -w 72 "URL"

JSON stilini garanti etmek istersen:
  glow -p -s /mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json -w 72 "URL"

Hazir dark stilini denemek istersen:
  glow -p -s dark -w 80 "URL"
  Bu daha standart dark gorunur, ama bu projedeki ozel goz-yormayan stil devre disi kalir.
```

Benim onerim: YouTube Shorts icin `goz-yormayan.json + -w 72`. Eger kayitta
yazi hala kucuk gorunurse `-w 64`; eger kod satirlari cok kiriliyorsa `-w 80`.

## Goz Yormayan Stil Ne Zaman Aktif Olur?

Glow'da stil secim onceligi basit:

```text
1. Komutta -s / --style yazildiysa en once o kullanilir.
2. -s yoksa config dosyasindaki style kullanilir.
3. Config de yoksa varsayilan auto stil kullanilir.
```

Bu yuzden su komut:

```bash
glow -p -w 72 "URL"
```

tek basina `goz-yormayan.json` dosyasini garanti etmez. Bu komutta `-s` yoktur,
bu yuzden Glow config'e bakar. Goz yormayan stilin otomatik secilmesi icin
Glow'un bu repo kokundeki `glow.yml` dosyasini okuyor olmasi gerekir.

Aktif olur:

```bash
GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow glow -p -w 72 "URL"
```

Bu durumda Glow su dosyayi okur:

```text
/mnt/samsung/orion-backup-local/projects/glow/glow.yml
```

O dosyanin icinde de su style ayari vardir:

```yaml
style: "/mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json"
```

Yani sonuc:

```text
glow -p -w 72 "URL"
-> GLOW_CONFIG_HOME dogruysa
-> glow.yml okunur
-> style goz-yormayan.json olur
-> goz yormayan ozel gorunum aktif olur
```

Config'i komutta acikca verirsen de aktif olur:

```bash
glow --config /mnt/samsung/orion-backup-local/projects/glow/glow.yml -p -w 72 "URL"
```

JSON stilini hic config'e bagli kalmadan garanti etmek istersen en kesin komut:

```bash
glow -p -s /mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json -w 72 "URL"
```

Aktif olmaz veya garanti olmaz:

```bash
glow -p -w 72 "URL"
```

Eger `GLOW_CONFIG_HOME` ayarli degilse ve sistem config'inde bu JSON yazmiyorsa
Glow repo kokundeki `glow.yml` dosyasini kendiliginden bulmaz. Bu durumda
`goz-yormayan.json` aktif olmayabilir; Glow sistem config'ini veya varsayilan
`auto` stilini kullanir.

Aktif olmaz:

```bash
glow -p -s dark -w 80 "URL"
```

Cunku `-s dark` komut satirinda acikca stil secmektir. Komut satiri stili,
config dosyasindaki `style: ".../goz-yormayan.json"` ayarini ezer.

Bu komutun sonucu:

```text
glow -p -s dark -w 80 "URL"
-> -s dark var
-> config'teki goz-yormayan.json yok sayilir
-> Glow'un hazir dark stili kullanilir
-> ozel goz yormayan JSON aktif olmaz
```

Kisa ezber:

```text
Goz yormayan otomatik:
  GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow glow -p -w 72 "URL"

Goz yormayan garanti:
  glow -p -s /mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json -w 72 "URL"

Goz yormayan kapali, hazir dark aktif:
  glow -p -s dark -w 80 "URL"
```

Benim onerilen Shorts komutum:

```bash
GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow glow -p -w 72 "URL"
```

Tek seferlik, en garanti Shorts komutum:

```bash
glow -p -s /mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json -w 72 "URL"
```

Tek kelimeyle kullanim:

```text
Rahat
```

Not:

```text
Goz yormayan stil aktifse ve -w 72 kullaniliyorsa Shorts icin kullanim rahattir.
Config ayari unutulursa veya komutta -s dark yazilirsa bu rahat gorunum garanti olmaz.
En temiz gunluk kullanim: GLOW_CONFIG_HOME ayarli olsun, komutta sadece glow -p -w 72 "URL" yaz.
En garanti tek seferlik kullanim: -s ile goz-yormayan.json yolunu acikca yaz.
```

## Uzun JSON Path Yazmadan Alias/Fonksiyon

En temiz cozum: `glow` komutunu zsh fonksiyonu yap. Boylece her normal
`glow` kullaniminda bu repo kokundeki `glow.yml` otomatik okunur ve
`goz-yormayan.json` aktif olur.

`~/.zshrc` icine ekle:

```bash
glow() {
  GLOW_CONFIG_HOME=/mnt/samsung/orion-backup-local/projects/glow /mnt/samsung/orion-backup-local/projects/glow/glow "$@"
}
```

Durum: bu fonksiyon `~/.zshrc` icine eklendi ve yeni zsh oturumunda aktif.

Sonra terminali yenile:

```bash
source ~/.zshrc
```

Artik normal kullanim:

```bash
glow -p -w 72 "URL"
```

Bu komutta uzun JSON path yazmazsin. Fonksiyon arkada `GLOW_CONFIG_HOME`
verdigi icin Glow su config'i okur:

```text
/mnt/samsung/orion-backup-local/projects/glow/glow.yml
```

O config icinde su ayar oldugu icin goz yormayan stil aktif olur:

```yaml
style: "/mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json"
```

Mantik:

```text
glow -p -w 72 "URL"
-> zsh fonksiyonu calisir
-> GLOW_CONFIG_HOME repo kokune ayarlanir
-> glow.yml okunur
-> goz-yormayan.json aktif olur
```

Ama sen bilerek `-s dark` yazarsan:

```bash
glow -p -s dark -w 80 "URL"
```

Bu sefer `-s dark` komut satiri flag'i config'teki JSON stilini ezer.
Yani mantiken dark aktif olur:

```text
glow -p -s dark -w 80 "URL"
-> zsh fonksiyonu yine glow.yml dosyasini buldurur
-> ama komutta -s dark var
-> -s dark config style ayarini ezer
-> goz-yormayan.json degil, hazir dark stil aktif olur
```

Kisa karar:

```text
Varsayilan goz yormayan:
  glow -p -w 72 "URL"

Bilerek dark:
  glow -p -s dark -w 80 "URL"

Bilerek light:
  glow -p -s light -w 80 "URL"
```

Buna alias yerine fonksiyon dememizin sebebi: `"$@"` ile yazdigin butun flag ve
URL'ler aynen korunur. Boylece normalde goz yormayan stil otomatik gelir, ama
`-s dark` gibi bilincli secimler yine calisir.

## Resimli Markdown

Normal `glow -p` ve `glow -t` markdown icindeki resmi gercek bitmap olarak gostermez; resmin alt metnini/linkini markdown gibi gosterir.

Gercek inline resim gostermek icin Kitty terminal gerekir. Bu forkta `--kitty-images` hem lokal resimleri hem de GitHub/raw markdown icindeki remote `http/https` resimleri indirip Kitty ile gostermeye calisir.

Kitty terminalde remote GitHub markdown icindeki resimleri gostermek icin:

~~~bash
glow --kitty-images -p -w 72 "GITHUB_MARKDOWN_LINKI"
~~~

Senin README linkin icin resimli deneme:

~~~bash
glow --kitty-images -p -w 72 "https://raw.githubusercontent.com/orioninsist/orioninsist.github.io/refs/heads/main/README.md"
~~~

GitHub `blob/.../*.md` linki de olur:

~~~bash
glow --kitty-images -p -w 72 "https://github.com/charmbracelet/glow/blob/master/README.md"
~~~

Lokal markdown icindeki lokal resimleri gostermek icin:

~~~bash
glow --kitty-images -p -w 72 README.md
~~~

Not: `--kitty-images` sadece CLI output icindir. `-t` TUI icinde gercek bitmap resim gosterimi yoktur.

Ornekler:

```bash
# GitHub repo README
glow -p -w 72 "https://github.com/charmbracelet/glow"

# Raw markdown dosyasi
glow -p -w 72 "https://raw.githubusercontent.com/charmbracelet/glow/master/README.md"

# GitHub blob markdown dosyasi
glow -p -w 72 "https://github.com/charmbracelet/glow/blob/master/README.md"

# Lokal markdown dosyasi
glow -p -w 72 README.md
```

## Komut Varyasyonlari

Normal en iyi okuma:

```bash
glow -p -w 72 "URL"
```

Pager olmadan terminale direkt bas:

```bash
glow -w 72 "URL"
```

TUI ile ac:

```bash
glow -t -w 72 "URL"
```

Repo README otomatik bul:

```bash
glow -p -w 72 "https://github.com/KULLANICI/REPO"
```

GitHub web markdown linki:

```bash
glow -p -w 72 "https://github.com/KULLANICI/REPO/blob/main/DOSYA.md"
```

Raw markdown linki:

```bash
glow -p -w 72 "https://raw.githubusercontent.com/KULLANICI/REPO/refs/heads/main/DOSYA.md"
```

Stdin'den oku:

```bash
curl -L "URL" | glow -p -w 72 -
```

Daha dar yazi alani:

```bash
glow -p -w 64 "URL"
```

Genis kod bloklari icin:

```bash
glow -p -w 120 "URL"
```

Stil override etmek icin:

```bash
glow -s /mnt/samsung/orion-backup-local/projects/glow/myconfig/goz-yormayan.json -p -w 72 "URL"
```

Tum gizli/ignored dosyalarla TUI:

```bash
glow -t -a -w 72 .
```

Lokal dizin TUI:

```bash
glow -t .
```

Lokal dosya TUI:

```bash
glow -t README.md
```

## Notlar

`-p` daha guzel okuma icin `$PAGER` kullanir. Bu makinede `~/.zshrc` icinde `PAGER="less -R"` ayarlandi.

`-w 72` YouTube Shorts/9:16 kayit icin dengeli genisliktir. Normal masaustu okuma icin `-w 100`, genis kod bloklari icin `-w 120` kullanilabilir.

`GLOW_CONFIG_HOME` verilmezse Glow sistem config dizinine bakar; repo kokundeki `glow.yml` otomatik okunmaz.

`-p` ve `-t` ayni komutta beraber kullanilmaz.
