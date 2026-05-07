# Glow Proje Analizi

Glow, terminalde Markdown görüntülemek için yazılmış bir CLI ve TUI uygulamasıdır. Temel amacı, Markdown dosyalarını düz metin olarak göstermek yerine stil, satır kırma, pager ve etkileşimli gezinme ile okunur hale getirmektir.

## Proje Ne Yapar?

Glow şu işleri yapar:

* Yerel Markdown dosyalarını render eder.
* Dizin içinde Markdown dosyaları bulur ve TUI içinde listeler.
* GitHub ve GitLab depo README'lerini otomatik bulup okur.
* HTTP ve HTTPS üzerinden gelen Markdown içeriklerini işler.
* stdin üzerinden gelen içeriği okuyabilir.
* İsteğe bağlı olarak çıktıyı pager içinde gösterir.
* TUI modunda dosya tarama, filtreleme, açma, yeniden yükleme ve dış editöre gönderme sağlar.

## Verilen Komutun Doğru Yazımı

Sorulan komutun doğru biçimi şöyledir:

```bash
glow -p -s dark -w 0 demo.md
```

`-p-s dark` yazımı doğru değildir. `-p`, `-s` ve `-w` ayrı bayraklardır; yan yana yazılırken tek bir birleşik bayrak gibi davranmazlar.

Bu komut `demo.md` dosyasını okur, `dark` stilini uygular, genişliği `0` olarak bırakır ve çıktıyı pager ile gösterir.

## Ornek Komutlarin Ayrintili Analizi

Kullanicinin verdigi ornekler ayni aileden gelir, fakat mod secimi ve gorunum davranisi farklidir:

```bash
glow -t abc.md
```

Bu komut `abc.md` dosyasini TUI icinde acar. `-t` veya `--tui`, Glow'un interaktif terminal arayuzunu kullanmasini ister. Dosya verildigi icin Glow once dosyayi okur, sonra bu icerigi TUI pager gorunumunde gosterir. Dosya yerel bir dosyaysa TUI icinden yeniden yukleme ve editorle acma gibi isler daha anlamli olur.

```bash
glow -p -s dark -w 100 orioninsist-cli-2.md
```

Bu komut `orioninsist-cli-2.md` dosyasini klasik CLI render modunda isler, sonucu `dark` stil ile bicimlendirir, satirlari en fazla `100` kolon civarinda sarar ve sonucu pager icinde acar. `-p` dis pager kullanir; sistemde `PAGER` ortam degiskeni varsa onu, yoksa `less -r` komutunu tercih eder.

Aralarindaki temel fark sudur:

| Komut | Mod | Stil | Genislik | Okuma sekli | Ne zaman iyi? |
|-------|-----|------|----------|-------------|---------------|
| `glow -t abc.md` | TUI | `auto` veya config | Otomatik/config | Glow'un kendi interaktif arayuzu | Uzun dosya okuma, dosya icinde rahat gezinme |
| `glow -p -s dark -w 100 orioninsist-cli-2.md` | CLI + pager | `dark` | `100` | `less -r` veya `$PAGER` | Tek dosyayi sabit, kontrollu gorunumle okumak |

## Tum Ana Flag Varyasyonlari ve Farklari

Glow flag'lerinde sira onemli degildir. Yani su iki komut ayni davranir:

```bash
glow -p -s dark -w 100 file.md
glow file.md -w 100 -s dark -p
```

Fakat her flag'in gorevi ayridir:

| Flag | Uzun hali | Deger alir mi? | Mod | Etki |
|------|-----------|----------------|-----|------|
| `-p` | `--pager` | Hayir | CLI | Render sonucunu dis pager icinde acar |
| `-t` | `--tui` | Hayir | TUI | Glow'un interaktif arayuzunu acar |
| `-s` | `--style` | Evet | CLI/TUI | Stil secer: `auto`, `dark`, `light`, `pink`, `dracula`, `notty`, `tokyonight` veya JSON dosyasi |
| `-w` | `--width` | Evet | CLI/TUI | Satir sarma genisligini belirler |
| `-a` | `--all` | Hayir | TUI | Gizli/ignore edilen dosyalari da listeler |
| `-l` | `--line-numbers` | Hayir | TUI | TUI dokuman gorunumunde satir numarasi gosterir |
| `-n` | `--preserve-new-lines` | Hayir | CLI/TUI | Markdown render ederken yeni satirlari korur |
| `--config` | `--config` | Evet | Genel | Ozel config dosyasi kullanir |
| `-h` | `--help` | Hayir | Genel | Yardim metnini gosterir |
| `-v` | `--version` | Hayir | Genel | Surum bilgisini gosterir |
| `-m` | `--mouse` | Hayir | TUI | Mouse wheel destegini acar; help'te gizlidir |

