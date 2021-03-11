# AWS EC2 の設定、 手順

```
# nginx, golangインストール
$ sudo amazon-linux-extras install nginx1.12
$ sudo amazon-linux-extras install golang1.11

# nginx起動
$ sudo systemctl start nginx.service

# nginx再起動
$ sudo nginx -s reload

# nginx自動起動
$ sudo systemctl enable nginx

# nginx設定ファイルを編集
$ sudo vi /etc/nginx/nginx.conf

    location / {
                proxy_pass http://127.0.0.1:8080;　接続先
    }

# MySQL必要パッケージのインストールと設定、 MySQLのリポジトリをyumに追加
$ sudo yum update
$ sudo yum remove mariadb-libs
$ sudo yum localinstall https://dev.mysql.com/get/mysql80-community-release-el7-3.noarch.rpm

# MySQLに必要なパッケージインストール
$ sudo yum install --enablerepo=mysql80-community mysql-community-server
$ sudo yum install --enablerepo=mysql80-community mysql-community-devel

# logファイルの生成
$ sudo touch /var/log/mysqld.log

# logファイルを開く
$ sudo less /var/log/mysqld.log

# mysqldがインスタンスの起動と同時に起動
$ sudo systemctl enable mysqld

# mysql設定
.env参照

# /varフォルダーにgit clone

# 実行
$ sudo go build main.go
$ ./main &
```
