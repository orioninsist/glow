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
*Bu dosya Antigravity tarafından Türkçe olarak oluşturulmuştur.*
