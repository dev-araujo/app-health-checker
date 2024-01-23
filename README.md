# app-health-checker

Esse é uma API feita em **Go**  que verifica se é possível estabelecer uma **conexão TCP bem-sucedida com o destino (site)** na porta especificada. 


## Instruções para rodar✨

No terminal, clone o projeto:

```
git clone https://github.com/dev-araujo/app-health-checker.git
```

Execute a aplicação:

```
go run .
```



## Como usar 🧑🏻‍💻: 

```
http://localhost:8080/healthcheck?domain={DOMINIO}
```


## Exemplo 💡: 

### Requisição :
```
http://localhost:8080/healthcheck?domain=google.com
```
### Resposta de sucesso:

```
{
  "status": "[SUCCESS]",
  "dominio": "google.com",
  "from": "192.168.1.104",
  "to": "142.250.79.14"
}
```

### Resposta de falha 

```
{
  "status": "ERROR",
  "dominio": "google.com",
  "error": "Error: dial tcp4: lookup google.com on 142.250.79.14: no such host"
}
```





## 🔨 Pacotes

- [gorilla/mux](https://github.com/gorilla/mux)

---




#### Author 👷

<img src="https://user-images.githubusercontent.com/97068163/149033991-781bf8b6-4beb-445a-913c-f05a76a28bfc.png" width="5%" alt="caricatura do autor desse repositório"/>

[![linkedin](https://img.shields.io/badge/LinkedIn-0077B5?style=for-the-badge&logo=linkedin&logoColor=white)](https://www.linkedin.com/in/araujocode/)
