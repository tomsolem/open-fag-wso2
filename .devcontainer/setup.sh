# bin/bash!
echo "Setting up the development container..."
sudo apt update -y
# PRETTY_NAME="Ubuntu 24.04.2 LTS"
sudo apt-get upgrade -y

# install the PostgreSQL CLI
echo "Installing PostgreSQL client..."
sudo apt-get install -y postgresql-client
# install software-properties-common
#sudo apt-get install -y software-properties-common
# install golang
#sudo add-apt-repository -y ppa:longsleep/golang-backports
#sudo apt-get update -y
#sudo apt-get install -y golang-1.24
echo "Installing Go 1.24.5..."
wget -qO /tmp/go1.24.5.linux-amd64.tar.gz https://go.dev/dl/go1.24.5.linux-amd64.tar.gz
# extract the tarball
sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf /tmp/go1.24.5.linux-amd64.tar.gz
#sudo tar -C /usr/local -xzf /tmp/go1.24.5.tar.gz

# Set environment variables for the current script's execution
echo "Setting up Go environment variables..."
export GOPATH=$HOME/go
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin

# set up the environment variables for Go in .zshrc and .bashrc
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.zshrc
echo "export GOPATH=\$HOME/go" >> ~/.zshrc
echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.zshrc

echo "export PATH=\$PATH:\$GOPATH/bin" >> ~/.bashrc
echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
echo "export GOPATH=\$HOME/go" >> ~/.bashrc
# source the .bashrc to apply changes
echo "Sourcing .bashrc..."
source ~/.bashrc
echo "checking Go version..."
go version
# install golang development tools
echo "Installing Go development tools..."
go install -v github.com/go-delve/delve/cmd/dlv@latest

# download the cli
wget -qO /tmp/fga.tar.gz https://github.com/openfga/cli/releases/download/v0.7.2/fga_0.7.2_linux_amd64.tar.gz
# extract the cli
tar -xzf /tmp/fga.tar.gz -C /tmp
# move the cli to /usr/local/bin
sudo mv /tmp/fga /usr/local/bin/fga
# make the cli executable
sudo chmod +x /usr/local/bin/fga