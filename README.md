Hello Server
------------

Sadece GET /hello isteklerine JSON formatında Merhaba Dünya
cevabını gönderir, ayrıca istekleri loglar.

Şu an için sadece 64 bit linux makinelerde çalışır. Docker, docker-compose, golang yüklü olması gerekiyor. Gerekli çalışma ortamının otomatize edilmesi daha sonraya kaldı.

###Kullanımı
Yukarıda belirtilen araçlar yüklendikten sonra proje dizininde `./start.sh` ile uygulama başlatılıyor. `./stop.sh` ile kapatılıyor.

Kibana için: http://localhost:5601
Elasticsearch için: http://localhost:9200
Weave Scope için: http://localhost:4040
Hello uygulama sunucusu için: http://localhost:8080/hello




