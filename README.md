# switchOS

O gerenciador de processos deve ser capaz de aplicar o algoritmo de escalonamento definido por meio de parâmetro pelo usuário do SO. O gerenciador de memória deve garantir que um processo não acesse as regiões de memória de um outro processo, e que o algoritmo de substituição de página seja adequadamente usado. E o gerenciador de entrada/saída deve ser responsável por administrar o algoritmo especificado para a busca em disco. Cada módulo será testado de acordo com as especificações determinadas abaixo. O sistema receberá como parâmetro um inteiro e um arquivo texto, por exemplo $ 1 processes.txt. O inteiro determina qual módulo deve ser ativado (no exemplo dado significa que será ativado o módulo de processos, pois foi o inteiro 1), e o arquivo texto (com extensão .txt) repassa os dados de entrada necessários para a execução do módulo escolhido.

Parâmetros de entrada:
1) Gerenciador de Processos
2) Gerenciador de Memória
3) Gerenciador de Entrada/Saída
