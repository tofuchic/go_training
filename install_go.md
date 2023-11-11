# Go インストール方法

## 手動インストール

- Go バージョンを確認してインストール
  - [Go](https://go.dev/dl/)

## 最新版の自動インストール

### Linux

```bash
latest_stable_go_version=$(curl -Lso- https://go.dev/dl/ | grep -m1 -P 'id="go\d+\.\d+\.\d+' | cut -f 4 -d '"')
wget https://dl.google.com/go/$latest_stable_go_version.linux-amd64.tar.gz
sudo tar -xvf $latest_stable_go_version.linux-amd64.tar.gz
rm -f $latest_stable_go_version.linux-amd64.tar.gz
sudo mv go /usr/local
echo "export GOROOT=/usr/local/go" >> ~/.zshrc
echo "export GOPATH=\$HOME/go" >> ~/.zshrc
echo "export PATH=\$GOPATH/bin:\$GOROOT/bin:\$PATH" >> ~/.zshrc
source ~/.zshrc
```

# Go バージョン更新方法

## 最新版の自動インストール

### Linux

```bash
sudo rm -rf /usr/local/go
latest_stable_go_version=$(curl -Lso- https://go.dev/dl/ | grep -m1 -P 'id="go\d+\.\d+\.\d+' | cut -f 4 -d '"')
wget https://dl.google.com/go/$latest_stable_go_version.linux-amd64.tar.gz
sudo tar -xvf $latest_stable_go_version.linux-amd64.tar.gz
rm -f $latest_stable_go_version.linux-amd64.tar.gz
sudo mv go /usr/local
source ~/.zshrc
```
