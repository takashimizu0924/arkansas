# ベースイメージ
ARG VERSION="latest"
FROM debian:$VERSION

# ベースイメージで実行するコマンド
# --no-install-recommendsコマンドで必要最小限のインストールを一度に複数行う
# 最後に不要なパッケージマネージャを削除(rm -rf /var/lib/apt/lists/*)する
RUN apt-get update && apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
    apt-utils \
    vim \
    git \
    curl \
    libpq-dev \
    python3 \
    python3-setuptools \
    build-essential python \
    python3-dev \
    python3-pip \
    wget && \
    apt-get -y clean && \
    rm -rf /var/lib/apt/lists/*

# TA-Libでpipを使用するため先にpipのアップグレードとpythonパッケージのインストール
RUN pip3 install --upgrade pip
RUN pip3 install matplotlib numpy pandas scipy sklearn

# TA-Libのソースファイルをダウンロード(Ver.0.4.0)して展開後にmakeしてインストールする
RUN wget http://prdownloads.sourceforge.net/ta-lib/ta-lib-0.4.0-src.tar.gz && \
    tar xvzf ta-lib-0.4.0-src.tar.gz && \
    cd ta-lib && \
    ./configure --prefix=/usr && \
    make && \
    make install && \
    bash -c "echo "/usr/local/lib64" >> /etc/ld.so.conf" && \
    /sbin/ldconfig && \
    /usr/local/bin/pip3 install ta-lib

# Workディレクトリの作成
RUN mkdir -p /app/scripts /app/log
WORKDIR /app

# コンテナの実行時に実行されるコマンドの指定
# CMD ["python3", "app/launch.py"]