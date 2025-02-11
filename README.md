# Quote Project

### **Proyektin Strukturu**

Bu proyekt, Docker üzərində işləyən iki moduldan ibarətdir:

#### **1. API**
Go (net/http) istifadə edilərək inkişaf etdirilən bu servis, `quotes.json` faylından sitatları təqdim edir. API, istifadəçilərin müxtəlif sitatlara asanlıqla daxil olmasını təmin edir.

- **Texnologiya:** Go (net/http)
- **Funksiyalar:**
  - Təsadüfi sitat göstərilməsi
  - Sitat ID-sinə əsasən axtarış
  - Səhifə üzrə sitat siyahısı
- **Port:** **5000**

#### **2. Web**
Web modulu, sadə HTML, CSS və JavaScript istifadə edilərək hazırlanmış bir interfeysdən ibarətdir. Veb səhifə, API-dən məlumat çəkərək istifadəçiyə sitatları vizual olaraq təqdim edir. İstifadəçi, sitatları təsadüfi olaraq görməklə yanaşı, ID ilə axtarış edə və səhifə üzrə sitat siyahılarını görə bilər.

- **Texnologiyalar:** HTML, CSS, JavaScript
- **Funksiyalar:**
  - API-dən təsadüfi sitat çəkmək
  - Sitat ID-sinə əsasən axtarış etmək
  - Səhifə üzrə sitat siyahılarını göstərmək
- **Port:** **80**



## 👨‍💻 Running quote-project
### 📥 Clone the Repository

```bash
git clone https://github.com/Hajiagha-Sadikhov/quote-project.git
cd quote-project
```


To run the project using Docker Compose, follow these step.
```bash

docker-compose up --build

```

To stop the running containers, use the following command.
```bash

docker-compose down


```


### **Go Tədris Mənbələri**
Go dilini öyrənmək üçün bəzi faydalı mənbələr:
- [Go Rəsmi Sənədləri](https://go.dev/doc/)
  Go dilinə aid rəsmi sənədlər və təlimatlar.

- [Go by Example](https://gobyexample.com/)
  Go dilini nümunələr vasitəsilə öyrənmək üçün əla bir qaynaq.

- [GoLang Təlimi by freeCodeCamp.org - YouTube](https://youtu.be/un6ZyFkqFKo?si=YdrLUV52lBoWbiOF)
  Go dili haqqında ətraflı video formatında məlumatlar.

---

### **Docker Tədris Mənbələri**
Docker öyrənmək üçün aşağıdakı mənbələr faydalıdır:
- [Docker Rəsmi Sənədləri](https://docs.docker.com/)
  Docker-ın rəsmi sənədləri, konteynerlər və digər mövzular haqqında ətraflı məlumat.

- [Docker Rəsmi Təlimatları](https://docs.docker.com/guides/)
  Docker ilə bağlı müxtəlif təlimatlar və ən yaxşı təcrübələr.

- [Docker Tutorial for Beginners - YouTube](https://www.youtube.com/watch?v=fqMOX6JJhGo)
  Başlayanlar üçün Docker video formatında məlumatlar.

- [Docker Mastery Kursu - Udemy](https://www.udemy.com/course/docker-mastery/)
  Tətbiq inkişaf etdiriciləri üçün Docker-ın əsasları və necə istifadə edəcəyiniz haqqında faydalı məlumatlar.

### **Nginx Üzrə Faydalı Mənbələr**

- [Nginx Rəsmi Sənədləri](https://nginx.org/en/docs/)  
  Nginx-in rəsmi sənədləri, konfiqurasiya, modullar və ən yaxşı təcrübələr haqqında ətraflı məlumat təqdim edir.

- [Nginx Başlanğıc Təlimatı](https://www.digitalocean.com/community/tutorials/an-intuitive-guide-to-nginx-configuration)  
  Nginx ilə işləməyə başlamaq üçün istifadəçilərə uyğun və asan başa düşülən bir təlimat.

- [Nginx CookBook](https://www.nginx.com/resources/wiki/)  
  Müxtəlif ssenarilər üçün Nginx-i effektiv şəkildə konfiqurasiya etmək üçün istifadə edilə bilən istifadə hallarının və ən yaxşı təcrübələrin toplusu.

- [Nginx Performansını Tuning Etmək](https://www.nginx.com/blog/tuning-nginx/)  
  Nginx ilə yüksək performanslı tətbiqləri işlətmək üçün performans optimizasiyası və ən yaxşı təcrübələr haqqında məlumat.


## License

This project is licensed under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0).

## 📝 Notes  
📌 Some parts of this documentation were improved with the help of AI tools to enhance clarity and structure.