Gecerli pratik varyasyonlar:

```bash
glow file.md
```

Dosyayi dogrudan terminale render eder. Hizli kontrol icin en basit kullanimdir.

```bash
glow -p file.md
```

Dosyayi pager icinde acar. Uzun dokumanlarda daha rahattir.

```bash
glow -t file.md
```

Dosyayi Glow TUI icinde acar. Interaktif okuma icin daha zengindir.

```bash
glow -s dark file.md
glow -s light file.md
glow -s auto file.md
glow -s goz-yormayan.json file.md
```

Stil secimini degistirir. `auto` terminal arka planina gore secim yapar. JSON dosyasi verilirse o dosyanin var olmasi gerekir.

```bash
glow -w 80 file.md
glow -w 100 file.md
glow -w 120 file.md
glow -w 0 file.md
```

Genisligi kontrol eder. `80` dar ve odakli, `100` dengeli, `120` genis ekran icin rahat, `0` ise acikca word wrap kapali demektir. Kod agirlikli dosyalarda `-w 0` bazen iyi olur; normal Markdown okumada genellikle fazla genis kacabilir.

```bash
glow -p -s dark -w 100 file.md
```

CLI + pager + koyu tema + kontrollu genislik kombinasyonudur. Tek dosya okumak icin en dengeli komutlardan biridir.

```bash
glow -t -s dark -w 100 -l file.md
```

TUI + koyu tema + 100 kolon + satir numarasi kombinasyonudur. Dokuman uzerinde gezinirken satir referansi gerekiyorsa guzel secimdir.

```bash
glow -t -a .
```

Mevcut dizini TUI ile acar ve normalde gizlenen/ignore edilen dosyalari da gosterir. Proje tarama icin gucludur, ama kalabalik repo'larda listeyi sisirebilir.

Gecersiz veya sorunlu varyasyonlar:

```bash
glow -t -p file.md
```

Gecersizdir. Kod `pager && tui` durumunda `cannot use both pager and tui` hatasi verir. Cunku `-t` Glow'un kendi TUI'sini, `-p` ise dis pager'i ister.

```bash
glow -s olmayan-stil file.md
```

Eger `olmayan-stil` adinda yerlesik stil yoksa Glow bunu JSON stil dosyasi yolu sanir. Dosya da yoksa hata verir.

```bash
glow -w abc file.md
```

Gecersizdir. `-w` sayisal `uint` deger bekler.

```bash
glow -p-s dark file.md
```

Dogru yazim degildir. `-p` ve `-s` ayri flag olarak yazilmalidir: `glow -p -s dark file.md`.

## En Iyi Gorunum Onerisi

Genel okuma icin en iyi ve dengeli gorunum:

```bash
glow -p -s dark -w 100 orioninsist-cli-2.md
```

Sebebi: `dark` tema uzun terminal okumalarinda gozu daha az yorar, `-w 100` satirlari cok daraltmadan okunur tutar, `-p` ise uzun dokumanda rahat yukari/asagi gezinme saglar.

Interaktif calismak, dosya taramak veya dokumani TUI icinde acip yeniden yuklemek istiyorsan en iyi secim:

```bash
glow -t -s dark -w 100 -l orioninsist-cli-2.md
```

Bu secim daha "uygulama gibi" hissettirir. Sadece okumak icin `-p` daha sade ve hizli; dosya uzerinde gezinmek ve proje icinde calismak icin `-t` daha kullanislidir.

## CLI Bayrakları (Tüm Requirement'ler)

Kodda tanımlı bayraklar şunlardır:

