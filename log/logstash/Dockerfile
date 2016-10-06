FROM logstash

# logstash yapılandırma dosyasını kopyala konteynır içine
COPY logstash.conf /etc/logstash/conf.d/

# entrypoint.sh dosyasını kopyala ve çalıştılabilir yap
COPY entrypoint.sh /
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

