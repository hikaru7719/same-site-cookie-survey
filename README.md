# same-site-cookie-survey

# setup

1. リポジトリの clone

```
git clone https://github.com/hikaru7719/same-site-cookie-survey
```

2. mkcert のインストール

```
brew install mkcert
```

3. mkcert のセットアップ

```
mkcert -install
```

4. key file, cert file の作成

```
mkcert "*.test.local"
```

5. /etc/hosts ファイルの更新

```
127.0.0.1 api.test.local
127.0.0.1 frontend.test.local
```

6. アプリケーションの起動

```
docker-compose up
```
