# 02_Version Control and Branch Managemen
1.	Pengertian  git 
Git merupakan salah  satu Versioning Control System popular yang digunakan untuk mengembangkan software secara bersama-sama.
2.	Repository merupakan sebutan folder project didalam git. Yang terdapat folder, file dan history yang biasa disebut dengan commit.
3.	Pengistalan Git
Penginstalan git dapat dilakukan di OS, MAC, WINDOWS, dan LINUX. 
4.	Setingg Up Git
	Congfit : git congfit –global user .name  “isi _username”
: git congfit –global user .name  “isi _email”
5.	Ada 2 cara untuk menset secara global, username dan email yang terdaftar di platform github, gitlab dan lainnya.
	Git init merupakan  perintah yang berfungsi untuk membuat folder local include dengan config git secara global.
	Clone : git clone <link ssh/link hhtp_github>
: git clone -b <nama _branch> <link ssh/link hhtp_github>
Yang berfungsi untuk mengclone repository online ke local dengan nama folder sesuai nama repo.
	Git staged Area, area disini biasa disebut perubahan. File-file atau perubahan yang kita lakukan akan diatur seperti melakukan commit dan push.
•	git add  untukmenambahkan file ke area siap untuk di commit
•	git commit untuk melakukan commit di branch yang kita lakukan perubahan, setiap commit ada history untuk mempermudah kita .
•	git commit –m “message” dalam melakukan commit harus meninggalkan pesan berupa informasi terkait apa yang sudah dilakukan didalam branch, seperti perubahan dan lainnya. Untuk memperjelas status log commit kita. Karena itu pesan atau message kita harus jelas.
•	git push origin <nama_branch> adlah untuk melakukan push branch local ke branch online dan mengupdatesemua perubahannya.
	Git diff dan stash
•	git diff --staged untuk melakukan check terhadap perubahan apa saja yang dilakukan di area staged.
•	git stash untuk melakukan pengambilan code yang di lakukan perubahan.
•	git stash pop adalah untuk melakukan paste atau meletakan code kita yang masuk ke stash. Stash juga punya log dan di urut berdasarkan stash terakhir.
	Perintah git lainnya 
•	git pull origin <nama_branch> untuk melakukan sinkrosinasi antara local dengan online branch, yang nantinya akan melakukan get branch online tersebut danmelakukan pembaharuan di local.
•	git log --oneline
•	git log --graph 
untuk melihat log dan history commit yang dilakukan dibranch , semua log akan terlihat.
•	git checkout <nama_branch>
•	git checkout  -b <nama_branch>
untuk melakukan perpindahan branch atau posisi branch. Flag –b untuk melakukan pidah sekaligus membuat branch baru di local.
•	git reset  --soft 
•	git reset  --hard 
untuk melakukan reset commit berdasarkan flagnya. Biasanya jarang digunakan karena baerbahaya untuk branch, sebaiknya kembali menggunakan reset soft.
•	git add <ssh/http_github> :untuk melakukan sinkronisasi link repo yang diambil.
•	git remote –v :untuk melakukan check agar mengetahui link atau posisi link repository online.
	Git Branching
•	git branch --list :untuk mengetahi branch apa saja yang sudah tercipta.
•	git branch < branch> :untuk menciptakan branch yang baru dan nama branch.
•	git branch –D < branch> :untuk menghapus branch secara spesifik
•	git branch –a :untuk list remote branch
•	git push –u original developer  : agar branch baru (developer) masuk ke repository digithub
•	git commit –m “tag h1 developer” : agar masuk ke arah sync changenya, jika sudah masuk klik ok agar masuk ke ropo di github
•	git merge <nama_branch> : di dalam branch akan menerima penggabungan dan nama_branch adalah nama branch yang akan digabungkan.
	Conflict terjadi ketika melakukan perubahan di branch feature/fiture diline yang sama dan kemudian melakukan ke merge ke development yang bertujuan menggabungkan 2 branch fitur yang berbeda, jika terdapat perubahan yangberbeda di line yang sama akan mendapat info conflict ketika merge.
	Solusinya langsung melakukan  perubahan di setiap commit dan lakukan perbaikan.
	Pull Request adlah suatu permintaan untuk menggabungkan atau merge kode yang kita modifikasi dengan repository lain.
