FROM kibana

ENV KIBANA_HOME=/opt/kibana

RUN kibana plugin --install elastic/sense
RUN kibana plugin --install elasticsearch/marvel
RUN kibana plugin -i tagcloud -u https://github.com/stormpython/tagcloud/archive/master.zip
RUN kibana plugin -i kibi_wordcloud_vis -u https://github.com/sirensolutions/kibi_wordcloud_vis/raw/4.5.3/target/kibi_wordcloud_vis-4.5.3.zip



