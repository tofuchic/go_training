# Go インストール方法

```bash
wget https://dl.google.com/go/go1.21.4.linux-amd64.tar.gz
sudo tar -xvf go1.21.4.linux-amd64.tar.gz
sudo mv go /usr/local
echo "export GOROOT=/usr/local/go" >> ~/.zshrc
echo "export GOPATH=\$HOME/go" >> ~/.zshrc
echo "export PATH=\$GOPATH/bin:\$GOROOT/bin:\$PATH" >> ~/.zshrc
source ~/.zshrc
```
