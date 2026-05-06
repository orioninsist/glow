# Glow Proje Analizi

Glow, terminal tabanlı bir Markdown okuyucusudur. Markdown dosyalarını terminal üzerinde "pizzazz" (canlılık/gösteriş) ile görüntülemek için tasarlanmıştır.

## Glow'un Temel Amacı (Olayı)
Glow, terminal kullanıcılarının Markdown dosyalarını (README, dökümantasyon vb.) sadece düz metin olarak değil, zengin bir görsel formatta okumasını sağlar. Modern terminal özelliklerini (renkler, grafikler, emoji desteği) kullanarak terminali daha güçlü bir döküman okuma platformuna dönüştürür.

## En Güzel Görünüm İçin İpuçları
Glow'da en iyi görünümü elde etmek için şu yöntemleri kullanabilirsin:

1.  **TUI Modu:** Terminalde sadece `glow` yazıp Enter'a basarsan, interaktif bir arayüz açılır. Burada dosyaları gezinebilir, arama yapabilir ve daha düzenli bir görünüm elde edebilirsin.
2.  **Stil Seçimi:** Terminal temanıza göre `-s` parametresini kullanabilirsin:
    *   `glow -s dark dosya.md` (Koyu tema için)
    *   `glow -s light dosya.md` (Açık tema için)
3.  **Özel Stiller:** Kendi JSON stil dosyalarını oluşturarak renkleri ve fontları tamamen kendine göre özelleştirebilirsin.

## Önemli Özellikler
*   **Uzaktan Okuma:** Doğrudan GitHub veya GitLab linklerini vererek dökümanları terminalde okuyabilirsin. Örneğin: `glow github.com/charmbracelet/glow`
*   **Arama:** Yerel dizindeki veya Git deposundaki tüm Markdown dosyalarını otomatik olarak bulur.
*   **Pager Desteği:** Uzun dosyaları `less` benzeri bir sistemle kolayca kaydırarak okumanı sağlar.
*   **Favoriler:** Sık okuduğun dökümanları işaretleyebilirsin.

## Ekstra Bir Şey Gerekli mi?
Hayır, Glow'u sistemine kurman yeterlidir. Go diliyle yazıldığı için oldukça hızlıdır ve ek bir bağımlılığa ihtiyaç duymaz.

---

## Özel JSON Stili (Göz Yormayan Tema)

Glow'u daha profesyonel ve göz yormayan (eye-friendly) bir hale getirmek için kendi JSON stil dosyanı oluşturabilirsin. Genellikle yazılımcıların tercih ettiği **Rosé Pine** veya **Catppuccin** benzeri pastel tonlar uzun süreli okumalarda gözü korur.

### Örnek: `goz-yormayan.json`
Aşağıdaki içeriği bir dosyaya kaydedip `glow -s goz-yormayan.json dosya.md` şeklinde kullanabilirsin:

```json
{
  "document": {
    "margin": 2
  },
  "h1": {
    "color": "#ebbcba",
    "bold": true
  },
  "h2": {
    "color": "#9ccfd8"
  },
  "text": {
    "color": "#e0def4"
  },
  "link": {
    "color": "#c4a7e7",
    "underline": true
  },
  "code": {
    "background_color": "#26233a",
    "color": "#ebbcba"
  },
  "code_block": {
    "margin": 2
  }
}
```

### JSON Kullanmanın Farkı Nedir?
*   **Tam Kontrol:** Başlıkların renginden, kod bloklarının arka planına kadar her şeyi terminalinden bağımsız ayarlarsın.
*   **Standardizasyon:** Ekibindeki herkes aynı JSON'u kullanırsa dökümanlar herkeste aynı görünür.
*   **Görsel Konfor:** Varsayılan "dark" veya "light" temalar bazen çok kontrastlı olabilir; JSON ile bu kontrastı yumuşatabilirsin.

## Projelerde ve VHS Demolarda Kullanım

### Tüm Projelerde Çalışır mı?
Evet, bu dosyayı bir kez hazırladıktan sonra bilgisayarındaki tüm Markdown projelerinde sorunsuz kullanabilirsin. `glow` sadece bir "görüntüleyici"dir, projendeki kodları veya dosyaları etkilemez.

### VHS Terminal Demo ve Glow Kullanımı
Eğer bir **VHS terminal kaydı** yapıyorsan (terminal demosu hazırlıyorsan), Glow'u ve bu özel JSON'u kullanmak harika bir fikir!

*   **Değişiklik Gerekli mi?** Hayır, VHS `.tape` dosyanın içinde sadece `Type "glow -s goz-yormayan.json README.md"` komutunu vermen yeterli olur.
*   **Demo Kalitesi:** Profesyonel bir dökümantasyon videosunda düz siyah-beyaz bir çıktı yerine, senin belirlediğin kurumsal veya estetik renklerin olması videonun kalitesini çok artırır.
*   **Glow Yapılandırması:** Eğer bu stili her zaman kullanmak istersen, `glow config` komutuyla açılan yapılandırma dosyasına `style: "/yol/to/goz-yormayan.json"` satırını ekleyebilirsin. Böylece her seferinde parametre yazmana gerek kalmaz.

---

## Yapılandırma ve Dosya Konumları (Linux/Pacman)

Glow'u pacman ile kurduğun için sistemindeki standart yollar şunlardır:

### 1. JSON Stil Dosyasını Nereye Koymalıyım?
JSON dosyasını istediğin herhangi bir klasöre koyabilirsin. Ancak düzenli olması için şu klasörü tercih etmeni öneririm:
*   `~/.config/glow/goz-yormayan.json`

### 2. Global Olarak Nasıl Aktif Edilir?
Glow'un her zaman bu stili kullanması için yapılandırma dosyasını (config) düzenlemelisin:
1.  Terminalde `glow config` komutunu çalıştır.
2.  Açılan dökümana şu satırı ekle (yolu kendi kullanıcına göre düzenle):
    ```yaml
    style: "/home/murat/.config/glow/goz-yormayan.json"
    ```
3.  Kaydedip çıktığında artık her `glow` komutu bu stili referans alacaktır.

### 3. Stil Dosyasını Silersem Ne Olur?
Eğer yapılandırma dosyasında (config) bir yol belirttiysen ve o dosyayı silersen, Glow o dosyayı bulamadığına dair bir hata mesajı verebilir. Bu durumda `glow config` komutuyla tekrar yapılandırma dosyasını açıp `style` değerini varsayılana döndürmen gerekir.

### 4. Varsayılan (Default) Değer Nedir?
Glow'un varsayılan stil değeri `auto`'dur. 
*   **Auto Modu:** Terminalinin arka plan rengini algılar. Eğer arka plan koyuysa `dark`, açıksa `light` temasını otomatik olarak seçer.
*   **Diğer Yerleşik Stiller:** `dark`, `light`, `pink`, `dracula`. Bunları JSON dosyası olmadan da `-s dracula` şeklinde kullanabilirsin.

---
*Bu dosya Antigravity tarafından Türkçe olarak güncellenmiştir.*
