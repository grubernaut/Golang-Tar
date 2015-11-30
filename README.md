Golang Sparse files and Tar
===========================

Proving that Golang, as of Go1.5.1 does not maintain sparse files during file compression with `compress/tar`.

Steps to reproduce:

```bash
vagrant up
```

Vagrant will:

* Install Go1.5.1

* Create a Sparse file on the Linux Filesystem by:
```bash
# truncate -s 512M sparse.img
```

* Verify that the created file has an apparent size of 512M, but an actual size of 0M
```bash
# ls -lash sparse.img
0 -rw-r--r-- 1 root root 512M Nov 30 21:29 sparse.img
```

* Run `compress.go`, which uses `archive/tar` to compress the sparse file both with and without the sparse type set in the archive header. This creates directories of `non_sparse/` and `sparse/` for non-sparse and sparse compression, named respectively.

* Extract both of the archives created from `compress.go`.
```bash
# tar -C non_sparse/ -xf non_sparse/non_sparse.tar
# tar -C sparse/ -xf sparse/sparse.tar
```

* Verifies that Go did NOT keep file sparseness regardless of whether the GNUSparse type was set in the archive header or not.
```bash
# ls -lash non_sparse/sparse.img
512M -rw-r--r-- 1 root root 512M Nov 30 21:29 non_sparse/sparse.img
# ls -lash sparse/sparse.img
513M -rw-r--r-- 1 root root 512M Nov 30 21:29 sparse/sparse.img
```

* Compresses the sparse file using GNU Tar, maintaining the sparseness of the file
```bash
tar -Scf tar/sparse.tar sparse.img
```

* Extract the archive created from GNU Tar.
```bash
tar -C tar/ -xf tar/sparse.tar
```

* Verify that compression via Tar, maintains the sparse file.
```bash
ls -lash tar/sparse.img
0 -rw-r--r-- 1 root root 512M Nov 30 21:29 tar/sparse.img
```
