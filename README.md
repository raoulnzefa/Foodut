# Foodut

## WEB MAKANAN RINGAN UMKM

### Anggota Kelompok :

1119007 – Timothy Ray
1119023 – Jedediah Fanuel
1119033 – Fedly Septian
1119038 – Elangel Neilea Shaday
1119045 – Rafael Christo Gracia

## Tools, Library, and Framework

### Bahasa Pemrograman

- [Go Language](https://golang.org/)
  <br>
  Go is an open source programming language that makes it easy to build simple, reliable, and efficient software

### DBMS

- [MariaDB](https://www.apachefriends.org/index.html)
  <br>
  The MariaDB we use is the one that comes bundled with the XAMPP app.

### Dukungan Framework

- [Vue.js](https://vuejs.org/)
  <br>
  Vue.js is a progressive, incrementally-adoptable JavaScript framework for building UI on the web

### Dukungan Template

- [Vuexy](https://pixinvent.com/demo/vuexy-vuejs-admin-dashboard-template/demo-1/dashboard/ecommerce)
  <br>
  ArchitectUI is one of the most popular admin dashboard templates ever released. It is used by thousands of developers to create webapps and SaaS totaling 100 million monthly active users.

### Model Deployment

- [AWS Elastic Beanstalk](https://aws.amazon.com/elasticbeanstalk/)
  <br>
  AWS Elastic Beanstalk is an easy-to-use service for deploying and scaling web applications and services developed with Java, .NET, PHP, Node.js, Python, Ruby, Go, and Docker on familiar servers such as Apache, Nginx, Passenger, and IIS.

### Web Application Architecture

- [Three-tier architecture](https://herbertograca.com/2017/08/03/layered-architecture/)
  <br>
  - User Interface (Presentation): The user interface, be it a web page, a CLI or a native desktop application;
  - Business logic (Domain): The logic that is the reason why the application exists;
  - Data source: The data persistence mechanism (DB), or communication with other applications.

## Start Application

### Backend

```
go run main.go
```

### Frontend

### Pindah folder frondend windows

```
cd frontend
```

#### Melakukan installasi VueJS

```
// Pastikan ada di direktori frontend
npm install
```

#### Melakukan development server

```
// Pastikan ada di direktori frontend
- npm run serve
```

<br>

## GIT (Version Control)

#### Mengatur Username & Email

```
git config --global user.name "Adi Budi "
git config --global user.email "adibudi@gmail.com"
* Cara diatas untuk mengatur disemua repositori, untuk spesifik hilangkan "--global"
```

#### Cara Clone Lokal Git

```
git clone https://github.com/elangelshaday/Foodut.git
```

#### Push

```
git add .
git commit -m "Your messages"
git push -u origin main
```

Jika terjadi error seperti ini :

```
To https://github.com/albertusangkuw/Cheems.git
 ! [rejected]        main -> main (non-fast-forward)
error: failed to push some refs to 'https://github.com/albertusangkuw/Cheems.git'
hint: Updates were rejected because the tip of your current branch is behind
hint: its remote counterpart. Integrate the remote changes (e.g.
hint: 'git pull ...') before pushing again.
hint: See the 'Note about fast-forwards' in 'git push --help' for details.
```

Maka selesaikan dengan :

```
git pull origin main
git push -u origin main
```

Jika masih terjadi error ketika pull origin main :
Maka cari file yang konflik kemudian edit pilih yang bermasalah
Kemudian jalankan

```
git add .
git commit -m "Merge to ..isi command..."
git push -u origin main
```

#### Pull

```
git pull origin main
```

#### Remove

```
git rm file --> file is your file name
            --> and then, use PUSH
```

#### Show Config

```
git config --list
```

#### Remove A Commit That Already Pushed

1. `git log` to find out the commit you want to revert
2. `git push origin +7f6d03:master` while `7f6d03` is the commit before the wrongly pushed commit. + was for force push

#### Merging without Auto-Merge

1. `git merge --no-commit --no-ff <local-branch>`
2. `git reset HEAD`
3. To see all diff: `git diff`