* `-p`, `--pager`:
  - **İşlev**: Çıktıyı pager içinde açar (CLI modu için)
  - **Varsayılan Pager**: `PAGER` tanımlı değilse `less -r` kullanılır
  - **Çakışma**: `-t` bayrağı ile birlikte kullanılamaz (TUI modunda pager desteklenmez)

* `-t`, `--tui`:
  - **İşlev**: CLI çıktısı yerine interaktif TUI açar
  - **Etki**: Dosya listesi, filtreleme, dış editör entegrasyonu sağlar
  - **Çakışma**: `-p` bayrağı ile birlikte kullanılamaz (TUI kendi arayüzü vardır)

* `-s`, `--style`:
  - **İşlev**: Stil adı veya JSON stil dosyası yolu belirtir
  - **Yerleşik Stiller**: `auto`, `dark`, `light`, `pink`, `dracula`, `notty`, `tokyonight` vb.
  - **Çoklu Kullanım**: `-s dark`, `-s light`, `-s /path/to/style.json` şeklinde kullanılabilir

* `-w`, `--width`:
  - **İşlev**: Word wrap genişliğini belirtir
  - **Varsayılan**: Bayrak verilmezse terminal genişliği otomatik algılanır (bulunamazsa 80 kolon)
  - **`-w 0` Anlamı**: Açıkça sarmayı devre dışı bırakır (sarma tamamen kapanır)

* `-a`, `--all`:
  - **İşlev**: TUI içinde gizli (nokta ile başlayan) ve ignore edilen dosyaları gösterir
  - **Etki**: Varsayılan olarak yok sayılan klasörler (node_modules, GOPATH vb.) da listeye eklenir

* `-l`, `--line-numbers`:
  - **İşlev**: TUI içinde satır numaralarını gösterir

* `-n`, `--preserve-new-lines`:
  - **İşlev**: Çıktıda yeni satır karakterlerini korur (varsayılan olarak normalize edilebilir)

* `-m`, `--mouse`:
  - **İşlev**: TUI için mouse wheel scroll desteğini açar
  - **Not**: Bu bayrak kullanıcı arayüzünde gizlidir

* `--config`:
  - **İşlev**: Kullanılacak yapılandırma dosyasını açıkça belirtir
  - **Desteklenen Format**: `.yml` veya `.yaml`

### Komut Analizi: `glow -t -p -s dark -w 0 demo.md`

**Sorun: Bu komut ÇAKIŞMADIR ve düzeltilmesi gerekir.**

| Bayrak | Değer | İşlev | Sorun |
|--------|-------|-------|-------|
| `-t` | - | TUI modunu aç | `-p` ile çakışır |
| `-p` | - | Pager aç | `-t` ile çakışır |
| `-s` | `dark` | Dark stil uygula | Sorun yok |
| `-w` | `0` | Word wrap kapat | Sorun yok |
| `demo.md` | - | Hedef dosya | Sorun yok |

**Çakışmanın Sebebi**:
- `-t` bayrağı TUI (Text User Interface) açar. TUI kendi arayüzüne sahiptir ve pager kullanmaz.
- `-p` bayrağı pager açar. Pager sadece CLI modunda çalışır.
- Bu ikisini aynı anda kullanmak imkansızdır.

**Düzeltilmiş Versiyonlar**:

1. **TUI Modu** (interaktif liste + arama):
   ```bash
   glow -t -s dark demo.md
   ```
   (Bayrak sırası önemli değildir: `glow -s dark -t demo.md` da aynıdır)

2. **CLI + Pager Modu** (doğrudan render + pager):
   ```bash
   glow -p -s dark -w 0 demo.md
   ```

3. **CLI Modu** (doğrudan terminal çıktısı):
   ```bash
   glow -s dark -w 0 demo.md
   ```

## Stil Sistemi

`-s` ile verilen değer yerleşik bir stil ise doğrudan kullanılır. Desteklenen yerleşik stiller arasında `auto`, `dark`, `light`, `pink`, `dracula`, `notty` ve `tokyonight` benzeri tanımlar bulunur. Değer yerleşik bir stil değilse Glow bunu JSON stil dosyası yolu olarak yorumlar ve dosyanın varlığını denetler.

