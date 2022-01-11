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
* [quick start](#quickstart)
  * [hello world](#hello-world)
  * [tipe data](#tipe-data)
  * * [boolean](#boolean)
  * * [number](#int-float)
  * * [string](#string)
  * [variable](#variable)
  * [constant](#constant)
  


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
## variable
variable adalah tempat untuk menamakan data, menyimpan data, inisialisasi data dll
ada beberapa cara di bahasa pemograman go untuk membuat variable, check this out:
```
func main()  {
	func main(){
	//method 1
	var name string

	name = "Jennie"
	fmt.Println(name)
	name = "Jennie Kim"
	fmt.Println(name)

	//method 2
	var name2 = "rose"
	fmt.Println(name2)
	name2 = "roseanne park"
	fmt.Println(name2)

	//method 3
	name3 := "jiso"
	fmt.Println(name3)
	name3 = "jisoya"
	fmt.Println(name3)

	//method 4
	var (
		firstName = "lalisa"
		lastname = "manoban"
	)
	fmt.Println(firstName)
	fmt.Println(lastname)
}
}
```
## constant
const sama seperti variable TAPI bedanya jika constan nilainya konsisten, contoh jika i = 10 lalu ada sebuah kondisi dimana i berubah menjadi angka lain atau tipe data lain misal i = 20, jika menggunakan const itu tidak bisa dilakukan karena constan ya konsisten lebih baik pake var jika tau nilainya akan berubah.

```
func main(){
	// single constant
	// const firstName string = "hindia"
	// const lastName = "pamungkas"
	// const value = 100

	// multiple constant
	const(
		firstName = "hindia"
		lastName = "pamungkas"
		value = 200
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

}
```

