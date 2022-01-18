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
    python3-pip && \
    apt-get -y clean && \
    rm -rf /var/lib/apt/lists/*

# 機械学習用のライブラリ・モジュールをインストールする
RUN pip3 install --upgrade pip
RUN pip3 install numpy pandas scipy sklearn matplotlib psycopg2

# Workディレクトリの作成
RUN mkdir -p /machine-learning/proc /machine-learning/log
WORKDIR /machine-learning/proc

# PostgreSQL接続用の環境変数設定
ENV DB_URL=postgresql://postgres:3d93792ce41f165941d34c70f7175ece@192.168.0.117:5432/streliztia_db

# コンテナの実行時に実行されるコマンドの指定(機械学習エンジンの起動)
CMD ["python3", "main/src/run.py"]