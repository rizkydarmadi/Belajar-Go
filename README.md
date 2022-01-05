<!--lint disable no-literal-urls-->

<p align="center">
  <a href="https://go.dev/">
    <img
      alt="golang"
      src="https://miro.medium.com/max/1200/0*1cn_mKQxd-Z0ikvT.png"
      width="400"
    />
  </a>
</p>

Go adalah bahasa pemograman open source yang sudah didukung oleh Google
<br>
[official documentation](https://go.dev/)
# Table of contents

* [instalasi](#instalasi)
* [getting started](#getting-started)
  * [hello world](#hello-world)
  * [tipe data](#tipe-data)
  * * [boolean](#boolean)
  * * [number](#int-float)
  * * [string](#string)
  


## instalasi
Perintah untuk menginstall di linux(ubuntu 20.04)
```console
$ sudo apt install golang-go
```
perintah untuk mengecek apakah sudah terinstal si Go nya?
```console
$ go version
```
gampang kan ckck

## quickstart
### hello world

perintah untuk cloning dari repo ini
pre requisite --> mesti banget punya pengetahuan minimal basic tentang version control/git
```console
$ git clone https://github.com/rizkydarmadi/Belajar-Go.git
```
perintah untuk pindah ke directory yang ada source codenya
```console
$ cd Basic/Helloworld
```
peritah untuk build dari source code ke binary file
```console
$ cd go build helloworld.go
```
lalu cek binary filenya
```console
$ ~/go/Basic/Helloworld$ ls
helloworld  helloworld.go
```
itu yang gak ada extensi .go adalah binary file yang akan kita run
yuk kita run
```console
$ ./helloworld 
hello world
```
jika sudah tampil tulisanya maka sudah berhasil di run
untuk cara cepatnya gini, tapi ini hanya untuk development ya bukan untuk production
```console
$ go run helloworld.go
hello world
```
## tipe-data
### boolean
```
func main()  {
	fmt.Println("benar = ",true)
	fmt.Println("salah = ",false)
}
```
### int-float
```
func main()  {
	fmt.Println("satu = ",1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima = ", 3.5)
}
```
### string
```
func main()  {
	fmt.Println("taylor")
	fmt.Println("taylor swift")
	fmt.Println("taylor alisson swift")

	fmt.Println(len("taylor")) //menampilkan jumlah string
	fmt.Println("taylor swift"[0]) // menampilkan string index ke 0 - tapi hanya byte
	fmt.Println("taylor swift"[1]) // menampilkan string index ke 1 - tapi hanya byte
}
```



