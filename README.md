## Mudanças que impactaram na especificação de requisitos
- Chaves geradas pela aplicação
- Trabalhando sempre com valores positivos, utilizando o tipo de operação para fazer a definição entre débito e crédito em carteira
- Preparado para vários tipos de documentos

## Modelagem de dados atualizada

### accounts
| Account_ID  | Document_Type | Document_Number |
| ------------- | ------------- | ------------- |
| 4e989872-6cc9-4746-a91c-8228afba818b  | CPF  | 39758803093 |

### operations
| Operation_ID  | Description | Finality |
| ------------- | ------------- | ------------- |
| 38cc4090-1b33-4f11-a8e2-cdb3a9d260ef  | COMPRA A VISTA  | DEBIT |
| 71cfe893-6cb5-40f8-a065-59a9ab39eb41  | COMPRA PARCELADA  | DEBIT |
| edf12281-cf0b-421f-9480-b1fb3c999ac7  | SAQUE  | DEBIT |
| 4cebc9d2-f3e1-4d75-bcec-bb30b3afce00  | PAGAMENTO  | CREDIT |

## Como rodar?
```sh
// rodando o projeto em desenvolvimento
make run

// rodando os testes
make test
```

## Exemplos de requests
```curl
// show account details
curl -X GET http://localhost:8080/accounts/4e989872-6cc9-4746-a91c-8228afba818b

// create account
curl -X POST \
  http://localhost:8080/accounts \
  -H 'Content-Length: 37' \
  -H 'Content-Type: application/json' \
  -d '{
	"document_number": "08732860900"
}'

// create transaction
curl -X POST \
  http://localhost:8080/transactions \
  -H 'Content-Length: 136' \
  -H 'Content-Type: application/json' \
  -d '{
"account_id": "4e989872-6cc9-4746-a91c-8228afba818b",
"operation_type_id": "38cc4090-1b33-4f11-a8e2-cdb3a9d260ef",
"amount": 123.45
}'
```



## O que eu queria ter feito mais?
- Mais testes
- Melhorar a camada de iteração entre request e caso de uso
- Logs :/
- Melhorar as tratativas de erros
