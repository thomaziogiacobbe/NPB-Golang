# NPB-Golang
Benchmarks NPB versão Go para a disciplina de TEC VII.

Basta baixar e instalar a versão mais recento do Go e para compilar executar o comando **`go build`** na pasta onde está localizado **main.go**. Será criado um executável chamado NPB-Golang.

Para executar qualquer benchmark, dois parâmetros obrigatórios devem ser especificados em ordem, sendo eles, o benchmark a ser executado e a classe do benchmark:

* Possíveis benchmarks: EP e IS
* Possíveis classes: S, W, A, B, C, D ou E

**_DETALHE:_** a especificação desses parametros obrigatórios segue o padrão do próprio NPB ao compilar os benchmarks, isso significa que para especificar deve-se seguir a seguinte forma `CLASS={classe}` onde `classe` é uma das possíveis classes acima.

Também existem parâmetros opcionais, sendo eles em ordem:

* `-f:` um arquivo será criado e nele será escrito o resultado do tempo de execução do benchmark, esse arquivo é acessado em modo **APPEND**, com isso multiplas execuções do mesmo benchmark vão escrever no mesmo arquivo. O nome padrão do arquivo segue o formato **{BENCHMARK}_{CLASSE}.txt**. É possível alterar o nome do arquivo de saida com um parametro opcional após `-f`, para esse parâmetro não é necessario indicar a extensão do arquivo

* `USE_BUCKETS:` parâmetro exclusivo para o benchmark **IS**, indica que o benchmark vai fazer o uso de buckets durante a execução, caso não seja especificado o benchmark vai ser executado sem buckets