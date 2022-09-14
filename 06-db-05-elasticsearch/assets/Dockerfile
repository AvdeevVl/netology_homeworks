FROM centos:7

EXPOSE 9200 9300

ENV ES_VER=8.4.1 

RUN cd /opt && \
    groupadd elasticsearch && \
    useradd -c "elasticsearch" -g elasticsearch elasticsearch &&\
    yum -y install wget && \
    wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-${ES_VER}-linux-x86_64.tar.gz && \
    wget https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-${ES_VER}-linux-x86_64.tar.gz.sha512 && \
    sha512sum -c elasticsearch-${ES_VER}-linux-x86_64.tar.gz.sha512 && \
    tar -xzf elasticsearch-${ES_VER}-linux-x86_64.tar.gz && \
        rm -f elasticsearch-${ES_VER}-linux-x86_64.tar.gz* && \
        mkdir /var/lib/data && chmod -R 777 /var/lib/data && \
	    chown -R elasticsearch:elasticsearch /opt/elasticsearch-${ES_VER} && \
    yum -y remove wget && \
    yum clean all

USER elasticsearch
WORKDIR /opt/elasticsearch-${ES_VER}/
COPY elasticsearch.yml  config/
ENTRYPOINT ["bin/elasticsearch"]