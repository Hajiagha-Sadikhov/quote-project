# Quote Project

### **Proyektin Strukturu**

Bu proyekt, Docker üzərində işləyən iki modüldən ibarətdir:

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


###To run the project using Docker Compose, follow these step.
```bash

docker-compose up --build

```

###To stop the running containers, use the following command.
```bash

docker-compose down


```


### **Go Tədris Mənbələri**
Go dilini öyrənmək üçün bəzi faydalı mənbələr:
- [Go Rəsmi Sənədləri](https://go.dev/doc/)
- [Go by Example](https://gobyexample.com/)
- [GoLang Təlimi by freeCodeCamp.org - YouTube](https://youtu.be/un6ZyFkqFKo?si=YdrLUV52lBoWbiOF)

---

### **Docker Tədris Mənbələri**
Docker öyrənmək üçün aşağıdakı mənbələr faydalıdır:
- [Docker Rəsmi Sənədləri](https://docs.docker.com/)
- [Docker Rəsmi Təlimatları](https://docs.docker.com/guides/)
- [Docker Tutorial for Beginners - YouTube](https://www.youtube.com/watch?v=fqMOX6JJhGo)
- [Docker Mastery Kursu - Udemy](https://www.udemy.com/course/docker-mastery/)

## License

This project is licensed under the [Apache License 2.0](https://www.apache.org/licenses/LICENSE-2.0).

