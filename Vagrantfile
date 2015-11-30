# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "hashicorp/precise64"
  config.vm.provision "shell", inline: <<-SHELL
    if [ ! -f /vagrant/go1.5.1.linux-amd64.tar.gz ]
      then
        echo "Fetching Go binaries. This may take a while...."
        wget -q -O /vagrant/go1.5.1.linux-amd64.tar.gz https://storage.googleapis.com/golang/go1.5.1.linux-amd64.tar.gz
        echo "Binaries Retrieved!"
    fi
    sudo tar -C /usr/local -xzf /vagrant/go1.5.1.linux-amd64.tar.gz
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
    mkdir -p /$HOME/src/go
    echo "export GOPATH=~/src/go" >> ~/.bashrc
    if [ -x /usr/local/go/bin/go ]
      then
        /usr/local/go/bin/go version
      else
        echo "Error installing Go"
	      exit 2
    fi
    export PATH=$PATH:/usr/local/go/bin
    export GOPATH=~/src/go
    cd
    cp /vagrant/compress.go ~/
    echo "Creating Sparse file"
    truncate -s 512M sparse.img
    echo "Proving file is truly sparse"
    ls -lash sparse.img
    go run compress.go
    tar -C non_sparse/ -xf non_sparse/non_sparse.tar
    tar -C sparse/ -xf sparse/sparse.tar
    echo "Proving non-sparse in Go gained size on disk"
    ls -lash non_sparse/sparse.img
    echo "Proving sparse in Go DID NOT keep file size on disk"
    ls -lash sparse/sparse.img
    echo "Compressing via tar w/ Sparse Flag set"
    mkdir -p tar
    tar -Scf tar/sparse.tar sparse.img
    tar -C tar/ -xf tar/sparse.tar
    echo "Proving sparse via tar DID keep file size on disk"
    ls -lash tar/sparse.img
  SHELL
end