`auto` kullanılırsa terminal arka planı algılanır ve koyuysa `dark`, açıksa `light` seçilir. Eğer çıktı terminal değilse ve kullanıcı `style` bayrağını elle vermediyse Glow `notty` stiline geçer.

`glow` ayrıca TUI tarafında `GLAMOUR_STYLE` ortam değişkenini de okuyabilir.

## Genişlik ve Render Davranışı

Render genişliği `-w` ile kontrol edilir. Kod akışı şöyledir:

* Bayrak verilmezse terminal genişliği algılanır.
* Terminal genişliği bulunamazsa varsayılan olarak `80` kolon kullanılır.
* TUI içinde pencere genişliği temel alınır, üst sınır olarak `120` kolon uygulanabilir.
* CLI tarafında `glamour.WithWordWrap(int(width))` ile render edilir.

Önemli ayrım: bayrak hiç verilmezse Glow kendi genişlik keşfini yapar; `-w 0` açıkça verilirse sarma kapatılır.

## Kaynak Türleri

Glow şu giriş türlerini destekler:

* `glow demo.md` ile yerel dosya.
* `glow -` ile stdin.
* `glow github.com/charmbracelet/glow` gibi GitHub deposu.
* `glow gitlab.com/owner/repo` gibi GitLab deposu.
* `glow https://host.tld/file.md` gibi HTTP/HTTPS URL'leri.
* Dizin verilirse uygun README dosyası aranır.

GitHub ve GitLab için hedef repo README'si otomatik bulunur. Kısa şema biçimleri `github://owner/repo` ve `gitlab://owner/repo` olarak da çözümlenir; custom hostname desteği yoktur.

## TUI ve CLI Ayrımı

Glow çalıştırıldığında davranış şu şekilde ayrılır:

* Argüman yoksa TUI açılır.
* Tek argüman bir dizinse TUI o dizinde açılır.
* Argüman dosya, URL veya `-` ise CLI modu çalışır.

CLI modunda içerik doğrudan render edilir. TUI modunda ise dosya listesi, arama, yardım ve doküman görüntüleme akışları aktif olur.

## TUI Özellikleri

TUI, sadece dosya gösteren bir arayüz değildir. Şunları yapar:

* Geçerli dizin ve alt dizinlerde Markdown dosyaları arar.
* Git deposu içindeyse depo içini tarar.
* Fuzzy filtreleme sağlar.
* Sekmeler arasında geçiş yapabilir.
* Dosyayı açabilir.
* Seçili Markdown dosyasını dış editörde açabilir.
* Açık dosyayı yeniden yükleyebilir.
* Dosya değişikliklerini izleyecek şekilde fsnotify kullanır.
* Yardım ekranı gösterir.

### TUI Kısayolları

Dosya listesi görünümünde:

* `j`, `k`, `↑`, `↓`: Seçim değiştirir.
* `home`, `g`: Listenin başına gider.
* `end`, `G`: Listenin sonuna gider.
* `tab`, `L`: Sonraki bölüme geçer.
* `shift+tab`, `H`: Önceki bölüme geçer.
* `enter`: Seçili dokümanı açar.
* `/`: Filtre modunu açar.
* `?`: Yardım görünümünü büyütür veya küçültür.
* `!`: Hata varsa hata görünümünü açar.
* `e`: Seçili dokümanı dış editörde açar.
* `F`: Dosya listesini yeniden tarar.
* `r`: Filtrelenmiş doküman görünümünde yardımcı değil; bu tuş pager tarafında yeniden yükleme içindir.
* `q`: Uygulamadan çıkar.
* `esc`: Filtre açıkken filtreyi iptal eder.

Filtre düzenleme modunda:

* `esc`: Filtreyi iptal eder.
* `enter`, `tab`, `shift+tab`, `ctrl+k`, `↑`, `ctrl+j`, `↓`: Filtreyi onaylar veya mevcut sonuçlara göre açma davranışı uygular.

## Pager Özellikleri

Doküman açıldıktan sonra pager görünümünde şu davranışlar vardır:

