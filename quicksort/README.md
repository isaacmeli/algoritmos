QUICKSORT

-------
O algoritmo quicksort(ordenação rápida) aplica o paradigma de divisão e conquista.

Devemos seguir três etapas para executar o processo:

- Divisão:
  - Particionar(reorganizar) o arranjo `A[p .. r]` em dois subarranjos(possivelmentes vazios) `A[p .. q-1]` e `A[q+1 .. r]` 
  tais que, cada elemento de `A[p .. q-1]` é menor ou igual a `A[q]` que, por sua vez,é menor ou iguala cada elemento de 
  `A[q+1 .. r]`. Calcular o indice `q` como parte desse procedimento de particionamento.
- Conquista:
  - Ordenar os dois subarranjos `A[p .. q-1]` e `A[q+1 .. r]` por chamadas recursivas a quicksort.
- Combinação:
  - Como os subarranjos já estão ordenados, não é neccessario nenhhum trabalho para combinalos.