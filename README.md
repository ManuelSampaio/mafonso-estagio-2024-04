
### Empresa Acme, Lda

# Menbros do grupo
1. Manuel Afonso
2. Novais David

## Instrução para rodar o programa 

1. Para gerar o  executavel entra no diretório/pasta da aplicação a partir da linha de comando:  
```
  go build -o ACME main.go
```
1.2 Para gerar o executavel em um diretorio/pasta diferente, a partir linha de comando:
  ```
 go build -o ~/[o_caminho_onde_pretendes_colocar_o_executavel]/[nome_do_executavel] main.go
 ex.: go build -o ~/Documents/n/Novais main.go   
```
2. para executar a aplicação deve executar o seguinte comando:
```
  ./ACME expedirMercadoria -i [identificador do produto] -q [quantidade do produto a expedir]
```
2.1 Ou 
```
  ./ACME expedirMercadoria --idProduto [identificador do produto] --quantidade [quantidade do produto a expedir]
```