* `k`, `↑`: Yukarı.
* `j`, `↓`: Aşağı.
* `b`, `pgup`: Sayfa yukarı.
* `f`, `pgdn`: Sayfa aşağı.
* `u`: Yarım sayfa yukarı.
* `d`: Yarım sayfa aşağı.
* `g`, `home`: Üste git.
* `G`, `end`: Alta git.
* `c`: İçeriği kopyala.
* `e`: Bu dokümanı dış editörde aç.
* `r`: Bu dokümanı yeniden yükle.
* `esc`: Dosya listesinin bulunduğu TUI ekranına dön.
* `q`: Çık.
* `?`: Pager yardımını aç/kapat.

Pager açıkken dosya değişirse Glow fsnotify ile bunu fark edip yeniden yükleyebilir.

## Render Hattı

Render süreci kabaca şöyledir:

1. Kaynak okunur.
2. Frontmatter varsa çıkarılır.
3. Kaynak Markdown değilse kod bloğu gibi sarılır.
4. Glamour renderer oluşturulur.
5. Seçilen stil, genişlik ve yeni satır davranışı uygulanır.
6. Pager, terminal çıktı veya TUI akışı çalışır.

Markdown olmayan dosyalar düz Markdown gibi değil, kod bloğu gibi render edilir. Bu yüzden örneğin `.txt` veya uzantısız bir dosya farklı biçimde gösterilebilir.

## Yapılandırma Dosyası

`glow config` komutu yapılandırma dosyasını açar. Dosya yoksa oluşturur. Desteklenen uzantılar `.yml` ve `.yaml`'dır.

Varsayılan yapılandırma örneğinde şu alanlar vardır:

* `style`
* `mouse`
* `pager`
* `width`
* `all`

Konfigürasyon dosyası bulunamazsa Glow varsayılan dizinlerde `glow.yml` oluşturmaya çalışır.

### Yapılandırma İçin Ortam Değişkenleri

`ui.Config` içinde şu ortam değişkenleri okunur:

* `GOPATH`
* `HOME`
* `GLAMOUR_STYLE`
* `GLOW_HIGH_PERFORMANCE_PAGER`
* `GLOW_ENABLE_GLAMOUR`

## Yerel Dosya Tarama ve Ignore Kuralları

Glow TUI içinde dosya tararken bazı yolları yok sayar:

* `node_modules`
* gizli klasörler `.*`
* `GOPATH`
* macOS üzerinde ayrıca `Library`

Bu sayede gereksiz veya büyük dizinler listeyi şişirmez.

## Yardım, Man Sayfası ve Gizli Komutlar

Glow'da `man` adlı gizli bir alt komut vardır. Bu komut man sayfası üretir, ancak normal kullanıcı akışında görünmez.

`glow --help` komutu genel kullanım bilgisini verir. TUI içindeki `?` ise uygulama içi yardım ekranını açar.

## Ne Yapamaz?

Glow'un kapsamı sınırlıdır. Şunları yapmaz:

* Genel amaçlı bir dosya editörü değildir.
* Markdown dışı içerikleri semantik olarak dönüştürmez; gerekirse kod bloğu gibi gösterir.
* GitHub/GitLab dışındaki özel depo hostlarını otomatik README keşfi için desteklemez.
* Yerleşik olarak görsel düzenleme veya belge üretme aracı değildir.
* Dosyaları kendi içinde kalıcı olarak düzenlemez; edit işlemi dış editör açılarak yapılır.
* Favori işaretleme veya kitaplık yönetimi yapmaz.

## Pratik Kullanım Örnekleri

```bash
glow README.md
glow -p README.md
glow -s dark README.md
glow -w 60 README.md
glow github.com/charmbracelet/glow
glow https://host.tld/file.md
glow -
```

## Sonuç

Bu proje, Markdown okuma deneyimini terminale taşır. Gücü, tek başına çıktı üretmesinden değil; TUI tarama, pager, dış editör, fuzzy filtre, GitHub/GitLab README çözümleme, stil sistemi ve otomatik yeniden yükleme gibi parçaları birlikte sunmasından gelir.

Özetle: Glow bir Markdown görüntüleyicisidir, ama yalnızca görüntüleyici değil; aynı zamanda terminal içi bir Markdown keşif ve okuma aracıdır.
